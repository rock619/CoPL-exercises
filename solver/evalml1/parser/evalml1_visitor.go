// Code generated from EvalML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML1
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by EvalML1Parser.
type EvalML1Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EvalML1Parser#question.
	VisitQuestion(ctx *QuestionContext) interface{}

	// Visit a parse tree produced by EvalML1Parser#eval.
	VisitEval(ctx *EvalContext) interface{}

	// Visit a parse tree produced by EvalML1Parser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by EvalML1Parser#IfExpr.
	VisitIfExpr(ctx *IfExprContext) interface{}

	// Visit a parse tree produced by EvalML1Parser#BinOpExpr.
	VisitBinOpExpr(ctx *BinOpExprContext) interface{}

	// Visit a parse tree produced by EvalML1Parser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by EvalML1Parser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by EvalML1Parser#NegIntExpr.
	VisitNegIntExpr(ctx *NegIntExprContext) interface{}

	// Visit a parse tree produced by EvalML1Parser#IntValue.
	VisitIntValue(ctx *IntValueContext) interface{}

	// Visit a parse tree produced by EvalML1Parser#BoolValue.
	VisitBoolValue(ctx *BoolValueContext) interface{}
}
