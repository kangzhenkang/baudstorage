package master

import (
	"github.com/tiglabs/baudstorage/proto"
	"github.com/tiglabs/baudstorage/raftstore"
)

/*
  this struct is used for master need to persist to raft store
*/

type MetadataFsm struct {
	raftstore.RaftStoreFsm
}

func (mf *MetadataFsm) CreateNameSpace(request proto.CreateNameSpaceRequest) (response proto.CreateNameSpaceResponse) {
	//	var (
	//		data     []byte
	//		err      error
	//		raftResp *raft.Future
	//	)
	//	//key:ns#name,value:nil
	//	kv := &kvp.Kv{Opt: OptSetNamespace}
	//	kv.K = encodeNameSpaceKey(request.Name)
	//	if data, err = pbproto.Marshal(kv); err != nil {
	//		err = fmt.Errorf("action[CreateNameSpace],marshal kv:%v,err:%v", kv, err.Error())
	//		goto errDeal
	//	}
	//	raftResp = mf.RaftStoreFsm.GetRaftServer().Submit(ClusterGroupID, data)
	//	if _, err = raftResp.Response(); err != nil {
	//		err = fmt.Errorf("action[CreateNameSpace],raft submit err:%v", err.Error())
	//		goto errDeal
	//	}
	//	response.Status = OK
	//	return
	//errDeal:
	//	response.Status = Failed
	//	response.Result = err.Error()
	return

}

func encodeNameSpaceKey(name string) string {
	return PrefixNameSpace + KeySeparator + name
}

func (mf *MetadataFsm) CreateMetaRange(request proto.CreateMetaRangeRequest) (response proto.CreateMetaRangeResponse) {
	return
}
