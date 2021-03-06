package html

import (
	"context"
	"errors"
	"fmt"

	"gopkg.in/xmlpath.v2"
)

type XPathExpr struct {
	Name       string `json:"name,omitempty"`
	Expression string `json:"expression,omitempty"`
	expression string
	ctx        context.Context
	root       *xmlpath.Node
}

func (c *XPathExpr) Exec(ctx context.Context) (interface{}, error) {
	return c.ExecReturnIter(ctx)
}

func (c *XPathExpr) GetRootNode(ctx context.Context) (*xmlpath.Node, error) {
	if c.root == nil {
		val := ctx.Value(RootNodeKey)
		node, ok := val.(*xmlpath.Node)
		if !ok || node == nil {
			return nil, errors.New("Unknown rootNode definition")
		}
		c.root = node
	}
	return c.root, nil
}

func (c *XPathExpr) ExecReturnNode(ctx context.Context) (interface{}, error) {
	node, err := c.GetRootNode(ctx)
	if err != nil {
		return nil, err
	}
	expr := c.Expression
	if expr[0] == '^' {
		expr = expr[1:]
	} else {
		parentExpression := GetParentXPathExpression(ctx)
		expr = fmt.Sprintf("%s%s", parentExpression, expr)
	}
	path, err := xmlpath.Compile(expr)
	if err != nil {
		return nil, err
	}
	c.expression = expr
	exist := path.Exists(node)
	if !exist {
		return nil, nil
	}
	iter := path.Iter(node)
	nodes := []*xmlpath.Node{}
	for iter.Next() {
		nodes = append(nodes, iter.Node())
	}

	return nodes, nil
}

func (c *XPathExpr) ExecReturnIter(ctx context.Context) (interface{}, error) {
	node, err := c.GetRootNode(ctx)
	if err != nil {
		return nil, err
	}

	expr := c.Expression
	if expr[0] == '^' {
		expr = expr[1:]
	} else {
		parentExpression := GetParentXPathExpression(ctx)
		expr = fmt.Sprintf("%s%s", parentExpression, expr)
	}
	// parentExpression := GetParentXPathExpression(ctx)
	// expr := fmt.Sprintf("%s%s", parentExpression, c.Expression)
	path, err := xmlpath.Compile(expr)
	if err != nil {
		return nil, err
	}
	c.expression = expr
	exist := path.Exists(node)
	if !exist {
		return nil, nil
	}
	iter := path.Iter(node)

	return iter, nil
}

func (c *XPathExpr) String() string {
	return c.expression
}

func (c *XPathExpr) GetAttribute(ctx context.Context, node *xmlpath.Node, key string) string {
	return GetAttribute(ctx, node, c, key)
}

func GetAttribute(ctx context.Context, node *xmlpath.Node, expr *XPathExpr, key string) string {

	query := fmt.Sprintf("%s/@%s", expr, key)
	attr, err := xmlpath.Compile(query)
	if err != nil || attr == nil {
		return ""
	}
	result, _ := attr.String(node)
	return result
}
