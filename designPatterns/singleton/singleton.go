package singleton

import (
	"sync"
)

// MySingleton is the type you want as a singleton.
type MySingleton struct {
	Data string
}

var (
	instance *MySingleton
	once     sync.Once
)

// Instance gives you the singleton instance.
func Instance() *MySingleton {
	once.Do(func() {
		instance = &MySingleton{Data: "Hello, Singleton!"}
	})
	return instance
}
