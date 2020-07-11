// Code generated from dsl.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // dsl

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 43, 169,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 7, 2, 37, 10, 2, 12, 2, 14, 2, 40, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5,
	6, 58, 10, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 6, 8, 65, 10, 8, 13, 8, 14,
	8, 66, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 76, 10, 10,
	12, 10, 14, 10, 79, 11, 10, 3, 10, 3, 10, 3, 10, 7, 10, 84, 10, 10, 12,
	10, 14, 10, 87, 11, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 94, 10,
	10, 12, 10, 14, 10, 97, 11, 10, 5, 10, 99, 10, 10, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 7, 12, 109, 10, 12, 12, 12, 14, 12,
	112, 11, 12, 3, 12, 5, 12, 115, 10, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3,
	14, 7, 14, 122, 10, 14, 12, 14, 14, 14, 125, 11, 14, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 5, 15, 141, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 5, 16, 153, 10, 16, 3, 17, 3, 17, 5, 17, 157, 10,
	17, 3, 17, 7, 17, 160, 10, 17, 12, 17, 14, 17, 163, 11, 17, 3, 17, 3, 17,
	5, 17, 167, 10, 17, 3, 17, 2, 2, 18, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20,
	22, 24, 26, 28, 30, 32, 2, 2, 2, 181, 2, 34, 3, 2, 2, 2, 4, 43, 3, 2, 2,
	2, 6, 45, 3, 2, 2, 2, 8, 49, 3, 2, 2, 2, 10, 57, 3, 2, 2, 2, 12, 59, 3,
	2, 2, 2, 14, 62, 3, 2, 2, 2, 16, 68, 3, 2, 2, 2, 18, 98, 3, 2, 2, 2, 20,
	100, 3, 2, 2, 2, 22, 110, 3, 2, 2, 2, 24, 116, 3, 2, 2, 2, 26, 118, 3,
	2, 2, 2, 28, 140, 3, 2, 2, 2, 30, 152, 3, 2, 2, 2, 32, 161, 3, 2, 2, 2,
	34, 38, 5, 4, 3, 2, 35, 37, 5, 4, 3, 2, 36, 35, 3, 2, 2, 2, 37, 40, 3,
	2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39, 3, 2, 2, 2, 39, 41, 3, 2, 2, 2, 40,
	38, 3, 2, 2, 2, 41, 42, 7, 2, 2, 3, 42, 3, 3, 2, 2, 2, 43, 44, 5, 10, 6,
	2, 44, 5, 3, 2, 2, 2, 45, 46, 7, 28, 2, 2, 46, 47, 7, 36, 2, 2, 47, 48,
	7, 36, 2, 2, 48, 7, 3, 2, 2, 2, 49, 50, 7, 28, 2, 2, 50, 51, 7, 36, 2,
	2, 51, 52, 7, 17, 2, 2, 52, 53, 5, 10, 6, 2, 53, 9, 3, 2, 2, 2, 54, 58,
	5, 12, 7, 2, 55, 58, 5, 14, 8, 2, 56, 58, 7, 39, 2, 2, 57, 54, 3, 2, 2,
	2, 57, 55, 3, 2, 2, 2, 57, 56, 3, 2, 2, 2, 58, 11, 3, 2, 2, 2, 59, 60,
	7, 3, 2, 2, 60, 61, 5, 26, 14, 2, 61, 13, 3, 2, 2, 2, 62, 64, 7, 4, 2,
	2, 63, 65, 5, 16, 9, 2, 64, 63, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 64,
	3, 2, 2, 2, 66, 67, 3, 2, 2, 2, 67, 15, 3, 2, 2, 2, 68, 69, 7, 35, 2, 2,
	69, 17, 3, 2, 2, 2, 70, 71, 7, 31, 2, 2, 71, 99, 5, 20, 11, 2, 72, 73,
	7, 29, 2, 2, 73, 77, 5, 20, 11, 2, 74, 76, 5, 18, 10, 2, 75, 74, 3, 2,
	2, 2, 76, 79, 3, 2, 2, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 99,
	3, 2, 2, 2, 79, 77, 3, 2, 2, 2, 80, 81, 7, 30, 2, 2, 81, 85, 5, 20, 11,
	2, 82, 84, 5, 18, 10, 2, 83, 82, 3, 2, 2, 2, 84, 87, 3, 2, 2, 2, 85, 83,
	3, 2, 2, 2, 85, 86, 3, 2, 2, 2, 86, 99, 3, 2, 2, 2, 87, 85, 3, 2, 2, 2,
	88, 89, 7, 30, 2, 2, 89, 99, 7, 36, 2, 2, 90, 91, 7, 30, 2, 2, 91, 95,
	7, 36, 2, 2, 92, 94, 5, 18, 10, 2, 93, 92, 3, 2, 2, 2, 94, 97, 3, 2, 2,
	2, 95, 93, 3, 2, 2, 2, 95, 96, 3, 2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 95,
	3, 2, 2, 2, 98, 70, 3, 2, 2, 2, 98, 72, 3, 2, 2, 2, 98, 80, 3, 2, 2, 2,
	98, 88, 3, 2, 2, 2, 98, 90, 3, 2, 2, 2, 99, 19, 3, 2, 2, 2, 100, 101, 7,
	36, 2, 2, 101, 102, 7, 7, 2, 2, 102, 103, 5, 22, 12, 2, 103, 104, 7, 8,
	2, 2, 104, 21, 3, 2, 2, 2, 105, 106, 5, 24, 13, 2, 106, 107, 7, 13, 2,
	2, 107, 109, 3, 2, 2, 2, 108, 105, 3, 2, 2, 2, 109, 112, 3, 2, 2, 2, 110,
	108, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 114, 3, 2, 2, 2, 112, 110,
	3, 2, 2, 2, 113, 115, 5, 24, 13, 2, 114, 113, 3, 2, 2, 2, 114, 115, 3,
	2, 2, 2, 115, 23, 3, 2, 2, 2, 116, 117, 5, 10, 6, 2, 117, 25, 3, 2, 2,
	2, 118, 123, 5, 28, 15, 2, 119, 120, 7, 26, 2, 2, 120, 122, 5, 28, 15,
	2, 121, 119, 3, 2, 2, 2, 122, 125, 3, 2, 2, 2, 123, 121, 3, 2, 2, 2, 123,
	124, 3, 2, 2, 2, 124, 27, 3, 2, 2, 2, 125, 123, 3, 2, 2, 2, 126, 127, 7,
	7, 2, 2, 127, 128, 5, 26, 14, 2, 128, 129, 7, 8, 2, 2, 129, 141, 3, 2,
	2, 2, 130, 141, 7, 33, 2, 2, 131, 141, 7, 26, 2, 2, 132, 141, 7, 32, 2,
	2, 133, 141, 7, 40, 2, 2, 134, 141, 7, 35, 2, 2, 135, 141, 7, 42, 2, 2,
	136, 141, 7, 34, 2, 2, 137, 141, 7, 36, 2, 2, 138, 141, 7, 41, 2, 2, 139,
	141, 7, 43, 2, 2, 140, 126, 3, 2, 2, 2, 140, 130, 3, 2, 2, 2, 140, 131,
	3, 2, 2, 2, 140, 132, 3, 2, 2, 2, 140, 133, 3, 2, 2, 2, 140, 134, 3, 2,
	2, 2, 140, 135, 3, 2, 2, 2, 140, 136, 3, 2, 2, 2, 140, 137, 3, 2, 2, 2,
	140, 138, 3, 2, 2, 2, 140, 139, 3, 2, 2, 2, 141, 29, 3, 2, 2, 2, 142, 143,
	7, 7, 2, 2, 143, 144, 5, 32, 17, 2, 144, 145, 7, 8, 2, 2, 145, 146, 7,
	5, 2, 2, 146, 147, 7, 11, 2, 2, 147, 148, 5, 26, 14, 2, 148, 149, 7, 12,
	2, 2, 149, 153, 3, 2, 2, 2, 150, 151, 7, 6, 2, 2, 151, 153, 5, 26, 14,
	2, 152, 142, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 153, 31, 3, 2, 2, 2, 154,
	157, 7, 36, 2, 2, 155, 157, 5, 10, 6, 2, 156, 154, 3, 2, 2, 2, 156, 155,
	3, 2, 2, 2, 157, 158, 3, 2, 2, 2, 158, 160, 7, 13, 2, 2, 159, 156, 3, 2,
	2, 2, 160, 163, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2,
	162, 166, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 164, 167, 7, 36, 2, 2, 165,
	167, 5, 10, 6, 2, 166, 164, 3, 2, 2, 2, 166, 165, 3, 2, 2, 2, 166, 167,
	3, 2, 2, 2, 167, 33, 3, 2, 2, 2, 17, 38, 57, 66, 77, 85, 95, 98, 110, 114,
	123, 140, 152, 156, 161, 166,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'when'", "'map'", "'=>'", "'lambda:'", "'('", "')'", "'['", "']'",
	"'{'", "'}'", "','", "';'", "'&&'", "'||'", "'='", "'if'", "'else'", "'while'",
	"'break'", "'read'", "'write'", "'int'", "'real'", "", "", "'var'", "'|'",
	"'.'", "'@'",
}
var symbolicNames = []string{
	"", "", "", "", "", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE",
	"COMMA", "SEMICOLON", "ANDAND", "OROR", "EQ", "IF", "ELSE", "WHILE", "BREAK",
	"READ", "WRITE", "INT", "REAL", "Operator_lit", "Duration_unit", "VAR",
	"PIPE", "DOT", "AT", "Integer", "RealNumber", "BooleanLiteral", "Reference",
	"Identifier", "WS", "Comment", "LineComment", "StringLiteral", "RegexLiteral",
	"DurationLiteral", "RoundLiteral",
}

