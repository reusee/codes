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
