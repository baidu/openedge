package cmd

import (
	"fmt"
	"os"

	"github.com/baidu/openedge/logger"
	"github.com/baidu/openedge/master"
	"github.com/baidu/openedge/sdk/openedge-go"
	"github.com/baidu/openedge/utils"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start openedge",
	Long:  ``,
	Run:   start,
}

func init() {
	startCmd.Flags().StringVarP(&workDir, "workdir", "w", "", "work directory of openedge")
	startCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config path of openedge")
	rootCmd.AddCommand(startCmd)
}

func start(cmd *cobra.Command, args []string) {
	err := startInternal()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startInternal() error {
	isOTA := utils.FileExists(openedge.DefaultBinBackupFile)
	cfg, err := checkInternal()
	log := logger.InitLogger(cfg.Logger, "openedge", "master")
	if isOTA {
		log = logger.New(cfg.OTALog, "type", openedge.OTAMST)
	}
	if err != nil {
		if isOTA {
			log = log.WithField(openedge.OTAKeyStep, openedge.OTARollingBack)
		}
		log.WithError(err).Infof("failed to start master")
		rberr := master.RollBackMST()
		if rberr != nil {
			log.WithField(openedge.OTAKeyStep, openedge.OTAFailure).WithError(rberr).Infof("failed to roll back")
			return fmt.Errorf("failed to start master: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		if isOTA {
			log.WithField(openedge.OTAKeyStep, openedge.OTARolledBack).Infof("master is rolled back")
		}
		return fmt.Errorf("failed to start master: %s", err.Error())
	}

	m, err := master.New(workDir, *cfg, Version)
	if err != nil {
		if isOTA {
			log = log.WithField(openedge.OTAKeyStep, openedge.OTARollingBack)
		}
		log.WithError(err).Infof("failed to start master")
		rberr := master.RollBackMST()
		if rberr != nil {
			log.WithField(openedge.OTAKeyStep, openedge.OTAFailure).WithError(rberr).Infof("failed to roll back")
			return fmt.Errorf("failed to start master: %s; failed to roll back: %s", err.Error(), rberr.Error())
		}
		if isOTA {
			log.WithField(openedge.OTAKeyStep, openedge.OTARolledBack).Infof("master is rolled back")
		}
		return fmt.Errorf("failed to start master: %s", err.Error())
	}
	defer m.Close()
	master.CommitMST()
	if isOTA {
		log.WithField(openedge.OTAKeyStep, openedge.OTAUpdated).Infof("master is updated")
	}
	return m.Wait()
}
