package ast

type Node interface {
	Eval(vars map[string]bool) bool
}

type Var struct {
	name string
	//bool
}

func (val Var) Eval(vars map[string]bool) bool {
	return vars[val.name]
	//return val.bool
}

type And struct {
	left, right Node
}
type Or struct {
	left, right Node
}

func (or Or) Eval(p map[string]bool) bool {
	return or.left.Eval(p) || or.right.Eval(p)
}

func (and And) Eval(p map[string]bool) bool {
	return and.left.Eval(p) || and.right.Eval(p)
}

type Not struct {
	not Node
}

func (n Not) Eval(p map[string]bool) bool {
	return !n.not.Eval(p)
}
