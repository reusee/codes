package assign

type T interface{}

func assign(v T, p *T) T {
	*p = v
	return v
}
