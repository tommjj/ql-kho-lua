package mapmutex

import (
	"sync"
)

// Mapmutex a simple mutex lock and unlock by key
type Mapmutex struct {
	m sync.Map
}

// Lock lock by key
func (mm *Mapmutex) Lock(key any) {
	l, _ := mm.m.LoadOrStore(key, &sync.Mutex{})
	l.(*sync.Mutex).Lock()
}

// UnLock unlock by key
func (mm *Mapmutex) UnLock(key any) {
	l, _ := mm.m.Load(key)
	l.(*sync.Mutex).Unlock()
}
