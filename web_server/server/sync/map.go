package sync

import "sync"

var syncMap = struct {
	sync.RWMutex
	m map[string]bool
}{m: make(map[string]bool)}

func SetKeyValue(key string) {
	syncMap.RLock()
	syncMap.m[key] = true
	syncMap.RUnlock()
}

func RemoveKey(key string) {
	syncMap.RLock()
	delete(syncMap.m, key)
	syncMap.RUnlock()
}

func Value(key string) bool {
	syncMap.RLock()
	n := syncMap.m[key]
	syncMap.RUnlock()
	return n
}
