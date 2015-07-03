package err

import "fmt"

var Pkg = "packageName"

type Err struct {
	Pkg  string
	Info string
	Err  error
}

func (e *Err) Error() string {
	return fmt.Sprintf("%s: %s\n%v", e.Pkg, e.Info, e.Err)
}

func makeErr(err error, info string) *Err {
	return &Err{
		Pkg:  Pkg,
		Info: info,
		Err:  err,
	}
}

func ce(err error, info string) (ret bool) {
	ret = true
	if err != nil {
		panic(makeErr(err, info))
	}
	ret = false
	return
}

func ct(err *error) {
	if p := recover(); p != nil {
		if e, ok := p.(error); ok {
			*err = e
		} else {
			panic(p)
		}
	}
}
