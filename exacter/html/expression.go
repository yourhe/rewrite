package html

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"strings"

	"gitlab.iyorhe.com/dr2am/dr2am-rewrite/exacter/dsl"
	"gitlab.iyorhe.com/dr2am/dr2am-rewrite/exacter/dsl/ast"
)

type IExpression interface {
	Exec(ctx context.Context) (interface{}, error)
}

func ParseExpresion(str string) (IExpression, error) {
	idx := strings.Index(str, "=")
	if idx < 0 {
		return nil, errors.New("Unknown expression definition")
	}
	expr := str[:idx]
	val := str[idx+1:]
	switch expr {
	case "xpath":
		return &XPathExpr{
			Name:       expr,
			Expression: val,
		}, nil
	case "set":
		return &SetExpr{
			Expr: Expr{
				Name:       expr,
				Expression: val,
			},
		}, nil
	case "json":
		return &JSONExpr{
			Expr: Expr{
				Name:       expr,
				Expression: val,
			},
		}, nil
	}
	return nil, errors.New("Unknown expression definition")
}

func ParseConditionExpression(str string) (*ConditionExpr, error) {
	node, err := ast.Parse(str)
	if err != nil {
		return nil, err
	}
	program, ok := node.(*ast.ProgramNode)
	if !ok || len(program.Nodes) == 0 {
		return nil, errors.New("script parser error : not defined ast.ProgramNode")
	}

	dslExpression, err := dsl.NewExpression(program.Nodes[0])
	if err != nil {
		return nil, err
	}
	return &ConditionExpr{
		Expr: Expr{
			Name:       "condition",
			Expression: str,
		},
		dslExpression: dslExpression,
	}, nil
}

func ParseMapExpression(str string) (*MapExpr, error) {
	var dslExpression dsl.Expression
	var err error
	if strings.HasPrefix(strings.TrimSpace(str), "map ") {
		var node ast.Node
		node, err = ast.Parse(str)
		if err != nil {
			goto out
		}
		program, ok := node.(*ast.ProgramNode)
		if !ok || len(program.Nodes) == 0 {
			err = errors.New("script parser error : not defined ast.ProgramNode")
			goto out
		}

		dslExpression, err = dsl.NewExpression(program.Nodes[0])
		// dslExpression = dsl.NewExpression()
	}
out:
	return &MapExpr{
		Name:          "map",
		Expression:    str,
		dslExpression: dslExpression,
	}, err
}

type Expr struct {
	Name       string `json:"name,omitempty"`
	Expression string `json:"expression,omitempty"`
}

type SetExpr struct {
	Expr
}

func (c *SetExpr) Exec(ctx context.Context) (interface{}, error) {
	return c.Expression, nil
}

type JSONExpr struct {
	Expr
}

func (c *JSONExpr) Exec(ctx context.Context) (interface{}, error) {

	var resource = map[string]interface{}{}
	body := GetBody(ctx)
	err := json.Unmarshal(body, &resource)
	ctx = context.WithValue(ctx, ResourcesKey, resource)
	return resource, err
}

type ConditionExpr struct {
	Expr
	dslExpression dsl.Expression
}

func (c *ConditionExpr) Exec(ctx context.Context) (interface{}, error) {
	stdReq := GetRequest(ctx)
	scope := dsl.NewScope()
	resources := GetReources(ctx)
	for key, val := range resources {
		scope.Set(key, val)
	}
	url := map[string]interface{}{
		"hostname": stdReq.URL.Hostname(),
		"host":     stdReq.URL.Host,
		"path":     stdReq.URL.Path,
		"query":    stdReq.URL.Query(),
	}
	request := map[string]interface{}{
		"url": url,
	}
	scope.Set("request", request)
	return c.dslExpression.EvalBool(scope)
}

type condition struct {
}

func NewConditon(str string) condition {
	return condition{}
}

type boolCondition struct {
	mustClauses    []condition
	mustNotClauses []condition
	filterClauses  []condition
	shouldClauses  []condition
}

func (b *boolCondition) Must(conditions ...condition) *boolCondition {
	b.mustClauses = append(b.mustClauses, conditions...)
	return b
}
func (b *boolCondition) Should(conditions ...condition) *boolCondition {
	b.mustClauses = append(b.mustClauses, conditions...)
	return b
}
func parseToCondition(str string) {
	var skipSpace = func(c byte) {

	}
	var read = func(src string) (string, int) {
		var buf = bytes.NewBuffer(nil)
		i := 0
		for ; i < len(str); i++ {
			c := str[i]
			switch c {
			case '?', '!':
				break
			}
			buf.WriteByte(c)
		}
		return buf.String(), buf.Len()
	}
	boolCondtion := &boolCondition{}
	strings.Index(str, "!")
	var stack = bytes.NewBuffer(nil)
	for i := 0; i < len(str); i++ {
		c := str[i]
		skipSpace(c)
		switch c {
		case '?':
			stack.Reset()
			e, l := read(str[i+1:])
			condition := NewConditon(e)
			boolCondtion.Should(condition)
			i += l
		case '!':
			stack.Reset()
			e, l := read(str[i+1:])
			condition := NewConditon(e)
			boolCondtion.Must(condition)
			i += l
		default:
			stack.WriteByte(c)
		}
	}
}
