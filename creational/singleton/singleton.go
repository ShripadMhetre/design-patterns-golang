package singleton

import "sync"

type Singleton struct {
	name string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	//if instance == nil {
	//	mutex.Lock()
	//	defer mutex.Unlock()
	//
	//	if instance == nil { // 👈 additional check
	//		instance = &Singleton{name: "Safe Golang Singleton"}
	//	}
	//}

	// shorter version of above code using "once.Do()"
	once.Do(func() {
		instance = &Singleton{name: "Safe Golang Singleton"}
	})

	return instance
}

func (s *Singleton) GetName() string {
	return s.name
}
