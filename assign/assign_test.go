package assign

import "testing"

func TestAssign(t *testing.T) {
	var p T
	n := assign(42, &p)
	if p.(int) != 42 || n != 42 {
		t.Fail()
	}
}
