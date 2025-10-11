package main

import "sync"

type RwMap struct {
	rw sync.RWMutex
	m  map[string]int
}

func NewRwMap() *RwMap {
	return &RwMap{
		m: make(map[string]int),
	}
}

func (m *RwMap) Set(key string, v int) {
	m.rw.Lock()
	defer m.rw.Unlock()

	m.m[key] = v
}

func (m *RwMap) Get(key string) (int, bool) {
	m.rw.Lock()
	defer m.rw.Unlock()

	v, ok := m.m[key]
	return v, ok
}
