package dsl

import (
	"gitlab.iyorhe.com/dr2am/dr2am-rewrite/exacter/dsl/ast"
)

type unboundFunc func(obj interface{}) (interface{}, error)
type Var struct {
	Value       interface{}
	Type        ast.ValueType
	Description string
}

// Evaluate a node using a stack machine in a given scope
func eval(n ast.Node, scope *Scope, stck *stack, predefinedVars, defaultVars map[string]Var, ignoreMissingVars bool) (err error) {
	switch node := n.(type) {
	case *ast.ProgramNode:
		for _, n := range node.Nodes {
			err = eval(n, scope, stck, predefinedVars, defaultVars, ignoreMissingVars)
			if err != nil {
				return
			}
			// Pop unused result
			if stck.Len() > 0 {
				ret := stck.Pop()
				if f, ok := ret.(unboundFunc); ok {
					// Call global function
					_, err := f(nil)
					if err != nil {
						return err
					}
				}
			}
		}
	case *ast.StringNode:
		stck.Push(node.Literal)
	default:
		stck.Push(node)
	}
	return nil
}
