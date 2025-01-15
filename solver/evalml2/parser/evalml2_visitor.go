// Code generated from EvalML2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML2
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by EvalML2Parser.
type EvalML2Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EvalML2Parser#question.
	VisitQuestion(ctx *QuestionContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#eval.
	VisitEval(ctx *EvalContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#IfExpr.
	VisitIfExpr(ctx *IfExprContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#LetExpr.
	VisitLetExpr(ctx *LetExprContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#VarExpr.
	VisitVarExpr(ctx *VarExprContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#BinOpExpr.
	VisitBinOpExpr(ctx *BinOpExprContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#NegIntExpr.
	VisitNegIntExpr(ctx *NegIntExprContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#defList.
	VisitDefList(ctx *DefListContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#def.
	VisitDef(ctx *DefContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#IntValue.
	VisitIntValue(ctx *IntValueContext) interface{}

	// Visit a parse tree produced by EvalML2Parser#BoolValue.
	VisitBoolValue(ctx *BoolValueContext) interface{}
}
