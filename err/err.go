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
		panic(Err{
			Pkg:  Pkg,
			Info: info,
			Err:  err,
		})
	}
	ret = false
	return
}

func ct(err *error) {
	if p := recover(); p != nil {
		*err = p.(error)
	}
}
