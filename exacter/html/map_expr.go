package html

import (
	"bytes"
	"context"
	"fmt"
	"net/url"
	"strings"
	"text/template"

	"gitlab.iyorhe.com/dr2am/dr2am-rewrite/exacter/dsl"
)

type MapExpr struct {
	Name          string `json:"name,omitempty"`
	Expression    string `json:"expression,omitempty"`
	expression    string
	ctx           context.Context
	dslExpression dsl.Expression

	// root       *xmlpath.Node
}

type warpMap map[string]interface{}
type warp struct {
	Interface interface{}
}

func newWarp(i interface{}) interface{} {
	switch v := i.(type) {
	case map[string]interface{}:
		return newMapWarpp(v)
	}
	return newSampleWarp(i)
}
func newSampleWarp(i interface{}) *warp {
	return &warp{Interface: i}
}
func newMapWarpp(i map[string]interface{}) map[string]interface{} {
	var result = make(map[string]interface{}, len(i))
	for k, v := range i {
		result[k] = newWarp(v)
	}
	return result
}
func (w *warp) String() string {
	return fmt.Sprintf("%v", w.Interface)
}
func (w *warp) GetChild(n int) interface{} {
	// fmt.Println("GetChild")
	if v, ok := w.Interface.([]interface{}); ok {
		if len(v) < n {
			return nil
		}
		// return newWarp(v[n])
		return v[n]
	}
	return nil
}
func (w *warp) URL() *url.URL {
	link := w.String()
	uri, err := url.Parse(link)
	if err == nil {
		return uri
	}
	return &url.URL{}
}

func (me *MapExpr) ExecMapExpr(ctx context.Context) (interface{}, error) {
	// stdReq := GetRequest(ctx)
	resources := GetReources(ctx)
	request := GetRequest(ctx)
	response := GetResponse(ctx)
	scope := dsl.NewScope()

	for key, val := range resources {
		// data[key] = val
		scope.Set(key, val)
	}
	scope.Set("request", request)
	scope.Set("response", response)
	// url := map[string]interface{}{
	// 	"hostname": stdReq.URL.Hostname(),
	// 	"host":     stdReq.URL.Host,
	// 	"path":     stdReq.URL.Path,
	// 	"query":    stdReq.URL.Query(),
	// }
	// request := map[string]interface{}{
	// 	"url": url,
	// }
	// for key, val := range data {

	// }
	// fmt.Println(me.Expression)
	return me.dslExpression.Eval(scope)
}
func (me *MapExpr) Exec(ctx context.Context) (resoult interface{}, err error) {
	defer func() {
		if rerr := recover(); rerr != nil {
			err = fmt.Errorf("html:exacter:mapExpr:%v\n", rerr)
		}
	}()
	resources := GetReources(ctx)
	request := GetRequest(ctx)
	response := GetResponse(ctx)

	data := map[string]interface{}{}

	for key, val := range resources {
		data[key] = newWarp(val)
	}
	data["request"] = request
	data["response"] = response
	if me.dslExpression != nil {
		return me.ExecMapExpr(ctx)
	}

	t, err := template.New("").Parse(me.Expression)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)

	err = t.Execute(buf, data)
	return buf.String(), err
}

// Special marker that a value is empty
var empty = new(interface{})

func Get(src map[string]interface{}, expression string) (interface{}, error) {
	var result interface{}

	return result, nil
}

func get(src map[string]interface{}, name string) (interface{}, error) {

	if v, ok := src[name]; ok && v != empty {
		switch c := v.(type) {
		case int:
			return int64(c), nil
		case int32:
			return int64(c), nil
		case map[string]interface{}:
			return c, nil
		case map[string][]string:
			result := map[string]interface{}{}
			for key, val := range c {
				result[key] = val
			}
			return result, nil
		case url.Values:
			result := map[string]interface{}{}
			for key, val := range c {
				result[key] = val
			}
			return result, nil

		}
		return v, nil
	}
	var possible []string
	for k := range src {
		possible = append(possible, k)
	}
	return nil, fmt.Errorf("name %q is undefined. Names in scope: %s", name, strings.Join(possible, ","))
}
