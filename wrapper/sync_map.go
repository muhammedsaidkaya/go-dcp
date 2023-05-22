package wrapper

import (
	"sync"

	jsoniter "github.com/json-iterator/go"
)

type SyncMap[K comparable, V any] struct {
	m sync.Map
}

func (m *SyncMap[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if !ok {
		return value, ok
	}

	return v.(V), ok
}

func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		return value, loaded
	}

	return v.(V), loaded
}

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	a, loaded := m.m.LoadOrStore(key, value)
	return a.(V), loaded
}

func (m *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *SyncMap[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

func (m *SyncMap[K, V]) ToMap() map[K]V {
	result := make(map[K]V)
	m.Range(func(key K, value V) bool {
		result[key] = value
		return true
	})
	return result
}

func (m *SyncMap[K, V]) MarshalJSON() ([]byte, error) {
	return jsoniter.Marshal(m.ToMap())
}

func (m *SyncMap[K, V]) UnmarshalJSON(data []byte) error {
	var result map[K]V
	if err := jsoniter.Unmarshal(data, &result); err != nil {
		return err
	}

	for k, v := range result {
		m.Store(k, v)
	}

	return nil
}