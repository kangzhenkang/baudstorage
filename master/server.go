package master

import (
	"fmt"
	"sync"

	"github.com/tiglabs/baudstorage/util/config"
	"github.com/tiglabs/baudstorage/util/log"
)

//config keys
const (
	Module      = "master"
	HttpPort    = "httpPort"
	LogDir      = "logDir"
	RootUrlPath = "/"
)

var gLog *log.Log

type Master struct {
	config  *config.Config
	cluster *Cluster
	wg      sync.WaitGroup
}

func (m *Master) Start(cfg *config.Config) (err error) {
	if err = m.parseConfig(cfg); err != nil {
		return
	}
	m.startHttpService()
	return nil
}

func (m *Master) parseConfig(cfg *config.Config) (err error) {
	logDir := cfg.GetString(LogDir)
	if logDir == "" {
		return fmt.Errorf("bad config file,logDir is null")
	}

	if gLog, err = log.NewLog(logDir, Module, 0); err != nil {
		return fmt.Errorf("failed to initLog, err %s ", err.Error())
	}
	return
}

func (m *Master) Shutdown() {
	panic("implement me")
}

func (m *Master) Sync() {
	panic("implement me")
}

func NewServer() *Master {
	return &Master{}
}
