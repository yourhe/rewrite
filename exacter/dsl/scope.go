package dsl

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"gitlab.iyorhe.com/dr2am/dr2am-rewrite/exacter/dsl/ast"
)

// maxArgs is used to specify the largest number of arguments that a
// builtin function can accept.
// Increment this value if you create a builtin function with more than
// the current value of maxArgs.
const (
	maxArgs = 4
)

type Domain [maxArgs]ast.ValueType

// Special marker that a value is empty
var empty = new(interface{})

//Scope Contains a set of variables references and their values.
type Scope struct {
	variables map[string]interface{}
	// dynamicMethods map[string]DynamicMethod
	// dynamicFuncs   map[string]*DynamicFunc
}
type ReadOnlyScope interface {
	Get(name string) (interface{}, error)
	References() []string
	// DynamicFunc(name string) *DynamicFunc
}

// Get returns the value of 'name'.
func (s *Scope) Get(name string) (interface{}, error) {
	names := strings.Split(name, ".")
	var result interface{}
	var err error
	var variables interface{} = s.variables
	getName := func(name string) (string, []string) {
		names := strings.Split(name, "[")
		return names[0], names[1:]
	}

	for _, id := range names {
		id, chains := getName(id)
		result, err = get(variables, id)
		if err != nil || result == nil {
			return result, err
		}
		for _, chainID := range chains {
			chainID = chainID[:len(chainID)-1]
			var isString = strings.HasPrefix(chainID, "\"")
			var isInt bool
			var Identifier bool
			var idx int
			if isString {
				chainID = chainID[1 : len(chainID)-1]
			}

			if !isString {
				var err error
				idx, err = strconv.Atoi(chainID)
				if err == nil {
					isInt = true
				} else {
					Identifier = true
					value, err := s.Get(chainID)
					if err != nil || value == nil {
						return nil, err
					}

					idx, err = strconv.Atoi(fmt.Sprintf("%s", value))
					if err == nil {
						isInt = true
					} else {
						isInt = false
						chainID = fmt.Sprintf("%s", value)
					}
				}
			}

			if Identifier {

			}

			// var isString = strings.HasPrefix(chainID, "\"")
			if isInt {
				// idx, err := strconv.Atoi(chainID)

				rvalue := reflect.ValueOf(result)
				if rvalue.Kind() == reflect.Slice {
					if idx < rvalue.Len() {
						result = rvalue.Index(idx).Interface()
					} else {
						result = nil
					}
				}

				continue

				// found, ok := result.([]string)
				// fmt.Println(found, ok)
				// if ok {
				// 	result = found[idx]
				// }

			}
			if found, ok := result.(map[string]interface{}); ok {
				result, err = get(found, chainID)
				if err != nil || result == nil {
					return result, err
				}
				continue
			}
			return result, err
		}
		// fmt.Println(ok)
		// found, ok := result.(map[string]interface{})
		// if ok {
		// }
		variables = result

	}
	// switch v := result.(type) {
	// case []string:
	// 	return strings.Join(v, ","), err
	// }
	if v, ok := result.(interface{ String() string }); ok {
		return v.String(), nil
	}
	return result, err
	return get(s.variables, name)
	if v, ok := s.variables[name]; ok && v != empty {
		switch c := v.(type) {
		case int:
			return int64(c), nil
		case int32:
			return int64(c), nil
		}
		return v, nil
	}
	var possible []string
	for k := range s.variables {
		possible = append(possible, k)
	}
	return nil, fmt.Errorf("name %q is undefined. Names in scope: %s", name, strings.Join(possible, ","))
}

func get(in interface{}, name string) (interface{}, error) {
	if src, ok := in.(map[string]interface{}); ok {

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
	}
	val := reflect.ValueOf(in)
	elem := reflect.Indirect(val)
	if elem.Kind() == reflect.Struct {
		fieldVal := elem.FieldByName(name)
		if !fieldVal.IsZero() {
			fieldVal := elem.FieldByName(name)
			return fieldVal.Interface(), nil
		}
		return elem.Interface(), nil
	}
	var possible []string
	// for k := range src {
	// 	possible = append(possible, k)
	// }
	// fmt.Println(src, name)

	return nil, fmt.Errorf("namea %q is undefined. Names in scope: %s", name, strings.Join(possible, ","))
}
func (s *Scope) References() []string {
	ks := []string{}
	for k := range s.variables {
		ks = append(ks, "\""+k+"\"")
	}
	return ks
}

// Whether a value has been set on the scope
func (s *Scope) Has(name string) bool {
	v, ok := s.variables[name]
	return ok && v != empty
}

// Set defines a name -> value pairing in the scope.
func (s *Scope) Set(name string, value interface{}) {
	// var result interface{}
	// switch v := value.(type) {
	// case int, int32:
	// 	result = int64(v)
	// default:
	// 	result = v
	// }
	s.variables[name] = value
}

// Reset all scope values to an empty state.
func (s *Scope) Reset() {
	// Scopes, are intended to be reused so do not free resources
	for name := range s.variables {
		s.Set(name, empty)
	}
}

//Initialize a new Scope object.
func NewScope() *Scope {
	return &Scope{
		variables: make(map[string]interface{}),
		// dynamicMethods: make(map[string]DynamicMethod),
		// dynamicFuncs:   make(map[string]*DynamicFunc),
	}
}

// ScopePool - pooling mechanism for Scope
// The idea behind scope pool is to pool scopes and to put them only
// the needed variables for execution.
type ScopePool interface {
	Get() *Scope
	Put(scope *Scope)

	ReferenceVariables() []string
}

type scopePool struct {
	referenceVariables []string
	pool               sync.Pool
}

// NewScopePool - creates new ScopePool for the given Node
func NewScopePool(referenceVariables []string) ScopePool {
	scopePool := &scopePool{
		referenceVariables: referenceVariables,
	}

	scopePool.pool = sync.Pool{
		New: func() interface{} {
			scope := NewScope()
			for _, refVariable := range scopePool.referenceVariables {
				scope.Set(refVariable, empty)
			}

			return scope
		},
	}

	return scopePool
}

func (s *scopePool) ReferenceVariables() []string {
	return s.referenceVariables
}

// Get - returns a scope from a pool with the needed reference variables
// (with nil values/old values) in the scope
func (s *scopePool) Get() *Scope {
	return s.pool.Get().(*Scope)
}

// Put - put used scope back to the pool
func (s *scopePool) Put(scope *Scope) {
	scope.Reset()
	s.pool.Put(scope)
}
