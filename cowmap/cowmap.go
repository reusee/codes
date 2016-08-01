package cowmap

import (
	"sync"
	"sync/atomic"
)

type Key interface{}

type Value interface{}

type CowMap struct {
	s atomic.Value
	l sync.Mutex
}

func (m *CowMap) Get(key Key) (value Value, ok bool) {
	v := m.s.Load()
	if v != nil {
		value, ok = v.(map[Key]Value)[key]
	}
	return
}

func (m *CowMap) Set(key Key, value Value) {
	m.l.Lock()
	defer m.l.Unlock()
	var fresh map[Key]Value
	old, ok := m.s.Load().(map[Key]Value)
	if ok {
		fresh = make(map[Key]Value, len(old)+1)
		for k, v := range old {
			fresh[k] = v
		}
	} else {
		fresh = make(map[Key]Value)
	}
	fresh[key] = value
	m.s.Store(fresh)
}

func (m *CowMap) Delete(keys ...Key) {
	toDelete := make(map[Key]struct{})
	for _, key := range keys {
		toDelete[key] = struct{}{}
	}
	m.l.Lock()
	defer m.l.Unlock()
	var fresh map[Key]Value
	old, ok := m.s.Load().(map[Key]Value)
	if ok {
		fresh = make(map[Key]Value, len(old)+1)
		for k, v := range old {
			if _, del := toDelete[k]; !del {
				fresh[k] = v
			}
		}
	} else {
		fresh = make(map[Key]Value)
	}
	m.s.Store(fresh)
}

func (m *CowMap) Clear() {
	m.l.Lock()
	defer m.l.Unlock()
	fresh := make(map[Key]Value)
	m.s.Store(fresh)
}
