package html

import (
	"context"
	"fmt"
	"strings"

	"github.com/coreos/pkg/multierror"

	"gopkg.in/xmlpath.v2"
)

type Command struct {
	Name             string      `json:"name,omitempty"`
	Command          string      `json:"command,omitempty"`
	NestingCommands  Commands    `json:"nesting_commands,omitempty"`
	ExpressionString string      `json:"expression,omitempty"`
	Value            string      `json:"value,omitempty"`
	expression       IExpression `json:"expression,omitempty"`
	err              error       `json:"err,omitempty"`
}

func (c *Command) Err() error {
	return c.err
}

func (c *Command) handleNesting(ctx context.Context) (interface{}, error) {
	var errs multierror.Error
	var resultMap map[string]interface{} = nil

	for i, command := range c.NestingCommands {

		result, err := command.Exec(ctx)

		if err != nil {
			errs = append(errs, err)
		}
		if result == nil {
			continue
		}
		if resultMap == nil {
			resultMap = make(map[string]interface{})
		}
		switch {
		case isKV(result):
			kv := result.(map[string]interface{})
			for key, val := range kv {
				resultMap[key] = String(val)
			}
		default:
			key := fmt.Sprintf("%d", i)
			resultMap[key] = String(result)

		}
	}
	if len(resultMap) == 0 {
		return nil, errs.AsError()
	}
	// fmt.Println(resultMap)
	return resultMap, errs.AsError()
}
func (c *Command) Exec(ctx context.Context) (interface{}, error) {
	expr, err := c.GetExpr()
	if err != nil {
		c.err = err
	}

	if expr != nil {
		var errs multierror.Error
		result, err := expr.Exec(ctx)

		if err != nil {
			errs = append(errs, err)

		}

		if result == nil {
			c.err = errs.AsError()
			return nil, c.err
		}

		switch c.Command {
		case "getAll":
			var results []interface{}
			// result, err := expr.Exec(ctx)
			// if err != nil {
			// 	errs = append(errs, err)
			// }
			iter := result.(*xmlpath.Iter)
			var i = 0
			for iter.Next() {
				i++
				node := iter.Node()
				if len(c.NestingCommands) > 0 {
					rangeExpr := fmt.Sprintf("%s[%d]", expr, i)

					childCtx := SetParentXPathExpression(ctx, rangeExpr)
					nestingResult, err := c.handleNesting(childCtx)
					if err != nil {
						errs = append(errs, err)
					}

					if nestingResult != nil {
						results = append(results, nestingResult)
					}
				} else {
					results = append(results, String(node))
				}
			}
			if len(results) > 0 {
				if c.Value != "" {
					return toKeyValue(c.Value, results), c.err
				}
			}
			c.err = errs.AsError()
			return results, c.err
		case "get":
			var nestingResult interface{}
			var err error
			if len(c.NestingCommands) > 0 {
				rangeExpr := fmt.Sprintf("%s", expr)
				childCtx := SetParentXPathExpression(ctx, rangeExpr)
				nestingResult, err = c.handleNesting(childCtx)
				if err != nil {
					errs = append(errs, err)
					c.err = errs.AsError()
				}
				if nestingResult == nil {
					return nil, c.err
				}

				if c.Value != "" {
					return toKeyValue(c.Value, nestingResult), c.err
				}
				return nestingResult, c.err

			}
			// GetAttribute(ctx, nil, expr.(*XPathExpr), "href")
			if c.Value != "" {
				return toKeyValue(c.Value, result), c.err
			}
			return String(result), c.err
		}

	}

	return nil, c.err
}

func isKV(in interface{}) bool {
	if in == nil {
		return false
	}
	_, ok := in.(map[string]interface{})
	return ok
}

func toKeyValue(key string, val interface{}) map[string]interface{} {
	return map[string]interface{}{
		key: String(val),
	}
}

func String(in interface{}) interface{} {
	if in == nil {
		return ""
	}
	switch v := in.(type) {
	case *xmlpath.Iter:
		if v.Next() {
			return strings.TrimSpace(v.Node().String())
		}
	case string:
		return strings.TrimSpace(v)
		// default:
		// 	return fmt.Sprintf("%s", in)
	}
	src := fmt.Sprintf("%s", in)
	return strings.TrimSpace(src)
}

func Iter(in interface{}) *xmlpath.Iter {
	switch v := in.(type) {
	case *xmlpath.Iter:
		return v
	default:
		return nil
	}
}

func (c *Command) GetExpr() (IExpression, error) {
	if c.expression != nil {
		return c.expression, nil
	}

	expression, err := ParseExpresion(c.ExpressionString)
	if err != nil {
		return nil, err
	}
	c.expression = expression
	return expression, err
}

type Commands []Command

func (c Commands) Exec(ctx context.Context) (interface{}, error) {
	var results []interface{}
	for _, command := range c {
		result, err := command.Exec(ctx)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
