package redisclone

import "sync"

type RedisCloneStore struct {
	mu   sync.RWMutex
	dict map[string]string
}

func (store *RedisCloneStore) StoreSet(key string, value string) bool {
	store.mu.Lock()

	store.dict[key] = value

	store.mu.Unlock()

	return true
}

func (store *RedisCloneStore) StoreGet(key string) string {
	store.mu.RLock()

	value, key_exists := store.dict[key]

	store.mu.RUnlock()

	if !key_exists {
		return "nil"
	}

	return value
}
