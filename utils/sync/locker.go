package sync

import (
	gosync "sync"
)

type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

var _ Locker = (*gosync.RWMutex)(nil)

type FakeLocker struct {
}

func (l FakeLocker) Lock() {

}

func (l FakeLocker) Unlock() {

}

func (l FakeLocker) RLock() {

}

func (l FakeLocker) RUnlock() {

}
