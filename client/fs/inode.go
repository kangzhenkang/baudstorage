package fs

import (
	"fmt"
	"os"
	"sync"
	"time"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"

	"github.com/tiglabs/baudstorage/proto"
)

type InodeCommon struct {
	sync.RWMutex
	super   *Super
	parent  *Dir
	blksize uint32
	nlink   uint32
}

type Inode struct {
	ino     uint64
	size    uint64
	mode    uint32 //Inode Type
	extents []string
	ctime   time.Time
	mtime   time.Time
	atime   time.Time
}

func (s *Super) InodeGet(ino uint64, inode *Inode) error {
	status, info, err := s.meta.InodeGet_ll(ino)
	err = ParseResult(status, err)
	if err != nil {
		return err
	}
	fillInode(inode, info)
	// sanity check
	if ino != inode.ino {
		panic(fmt.Sprintf("Request ino: %v , Response ino: %v ", ino, inode.ino))
	}
	return nil
}

func fillInode(inode *Inode, info *proto.InodeInfo) {
	inode.ino = info.Inode
	inode.mode = info.Type
	inode.size = info.Size
	inode.extents = info.Extents
	inode.ctime = info.CreateTime
	inode.atime = info.AccessTime
	inode.mtime = info.ModifyTime
	//TODO: fill more fields
}

func fillAttr(attr *fuse.Attr, n fs.Node) {
	var inode *Inode
	switch v := n.(type) {
	case *Dir:
		inode = &v.inode
		attr.Nlink = v.nlink
		attr.BlockSize = v.blksize
		attr.Mode = os.ModeDir | os.ModePerm
	case *File:
		inode = &v.inode
		attr.Nlink = v.nlink
		attr.BlockSize = v.blksize
		attr.Mode = os.ModePerm
	default:
	}

	if inode == nil {
		return
	}

	attr.Inode = inode.ino
	attr.Size = inode.size
	attr.Blocks = attr.Size >> 9 // In 512 bytes
	attr.Atime = inode.atime
	attr.Ctime = inode.ctime
	attr.Mtime = inode.mtime
}
