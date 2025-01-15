// Code generated from EvalML3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML3
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by EvalML3Parser.
type EvalML3Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EvalML3Parser#question.
	VisitQuestion(ctx *QuestionContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#eval.
	VisitEval(ctx *EvalContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#IfExpr.
	VisitIfExpr(ctx *IfExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#LetExpr.
	VisitLetExpr(ctx *LetExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#LetRecExpr.
	VisitLetRecExpr(ctx *LetRecExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#VarExpr.
	VisitVarExpr(ctx *VarExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#AppExpr.
	VisitAppExpr(ctx *AppExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#BinOpExpr.
	VisitBinOpExpr(ctx *BinOpExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#NegIntExpr.
	VisitNegIntExpr(ctx *NegIntExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#FunExpr.
	VisitFunExpr(ctx *FunExprContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#defList.
	VisitDefList(ctx *DefListContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#def.
	VisitDef(ctx *DefContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#fun.
	VisitFun(ctx *FunContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#recFun.
	VisitRecFun(ctx *RecFunContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#IntValue.
	VisitIntValue(ctx *IntValueContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#BoolValue.
	VisitBoolValue(ctx *BoolValueContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#FunValue.
	VisitFunValue(ctx *FunValueContext) interface{}

	// Visit a parse tree produced by EvalML3Parser#RecFunValue.
	VisitRecFunValue(ctx *RecFunValueContext) interface{}
}
