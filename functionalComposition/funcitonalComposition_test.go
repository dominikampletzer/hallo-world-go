package functionalComposition

import (
	"fmt"
	"testing"
)

func TestFunctionalComposition(t *testing.T) {
	var c = compose(square, square)(2)
	fmt.Println(c)

	a := intSequence()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := intSequence()
	fmt.Println(b())
	fmt.Println(a())
	fmt.Println(b())
}
