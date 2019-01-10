package sync

import "vue-admin/web_server/concurrent-map"

var (
	syncMap = cmap.New()
)

func SetKeyValue(key string) {
	syncMap.Set(key, true)
}

func RemoveKey(key string) {
	syncMap.Remove(key)
}

func Value(key string) bool {
	_, b := syncMap.Get(key)
	return b
}
