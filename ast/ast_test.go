package ast

import (
	"fmt"
	"testing"
)

func TestAst(t *testing.T) {
	ast := Or{And{Var{"A"}, Var{"B"}}, Var{"C"}}

	vars := map[string]bool{"A": true, "B": false, "C": true}
	result := ast.Eval(vars)
	//fmt.Printf("expression := %v\n", ast)
	//fmt.Printf("vars := %v\n", vars)
	//fmt.Printf("result := %v\n", result)
	fmt.Println(result)

	vars = map[string]bool{"A": true, "B": true, "C": false}
	result = ast.Eval(vars)
	//fmt.Printf("expression := %v\n", ast)
	//fmt.Printf("vars := %v\n", vars)
	//fmt.Printf("result := %v\n", result)
	//ma := make(map[string]bool)
	////TODO
	//ma = append(ma, Var{true})
	//ma["b"] = false
	//ma["c"] = true

}
