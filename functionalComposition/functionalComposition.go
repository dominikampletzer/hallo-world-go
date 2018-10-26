package functionalComposition

type Function func(Any) Any
type Any interface{}

var square = func(a Any) Any {
	return a.(int) * a.(int)
}

var compose = func(a, b Function) Function {
	return func(x Any) Any {
		return a(b(x))
	}
}

var intSequence = func() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