var ruleNames = []string{
	"program", "statement", "typeDeclaration", "declaration", "expression",
	"condition", "mapExpr", "mapprimay", "chain", "function", "parameters",
	"parameter", "primaryExpr", "primary", "lambda", "lambdaparameters",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type dslParser struct {
	*antlr.BaseParser
}

func NewdslParser(input antlr.TokenStream) *dslParser {
	this := new(dslParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "dsl.g4"

	return this
}

// dslParser tokens.
const (
	dslParserEOF             = antlr.TokenEOF
	dslParserT__0            = 1
	dslParserT__1            = 2
	dslParserT__2            = 3
	dslParserT__3            = 4
	dslParserLPAREN          = 5
	dslParserRPAREN          = 6
	dslParserLBRACK          = 7
	dslParserRBRACK          = 8
	dslParserLBRACE          = 9
	dslParserRBRACE          = 10
	dslParserCOMMA           = 11
	dslParserSEMICOLON       = 12
	dslParserANDAND          = 13
	dslParserOROR            = 14
	dslParserEQ              = 15
	dslParserIF              = 16
	dslParserELSE            = 17
	dslParserWHILE           = 18
	dslParserBREAK           = 19
	dslParserREAD            = 20
	dslParserWRITE           = 21
	dslParserINT             = 22
	dslParserREAL            = 23
	dslParserOperator_lit    = 24
	dslParserDuration_unit   = 25
	dslParserVAR             = 26
	dslParserPIPE            = 27
	dslParserDOT             = 28
	dslParserAT              = 29
	dslParserInteger         = 30
	dslParserRealNumber      = 31
	dslParserBooleanLiteral  = 32
	dslParserReference       = 33
	dslParserIdentifier      = 34
	dslParserWS              = 35
	dslParserComment         = 36
	dslParserLineComment     = 37
	dslParserStringLiteral   = 38
	dslParserRegexLiteral    = 39
	dslParserDurationLiteral = 40
	dslParserRoundLiteral    = 41
)

// dslParser rules.
const (
	dslParserRULE_program          = 0
	dslParserRULE_statement        = 1
	dslParserRULE_typeDeclaration  = 2
	dslParserRULE_declaration      = 3
	dslParserRULE_expression       = 4
	dslParserRULE_condition        = 5
	dslParserRULE_mapExpr          = 6
	dslParserRULE_mapprimay        = 7
	dslParserRULE_chain            = 8
	dslParserRULE_function         = 9
	dslParserRULE_parameters       = 10
	dslParserRULE_parameter        = 11
	dslParserRULE_primaryExpr      = 12
	dslParserRULE_primary          = 13
	dslParserRULE_lambda           = 14
	dslParserRULE_lambdaparameters = 15
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(dslParserEOF, 0)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, dslParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Statement()
	}
	p.SetState(36)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == dslParserT__0 || _la == dslParserT__1 || _la == dslParserLineComment {
		{
			p.SetState(33)
			p.Statement()
		}

		p.SetState(38)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(39)
		p.Match(dslParserEOF)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, dslParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Expression()
	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(dslParserVAR, 0)
}

func (s *TypeDeclarationContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(dslParserIdentifier)
}

func (s *TypeDeclarationContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(dslParserIdentifier, i)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, dslParserRULE_typeDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(dslParserVAR)
	}
	{
		p.SetState(44)
		p.Match(dslParserIdentifier)
	}
	{
		p.SetState(45)
		p.Match(dslParserIdentifier)
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) VAR() antlr.TerminalNode {
	return s.GetToken(dslParserVAR, 0)
}

func (s *DeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(dslParserIdentifier, 0)
}

func (s *DeclarationContext) EQ() antlr.TerminalNode {
	return s.GetToken(dslParserEQ, 0)
}

func (s *DeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, dslParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(dslParserVAR)
	}
	{
		p.SetState(48)
		p.Match(dslParserIdentifier)
	}
	{
		p.SetState(49)
		p.Match(dslParserEQ)
	}
	{
		p.SetState(50)
		p.Expression()
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ExpressionContext) MapExpr() IMapExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapExprContext)
}

