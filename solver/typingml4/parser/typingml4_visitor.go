// Code generated from TypingML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TypingML4
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TypingML4Parser.
type TypingML4Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TypingML4Parser#judgement.
	VisitJudgement(ctx *JudgementContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#BoolType.
	VisitBoolType(ctx *BoolTypeContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#ListType.
	VisitListType(ctx *ListTypeContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#FunType.
	VisitFunType(ctx *FunTypeContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#ParenType.
	VisitParenType(ctx *ParenTypeContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#IntType.
	VisitIntType(ctx *IntTypeContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#IfExpr.
	VisitIfExpr(ctx *IfExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#LetExpr.
	VisitLetExpr(ctx *LetExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#LetRecExpr.
	VisitLetRecExpr(ctx *LetRecExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#AppExpr.
	VisitAppExpr(ctx *AppExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#EmptyListExpr.
	VisitEmptyListExpr(ctx *EmptyListExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#ConsExpr.
	VisitConsExpr(ctx *ConsExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#VarExpr.
	VisitVarExpr(ctx *VarExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#BinOpExpr.
	VisitBinOpExpr(ctx *BinOpExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#ParenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#FunExpr.
	VisitFunExpr(ctx *FunExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#MatchExpr.
	VisitMatchExpr(ctx *MatchExprContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#env.
	VisitEnv(ctx *EnvContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#bind.
	VisitBind(ctx *BindContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#emptyPattern.
	VisitEmptyPattern(ctx *EmptyPatternContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#consPattern.
	VisitConsPattern(ctx *ConsPatternContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#fun.
	VisitFun(ctx *FunContext) interface{}

	// Visit a parse tree produced by TypingML4Parser#recFun.
	VisitRecFun(ctx *RecFunContext) interface{}
}
