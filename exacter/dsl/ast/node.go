package ast

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	octal   = 8
	decimal = 10
)

type Position interface {
	Position() int // byte position of start of node in full original input string
	Line() int
	Char() int
	Set(pos, line, char int)
}
type PositionSetter interface {
	SetPosition(pos, line, char int)
}

type Node interface {
	Position
	String() string
	Format(buf *bytes.Buffer, indent string, onNewLine bool)
	// Report whether to nodes are functionally equal, ignoring position and comments
	Equal(interface{}) bool

	json.Marshaler
	json.Unmarshaler
	// unmarshal(JSONNode) error
}
type position struct {
	pos  int
	line int
	char int
}

func (p position) Position() int {
	return p.pos
}
func (p position) Line() int {
	return p.line
}
func (p position) Char() int {
	return p.char
}
func (p position) String() string {
	return fmt.Sprintf("%dl%dc%d", p.pos, p.line, p.char)
}
func (p *position) Set(pos, line, char int) {
	(p).pos = pos
	(p).line = line
	(p).char = char
}

type ProgramNode struct {
	position
	Nodes []Node
}

// p *ProgramNode gitlab.iyorhe.com/dr2am/dr2am-rewrite/exacter/dsl/ast.Node

func (pn *ProgramNode) String() string {
	return fmt.Sprintf("ProgramNode@%v{%v}", pn.position, pn.Nodes)
}

func (pn *ProgramNode) Format(buf *bytes.Buffer, indent string, onNewLine bool) {
	panic("not implt")
}

func (pn *ProgramNode) Equal(interface{}) bool {
	panic("not implt")
}

// pn *ProgramNode json.Marshaler
func (pn *ProgramNode) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

//pn *ProgramNode json.Unmarshaler
func (pn *ProgramNode) UnmarshalJSON([]byte) error {
	panic("not implemented")
}

func (pn *ProgramNode) Add(node Node) {
	pn.Nodes = append(pn.Nodes, node)

}

func NewProgramNode(options ...Option) *ProgramNode {
	prog := ProgramNode{
		Nodes: []Node{},
	}
	for _, opt := range options {
		opt(&prog)
	}
	return &prog

}

// BinaryNode holds two arguments and an operator.
type BinaryNode struct {
	position
	Left      Node
	Right     Node
	Operator  TokenType
	Comment   *CommentNode
	Parens    bool
	MultiLine bool
}

func NewBinaryNode(left, right Node, multiLine bool, options ...Option) *BinaryNode {
	bn := &BinaryNode{
		Left:  left,
		Right: right,
		// Operator:  op,
		MultiLine: multiLine,
	}
	for _, opt := range options {
		opt(bn)
	}
	return bn
}

func (n *BinaryNode) SetOperator(tt TokenType) {
	n.Operator = tt
}
func (n *BinaryNode) SetParens(parens bool) {
	n.Parens = parens
}
func (n *BinaryNode) String() string {
	return fmt.Sprintf("BinaryNode@%v{p:%v m:%v %v %v %v}%v", n.position, n.Parens, n.MultiLine, n.Left, n.Operator, n.Right, n.Comment)

}
func (n *BinaryNode) Format(buf *bytes.Buffer, indent string, onNewLine bool) {
	if n.Comment != nil {
		n.Comment.Format(buf, indent, onNewLine)
		onNewLine = true
	}
	writeIndent(buf, indent, onNewLine)
	if n.Parens {
		buf.WriteByte('(')
		indent += indentStep
	}
	n.Left.Format(buf, indent, false)
	// _, ok := n.Left.(*BinaryNode)
	// if ok {
	// 	buf.WriteByte(' ')
	// }

	if n.Operator == TokenAnd || n.Operator == TokenOr {
		buf.WriteByte(' ')
	}
	buf.WriteString(n.Operator.String())
	if n.Operator == TokenAnd || n.Operator == TokenOr {
		buf.WriteByte(' ')
	}
	if n.MultiLine {
		buf.WriteByte('\n')
	} else {
		// _, ok := n.Right.(*BinaryNode)
		// if ok {
		// 	buf.WriteByte(' ')
		// }
	}
	n.Right.Format(buf, indent, n.MultiLine)
	if n.Parens {
		buf.WriteByte(')')
	}
}

