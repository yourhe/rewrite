package dsl

import (
	"fmt"
	"regexp"
	"time"

	"gitlab.iyorhe.com/dr2am/dr2am-rewrite/exacter/dsl/ast"
)

// ErrTypeGuardFailed is returned when a speicifc value type is requested thorugh NodeEvaluator (for example: "Float64Value")
// when the node doesn't support the given type, for example "Float64Value" is called on BoolNode
type ErrTypeGuardFailed struct {
	RequestedType ast.ValueType
	ActualType    ast.ValueType
}

func (e ErrTypeGuardFailed) Error() string {
	return fmt.Sprintf("TypeGuard: expression returned unexpected type %s, expected %s", e.ActualType, e.RequestedType)
}

// Expression is interface that describe expression with state and
// it's evaluation.
type Expression interface {
	Reset()

	Type(scope ReadOnlyScope) (ast.ValueType, error)

	EvalFloat(scope *Scope) (float64, error)
	EvalInt(scope *Scope) (int64, error)
	EvalString(scope *Scope) (string, error)
	EvalBool(scope *Scope) (bool, error)
	EvalDuration(scope *Scope) (time.Duration, error)

	Eval(scope *Scope) (interface{}, error)

	// Return a copy of the expression but with a Reset state.
	CopyReset() Expression
}

// A callable function from within the expression
type Func interface {
	Reset()
	Call(...interface{}) (interface{}, error)
	Signature() map[Domain]ast.ValueType
}

// Lookup for functions
type Funcs map[string]Func

// NodeEvaluator provides a generic way for trying to fetch
// node value, if a speicifc type is requested (so Value isn't called, the *Value is called) ErrTypeGuardFailed must be returned
type NodeEvaluator interface {
	EvalFloat(scope *Scope, executionState ExecutionState) (float64, error)
	EvalInt(scope *Scope, executionState ExecutionState) (int64, error)
	EvalString(scope *Scope, executionState ExecutionState) (string, error)
	EvalBool(scope *Scope, executionState ExecutionState) (bool, error)
	EvalRegex(scope *Scope, executionState ExecutionState) (*regexp.Regexp, error)
	EvalTime(scope *Scope, executionState ExecutionState) (time.Time, error)
	EvalDuration(scope *Scope, executionState ExecutionState) (time.Duration, error)
	// EvalMissing(scope *Scope, executionState ExecutionState) (*ast.Missing, error)

	// Type returns the type of ast.ValueType
	Type(scope ReadOnlyScope) (ast.ValueType, error)
	// Whether the type returned by the node can change.
	IsDynamic() bool
}

type expression struct {
	nodeEvaluator  NodeEvaluator
	executionState ExecutionState
}

// se *expression Expression

func NewExpression(node ast.Node) (Expression, error) {
	nodeEvaluator, err := createNodeEvaluator(node)
	if err != nil {
		return nil, err
	}

	return &expression{
		nodeEvaluator:  nodeEvaluator,
		executionState: CreateExecutionState(),
		// executionState: CreateExecutionState(),
	}, nil
}
func (se *expression) Reset() {
	se.executionState.ResetAll()
	// panic("not impleated")
}

func (se *expression) Type(scope ReadOnlyScope) (ast.ValueType, error) {
	return se.nodeEvaluator.Type(scope)
}

func (se *expression) EvalFloat(scope *Scope) (float64, error) {
	return se.nodeEvaluator.EvalFloat(scope, se.executionState)
}
func (se *expression) EvalInt(scope *Scope) (int64, error) {
	return se.nodeEvaluator.EvalInt(scope, se.executionState)
}
func (se *expression) EvalString(scope *Scope) (string, error) { panic("not impleated") }
func (se *expression) EvalBool(scope *Scope) (bool, error) {
	return se.nodeEvaluator.EvalBool(scope, se.executionState)
}
func (se *expression) EvalDuration(scope *Scope) (time.Duration, error) {
	return se.nodeEvaluator.EvalDuration(scope, se.executionState)
}

