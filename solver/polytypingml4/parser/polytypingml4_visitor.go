// Code generated from PolyTypingML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PolyTypingML4
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by PolyTypingML4Parser.
type PolyTypingML4Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by PolyTypingML4Parser#eval.
	VisitEval(ctx *EvalContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#var.
	VisitVar(ctx *VarContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#tVar.
	VisitTVar(ctx *TVarContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#BoolType.
	VisitBoolType(ctx *BoolTypeContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#TypeVar.
	VisitTypeVar(ctx *TypeVarContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#ListType.
	VisitListType(ctx *ListTypeContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#FunType.
	VisitFunType(ctx *FunTypeContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#ParenType.
	VisitParenType(ctx *ParenTypeContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#IntType.
	VisitIntType(ctx *IntTypeContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#tyScheme.
	VisitTyScheme(ctx *TySchemeContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#env.
	VisitEnv(ctx *EnvContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#bind.
	VisitBind(ctx *BindContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#BoolExp.
	VisitBoolExp(ctx *BoolExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#ConsExp.
	VisitConsExp(ctx *ConsExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#FunExp.
	VisitFunExp(ctx *FunExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#BinOpExp.
	VisitBinOpExp(ctx *BinOpExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#LetRecExp.
	VisitLetRecExp(ctx *LetRecExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#IfExp.
	VisitIfExp(ctx *IfExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#MatchExp.
	VisitMatchExp(ctx *MatchExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#AppExp.
	VisitAppExp(ctx *AppExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#ParenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#LetExp.
	VisitLetExp(ctx *LetExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#VarExp.
	VisitVarExp(ctx *VarExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#IntExp.
	VisitIntExp(ctx *IntExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#NilExp.
	VisitNilExp(ctx *NilExpContext) interface{}

	// Visit a parse tree produced by PolyTypingML4Parser#fun.
	VisitFun(ctx *FunContext) interface{}
}
