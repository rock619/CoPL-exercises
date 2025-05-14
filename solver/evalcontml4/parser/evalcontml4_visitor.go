// Code generated from EvalContML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalContML4
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by EvalContML4Parser.
type EvalContML4Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EvalContML4Parser#eval.
	VisitEval(ctx *EvalContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#BoolValue.
	VisitBoolValue(ctx *BoolValueContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#FunValue.
	VisitFunValue(ctx *FunValueContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#RecFunValue.
	VisitRecFunValue(ctx *RecFunValueContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#ContValue.
	VisitContValue(ctx *ContValueContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#IntValue.
	VisitIntValue(ctx *IntValueContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#NilValue.
	VisitNilValue(ctx *NilValueContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#ConsValue.
	VisitConsValue(ctx *ConsValueContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#env.
	VisitEnv(ctx *EnvContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#bind.
	VisitBind(ctx *BindContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#BoolExp.
	VisitBoolExp(ctx *BoolExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#ConsExp.
	VisitConsExp(ctx *ConsExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#FunExp.
	VisitFunExp(ctx *FunExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#BinOpExp.
	VisitBinOpExp(ctx *BinOpExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#LetRecExp.
	VisitLetRecExp(ctx *LetRecExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#LetCCExp.
	VisitLetCCExp(ctx *LetCCExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#MatchExp.
	VisitMatchExp(ctx *MatchExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#IfExp.
	VisitIfExp(ctx *IfExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#AppExp.
	VisitAppExp(ctx *AppExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#ParenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#LetExp.
	VisitLetExp(ctx *LetExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#VarExp.
	VisitVarExp(ctx *VarExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#IntExp.
	VisitIntExp(ctx *IntExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#NilExp.
	VisitNilExp(ctx *NilExpContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#fun.
	VisitFun(ctx *FunContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#recFun.
	VisitRecFun(ctx *RecFunContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#TerminalCont.
	VisitTerminalCont(ctx *TerminalContContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#ExpCont.
	VisitExpCont(ctx *ExpContContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#ValueCont.
	VisitValueCont(ctx *ValueContContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#IfCont.
	VisitIfCont(ctx *IfContContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#LetCont.
	VisitLetCont(ctx *LetContContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#AppExpCont.
	VisitAppExpCont(ctx *AppExpContContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#AppValueCont.
	VisitAppValueCont(ctx *AppValueContContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#ConsExpCont.
	VisitConsExpCont(ctx *ConsExpContContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#ConsValueCont.
	VisitConsValueCont(ctx *ConsValueContContext) interface{}

	// Visit a parse tree produced by EvalContML4Parser#MatchCont.
	VisitMatchCont(ctx *MatchContContext) interface{}
}
