package node

import (
	"encoding/json"
	"os"
	"runtime"
	"time"

	"github.com/baetyl/baetyl-go/v2/context"
	"github.com/baetyl/baetyl-go/v2/errors"
	"github.com/baetyl/baetyl-go/v2/http"
	"github.com/baetyl/baetyl-go/v2/log"
	"github.com/baetyl/baetyl-go/v2/mqtt"
	v1 "github.com/baetyl/baetyl-go/v2/spec/v1"
	"github.com/baetyl/baetyl-go/v2/utils"
	routing "github.com/qiangxue/fasthttp-routing"
	bh "github.com/timshannon/bolthold"
	bolt "go.etcd.io/bbolt"
)

const OfflineDuration = 40 * time.Second
const KeyNodeProps = "nodeProps"

// Node node
type Node struct {
	tomb  utils.Tomb
	log   *log.Logger
	id    []byte
	store *bh.Store
	mqtt  *mqtt.Client
}

// NewNode create a node with shadow
func NewNode(store *bh.Store) (*Node, error) {
	m := &v1.Node{
		CreationTimestamp: time.Now(),
		Desire:            v1.Desire{},
		Report: v1.Report{
			"core": v1.CoreInfo{
				GoVersion:   runtime.Version(),
				BinVersion:  utils.VERSION,
				GitRevision: utils.REVISION,
			},
			"node":      nil,
			"nodestats": nil,
			"apps":      nil,
			"sysapps":   nil,
			"appstats":  nil,
		},
	}
	n := &Node{
		id:    []byte("baetyl-edge-node"),
		store: store,
		log:   log.With(log.Any("core", "node")),
	}
	err := n.insert(m)
	if err != nil && errors.Cause(err) != bh.ErrKeyExists {
		return nil, errors.Trace(err)
	}
	// report some core info
	_, err = n.Report(m.Report)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return n, nil
}

// Get returns node model
func (n *Node) Get() (m *v1.Node, err error) {
	err = n.store.Bolt().View(func(tx *bolt.Tx) error {
		b := tx.Bucket(n.id)
		prev := b.Get(n.id)
		if len(prev) == 0 {
			return errors.Trace(bh.ErrNotFound)
		}
		m = &v1.Node{}
		return errors.Trace(json.Unmarshal(prev, m))
	})
	return
}

// Desire update shadow desired data, then return the delta of desired and reported data
func (n *Node) Desire(desired v1.Desire) (delta v1.Desire, err error) {
	err = n.store.Bolt().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(n.id)
		prev := bucket.Get(n.id)
		if len(prev) == 0 {
			return errors.Trace(bh.ErrNotFound)
		}
		m := &v1.Node{}
		err := json.Unmarshal(prev, m)
		if err != nil {
			return errors.Trace(err)
		}
		if m.Desire == nil {
			m.Desire = desired
		} else {
			err = m.Desire.Merge(desired)
			if err != nil {
				return errors.Trace(err)
			}
		}
		curr, err := json.Marshal(m)
		if err != nil {
			return errors.Trace(err)
		}
		err = bucket.Put(n.id, curr)
		if err != nil {
			return errors.Trace(err)
		}
		delta, err = m.Desire.Diff(m.Report)
		return errors.Trace(err)
	})
	return
}

// Report update shadow reported data, then return the delta of desired and reported data
func (n *Node) Report(reported v1.Report) (delta v1.Desire, err error) {
	err = n.store.Bolt().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(n.id)
		prev := bucket.Get(n.id)
		if len(prev) == 0 {
			return errors.Trace(bh.ErrNotFound)
		}
		m := &v1.Node{}
		err := json.Unmarshal(prev, m)
		if err != nil {
			return errors.Trace(err)
		}
		if m.Report == nil {
			m.Report = reported
		} else {
			err = m.Report.Merge(reported)
			if err != nil {
				return errors.Trace(err)
			}
		}
		curr, err := json.Marshal(m)
		if err != nil {
			return errors.Trace(err)
		}
		err = bucket.Put(n.id, curr)
		if err != nil {
			return errors.Trace(err)
		}
		delta, err = m.Desire.Diff(m.Report)
		return errors.Trace(err)
	})
	return
}

// GetStatus get status
// TODO: add an error handling middleware like baetyl-cloud @chensheng
func (n *Node) GetStats(ctx *routing.Context) (interface{}, error) {
	node, err := n.Get()
	if err != nil {
		return nil, errors.Trace(err)
	}
	node.Name = os.Getenv(context.KeyNodeName)
	view, err := node.View(OfflineDuration)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return view, nil
}

func (n *Node) GetNodeProperties(ctx *routing.Context) (interface{}, error) {
	node, err := n.Get()
	if err != nil {
		return nil, errors.Trace(err)
	}
	if node.Desire == nil {
		return nil, nil
	}
	return node.Desire[KeyNodeProps], nil
}

func (n *Node) UpdateNodeProperties(ctx *routing.Context) (interface{}, error) {
	node, err := n.Get()
	if err != nil {
		return nil, errors.Trace(err)
	}
	var delta v1.Delta
	err = json.Unmarshal(ctx.Request.Body(), &delta)
	if err != nil {
		http.RespondMsg(ctx, 500, "UnknownError", err.Error())
		return nil, errors.Trace(err)
	}
	for _, v := range delta {
		if _, ok := v.(string); !ok {
			return nil, errors.Trace(errors.New("value is not string"))
		}
	}
	var oldReport v1.Report
	if node.Report == nil {
		node.Report = map[string]interface{}{}
	}
	reportVal := node.Report[KeyNodeProps]
	if reportVal == nil {
		reportVal = map[string]interface{}{}
	}
	oldReport, ok := reportVal.(map[string]interface{})
	if !ok {
		return nil, errors.Trace(errors.New("old node props is invalid"))
	}
	newReport, err := oldReport.Patch(delta)
	if err != nil {
		return nil, errors.Trace(err)
	}
	node.Report[KeyNodeProps] = map[string]interface{}(newReport)
	if _, err = n.Report(node.Report); err != nil {
		return nil, errors.Trace(err)
	}
	return newReport, nil
}

// Get insert the whole shadow data
func (n *Node) insert(m *v1.Node) error {
	return n.store.Bolt().Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(n.id)
		if err != nil {
			return errors.Trace(err)
		}
		data := b.Get(n.id)
		if len(data) != 0 {
			return errors.Trace(bh.ErrKeyExists)
		}
		data, err = json.Marshal(m)
		if err != nil {
			return errors.Trace(err)
		}
		return errors.Trace(b.Put(n.id, data))
	})
}