func (n *BinaryNode) Equal(o interface{}) bool {
	if on, ok := o.(*BinaryNode); ok {
		return n.Operator == on.Operator &&
			n.Left.Equal(on.Left) &&
			n.Right.Equal(on.Right)
	}
	return false

}

// n *BinaryNode json.Marshaler
func (n *BinaryNode) MarshalJSON() ([]byte, error) {
	// panic("not implemented")
	props := map[string]interface{}{}
	props["typeof"] = "binary"
	props["left"] = n.Left
	props["right"] = n.Right
	props["operator"] = n.Operator
	props["operatorStr"] = n.Operator.String()
	props["parens"] = n.Parens
	return json.Marshal(&props)
}

//n *BinaryNode json.Unmarshaler
func (n *BinaryNode) UnmarshalJSON([]byte) error {
	panic("not implemented")
}

// Hold the contents of a comment
type CommentNode struct {
	position
	Comments []string
}

func (n *CommentNode) String() string {
	return fmt.Sprintf("CommentNode@%v{%v}", n.position, n.Comments)

}
func (n *CommentNode) Format(buf *bytes.Buffer, indent string, onNewLine bool) {
	panic("not implt")
}

func (n *CommentNode) Equal(interface{}) bool {
	panic("not implt")
}

func (n *CommentNode) MarshalJSON() ([]byte, error) {
	panic("not implemented")
}

func (n *CommentNode) UnmarshalJSON([]byte) error {
	panic("not implemented")
}
func (n *CommentNode) CommentString() string {
	return strings.Join(n.Comments, "\n")
}

// numberNode holds a number: signed or unsigned integer or float.
// The value is parsed and stored under all the types that can represent the value.
// This simulates in a small amount of code the behavior of Go's ideal constants.
type NumberNode struct {
	position
	IsInt   bool    // Number has an integral value.
	IsFloat bool    // Number has a floating-point value.
	Int64   int64   // The integer value.
	Float64 float64 // The floating-point value.
	Base    int     // The base of an integer value.
	Comment *CommentNode
}

//NewNumber create a new number from a text string
func NewNumberNode(text string, options ...Option) (*NumberNode, error) {
	if text == "" {
		return nil, errors.New("invalid number literal, empty string")
	}
	n := &NumberNode{
		// position: p,
		// Comment:  c,
	}

	for _, opt := range options {
		opt(n)
	}

	if s := strings.IndexRune(text, '.'); s != -1 {
		f, err := strconv.ParseFloat(text, 64)
		if err == nil {
			n.IsFloat = true
			n.Float64 = f
			if n.Float64 < 0 {
				return nil, errors.New("parser should not allow for negative number nodes")
			}
		}
	} else {
		if text[0] == '0' && len(text) > 1 {
			// We have an octal number
			n.Base = octal
		} else {
			n.Base = decimal
		}
		i, err := strconv.ParseInt(text, n.Base, 64)
		if err == nil {
			n.IsInt = true
			n.Int64 = i
		}
		if n.Int64 < 0 {
			return nil, errors.New("parser should not allow for negative number nodes")
		}
	}

	if !n.IsInt && !n.IsFloat {
		return nil, fmt.Errorf("illegal number syntax: %q", text)
	}
	return n, nil
}

func (n *NumberNode) String() string {
	if n.IsInt {
		return fmt.Sprintf("NumberNode@%v{%di,b%d}%v", n.position, n.Int64, n.Base, n.Comment)
	}
	return fmt.Sprintf("NumberNode@%v{%f,b%d}%v", n.position, n.Float64, n.Base, n.Comment)
}

