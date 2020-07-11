package ast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
)

func TestAst(t *testing.T) {
	// input := antlr.NewInputStream(testdata)
	// lexer := parser.NewdslLexer(input)
	// stream := antlr.NewCommonTokenStream(lexer, 0)
	// p := parser.NewdslParser(stream)
	// p.BuildParseTrees = true
	// tree := p.Program()
	// builder := NewAstBuilder()
	// builder.Visit(tree)
	fmt.Println(testdata)
	node, err := Parse(testdata)
	if err != nil {
		t.Fatal(err)
	}
	// fmt.Println(node.(*ProgramNode).Nodes[0])
	// fmt.Println(node.(*ProgramNode).Nodes[0].MarshalJSON())
	printJSON(node.(*ProgramNode).Nodes[0])
	buf := bytes.NewBuffer(nil)
	node.(*ProgramNode).Nodes[0].Format(buf, "", true)
	fmt.Println(buf.String())

}
func printJSON(src interface{}) {
	bs, err := json.MarshalIndent(src, "", "  ")
	fmt.Println(string(bs), err)
}

// var testdata = `when domain is "a" and domain is "a"`
// var testdata = `when domain:"ab" AND (domain:"b" OR domain:"c" OR domain:"c") `

// var testdata = `when request.url.hostname:"d.wanfangdata.com.cn" AND ( request.url.path=~/\/Detail\/.+/ OR request.url.path=~/\/Detail\/.+/  )`

var testdata = `map ddf[sdfsdf]`

// var testdata = `when request.host.c :"ab" `

// A->B: A.data> B.createdata
// A->A: A
// "A->B": {

// }

// A {data=>[xxx,xxx,xx]}
// A { }
// Entity

// relactionExpression Entity -> Entity

// mapExpression Entity.proty > Entity.proty

// exctExpression relactionExpression : Entity.proty > Entity.proty
