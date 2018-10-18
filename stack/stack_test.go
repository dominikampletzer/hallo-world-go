package stack

import (
	"github.com/dominikampletzer/hello/rational"
	"testing"
)

func TestStack(t *testing.T) {
	var a = NewStack()
	a.Push(rational.NewRational(2, 4))
	a.Push(rational.NewRational(6, 4))
	c := rational.NewRational(1, 1)
	for idx := 0; idx < a.Size(); idx++ {
		c = c.Add(a.GetAt(idx).(rational.Rational))
	}
	if c != rational.NewRational(3, 1) {
		t.Error("Summe passt nicht")
	}
}