func (n *NumberNode) Format(buf *bytes.Buffer, indent string, onNewLine bool) {
	if n.Comment != nil {
		n.Comment.Format(buf, indent, onNewLine)
		onNewLine = true
	}
	writeIndent(buf, indent, onNewLine)
	if n.IsInt {
		if n.Base == octal {
			buf.WriteByte('0')
		}
		buf.WriteString(strconv.FormatInt(n.Int64, n.Base))
	} else {
		s := strconv.FormatFloat(n.Float64, 'f', -1, 64)
		if strings.IndexRune(s, '.') == -1 {
			s += ".0"
		}
		buf.WriteString(s)
	}
}
func (n *NumberNode) SetComment(c *CommentNode) {
	n.Comment = c
}

func (n *NumberNode) Equal(o interface{}) bool {
	if on, ok := o.(*NumberNode); ok {
		return n.IsInt == on.IsInt &&
			n.IsFloat == on.IsFloat &&
			n.Int64 == on.Int64 &&
			n.Float64 == on.Float64
	}
	return false
}

// MarshalJSON converts the node to JSON with an additional
// typeOf field.
func (n *NumberNode) MarshalJSON() ([]byte, error) {

	return json.Marshal(n)
}

// UnmarshalJSON converts JSON bytes to a Number node.
func (n *NumberNode) UnmarshalJSON(data []byte) error {
	var props NumberNode
	err := json.Unmarshal(data, &props)
	if err != nil {
		return err
	}
	return nil
}

//Holds the textual representation of a string literal
type StringNode struct {
	position
	Literal      string // The string literal
	TripleQuotes bool
	Comment      *CommentNode
}

// MarshalJSON converts the node to JSON with an additional
// typeOf field.
func (n *StringNode) MarshalJSON() ([]byte, error) {
	props := map[string]interface{}{}
	props["typeof"] = "string"
	props["literal"] = n.Literal
	return json.Marshal(&props)
	// return json.Marshal(n)
}

// UnmarshalJSON converts JSON bytes to a StringNode
func (n *StringNode) UnmarshalJSON(data []byte) error {
	var props StringNode
	err := json.Unmarshal(data, &props)
	if err != nil {
		return err
	}
	return nil
}

func NewStringNode(txt string, options ...Option) *StringNode {

	tripleQuotes := false
	// Remove leading and trailing quotes
	var literal string
	if len(txt) >= 6 && txt[0:3] == "'''" {
		literal = txt[3 : len(txt)-3]
		tripleQuotes = true
	} else {
		literal = txt[1 : len(txt)-1]
		quote := txt[0]
		// Unescape quotes
		var buf bytes.Buffer
		buf.Grow(len(literal))
		last := 0
		for i := 0; i < len(literal)-1; i++ {
			if literal[i] == '\\' && literal[i+1] == quote {
				buf.Write([]byte(literal[last:i]))
				i++
				last = i
			}
		}
		buf.Write([]byte(literal[last:]))
		literal = buf.String()
	}

	sn := &StringNode{
		Literal:      literal,
		TripleQuotes: tripleQuotes,
	}
	for _, opt := range options {
		opt(sn)
	}
	return sn
}

func (n *StringNode) String() string {
	return fmt.Sprintf("StringNode@%v{%s}%v", n.position, n.Literal, n.Comment)
}

func (n *StringNode) Format(buf *bytes.Buffer, indent string, onNewLine bool) {
	if n.Comment != nil {
		n.Comment.Format(buf, indent, onNewLine)
		onNewLine = true
	}
	writeIndent(buf, indent, onNewLine)
	if n.TripleQuotes {
		buf.WriteString("'''")
	} else {
		buf.WriteByte('\'')
	}
	if n.TripleQuotes {
		buf.WriteString(n.Literal)
	} else {
		for _, c := range n.Literal {
			if c == '\'' {
				buf.WriteByte('\\')
			}
			buf.WriteRune(c)
		}
	}
	if n.TripleQuotes {
		buf.WriteString("'''")
	} else {
		buf.WriteByte('\'')
	}
}
func (n *StringNode) SetComment(c *CommentNode) {
	n.Comment = c
}

