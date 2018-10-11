package hello

import (
	"testing"
)

func TestRational(t *testing.T) {
	a := NewRational(1, 2)
	b := NewRational(2, 4)

	if !(a == b) {
		t.Error("Equals geht nicht. Var sind gleich", a, b)
	}

	c := a.Add(b)
	if !(c.nummerator == 1 && c.deminator == 1) {
		t.Error("Add hat probleme")
	}

	d := NewRational(1, 4)
	if a.Multilicate(b) != d {
		t.Error("Multiple hat probleme")
	}
}
