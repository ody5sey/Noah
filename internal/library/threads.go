package library

import (
	"sync"
)

/*
实现一个线程安全的slice
*/

type ResultPath struct {
	Code    int
	Address string
	Title   string
	Header  string
}

type ThreadsSlice struct {
	locker *sync.RWMutex
	slices []ResultPath
}

func (ts *ThreadsSlice) Add(element ResultPath) {
	ts.locker.Lock()
	ts.slices = append(ts.slices, element)
	ts.locker.Unlock()
}

func (ts *ThreadsSlice) Get() []ResultPath {
	return ts.slices
}

func NewSlice() *ThreadsSlice {
	return &ThreadsSlice{
		locker: &sync.RWMutex{},
		slices: []ResultPath{},
	}

}
