package metanode

import (
	"encoding/json"
	"io"

	"github.com/google/btree"
)

type MetaRangeSnapshot struct {
	Op uint32 `json:"op"`
	K  []byte `json:"k"`
	V  []byte `json:"v"`
}

func (s *MetaRangeSnapshot) Encode() ([]byte, error) {
	return json.Marshal(s)
}

func (s *MetaRangeSnapshot) Decode(data []byte) error {
	return json.Unmarshal(data, s)
}

func NewMetaRangeSnapshot(op uint32, key, value []byte) *MetaRangeSnapshot {
	return &MetaRangeSnapshot{
		Op: op,
		K:  key,
		V:  value,
	}
}

type SnapshotIterator struct {
	applyID    uint64
	cur        int
	curItem    btree.Item
	inoLen     int
	inodeTree  *btree.BTree
	dentryLen  int
	dentryTree *btree.BTree
	total      int
}

func NewSnapshotIterator(applyID uint64, ino, den *btree.BTree) *SnapshotIterator {
	si := new(SnapshotIterator)
	si.applyID = applyID
	si.inodeTree = ino
	si.dentryTree = den
	si.cur = 1
	si.inoLen = ino.Len()
	si.dentryLen = den.Len()
	si.total = si.inoLen + si.dentryLen
	return si
}

func (si *SnapshotIterator) ApplyIndex() uint64 {
	return si.applyID
}

func (si *SnapshotIterator) Close() {
	return
}

func (si *SnapshotIterator) Next() (data []byte, err error) {
	if si.cur > si.total {
		err = io.EOF
		return
	}
	// ascend inode tree
	if si.cur <= si.inoLen {
		si.inodeTree.AscendGreaterOrEqual(si.curItem, func(i btree.Item) bool {
			ino := i.(*Inode)
			if si.curItem == ino {
				return true
			}
			si.curItem = ino
			snap := NewMetaRangeSnapshot(opCreateInode, ino.GetKeyBytes(),
				ino.GetValueBytes())
			data, err = snap.Encode()
			si.cur++
			return false
		})
		return
	}

	//ascend dentry tree
	if si.cur == (si.inoLen + 1) {
		si.curItem = nil
	}
	si.dentryTree.AscendGreaterOrEqual(si.curItem, func(i btree.Item) bool {
		dentry := i.(*Dentry)
		if si.curItem == dentry {
			return true
		}
		si.curItem = dentry
		snap := NewMetaRangeSnapshot(opCreateDentry, dentry.GetKeyBytes(),
			dentry.GetValueBytes())
		data, err = snap.Encode()
		si.cur++
		return false
	})
	return
}