func (s *ExpressionContext) LineComment() antlr.TerminalNode {
	return s.GetToken(dslParserLineComment, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, dslParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case dslParserT__0:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Condition()
		}

	case dslParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(53)
			p.MapExpr()
		}

	case dslParserLineComment:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(54)
			p.Match(dslParserLineComment)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, dslParserRULE_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(dslParserT__0)
	}
	{
		p.SetState(58)
		p.PrimaryExpr()
	}

	return localctx
}

// IMapExprContext is an interface to support dynamic dispatch.
type IMapExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapExprContext differentiates from other interfaces.
	IsMapExprContext()
}

type MapExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapExprContext() *MapExprContext {
	var p = new(MapExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_mapExpr
	return p
}

func (*MapExprContext) IsMapExprContext() {}

func NewMapExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapExprContext {
	var p = new(MapExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_mapExpr

	return p
}

func (s *MapExprContext) GetParser() antlr.Parser { return s.parser }

func (s *MapExprContext) AllMapprimay() []IMapprimayContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMapprimayContext)(nil)).Elem())
	var tst = make([]IMapprimayContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMapprimayContext)
		}
	}

	return tst
}

func (s *MapExprContext) Mapprimay(i int) IMapprimayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapprimayContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMapprimayContext)
}

