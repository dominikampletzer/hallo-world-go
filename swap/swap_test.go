package swap

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T) {
	var a, b = 1, 2
	swap0(&a, &b)
	if !(a == 2) || !(b == 1) {
		fmt.Println("swap0 klappt nicht", a, b)
	}

	a, b = 1, 2
	var pa, pb = &a, &b
	swap1(&pa, &pb)
	if !(a == 2) || !(b == 1) {
		fmt.Println("swap1 klappt nicht", a, b)
	}

	a, b = 1, 2
	a, b = swap2(a, b)
	if !(a == 2) || !(b == 1) {
		fmt.Println("swap2 klappt nicht", a, b)
	}

	a, b = 1, 2
	swap3(a, b)
	if !(a == 2) || !(b == 1) {
		fmt.Println("swap3 klappt nicht", a, b)
	}
}
