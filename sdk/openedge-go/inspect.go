package openedge

import (
	"time"

	"github.com/baidu/openedge/utils"
)

// Inspect all openedge information and status inspected
type Inspect struct {
	// exception information
	Error    string    `json:"error"`
	// inspect time
	Time     time.Time `json:"time"`
	// software information
	Software Software  `json:"software,omitempty"`
	// hardware information
	Hardware Hardware  `json:"hardware,omitempty"`
	// service information, including service name, instance running status, etc.
	Services Services  `json:"services,omitempty"`
	// storage volume information, currently not implemented
	// Volumes  []VolumeStatus `json:"volumes,omitempty"`
}

// Software software information
type Software struct {
	// operating system information of host
	OS          string `json:"os,omitempty"`
	// CPU information of host
	Arch        string `json:"arch,omitempty"`
	// OpenEdge running mode of application services
	Mode        string `json:"mode,omitempty"`
    // OpenEdge compiled Golang version
	GoVersion   string `json:"go_version,omitempty"`
	// OpenEdge release version
	BinVersion  string `json:"bin_version,omitempty"`
	// OpenEdge loaded application configuration version
	ConfVersion string `json:"conf_version,omitempty"`
}

// Hardware hardware information
type Hardware struct {
	// memory usage information of host
	MemInfo  *utils.MemInfo  `json:"mem_stats,omitempty"`
	// CPU usage information of host
	CPUInfo  *utils.CPUInfo  `json:"cpu_stats,omitempty"`
	// disk usage information of host
	DiskInfo *utils.DiskInfo `json:"disk_stats,omitempty"`
	// CPU usage information of host
	GPUInfo  []utils.GPUInfo `json:"gpu_stats,omitempty"`
}

// Services all services' information
type Services []ServiceStatus

// ServiceStatus service status
type ServiceStatus struct {
	Name      string           `json:"name,omitempty"`
	Instances []InstanceStatus `json:"instances,omitempty"`
}

// InstanceStatus service instance status
type InstanceStatus map[string]interface{}

// NewInspect create a new information inspected
func NewInspect() *Inspect {
	return &Inspect{
		Services: Services{},
	}
}

// NewServiceStatus create a new service status
func NewServiceStatus(name string) ServiceStatus {
	return ServiceStatus{
		Name:      name,
		Instances: []InstanceStatus{},
	}
}
