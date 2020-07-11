package ast

import (
	"bytes"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"gitlab.iyorhe.com/dr2am/dr2am-rewrite/exacter/dsl/parser"
)

type AST struct {
}

// type Node struct {
// }
type AstBuilder struct {
	parser.BasedslVisitor
	Root AST
}

func (v *AstBuilder) error(err error) {
	v.errorf("%s", err)
}

// errorf formats the error and terminates processing.
func (p *AstBuilder) errorf(format string, args ...interface{}) {
	format = fmt.Sprintf("parser: %s", format)
	panic(fmt.Errorf(format, args...))
}
func NewAstBuilder() *AstBuilder {
	return &AstBuilder{}
}
func (v *AstBuilder) Visit(tree antlr.ParseTree) interface{} {
	token := tree.(*parser.ProgramContext).GetStart()
	// fmt.Println(token)
	pos, line, char := GetPositionByToken(token)
	prog := NewProgramNode(SetPosition(pos, line, char))

	for _, children := range tree.GetChildren() {
		switch ctx := children.(type) {
		case *parser.StatementContext:
			v.VisitStatement(ctx, prog)
		default:
			// fmt.Printf("%T\n", ctx)

		}
	}

	return prog
}

func (v *AstBuilder) VisitStatement(ctx *parser.StatementContext, prog *ProgramNode) {
	// return v.VisitChildren(ctx)
	for _, state := range ctx.GetChildren() {
		switch c := state.(type) {
		// case *parser.DeclarationContext:
		// 	prog.Add(v.VisitDeclaration(c))
		// case *parser.TypeDeclarationContext:

		case *parser.ExpressionContext:
			node := v.VisitExpression(c)
			prog.Add(node)

		default:
			// fmt.Println(c)
			fmt.Printf("VisitStatement %#v %t\n", state, state)
		}
	}

	// fmt.Printf("%#v\n", ctx.GetRuleContext().GetText())

}

func (v *AstBuilder) VisitExpression(ctx *parser.ExpressionContext) Node {
	// fmt.Printf("%#v", ctx.Identifier())
	// if identifier != nil {

	// chains := ctx.AllChain()
	// //Identifier
	// id := ctx.Identifier().GetText()

	// token := ctx.GetStart()
	// pos, line, char := GetPositionByToken(token)
	// lhs := NewIdentifierNode(id, SetPosition(pos, line, char))

	// 	return v.VisitChain(lhs, chains...)
	// }

	// function := ctx.Function()
	// if function != nil {
	// 	chains := ctx.AllChain()
	// 	op := ctx.GetStart()
	// 	var ft FuncType
	// 	switch op.GetTokenType() {
	// 	case dslParserPIPE:
	// 		ft = ChainFunc
	// 	case dslParserAT:
	// 		ft = DynamicFunc
	// 	case dslParserDOT:
	// 		ft = PropertyFunc
	// 	}

	// 	lhs := v.VisitFunction(function.(*parser.FunctionContext), ft)
	// 	return v.VisitChain(lhs, chains...)
	// 	// return nil
	// }

	// primary := ctx.PrimaryExpr()
	// if primary != nil {

	// 	return v.VisitPrimaryExpr(primary.(*parser.PrimaryExprContext))
	// }

	// lambda := ctx.Lambda()
	// if lambda != nil {
	// 	return v.VisitLambda(lambda.(*parser.LambdaContext))
	// }
	// reference := ctx.Reference()

	// if reference != nil {
	// 	return v.VisitReference(reference)
	// }
	condition := ctx.Condition()
	if condition != nil {

		return v.VisitCondition(condition.(*parser.ConditionContext))
	}

	mapExpr := ctx.MapExpr()
	if mapExpr != nil {

		return v.VisitMapExpr(mapExpr.(*parser.MapExprContext))
	}
	return nil
}

func (v *AstBuilder) VisitReference(ctx antlr.TerminalNode) Node {
	// fmt.Println(ctx.GetChildren())
	return nil
}

