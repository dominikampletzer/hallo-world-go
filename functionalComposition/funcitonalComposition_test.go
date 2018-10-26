package functionalComposition

import (
	"fmt"
	"testing"
)

func TestFunctionalComposition(t *testing.T) {
	var c = compose(square, square)(2)
	fmt.Println(c)
}
