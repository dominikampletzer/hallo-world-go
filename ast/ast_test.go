package ast

import (
	"testing"
)

func TestAst(t *testing.T) {

	ma := make(map[string]bool)

	ma = append(ma, Var{true})
	ma["b"] = false
	ma["c"] = true

}