func (se *expression) Eval(scope *Scope) (interface{}, error) {
	if v, ok := se.nodeEvaluator.(interface {
		Eval(*Scope, ExecutionState) (interface{}, error)
	}); ok {
		return v.Eval(scope, se.executionState)
	}
	panic("not impleated")

	return nil, nil
}

// Return a copy of the expression but with a Reset state.
func (se *expression) CopyReset() Expression {
	return &expression{
		nodeEvaluator:  se.nodeEvaluator,
		executionState: CreateExecutionState(),
	}

}
func createNodeEvaluator(n ast.Node) (NodeEvaluator, error) {
	switch node := n.(type) {

	// case *ast.BoolNode:
	// 	return &EvalBoolNode{Node: node}, nil

	// case *ast.NumberNode:
	// 	switch {
	// 	case node.IsFloat:
	// 		return &EvalFloatNode{Float64: node.Float64}, nil

	// 	case node.IsInt:
	// 		return &EvalIntNode{Int64: node.Int64}, nil

	// 	default:
	// 		// We wouldn't reach ever, unless there is bug in tick parsing ;)
	// 		return nil, errors.New("Invalid NumberNode: Not float or int")
	// 	}

	// case *ast.DurationNode:
	// 	return &EvalDurationNode{Duration: node.Dur}, nil

	case *ast.StringNode:
		return &EvalStringNode{Node: node}, nil

	case *ast.RegexNode:
		return &EvalRegexNode{Node: node}, nil

	case *ast.BinaryNode:
		bn, err := NewEvalBinaryNode(node)
		if err != nil {
			return nil, err
		}
		return bn, nil

	case *ast.ReferenceNode:
		return &EvalReferenceNode{Node: node}, nil
	case *ast.IdentifierNode:
		ern := &EvalReferenceNode{
			Node: ast.NewReference(node.Ident),
		}
		return ern, nil

	// case *ast.FunctionNode:
	// 	return NewEvalFunctionNode(node)

	// case *ast.UnaryNode:
	// 	return NewEvalUnaryNode(node)

	case *ast.ConditionNode:
		return NewEvalConditionNode(node)

	case *ast.MapNode:
		return NewEvalMapNode(node)
	}

	return nil, fmt.Errorf("Given node type is not valid evaluation node: %T", n)
}

type EvalConditionNode struct {
	nodeEvaluator   NodeEvaluator
	constReturnType ast.ValueType
	state           ExecutionState
}

func NewEvalConditionNode(condition *ast.ConditionNode) (*EvalConditionNode, error) {
	nodeEvaluator, err := createNodeEvaluator(condition.Expression)
	if err != nil {
		return nil, fmt.Errorf("Failed to handle node: %v", err)
	}

	return &EvalConditionNode{
		nodeEvaluator:   nodeEvaluator,
		constReturnType: getConstantNodeType(condition.Expression),
		// Create an independent state for this expression
		state: CreateExecutionState(),
	}, nil
}

func (n *EvalConditionNode) String() string {
	return fmt.Sprintf("%s", n.nodeEvaluator)
}

func (n *EvalConditionNode) Type(scope ReadOnlyScope) (ast.ValueType, error) {
	if n.constReturnType == ast.InvalidType {
		// We are dynamic and we need to figure out our type
		// Do NOT cache this result in n.returnType since it can change.
		return n.nodeEvaluator.Type(scope)
	}
	return n.constReturnType, nil
}

func (n *EvalConditionNode) IsDynamic() bool {
	return n.nodeEvaluator.IsDynamic()
}

func (n *EvalConditionNode) EvalRegex(scope *Scope, _ ExecutionState) (*regexp.Regexp, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return nil, err
	}
	if typ == ast.TRegex {
		return n.nodeEvaluator.EvalRegex(scope, n.state)
	}

	return nil, ErrTypeGuardFailed{RequestedType: ast.TRegex, ActualType: typ}
}

