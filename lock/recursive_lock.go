package lock

import (
	"sync"
	"sync/atomic"
)

type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int32
}

func (m *RecursiveMutex) Lock() {
	gid := GetGoroutineId() // Function to get the current goroutine ID (pseudo-code)
	if m.owner == gid {
		atomic.AddInt32(&(m.recursion), 1)
	} else {
		m.Mutex.Lock()
		atomic.StoreInt64(&(m.owner), gid)
		atomic.StoreInt32(&(m.recursion), 1)
	}
}

func (m *RecursiveMutex) Unlock() {
	gid := GetGoroutineId() // Function to get the current goroutine ID (pseudo-code)
	if m.owner == gid {
		atomic.AddInt32(&(m.recursion), -1)
		if m.recursion == 0 {
			atomic.StoreInt64(&(m.owner), 0)
			m.Mutex.Unlock()
		}
	} else {
		panic("Attempting to unlock a mutex not owned by the calling goroutine")
	}
}
