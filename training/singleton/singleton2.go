package singleton

import "sync"

var once sync.Once

var instance2 *Single

func GetInstance2() *Single {
	if instance2 == nil {
		once.Do(func() {
			instance2 = &Single{}
		})
	}

	return instance2
}