func (n *EvalConditionNode) EvalTime(scope *Scope, _ ExecutionState) (time.Time, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return time.Time{}, err
	}
	return time.Time{}, ErrTypeGuardFailed{RequestedType: ast.TTime, ActualType: typ}
}

func (n *EvalConditionNode) EvalDuration(scope *Scope, _ ExecutionState) (time.Duration, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return 0, err
	}
	if typ == ast.TDuration {
		return n.nodeEvaluator.EvalDuration(scope, n.state)
	}

	return 0, ErrTypeGuardFailed{RequestedType: ast.TDuration, ActualType: typ}
}

func (n *EvalConditionNode) EvalString(scope *Scope, _ ExecutionState) (string, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return "", err
	}
	if typ == ast.TString {
		return n.nodeEvaluator.EvalString(scope, n.state)
	}

	return "", ErrTypeGuardFailed{RequestedType: ast.TString, ActualType: typ}
}

func (n *EvalConditionNode) EvalFloat(scope *Scope, _ ExecutionState) (float64, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return 0, err
	}
	if typ == ast.TFloat {
		return n.nodeEvaluator.EvalFloat(scope, n.state)
	}

	return 0, ErrTypeGuardFailed{RequestedType: ast.TFloat, ActualType: typ}
}

func (n *EvalConditionNode) EvalInt(scope *Scope, _ ExecutionState) (int64, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return 0, err
	}
	if typ == ast.TInt {
		return n.nodeEvaluator.EvalInt(scope, n.state)
	}

	return 0, ErrTypeGuardFailed{RequestedType: ast.TInt, ActualType: typ}
}

func (n *EvalConditionNode) EvalBool(scope *Scope, _ ExecutionState) (bool, error) {
	typ, err := n.Type(scope)
	if err != nil {
		return false, err
	}
	if typ == ast.TBool {
		return n.nodeEvaluator.EvalBool(scope, n.state)
	}

	return false, ErrTypeGuardFailed{RequestedType: ast.TBool, ActualType: typ}
}

// func (n *EvalLambdaNode) EvalMissing(scope *Scope, _ ExecutionState) (*ast.Missing, error) {
// 	typ, err := n.Type(scope)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if typ == ast.TMissing {
// 		return n.nodeEvaluator.EvalMissing(scope, n.state)
// 	}

// 	return nil, ErrTypeGuardFailed{RequestedType: ast.TBool, ActualType: typ}
// }

// getCostantNodeType - Given a ast.Node we want to know it's return type
// this method does exactly this, few examples:
// *) StringNode -> ast.TString
// *) UnaryNode -> we base the type by the node type
func getConstantNodeType(n ast.Node) ast.ValueType {
	switch node := n.(type) {
	case *ast.NumberNode:
		if node.IsInt {
			return ast.TInt
		}

		if node.IsFloat {
			return ast.TFloat
		}
	// case *ast.DurationNode:
	// 	return ast.TDuration
	case *ast.StringNode:
		return ast.TString
	// case *ast.BoolNode:
	// 	return ast.TBool
	// case *ast.RegexNode:
	// 	return ast.TRegex

	// case *ast.UnaryNode:
	// 	// If this is comparison operator we know for sure the output must be boolean
	// 	if node.Operator == ast.TokenNot {
	// 		return ast.TBool
	// 	}

	// 	// Could be float int or duration
	// 	if node.Operator == ast.TokenMinus {
	// 		return getConstantNodeType(node.Node)
	// 	}

	// case *ast.BinaryNode:
	// 	// Check first using only the operator
	// 	if ast.IsCompOperator(node.Operator) || ast.IsLogicalOperator(node.Operator) {
	// 		return ast.TBool
	// 	}
	// 	leftType := getConstantNodeType(node.Left)
	// 	rightType := getConstantNodeType(node.Right)
	// 	// Check known constant types
	// 	return binaryConstantTypes[operationKey{operator: node.Operator, leftType: leftType, rightType: rightType}]
	case *ast.ConditionNode:
		return getConstantNodeType(node.Expression)
	}

	return ast.InvalidType
}
