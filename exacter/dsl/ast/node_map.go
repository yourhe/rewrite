package ast

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Represents the beginning of a lambda expression
type MapNode struct {
	position
	Expression Node
	Comment    *CommentNode
}

func NewMapNode(node Node, options ...Option) *MapNode {
	ln := &MapNode{
		Expression: node,
	}
	for _, opt := range options {
		opt(ln)
	}
	return ln
}

// n *LambdaNode gitlab.iyorhe.com/dr2am/dr2am-script-lang/ast.Node
// n *LambdaNode Node
func (n *MapNode) String() string {
	return fmt.Sprintf("MapNode@%v{%v}%v", n.position, n.Expression, n.Comment)
}
func (n *MapNode) Format(buf *bytes.Buffer, indent string, onNewLine bool) {
	if n.Comment != nil {
		n.Comment.Format(buf, indent, onNewLine)
		onNewLine = true
	}
	writeIndent(buf, indent, onNewLine)
	buf.WriteString("map ")
	n.Expression.Format(buf, indent, onNewLine)
}

func (n *MapNode) Equal(o interface{}) bool {
	if on, ok := o.(*MapNode); ok {
		return (n == nil && on == nil) || n.Expression.Equal(on.Expression)
	}
	return false
}

// n *LambdaNode json.Marshaler
func (n *MapNode) MarshalJSON() ([]byte, error) {
	// panic("not implemented")
	props := map[string]interface{}{}
	props["typeof"] = "map"
	props["expression"] = n.Expression
	return json.Marshal(&props)
	// return []byte(`{
	// 	"command":"condition",
	// 	"Expression":{}
	// 	}`), nil
}

//n *LambdaNode json.Unmarshaler
func (n *MapNode) UnmarshalJSON([]byte) error {
	panic("not implemented")
}
