package hello

import (
	"testing"
)

func TestStack(t *testing.T) {
	var a = NewStack()
	a.Push(NewRational(2, 4))
	a.Push(NewRational(6, 4))
	c := NewRational(1, 1)
	for idx := 0; idx < a.Size(); idx++ {
		c = c.Add(a.GetAt(idx).(Rational))
	}
	if c != NewRational(3, 1) {
		t.Error("Summe passt nicht")
	}
}
