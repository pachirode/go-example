package main

import (
	"hash/maphash"
	"sync"
)

// 读写锁会使得程序的性能急速下降，可以采取分片的思想，减小锁的颗粒度

var seed = maphash.MakeSeed()

func hashKey(key string) uint64 {
	return maphash.String(seed, key)
}

type ShardingMap struct {
	locks  []sync.RWMutex
	shards []map[string]int
}

func NewShardingMap(size int) *ShardingMap {
	sm := &ShardingMap{
		locks:  make([]sync.RWMutex, size),
		shards: make([]map[string]int, size),
	}

	for i := 0; i < size; i++ {
		sm.shards[i] = make(map[string]int)
	}

	return sm
}

func (m *ShardingMap) getSharedIdx(key string) uint64 {
	hash := hashKey(key)
	return hash % uint64(len(m.shards))
}

func (m *ShardingMap) Set(key string, value int) {
	idx := m.getSharedIdx(key)
	m.locks[idx].Lock()
	defer m.locks[idx].Unlock()
	m.shards[idx][key] = value
}

func (m *ShardingMap) Get(key string) (int, bool) {
	idx := m.getSharedIdx(key)
	m.locks[idx].RLock()
	defer m.locks[idx].RUnlock()
	value, ok := m.shards[idx][key]
	return value, ok
}
