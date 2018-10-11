package hello

import "fmt"

func swap0(x *int, y *int) { // ändert pointer
	*x, *y = *y, *x
}
func swap1(x **int, y **int) {
	*x, *y = *y, *x
}

func swap2(x int, y int) (int, int) {
	return y, x
}
func swap3(x int, y int) { // darf nicht gehen!!
	x, y = y, x
}

var a = 1

var b = 2

func nix() {

	fmt.Println("start", a, b)
	swap0(&a, &b)
	fmt.Println(a, b)
	a, b = 1, 2
	var pa, pb = &a, &b
	swap1(&pa, &pb)
	fmt.Println("dreht nur die pointer", pa, pb, a, b)

	a, b = 1, 2
	a, b = swap2(a, b)

	fmt.Println(a, b)
	a, b = 1, 2
	swap3(a, b)
	fmt.Println("geht nicht", a, b)
}
