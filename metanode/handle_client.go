package metanode

import (
	"encoding/json"
	"net"

	"github.com/tiglabs/baudstorage/proto"
)

// Handle OpCreate Inode
func (m *MetaNode) opCreateInode(conn net.Conn, p *Packet) (err error) {
	req := &CreateInoReq{}
	if err = json.Unmarshal(p.Data, req); err != nil {
		return
	}
	mr, err := m.metaRangeManager.LoadMetaRange(req.Namespace)
	if err != nil {
		return
	}
	resp, err := mr.CreateInode(req)
	if err != nil {
		return
	}
	// Reply operation result to client though TCP connection.
	err = m.replyToClient(conn, p, resp)
	return
}

// Handle OpCreate
func (m *MetaNode) opCreateDentry(conn net.Conn, p *Packet) (err error) {
	req := &CreateDentryReq{}
	if err = json.Unmarshal(p.Data, req); err != nil {
		return
	}
	mr, err := m.metaRangeManager.LoadMetaRange(req.Namespace)
	if err != nil {
		return err
	}
	resp, err := mr.CreateDentry(req)
	if err != nil {
		return
	}
	// Reply operation result to client though TCP connection.
	err = m.replyToClient(conn, p, resp)
	return
}

// Handle OpDelete Dentry
func (m *MetaNode) opDeleteDentry(conn net.Conn, p *Packet) (err error) {
	req := &DeleteDentryReq{}
	if err = json.Unmarshal(p.Data, req); err != nil {
		return
	}
	mr, err := m.metaRangeManager.LoadMetaRange(req.Namespace)
	if err != nil {
		return
	}
	resp, err := mr.DeleteDentry(req)
	if err != nil {
		return
	}
	// Reply operation result to client though TCP connection.
	err = m.replyToClient(conn, p, resp)
	return
}

func (m *MetaNode) opDeleteInode(conn net.Conn, p *Packet) (err error) {
	req := &DeleteInoReq{}
	if err = json.Unmarshal(p.Data, req); err != nil {
		return
	}
	mr, err := m.metaRangeManager.LoadMetaRange(req.Namespace)
	if err != nil {
		return
	}
	resp, err := mr.DeleteInode(req)
	if err != nil {
		return
	}
	err = m.replyToClient(conn, p, resp)
	return
}

// Handle OpReadDir
func (m *MetaNode) opReadDir(conn net.Conn, p *Packet) (err error) {
	req := &proto.ReadDirRequest{}
	if err = json.Unmarshal(p.Data, req); err != nil {
		return
	}
	mr, err := m.metaRangeManager.LoadMetaRange(req.Namespace)
	if err != nil {
		return
	}
	resp, err := mr.ReadDir(req)
	if err != nil {
		return
	}
	// Reply operation result to client though TCP connection.
	err = m.replyToClient(conn, p, resp)
	return
}

// Handle OpOpen
func (m *MetaNode) opOpen(conn net.Conn, p *Packet) (err error) {
	req := &proto.OpenRequest{}
	if err = json.Unmarshal(p.Data, req); err != nil {
		return
	}
	mr, err := m.metaRangeManager.LoadMetaRange(req.Namespace)
	if err != nil {
		return
	}
	resp, err := mr.Open(req)
	if err != nil {
		return
	}
	// Reply operation result to client though TCP connection.
	err = m.replyToClient(conn, p, resp)
	return
}
