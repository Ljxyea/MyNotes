package singleton

import "sync"

var lock = new(sync.Mutex)

type Single struct {
}

var instance *Single

func GetInstance() *Single {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Single{}
		}
	}
	return instance
}
