// Code generated from dsl.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // dsl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasedslListener is a complete listener for a parse tree produced by dslParser.
type BasedslListener struct{}

var _ dslListener = &BasedslListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasedslListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasedslListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasedslListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasedslListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasedslListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasedslListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasedslListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasedslListener) ExitStatement(ctx *StatementContext) {}

// EnterTypeDeclaration is called when production typeDeclaration is entered.
func (s *BasedslListener) EnterTypeDeclaration(ctx *TypeDeclarationContext) {}

// ExitTypeDeclaration is called when production typeDeclaration is exited.
func (s *BasedslListener) ExitTypeDeclaration(ctx *TypeDeclarationContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BasedslListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BasedslListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasedslListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasedslListener) ExitExpression(ctx *ExpressionContext) {}

// EnterCondition is called when production condition is entered.
func (s *BasedslListener) EnterCondition(ctx *ConditionContext) {}

// ExitCondition is called when production condition is exited.
func (s *BasedslListener) ExitCondition(ctx *ConditionContext) {}

// EnterMapExpr is called when production mapExpr is entered.
func (s *BasedslListener) EnterMapExpr(ctx *MapExprContext) {}

// ExitMapExpr is called when production mapExpr is exited.
func (s *BasedslListener) ExitMapExpr(ctx *MapExprContext) {}

// EnterMapprimay is called when production mapprimay is entered.
func (s *BasedslListener) EnterMapprimay(ctx *MapprimayContext) {}

// ExitMapprimay is called when production mapprimay is exited.
func (s *BasedslListener) ExitMapprimay(ctx *MapprimayContext) {}

// EnterChain is called when production chain is entered.
func (s *BasedslListener) EnterChain(ctx *ChainContext) {}

// ExitChain is called when production chain is exited.
func (s *BasedslListener) ExitChain(ctx *ChainContext) {}

// EnterFunction is called when production function is entered.
func (s *BasedslListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BasedslListener) ExitFunction(ctx *FunctionContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BasedslListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BasedslListener) ExitParameters(ctx *ParametersContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BasedslListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasedslListener) ExitParameter(ctx *ParameterContext) {}

// EnterPrimaryExpr is called when production primaryExpr is entered.
func (s *BasedslListener) EnterPrimaryExpr(ctx *PrimaryExprContext) {}

// ExitPrimaryExpr is called when production primaryExpr is exited.
func (s *BasedslListener) ExitPrimaryExpr(ctx *PrimaryExprContext) {}

// EnterPrimary is called when production primary is entered.
func (s *BasedslListener) EnterPrimary(ctx *PrimaryContext) {}

// ExitPrimary is called when production primary is exited.
func (s *BasedslListener) ExitPrimary(ctx *PrimaryContext) {}

// EnterLambda is called when production lambda is entered.
func (s *BasedslListener) EnterLambda(ctx *LambdaContext) {}

// ExitLambda is called when production lambda is exited.
func (s *BasedslListener) ExitLambda(ctx *LambdaContext) {}

// EnterLambdaparameters is called when production lambdaparameters is entered.
func (s *BasedslListener) EnterLambdaparameters(ctx *LambdaparametersContext) {}

// ExitLambdaparameters is called when production lambdaparameters is exited.
func (s *BasedslListener) ExitLambdaparameters(ctx *LambdaparametersContext) {}
