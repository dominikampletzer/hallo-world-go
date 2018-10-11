package hello

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	return new(Stack)
}

func (stack *Stack) Push(el interface{}) {
	stack.data = append(stack.data, el)
}
func (stack *Stack) Pop() interface{} {
	var b = stack.data[len(stack.data)-1]
	stack.data = stack.data[0 : len(stack.data)-1]
	return b
}
func (stack *Stack) Size() int {
	return len(stack.data)
}
func (stack *Stack) GetAt(idx int) interface{} {
	return stack.data[idx]
}