func (v *AstBuilder) VisitMapExpr(ctx *parser.MapExprContext) Node {
	mapPrimary := ctx.Mapprimay(0)
	reference := mapPrimary.(*parser.MapprimayContext).Reference()
	return NewMapNode(NewReference(reference.GetText()))
}
func (v *AstBuilder) VisitCondition(ctx *parser.ConditionContext) Node {
	token := ctx.GetStart()
	pos, line, char := GetPositionByToken(token)
	primaryExpr := ctx.PrimaryExpr()
	if primaryExpr != nil {
		return NewConditionNode(v.VisitPrimaryExpr(primaryExpr.(*parser.PrimaryExprContext)), SetPosition(pos, line, char))

	}
	return nil
}

func (v *AstBuilder) VisitPrimary(ctx *parser.PrimaryContext) Node {
	token := ctx.GetStart()
	pos, line, char := GetPositionByToken(token)
	//Identifier
	id := token.GetText()
	if id == "(" {
		trees := ctx.GetChild(1).GetChildren()
		lhs := v.VisitPrimary(trees[0].(*parser.PrimaryContext))
		node := v.precedence(lhs, 0, trees[1:])
		node.(*BinaryNode).Parens = true
		return node
	}

	if ctx.Integer() != nil || ctx.RealNumber() != nil {
		node, err := NewNumberNode(id, SetPosition(pos, line, char))
		if err != nil {
			v.error(err)
		}
		return node
	}
	// if ref := ctx.Reference(); ref != nil {

	// 	return NewReference(id[1:], SetPosition(pos, line, char))
	// }
	if ctx.StringLiteral() != nil {
		return NewStringNode(id, SetPosition(pos, line, char))
	}

	// if ctx.DurationLiteral() != nil {
	// 	n, err := NewDurationNode(id, SetPosition(pos, line, char))
	// 	if err != nil {
	// 		v.error(err)
	// 	}
	// 	return n
	// }
	// if ctx.BooleanLiteral() != nil {
	// 	n, err := NewBoolNode(id, SetPosition(pos, line, char))
	// 	if err != nil {
	// 		v.error(err)
	// 	}
	// 	return n
	// }
	if ctx.Identifier() != nil {
		return NewIdentifierNode(id, SetPosition(pos, line, char))
	}
	if ctx.Reference() != nil {
		return NewReference(id, SetPosition(pos, line, char))

	}
	if ctx.RegexLiteral() != nil {
		rn, err := NewRegexNode(id, SetPosition(pos, line, char))
		if err != nil {
			v.error(err)
		}
		return rn

	}

	// if ctx.GetLeft() != nil {
	// 	fmt.Println("LEFT", ctx.GetLeft())
	// 	fmt.Println("LEFT", ctx.GetLeft().GetChildren())
	// 	fmt.Println("RIGHT", ctx.GetRight())
	// 	fmt.Println("RIGHT", ctx.GetRight())
	// 	fmt.Println("Op", ctx.GetOp().GetText())

	// 	op := stringTokenMap[ctx.GetOp().GetText()]
	// 	l := v.VisitPrimary(ctx.GetLeft().(*parser.PrimaryContext))
	// 	r := v.VisitPrimary(ctx.GetRight().(*parser.PrimaryContext))
	// 	return NewBinaryNode(l, r, false, SetOperator(op))
	// }

	// if ctx.LPAREN() != nil {
	// 	fmt.Println("GGG", ctx.GetChildren())
	// 	fmt.Printf("%#v", ctx)
	// 	lhs := v.VisitPrimary(ctx.GetChildren()[1].(*parser.PrimaryContext))

	// 	return v.precedence(lhs, 0, nil)
	// }

	return nil
}
func (v *AstBuilder) VisitPrimaryExprNew(ctx *parser.PrimaryExprContext) Node {
	trees := ctx.GetChildren()
	// fmt.Println("GE", trees)
	// fmt.Printf("%v", trees[4])
	// fmt.Printf("GG%v", trees[4].GetChildren())
	// fmt.Println(trees[0].GetChildren())
	lhs := v.VisitPrimary(trees[0].(*parser.PrimaryContext))
	// if lhs == nil {
	// 	return v.precedence(lhs, 0, trees[1+1:])
	// }
	return v.precedence(lhs, 0, trees[1:])
}
func (v *AstBuilder) VisitPrimaryExpr(ctx *parser.PrimaryExprContext) Node {
	trees := ctx.GetChildren()
	// fmt.Println("G", trees)
	// fmt.Printf("%#v\n", trees[0])
	// fmt.Println("GG", trees[8].GetChildren())
	// fmt.Println(trees[0].GetChildren())
	// fmt.Println(ctx.AllPrimary())
	// fmt.Println(ctx.AllPrimaryExpr()[0].GetChildren())
	// fmt.Println(trees[0].(*parser.PrimaryContext))
	lhs := v.VisitPrimary(trees[0].(*parser.PrimaryContext))
	// if lhs == nil {
	// 	return v.precedence(lhs, 0, trees[1+1:])
	// }
	return v.precedence(lhs, 0, trees[1:])
}

