package sync

import (
	gosync "sync"
)

// Locker define an abstract locker interface
type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

var _ Locker = (*gosync.RWMutex)(nil)

// FakeLocker is a fack locker
type FakeLocker struct {
}

// Lock does nothing
func (l FakeLocker) Lock() {

}

// Unlock does nothing
func (l FakeLocker) Unlock() {

}

// RLock does nothing
func (l FakeLocker) RLock() {

}

// RUnlock does nothing
func (l FakeLocker) RUnlock() {

}
