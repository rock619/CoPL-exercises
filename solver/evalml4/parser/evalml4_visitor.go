// Code generated from EvalML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML4
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by EvalML4Parser.
type EvalML4Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EvalML4Parser#question.
	VisitQuestion(ctx *QuestionContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#eval.
	VisitEval(ctx *EvalContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#IfExpr.
	VisitIfExpr(ctx *IfExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#LetExpr.
	VisitLetExpr(ctx *LetExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#LetRecExpr.
	VisitLetRecExpr(ctx *LetRecExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#AppExpr.
	VisitAppExpr(ctx *AppExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#EmptyListExpr.
	VisitEmptyListExpr(ctx *EmptyListExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#ConsExpr.
	VisitConsExpr(ctx *ConsExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#VarExpr.
	VisitVarExpr(ctx *VarExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#BinOpExpr.
	VisitBinOpExpr(ctx *BinOpExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#FunExpr.
	VisitFunExpr(ctx *FunExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#MatchExpr.
	VisitMatchExpr(ctx *MatchExprContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#defList.
	VisitDefList(ctx *DefListContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#def.
	VisitDef(ctx *DefContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#emptyPattern.
	VisitEmptyPattern(ctx *EmptyPatternContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#consPattern.
	VisitConsPattern(ctx *ConsPatternContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#fun.
	VisitFun(ctx *FunContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#recFun.
	VisitRecFun(ctx *RecFunContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#EmptyListValue.
	VisitEmptyListValue(ctx *EmptyListValueContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#BoolValue.
	VisitBoolValue(ctx *BoolValueContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#FunValue.
	VisitFunValue(ctx *FunValueContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#RecFunValue.
	VisitRecFunValue(ctx *RecFunValueContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#IntValue.
	VisitIntValue(ctx *IntValueContext) interface{}

	// Visit a parse tree produced by EvalML4Parser#ConsValue.
	VisitConsValue(ctx *ConsValueContext) interface{}
}
