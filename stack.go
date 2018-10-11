package hello

type Stack struct {
	data interface{}
}

func (stack *Stack) Push(el interface{}) {
	stack = append(stack, el)
}
func (stack *Stack) Pop() interface{} {

}
func (stack *Stack) Size() int {

}
func (stack *Stack) GetAt(idx int) interface{} {

}
