package singleton

import "sync"

// LazySingleton 懒汉式单例
type LazySingleton struct{}

var (
	lazySingleton *LazySingleton
	once          sync.Once
)

// GetLazyInstance 获取实例
func GetLazyInstance() *LazySingleton {
	if lazySingleton == nil {
		once.Do(func() {
			lazySingleton = &LazySingleton{}
		})
	}
	return lazySingleton
}
