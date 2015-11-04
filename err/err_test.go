package err

import (
	"errors"
	"testing"
)

func TestCatch(t *testing.T) {
	var err error
	func() {
		defer ct(&err)
		e := errors.New("foobar")
		ce(e, "FOO")
	}()
	if err == nil || err.Error() != "packageName: FOO\nfoobar" {
		t.Fail()
	}
}

func TestCatch2(t *testing.T) {
	err := func() (err error) {
		defer ct(&err)
		return
	}()
	if err != nil {
		t.Fail()
	}
}

func TestCatch3(t *testing.T) {
	err := func() (err error) {
		defer ct(&err)
		return errors.New("Err")
	}()
	if err == nil || err.Error() != "Err" {
		t.Fail()
	}
}

func TestNoError(t *testing.T) {
	var err error
	func() {
		defer ct(&err)
		ce(nil, "test")
	}()
	if err != nil {
		t.Fail()
	}
}

func TestFormat(t *testing.T) {
	err := me(nil, "foo")
	if err.Error() != "packageName: foo" {
		t.Fatal("format error")
	}
	err = me(nil, "foo%s", "bar")
	if err.Error() != "packageName: foobar" {
		t.Fatal("format error")
	}
}

func TestPanic(t *testing.T) {
	var err error
	func() {
		defer func() {
			p := recover()
			if ps, ok := p.(string); !ok || ps != "foo" {
				t.Fail()
			}
			if err != nil {
				t.Fail()
			}
		}()
		func() {
			defer ct(&err)
			panic("foo")
		}()
	}()
}

func BenchmarkCatchError(b *testing.B) {
	var err error
	e := errors.New("foo")
	for i := 0; i < b.N; i++ {
		func() {
			defer ct(&err)
			ce(e, "bench")
		}()
	}
}

func BenchmarkNoError(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		func() {
			defer ct(&err)
			ce(nil, "bench")
		}()
	}
}
