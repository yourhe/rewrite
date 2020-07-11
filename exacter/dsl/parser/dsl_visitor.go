// Code generated from dsl.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // dsl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by dslParser.
type dslVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by dslParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by dslParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by dslParser#typeDeclaration.
	VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{}

	// Visit a parse tree produced by dslParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by dslParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by dslParser#condition.
	VisitCondition(ctx *ConditionContext) interface{}

	// Visit a parse tree produced by dslParser#mapExpr.
	VisitMapExpr(ctx *MapExprContext) interface{}

	// Visit a parse tree produced by dslParser#mapprimay.
	VisitMapprimay(ctx *MapprimayContext) interface{}

	// Visit a parse tree produced by dslParser#chain.
	VisitChain(ctx *ChainContext) interface{}

	// Visit a parse tree produced by dslParser#function.
	VisitFunction(ctx *FunctionContext) interface{}

	// Visit a parse tree produced by dslParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by dslParser#parameter.
	VisitParameter(ctx *ParameterContext) interface{}

	// Visit a parse tree produced by dslParser#primaryExpr.
	VisitPrimaryExpr(ctx *PrimaryExprContext) interface{}

	// Visit a parse tree produced by dslParser#primary.
	VisitPrimary(ctx *PrimaryContext) interface{}

	// Visit a parse tree produced by dslParser#lambda.
	VisitLambda(ctx *LambdaContext) interface{}

	// Visit a parse tree produced by dslParser#lambdaparameters.
	VisitLambdaparameters(ctx *LambdaparametersContext) interface{}
}
