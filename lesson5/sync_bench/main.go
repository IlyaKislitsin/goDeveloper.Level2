// здесь конечно ещё бы всё по разным файлам ранести...
package main

import (
	"sync"
)

// SetInterface
type SetInterface interface {
	Add(i int)
	Has(i int) bool
}

// Set struct implements SetInterface
type Set struct {
	sync.Mutex
	testMap map[int]struct{}
}

// NewSet create instances new Set struct
func NewSet() *Set {
	return &Set{
		testMap: map[int]struct{}{},
	}
}

// Add item to Set.testMap using Mutex.Lock
func (s *Set) Add(i int) {
	s.Lock()
	s.testMap[i] = struct{}{}
	s.Unlock()
}

// Has check item into Set.testMap using Mutex.Lock
func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.testMap[i]
	return ok
}

// RWSet struct implements SetInterface
type RWSet struct {
	sync.RWMutex
	testMap map[int]struct{}
}

// NewRWSet create instances new RWSet struct
func NewRWSet() *RWSet {
	return &RWSet{
		testMap: map[int]struct{}{},
	}
}

// Add item to RWSet.testMap using RWMutex.Lock
func (s *RWSet) Add(i int) {
	s.Lock()
	s.testMap[i] = struct{}{}
	s.Unlock()
}

// Has check item into RWSet.testMap using RWMutex.RLock
func (s *RWSet) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.testMap[i]
	return ok
}

// MapSet struct implements SetInterface
type MapSet struct {
	sync.Map
}

// NewMapSet create instances new MapSet struct
func NewMapSet() *MapSet {
	return &MapSet{
		Map: sync.Map{},
	}
}

// Add item to MapSet.Map
func (s *MapSet) Add(i int) {
	s.Map.Store(i, struct{}{})
}

// Has check item into MapSet.Map
func (s *MapSet) Has(i int) bool {
	_, ok := s.Map.Load(i)
	return ok
}