func (n *StringNode) Equal(o interface{}) bool {
	if on, ok := o.(*StringNode); ok {
		return n.Literal == on.Literal
	}
	return false
}

//Holds the textual representation of an identifier
type IdentifierNode struct {
	position
	Ident   string // The identifier
	Comment *CommentNode
}

func NewIdentifierNode(ident string, options ...Option) *IdentifierNode {
	node := IdentifierNode{
		Ident: ident,
	}
	for _, opt := range options {
		opt(&node)
	}
	return &node
}

func (n *IdentifierNode) String() string {
	return fmt.Sprintf("IdentifierNode@%v{%v}%v", n.position, n.Ident, n.Comment)
}
func (n *IdentifierNode) Format(buf *bytes.Buffer, indent string, onNewLine bool) {
	if n.Comment != nil {
		n.Comment.Format(buf, indent, onNewLine)
		onNewLine = true
	}
	writeIndent(buf, indent, onNewLine)
	buf.WriteString(n.Ident)
}

func (n *IdentifierNode) Equal(interface{}) bool {
	panic("not implt")
}

// n *IdentifierNode json.Marshaler

func (n *IdentifierNode) MarshalJSON() ([]byte, error) {
	props := map[string]interface{}{}
	props["typeof"] = "identifier"
	props["ident"] = n.Ident
	return json.Marshal(&props)

}

//n *IdentifierNode json.Unmarshaler
func (n *IdentifierNode) UnmarshalJSON([]byte) error {
	panic("not implemented")
}

// Represents the beginning of a lambda expression
type ConditionNode struct {
	position
	Expression Node
	Comment    *CommentNode
}

func NewConditionNode(node Node, options ...Option) *ConditionNode {
	ln := &ConditionNode{
		Expression: node,
	}
	for _, opt := range options {
		opt(ln)
	}
	return ln
}

// n *LambdaNode gitlab.iyorhe.com/dr2am/dr2am-script-lang/ast.Node
// n *LambdaNode Node
func (n *ConditionNode) String() string {
	return fmt.Sprintf("ConditionNode@%v{%v}%v", n.position, n.Expression, n.Comment)
}
func (n *ConditionNode) Format(buf *bytes.Buffer, indent string, onNewLine bool) {
	if n.Comment != nil {
		n.Comment.Format(buf, indent, onNewLine)
		onNewLine = true
	}
	writeIndent(buf, indent, onNewLine)
	buf.WriteString("when ")
	n.Expression.Format(buf, indent, onNewLine)
}

func (n *ConditionNode) Equal(o interface{}) bool {
	if on, ok := o.(*ConditionNode); ok {
		return (n == nil && on == nil) || n.Expression.Equal(on.Expression)
	}
	return false
}

// n *LambdaNode json.Marshaler
func (n *ConditionNode) MarshalJSON() ([]byte, error) {
	// panic("not implemented")
	props := map[string]interface{}{}
	props["typeof"] = "condition"
	props["expression"] = n.Expression
	return json.Marshal(&props)
	// return []byte(`{
	// 	"command":"condition",
	// 	"Expression":{}
	// 	}`), nil
}

//n *LambdaNode json.Unmarshaler
func (n *ConditionNode) UnmarshalJSON([]byte) error {
	panic("not implemented")
}

//Holds the textual representation of an identifier
type ReferenceNode struct {
	position
	Reference string // The field reference
	Comment   *CommentNode
}

