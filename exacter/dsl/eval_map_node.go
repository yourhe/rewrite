package dsl

import (
	"fmt"
	"regexp"
	"time"

	"gitlab.iyorhe.com/dr2am/dr2am-rewrite/exacter/dsl/ast"
)

type EvalMapNode struct {
	nodeEvaluator   NodeEvaluator
	constReturnType ast.ValueType
	state           ExecutionState
}

func NewEvalMapNode(mapExpr *ast.MapNode) (*EvalMapNode, error) {
	nodeEvaluator, err := createNodeEvaluator(mapExpr.Expression)
	if err != nil {
		return nil, fmt.Errorf("Failed to handle node: %v", err)
	}

	return &EvalMapNode{
		nodeEvaluator:   nodeEvaluator,
		constReturnType: getConstantNodeType(mapExpr.Expression),
		// Create an independent state for this expression
		state: CreateExecutionState(),
	}, nil
}

func (n *EvalMapNode) String() string {
	return fmt.Sprintf("%s", n.nodeEvaluator)
}

func (n *EvalMapNode) Type(scope ReadOnlyScope) (ast.ValueType, error) {
	if n.constReturnType == ast.InvalidType {
		// We are dynamic and we need to figure out our type
		// Do NOT cache this result in n.returnType since it can change.
		return n.nodeEvaluator.Type(scope)
	}
	return n.constReturnType, nil
}

func (n *EvalMapNode) IsDynamic() bool {
	return n.nodeEvaluator.IsDynamic()
}

func (n *EvalMapNode) EvalRegex(scope *Scope, _ ExecutionState) (*regexp.Regexp, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return nil, err
	}
	if typ == ast.TRegex {
		return n.nodeEvaluator.EvalRegex(scope, n.state)
	}

	return nil, ErrTypeGuardFailed{RequestedType: ast.TRegex, ActualType: typ}
}

func (n *EvalMapNode) EvalTime(scope *Scope, _ ExecutionState) (time.Time, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return time.Time{}, err
	}
	return time.Time{}, ErrTypeGuardFailed{RequestedType: ast.TTime, ActualType: typ}
}

func (n *EvalMapNode) EvalDuration(scope *Scope, _ ExecutionState) (time.Duration, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return 0, err
	}
	if typ == ast.TDuration {
		return n.nodeEvaluator.EvalDuration(scope, n.state)
	}

	return 0, ErrTypeGuardFailed{RequestedType: ast.TDuration, ActualType: typ}
}

func (n *EvalMapNode) EvalString(scope *Scope, _ ExecutionState) (string, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return "", err
	}
	if typ == ast.TString {
		return n.nodeEvaluator.EvalString(scope, n.state)
	}

	return "", ErrTypeGuardFailed{RequestedType: ast.TString, ActualType: typ}
}

func (n *EvalMapNode) EvalFloat(scope *Scope, _ ExecutionState) (float64, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return 0, err
	}
	if typ == ast.TFloat {
		return n.nodeEvaluator.EvalFloat(scope, n.state)
	}

	return 0, ErrTypeGuardFailed{RequestedType: ast.TFloat, ActualType: typ}
}

func (n *EvalMapNode) EvalInt(scope *Scope, _ ExecutionState) (int64, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return 0, err
	}
	if typ == ast.TInt {
		return n.nodeEvaluator.EvalInt(scope, n.state)
	}

	return 0, ErrTypeGuardFailed{RequestedType: ast.TInt, ActualType: typ}
}

func (n *EvalMapNode) EvalBool(scope *Scope, _ ExecutionState) (bool, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return false, err
	}
	if typ == ast.TBool {
		return n.nodeEvaluator.EvalBool(scope, n.state)
	}

	return false, ErrTypeGuardFailed{RequestedType: ast.TBool, ActualType: typ}
}

func (n *EvalMapNode) Eval(scope *Scope, _ ExecutionState) (interface{}, error) {
	// fmt.Println(n.nodeEvaluator)
	// fmt.Println(n.nodeEvaluator, typ)

	if v, ok := n.nodeEvaluator.(interface {
		Eval(scope *Scope, _ ExecutionState) (interface{}, error)
	}); ok {
		return v.Eval(scope, n.state)
	}
	typ, err := n.Type(scope)
	if err != nil {
		return false, err
	}
	return false, ErrTypeGuardFailed{RequestedType: ast.TBool, ActualType: typ}
}