// Position return pos,line,char
func GetPositionByToken(token antlr.Token) (int, int, int) {
	return token.GetColumn(), token.GetLine(), token.GetStop() - token.GetStart() + 1
}
func Parse(script string) (node Node, err error) {
	//TODO:recover
	defer func() {
		if rerr := recover(); rerr != nil {
			err = rerr.(error)
		}
	}()
	input := antlr.NewInputStream(script)
	lexer := parser.NewdslLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewdslParser(stream)
	p.BuildParseTrees = true
	tree := p.Program()
	visitor := NewAstBuilder()
	node = visitor.Visit(tree).(Node)
	return node, nil
}

// Operator Precedence parsing
var precedence = [...]int{
	TokenOr:            0,
	TokenAnd:           1,
	TokenEqual:         2,
	TokenNotEqual:      2,
	TokenMatch:         2,
	TokenRegexEqual:    2,
	TokenRegexNotEqual: 2,
	TokenGreater:       3,
	TokenGreaterEqual:  3,
	TokenLess:          3,
	TokenLessEqual:     3,
	TokenPlus:          4,
	TokenMinus:         4,
	TokenMult:          5,
	TokenDiv:           5,
	TokenMod:           5,
}

// parse the expression considering operator precedence.
// https://en.wikipedia.org/wiki/Operator-precedence_parser#Pseudo-code
func (v *AstBuilder) precedence(lhs Node, minP int, trees []antlr.Tree) Node {
	if len(trees) == 0 {
		return lhs
	}
	var rhs Node
	var node antlr.Tree
	node, trees = pop(trees)

	op := stringTokenMap[node.(*antlr.TerminalNodeImpl).GetText()]
	if IsExprOperator(op) {
		rhs = v.VisitPrimary(trees[0].(*parser.PrimaryContext))
		_, trees = pop(trees)
		if len(trees) > 0 {
			look := stringTokenMap[trees[0].(*antlr.TerminalNodeImpl).GetText()]

			if IsExprOperator(look) {
				if precedence[look] > precedence[op] {
					rhs = v.precedence(rhs, precedence[look], trees)
					trees = nil
				}
			}
		}
		options := []Option{SetOperator(op)}
		// if minP == 1 {
		// 	options = append(options, SetParens(true))
		// }
		lhs = NewBinaryNode(lhs, rhs, false, options...)
	}

	if len(trees) == 0 {
		return lhs
	}
	return v.precedence(lhs, 0, trees)

}

func pop(t []antlr.Tree) (antlr.Tree, []antlr.Tree) {
	p := t[0]
	t = t[1:]
	return p, t
}

func writeIndent(buf *bytes.Buffer, indent string, onNewLine bool) {
	if onNewLine {
		buf.WriteString(indent)
	}
}

func IsExprOperator(t TokenType) bool {
	return t >= TokenNot && t <= TokenOr
	// return true

}
