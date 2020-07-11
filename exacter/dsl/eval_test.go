package dsl

import (
	"fmt"
	"testing"

	"gitlab.iyorhe.com/dr2am/dr2am-rewrite/exacter/dsl/ast"
)

func Test_eval(t *testing.T) {
	script := `when request.url.hostname2:"wwwd" AND (request.url.hostname2=~/w/ OR request.url.hostname2=~/c.*/ )`
	node, err := ast.Parse(script)
	if err != nil {
		t.Fatal(err)
	}
	// s := &stack{}
	// eval(node, nil, s, nil, nil, false)
	// fmt.Println(node)
	expr, err := NewExpression(node.(*ast.ProgramNode).Nodes[0])
	if err != nil {
		t.Fatal(err)
	}
	scope := NewScope()
	scope.Set("request", map[string]interface{}{
		"url": map[string]interface{}{
			"hostname":  "wwwd",
			"hostname2": "wwwd",
		},
	})
	scope.Set("domain.d", "www")
	fmt.Println(expr.EvalBool(scope))
}

func Test_MapEval(t *testing.T) {
	script := `map request[a]`
	node, err := ast.Parse(script)
	if err != nil {
		t.Fatal(err)
	}
	// s := &stack{}
	// eval(node, nil, s, nil, nil, false)
	// fmt.Println(node)
	expr, err := NewExpression(node.(*ast.ProgramNode).Nodes[0])
	if err != nil {
		t.Fatal(err)
	}
	scope := NewScope()
	scope.Set("request", map[string]interface{}{
		"url": map[string]interface{}{
			"hostname":  "wwwd",
			"hostname2": []string{"wwwd)"},
		},
	})
	scope.Set("domain", []string{"wwwd)"})
	scope.Set("a", "url")
	fmt.Println(expr.Eval(scope))
}