func (s *MapExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterMapExpr(s)
	}
}

func (s *MapExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitMapExpr(s)
	}
}

func (s *MapExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitMapExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) MapExpr() (localctx IMapExprContext) {
	localctx = NewMapExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, dslParserRULE_mapExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(dslParserT__1)
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == dslParserReference {
		{
			p.SetState(61)
			p.Mapprimay()
		}

		p.SetState(64)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMapprimayContext is an interface to support dynamic dispatch.
type IMapprimayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapprimayContext differentiates from other interfaces.
	IsMapprimayContext()
}

type MapprimayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapprimayContext() *MapprimayContext {
	var p = new(MapprimayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_mapprimay
	return p
}

func (*MapprimayContext) IsMapprimayContext() {}

func NewMapprimayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapprimayContext {
	var p = new(MapprimayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_mapprimay

	return p
}

func (s *MapprimayContext) GetParser() antlr.Parser { return s.parser }

func (s *MapprimayContext) Reference() antlr.TerminalNode {
	return s.GetToken(dslParserReference, 0)
}

func (s *MapprimayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapprimayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapprimayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterMapprimay(s)
	}
}

func (s *MapprimayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitMapprimay(s)
	}
}

func (s *MapprimayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitMapprimay(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Mapprimay() (localctx IMapprimayContext) {
	localctx = NewMapprimayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, dslParserRULE_mapprimay)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(66)
		p.Match(dslParserReference)
	}

	return localctx
}

// IChainContext is an interface to support dynamic dispatch.
type IChainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsChainContext differentiates from other interfaces.
	IsChainContext()
}

type ChainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyChainContext() *ChainContext {
	var p = new(ChainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_chain
	return p
}

func (*ChainContext) IsChainContext() {}

func NewChainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ChainContext {
	var p = new(ChainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_chain

	return p
}

func (s *ChainContext) GetParser() antlr.Parser { return s.parser }

func (s *ChainContext) AT() antlr.TerminalNode {
	return s.GetToken(dslParserAT, 0)
}

func (s *ChainContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *ChainContext) PIPE() antlr.TerminalNode {
	return s.GetToken(dslParserPIPE, 0)
}

func (s *ChainContext) AllChain() []IChainContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IChainContext)(nil)).Elem())
	var tst = make([]IChainContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IChainContext)
		}
	}

	return tst
}

func (s *ChainContext) Chain(i int) IChainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IChainContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IChainContext)
}

func (s *ChainContext) DOT() antlr.TerminalNode {
	return s.GetToken(dslParserDOT, 0)
}

func (s *ChainContext) Identifier() antlr.TerminalNode {
	return s.GetToken(dslParserIdentifier, 0)
}

func (s *ChainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ChainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ChainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterChain(s)
	}
}

func (s *ChainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitChain(s)
	}
}

func (s *ChainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitChain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Chain() (localctx IChainContext) {
	localctx = NewChainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, dslParserRULE_chain)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(68)
			p.Match(dslParserAT)
		}
		{
			p.SetState(69)
			p.Function()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(70)
			p.Match(dslParserPIPE)
		}
		{
			p.SetState(71)
			p.Function()
		}
		p.SetState(75)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(72)
					p.Chain()
				}

			}
			p.SetState(77)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(78)
			p.Match(dslParserDOT)
		}
		{
			p.SetState(79)
			p.Function()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(80)
					p.Chain()
				}

			}
			p.SetState(85)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(86)
			p.Match(dslParserDOT)
		}
		{
			p.SetState(87)
			p.Match(dslParserIdentifier)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(88)
			p.Match(dslParserDOT)
		}
		{
			p.SetState(89)
			p.Match(dslParserIdentifier)
		}
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(90)
					p.Chain()
				}

			}
			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
		}

	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(dslParserIdentifier, 0)
}

func (s *FunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(dslParserLPAREN, 0)
}

func (s *FunctionContext) Parameters() IParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(dslParserRPAREN, 0)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (s *FunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, dslParserRULE_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(98)
		p.Match(dslParserIdentifier)
	}
	{
		p.SetState(99)
		p.Match(dslParserLPAREN)
	}
	{
		p.SetState(100)
		p.Parameters()
	}
	{
		p.SetState(101)
		p.Match(dslParserRPAREN)
	}

	return localctx
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_parameters
	return p
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllParameter() []IParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IParameterContext)(nil)).Elem())
	var tst = make([]IParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IParameterContext)
		}
	}

	return tst
}

func (s *ParametersContext) Parameter(i int) IParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, dslParserRULE_parameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(103)
				p.Parameter()
			}
			{
				p.SetState(104)
				p.Match(dslParserCOMMA)
			}

		}
		p.SetState(110)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == dslParserT__0 || _la == dslParserT__1 || _la == dslParserLineComment {
		{
			p.SetState(111)
			p.Parameter()
		}

	}

	return localctx
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_parameter
	return p
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (s *ParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, dslParserRULE_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Expression()
	}

	return localctx
}

