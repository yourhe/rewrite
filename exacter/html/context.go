package html

import (
	"context"
	"net/http"

	"gopkg.in/xmlpath.v2"
)

type contextKey string

const ValueKey = contextKey("value")
const RootNodeKey = contextKey("rootNode")
const RequestKey = contextKey("request")
const ResponseKey = contextKey("response")
const BodyKey = contextKey("body")
const ResoultSetKey = contextKey("resoultSet")
const ParentXPathExpressionKey = contextKey("parentXPathExpression")
const ResourcesKey = contextKey("resource")

// const JSONReourcesKey = contextKey("json_resource")

func SetReources(ctx context.Context, resource interface{}) context.Context {
	return context.WithValue(ctx, ResourcesKey, resource)
}

func AddReources(ctx context.Context, resources map[string]interface{}) context.Context {
	if resources == nil || len(resources) == 0 {
		return ctx
	}
	old := GetReources(ctx)
	if old != nil {

		for k, v := range resources {
			old[k] = v
		}
		return context.WithValue(ctx, ResourcesKey, old)
	}
	return context.WithValue(ctx, ResourcesKey, resources)
}
func GetReources(ctx context.Context) map[string]interface{} {
	val := ctx.Value(ResourcesKey)
	if val == nil {
		return nil
	}
	iter, ok := val.(map[string]interface{})
	if ok {
		return iter
	}
	return nil
}

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

func GetRequest(ctx context.Context) *http.Request {
	val := ctx.Value(RequestKey)
	if val == nil {
		return nil
	}
	req, ok := val.(*http.Request)
	if ok {
		return req
	}
	return nil
}

func GetResponse(ctx context.Context) *http.Response {
	val := ctx.Value(ResponseKey)
	if val == nil {
		return nil
	}
	resp, ok := val.(*http.Response)
	if ok {
		return resp
	}
	return nil
}

func GetBody(ctx context.Context) []byte {
	val := ctx.Value(BodyKey)
	if val == nil {
		return nil
	}
	body, ok := val.([]byte)
	if ok {
		return body
	}
	return nil
}
