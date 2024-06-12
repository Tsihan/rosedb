package index

import (
	art "github.com/plar/go-adaptive-radix-tree"
	"github.com/rosedblabs/wal"
	"sync"
)

type MemoryARTTree struct {
	tree art.Tree
	lock *sync.RWMutex
}

func newARTTree() *MemoryARTTree {
	return &MemoryARTTree{
		tree: art.New(),
		lock: new(sync.RWMutex),
	}
}

func (mt *MemoryARTTree) Put(key []byte, position *wal.ChunkPosition) *wal.ChunkPosition {
	mt.lock.Lock()
	defer mt.lock.Unlock()

	oldValue, _ := mt.tree.Insert(key, position)
	if oldValue != nil {
		return oldValue.(*wal.ChunkPosition)
	}
	return nil
}

func (mt *MemoryARTTree) Get(key []byte) *wal.ChunkPosition {
	mt.lock.RLock()
	defer mt.lock.RUnlock()
	value, found := mt.tree.Search(key)
	if !found {
		return nil
	}
	if value != nil {
		return value.(*wal.ChunkPosition)
	}
	return nil
}

func (mt *MemoryARTTree) Delete(key []byte) (*wal.ChunkPosition, bool) {
	mt.lock.Lock()
	defer mt.lock.Unlock()

	value, _ := mt.tree.Delete(key)
	if value != nil {
		return value.(*wal.ChunkPosition), true
	}
	return nil, false
}

func (mt *MemoryARTTree) Size() int {
	println("MemoryARTTree.Size()")
	return mt.tree.Size()
}

func (mt *MemoryARTTree) Ascend(handleFn func(key []byte, position *wal.ChunkPosition) (bool, error)) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	//mt.tree.Ascend(func(i btree.Item) bool {
	//	cont, err := handleFn(i.(*item).key, i.(*item).pos)
	//	if err != nil {
	//		return false
	//	}
	//	return cont
	//})
}

func (mt *MemoryARTTree) Descend(handleFn func(key []byte, position *wal.ChunkPosition) (bool, error)) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	//mt.tree.Descend(func(i btree.Item) bool {
	//	cont, err := handleFn(i.(*item).key, i.(*item).pos)
	//	if err != nil {
	//		return false
	//	}
	//	return cont
	//})
}

func (mt *MemoryARTTree) AscendRange(startKey, endKey []byte, handleFn func(key []byte, position *wal.ChunkPosition) (bool, error)) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	//mt.tree.AscendRange(&item{key: startKey}, &item{key: endKey}, func(i btree.Item) bool {
	//	cont, err := handleFn(i.(*item).key, i.(*item).pos)
	//	if err != nil {
	//		return false
	//	}
	//	return cont
	//})
}

func (mt *MemoryARTTree) DescendRange(startKey, endKey []byte, handleFn func(key []byte, position *wal.ChunkPosition) (bool, error)) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	//mt.tree.DescendRange(&item{key: startKey}, &item{key: endKey}, func(i btree.Item) bool {
	//	cont, err := handleFn(i.(*item).key, i.(*item).pos)
	//	if err != nil {
	//		return false
	//	}
	//	return cont
	//})
}

func (mt *MemoryARTTree) AscendGreaterOrEqual(key []byte, handleFn func(key []byte, position *wal.ChunkPosition) (bool, error)) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	//mt.tree.AscendGreaterOrEqual(&item{key: key}, func(i btree.Item) bool {
	//	cont, err := handleFn(i.(*item).key, i.(*item).pos)
	//	if err != nil {
	//		return false
	//	}
	//	return cont
	//})
}

func (mt *MemoryARTTree) DescendLessOrEqual(key []byte, handleFn func(key []byte, position *wal.ChunkPosition) (bool, error)) {
	mt.lock.RLock()
	defer mt.lock.RUnlock()

	//mt.tree.DescendLessOrEqual(&item{key: key}, func(i btree.Item) bool {
	//	cont, err := handleFn(i.(*item).key, i.(*item).pos)
	//	if err != nil {
	//		return false
	//	}
	//	return cont
	//})
}