func NewReference(txt string, options ...Option) *ReferenceNode {
	literal := txt
	// Remove leading and trailing quotes

	if strings.HasPrefix(txt, "@") {
		literal = txt[1 : len(txt)-1]
	}

	// Unescape quotes
	var buf bytes.Buffer
	buf.Grow(len(literal))
	last := 0
	for i := 0; i < len(literal)-1; i++ {
		if literal[i] == '\\' && literal[i+1] == '"' {
			buf.Write([]byte(literal[last:i]))
			i++
			last = i
		}
	}
	buf.Write([]byte(literal[last:]))
	literal = buf.String()

	rn := &ReferenceNode{
		Reference: literal,
	}
	for _, opt := range options {
		opt(rn)
	}
	return rn
}
func (n *ReferenceNode) String() string {
	return fmt.Sprintf("ReferenceNode@%v{%s}%v", n.position, n.Reference, n.Comment)
}
func (n *ReferenceNode) Format(buf *bytes.Buffer, indent string, onNewLine bool) {
	if n.Comment != nil {
		n.Comment.Format(buf, indent, onNewLine)
		onNewLine = true
	}
	writeIndent(buf, indent, onNewLine)
	buf.WriteString(n.Reference)
}

func (n *ReferenceNode) Equal(o interface{}) bool {
	if on, ok := o.(*ReferenceNode); ok {
		return n.Reference == on.Reference
	}
	return false

}

// n *ReferenceNode json.Marshaler
func (n *ReferenceNode) MarshalJSON() ([]byte, error) {
	props := map[string]interface{}{}
	props["typeof"] = "reference"
	props["reference"] = n.Reference
	return json.Marshal(&props)
}

//n *ReferenceNode json.Unmarshaler
func (n *ReferenceNode) UnmarshalJSON([]byte) error {
	panic("not implemented")
}

//Holds the textual representation of a regex literal
type RegexNode struct {
	position
	Regex   *regexp.Regexp
	Literal string
	Comment *CommentNode
}

// MarshalJSON converts the node to JSON with an additional
// typeOf field.
func (n *RegexNode) MarshalJSON() ([]byte, error) {
	props := map[string]interface{}{}
	props["typeof"] = "regexp"
	props["literal"] = n.Literal
	return json.Marshal(&props)
}

// UnmarshalJSON converts JSON bytes to a RegexNode
func (n *RegexNode) UnmarshalJSON(data []byte) error {
	panic("not implemented")
}

func NewRegexNode(txt string, options ...Option) (*RegexNode, error) {
	// Remove leading and trailing quotes
	literal := txt[1 : len(txt)-1]
	// Unescape slashes '/'
	var buf bytes.Buffer
	buf.Grow(len(literal))
	last := 0
	for i := 0; i < len(literal)-1; i++ {
		if literal[i] == '\\' && literal[i+1] == '/' {
			buf.Write([]byte(literal[last:i]))
			i++
			last = i
		}
	}
	buf.Write([]byte(literal[last:]))
	unescaped := buf.String()

	r, err := regexp.Compile(unescaped)
	if err != nil {
		return nil, err
	}
	node := RegexNode{
		// position: p,
		Regex:   r,
		Literal: literal,
		// Comment:  c,
	}
	for _, opt := range options {
		opt(&node)
	}
	return &node, nil
}

func (n *RegexNode) String() string {
	return fmt.Sprintf("RegexNode@%v{%v}%v", n.position, n.Regex, n.Comment)
}

func (n *RegexNode) Format(buf *bytes.Buffer, indent string, onNewLine bool) {
	if n.Comment != nil {
		n.Comment.Format(buf, indent, onNewLine)
		onNewLine = true
	}
	writeIndent(buf, indent, onNewLine)
	buf.WriteByte('/')
	buf.WriteString(n.Literal)
	buf.WriteByte('/')
}

func (n *RegexNode) SetComment(c *CommentNode) {
	n.Comment = c
}
func (n *RegexNode) Equal(o interface{}) bool {
	if on, ok := o.(*RegexNode); ok {
		return n.Regex.String() == on.Regex.String()
	}
	return false
}
