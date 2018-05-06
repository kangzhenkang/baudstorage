package metanode

import (
	"github.com/tiglabs/baudstorage/raftopt"
	"github.com/tiglabs/raft"
)

// StartRaftServer init address resolver and raft server instance.
func (m *MetaNode) startRaftServer() (err error) {
	var resolver *raftopt.Resolver
	resolver = raftopt.NewResolver()
	var server *raft.RaftServer
	raftConfig := raft.DefaultConfig()
	raftConfig.Resolver = resolver
	if server, err = raft.NewRaftServer(raftConfig); err != nil {
		return
	}
	m.raftResolver = resolver
	m.raftServer = server
	return
}

// StopRaftServer stop raft server instance if possible.
func (m *MetaNode) stopRaftServer() {
	if m.raftServer != nil {
		m.raftServer.Stop()
	}
}