// IPrimaryExprContext is an interface to support dynamic dispatch.
type IPrimaryExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryExprContext differentiates from other interfaces.
	IsPrimaryExprContext()
}

type PrimaryExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExprContext() *PrimaryExprContext {
	var p = new(PrimaryExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_primaryExpr
	return p
}

func (*PrimaryExprContext) IsPrimaryExprContext() {}

func NewPrimaryExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExprContext {
	var p = new(PrimaryExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_primaryExpr

	return p
}

func (s *PrimaryExprContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExprContext) AllPrimary() []IPrimaryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPrimaryContext)(nil)).Elem())
	var tst = make([]IPrimaryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPrimaryContext)
		}
	}

	return tst
}

func (s *PrimaryExprContext) Primary(i int) IPrimaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPrimaryContext)
}

func (s *PrimaryExprContext) AllOperator_lit() []antlr.TerminalNode {
	return s.GetTokens(dslParserOperator_lit)
}

func (s *PrimaryExprContext) Operator_lit(i int) antlr.TerminalNode {
	return s.GetToken(dslParserOperator_lit, i)
}

func (s *PrimaryExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitPrimaryExpr(s)
	}
}

func (s *PrimaryExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitPrimaryExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) PrimaryExpr() (localctx IPrimaryExprContext) {
	localctx = NewPrimaryExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, dslParserRULE_primaryExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Primary()
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == dslParserOperator_lit {
		{
			p.SetState(117)
			p.Match(dslParserOperator_lit)
		}
		{
			p.SetState(118)
			p.Primary()
		}

		p.SetState(123)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *PrimaryContext) RealNumber() antlr.TerminalNode {
	return s.GetToken(dslParserRealNumber, 0)
}

func (s *PrimaryContext) Operator_lit() antlr.TerminalNode {
	return s.GetToken(dslParserOperator_lit, 0)
}

func (s *PrimaryContext) Integer() antlr.TerminalNode {
	return s.GetToken(dslParserInteger, 0)
}

func (s *PrimaryContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(dslParserStringLiteral, 0)
}

func (s *PrimaryContext) Reference() antlr.TerminalNode {
	return s.GetToken(dslParserReference, 0)
}

func (s *PrimaryContext) DurationLiteral() antlr.TerminalNode {
	return s.GetToken(dslParserDurationLiteral, 0)
}

func (s *PrimaryContext) BooleanLiteral() antlr.TerminalNode {
	return s.GetToken(dslParserBooleanLiteral, 0)
}

func (s *PrimaryContext) Identifier() antlr.TerminalNode {
	return s.GetToken(dslParserIdentifier, 0)
}

func (s *PrimaryContext) RegexLiteral() antlr.TerminalNode {
	return s.GetToken(dslParserRegexLiteral, 0)
}

func (s *PrimaryContext) RoundLiteral() antlr.TerminalNode {
	return s.GetToken(dslParserRoundLiteral, 0)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, dslParserRULE_primary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(138)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case dslParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(124)
			p.Match(dslParserLPAREN)
		}
		{
			p.SetState(125)
			p.PrimaryExpr()
		}
		{
			p.SetState(126)
			p.Match(dslParserRPAREN)
		}

	case dslParserRealNumber:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(128)
			p.Match(dslParserRealNumber)
		}

	case dslParserOperator_lit:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(129)
			p.Match(dslParserOperator_lit)
		}

	case dslParserInteger:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(130)
			p.Match(dslParserInteger)
		}

	case dslParserStringLiteral:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(131)
			p.Match(dslParserStringLiteral)
		}

	case dslParserReference:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(132)
			p.Match(dslParserReference)
		}

	case dslParserDurationLiteral:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(133)
			p.Match(dslParserDurationLiteral)
		}

	case dslParserBooleanLiteral:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(134)
			p.Match(dslParserBooleanLiteral)
		}

	case dslParserIdentifier:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(135)
			p.Match(dslParserIdentifier)
		}

	case dslParserRegexLiteral:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(136)
			p.Match(dslParserRegexLiteral)
		}

	case dslParserRoundLiteral:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(137)
			p.Match(dslParserRoundLiteral)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILambdaContext is an interface to support dynamic dispatch.
type ILambdaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambdaContext differentiates from other interfaces.
	IsLambdaContext()
}

type LambdaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaContext() *LambdaContext {
	var p = new(LambdaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_lambda
	return p
}

func (*LambdaContext) IsLambdaContext() {}

func NewLambdaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaContext {
	var p = new(LambdaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_lambda

	return p
}

func (s *LambdaContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(dslParserLPAREN, 0)
}

func (s *LambdaContext) Lambdaparameters() ILambdaparametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambdaparametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILambdaparametersContext)
}

func (s *LambdaContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(dslParserRPAREN, 0)
}

func (s *LambdaContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(dslParserLBRACE, 0)
}

func (s *LambdaContext) PrimaryExpr() IPrimaryExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimaryExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimaryExprContext)
}

func (s *LambdaContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(dslParserRBRACE, 0)
}

func (s *LambdaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterLambda(s)
	}
}

func (s *LambdaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitLambda(s)
	}
}

func (s *LambdaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitLambda(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Lambda() (localctx ILambdaContext) {
	localctx = NewLambdaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, dslParserRULE_lambda)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(150)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case dslParserLPAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Match(dslParserLPAREN)
		}
		{
			p.SetState(141)
			p.Lambdaparameters()
		}
		{
			p.SetState(142)
			p.Match(dslParserRPAREN)
		}
		{
			p.SetState(143)
			p.Match(dslParserT__2)
		}
		{
			p.SetState(144)
			p.Match(dslParserLBRACE)
		}
		{
			p.SetState(145)
			p.PrimaryExpr()
		}
		{
			p.SetState(146)
			p.Match(dslParserRBRACE)
		}

	case dslParserT__3:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Match(dslParserT__3)
		}
		{
			p.SetState(149)
			p.PrimaryExpr()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILambdaparametersContext is an interface to support dynamic dispatch.
type ILambdaparametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLambdaparametersContext differentiates from other interfaces.
	IsLambdaparametersContext()
}

type LambdaparametersContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLambdaparametersContext() *LambdaparametersContext {
	var p = new(LambdaparametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = dslParserRULE_lambdaparameters
	return p
}

func (*LambdaparametersContext) IsLambdaparametersContext() {}

func NewLambdaparametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaparametersContext {
	var p = new(LambdaparametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = dslParserRULE_lambdaparameters

	return p
}

func (s *LambdaparametersContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaparametersContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(dslParserIdentifier)
}

func (s *LambdaparametersContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(dslParserIdentifier, i)
}

func (s *LambdaparametersContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LambdaparametersContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LambdaparametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaparametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaparametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.EnterLambdaparameters(s)
	}
}

func (s *LambdaparametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(dslListener); ok {
		listenerT.ExitLambdaparameters(s)
	}
}

func (s *LambdaparametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case dslVisitor:
		return t.VisitLambdaparameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *dslParser) Lambdaparameters() (localctx ILambdaparametersContext) {
	localctx = NewLambdaparametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, dslParserRULE_lambdaparameters)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(154)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case dslParserIdentifier:
				{
					p.SetState(152)
					p.Match(dslParserIdentifier)
				}

			case dslParserT__0, dslParserT__1, dslParserLineComment:
				{
					p.SetState(153)
					p.Expression()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}
			{
				p.SetState(156)
				p.Match(dslParserCOMMA)
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext())
	}
	p.SetState(164)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case dslParserIdentifier:
		{
			p.SetState(162)
			p.Match(dslParserIdentifier)
		}

	case dslParserT__0, dslParserT__1, dslParserLineComment:
		{
			p.SetState(163)
			p.Expression()
		}

	case dslParserRPAREN:

	default:
	}

	return localctx
}
