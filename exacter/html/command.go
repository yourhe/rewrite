package html

import (
	"context"
	"fmt"
	"strings"

	"github.com/coreos/pkg/multierror"

	"gopkg.in/xmlpath.v2"
)

type Command struct {
	ID               string      `json:"id,omitempty"`
	Name             string      `json:"name,omitempty"`
	Command          string      `json:"command,omitempty"`
	NestingCommands  Commands    `json:"nesting_commands,omitempty"`
	ExpressionString string      `json:"expression,omitempty"`
	Value            string      `json:"value,omitempty"`
	expression       IExpression `json:"expression,omitempty"`
	err              error       `json:"err,omitempty"`
	resource         interface{}
}

func (c *Command) Err() error {
	return c.err
}

func (c *Command) handleNesting(ctx context.Context) (interface{}, error) {
	var errs multierror.Error
	var resultMap map[string]interface{} = nil

	for i, command := range c.NestingCommands {
		// command.resource = c.resource
		if command.Command == "map" {
			// fmt.Println("handleNesting", command.Command)
			ctx = AddReources(ctx, resultMap)
			// fmt.Println(resultMap)
		}
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
	// fmt.Println("Exec", c.Command, GetReources(ctx))
	_, expr, err := c.GetExpr()
	if err != nil {
		c.err = err
	}

	if expr != nil {
		var errs multierror.Error
		// ctx = SetReources(ctx, c.resource)
		result, err := expr.Exec(ctx)

		// fmt.Println("Execd", c.Command, GetReources(ctx))
		// fmt.Println("Execd", result, GetReources(ctx))

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
				childCtx = SetReources(childCtx, result)
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
		case "condition":
			if allow, ok := result.(bool); !ok ||
				!allow {
				return nil, c.err
			}
			var nestingResult interface{}
			var err error
			if len(c.NestingCommands) > 0 {
				// rangeExpr := fmt.Sprintf("%s", expr)
				// childCtx := SetParentXPathExpression(ctx, rangeExpr)
				nestingResult, err = c.handleNesting(ctx)
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

		case "map":
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
		key: val,
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
	case []interface{}:
		return in
	case *xmlpath.Node:
		return strings.TrimSpace(fmt.Sprintf("%s", in))
	default:
		return in
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

func (c *Command) GetExpr() (context.Context, IExpression, error) {
	ctx := context.Background()
	if c.expression != nil {
		return ctx, c.expression, nil
	}
	var expression IExpression
	var err error
	switch c.Command {
	case "condition":
		expression, err = ParseConditionExpression(c.ExpressionString)
	case "map":

		mapExpr, errx := ParseMapExpression(c.ExpressionString)
		expression = mapExpr
		err = errx
	default:
		expression, err = ParseExpresion(c.ExpressionString)

	}

	if err != nil {
		return ctx, nil, err
	}
	c.expression = expression
	return ctx, expression, err
}

type Commands []Command

func (c Commands) Exec(ctx context.Context) (interface{}, error) {
	var results []interface{}
	var errs multierror.Error
	for _, command := range c {
		// fmt.Println(ctx)
		// command.resource = &results
		result, err := command.Exec(ctx)
		// ctx = context.WithValue(ctx, ResourcesKey, "aaaaa")
		// fmt.Println(ctx)
		// fmt.Println(err)
		// fmt.Println("commands Exec")
		// fmt.Println(command.Command, command.ExpressionString)
		if err != nil {
			errs = append(errs, err)
			// 	return nil, err
		}
		results = append(results, result)
		// fmt.Println("command.resource", command.resource)
	}
	return results, errs.AsError()
}

func NewCommand(command string, expression string, value string, commands ...Command) Command {
	return Command{
		Command:          command,
		ExpressionString: expression,
		Value:            value,
		ID:               NewID(),
		NestingCommands:  commands,
	}
}
