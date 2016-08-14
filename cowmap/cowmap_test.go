package cowmap

import (
	"testing"
)

func TestAll(t *testing.T) {
	m := new(CowMap)
	_, ok := m.Get("foo")
	if ok {
		t.Fail()
	}
	m.Set("foo", "bar")
	m.Set("foo", "bar")
	if len(m.s.Load().(map[Key]Value)) != 1 {
		t.Fail()
	}
	v, ok := m.Get("foo")
	if !ok {
		t.Fail()
	}
	if v != "bar" {
		t.Fail()
	}
	_, ok = m.Get("bar")
	if ok {
		t.Fail()
	}
	m.Delete("foo")
	if len(m.s.Load().(map[Key]Value)) != 0 {
		t.Fail()
	}

	m.Set("FOO", "FOO")
	m.Set("BAR", "BAR")
	m.Clear()
	if len(m.s.Load().(map[Key]Value)) != 0 {
		t.Fail()
	}

	m.Set("foo", "foo")
	m.Set("bar", "bar")
	n := 0
	m.IterKeys(func(key Key) {
		n++
	})
	if n != 2 {
		t.Fail()
	}
}
