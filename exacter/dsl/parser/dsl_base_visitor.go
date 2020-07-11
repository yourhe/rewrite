// Code generated from dsl.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // dsl

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BasedslVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasedslVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitTypeDeclaration(ctx *TypeDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitCondition(ctx *ConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitMapExpr(ctx *MapExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitMapprimay(ctx *MapprimayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitChain(ctx *ChainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitFunction(ctx *FunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitPrimaryExpr(ctx *PrimaryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitPrimary(ctx *PrimaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasedslVisitor) VisitLambdaparameters(ctx *LambdaparametersContext) interface{} {
	return v.VisitChildren(ctx)
}
