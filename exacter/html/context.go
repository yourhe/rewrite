package html

import (
	"context"

	"gopkg.in/xmlpath.v2"
)

type contextKey string

const ValueKey = contextKey("value")
const RootNodeKey = contextKey("rootNode")
const ResoultSetKey = contextKey("resoultSet")
const ParentXPathExpressionKey = contextKey("parentXPathExpression")

func SetResoultSet(ctx context.Context, resoult *xmlpath.Iter) context.Context {
	return context.WithValue(ctx, ResoultSetKey, resoult)
}

func GetResoutSet(ctx context.Context) *xmlpath.Iter {
	val := ctx.Value(ResoultSetKey)
	if val == nil {
		return nil
	}
	iter, ok := val.(*xmlpath.Iter)
	if ok {
		return iter
	}
	return nil
}

func SetParentXPathExpression(ctx context.Context, expression string) context.Context {
	return context.WithValue(ctx, ParentXPathExpressionKey, expression)
}

func GetParentXPathExpression(ctx context.Context) string {
	val := ctx.Value(ParentXPathExpressionKey)
	if val == nil {
		return ""
	}
	expression, ok := val.(string)
	if ok {
		return expression
	}
	return ""
}
