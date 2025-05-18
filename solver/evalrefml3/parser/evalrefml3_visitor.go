// Code generated from EvalRefML3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalRefML3
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by EvalRefML3Parser.
type EvalRefML3Visitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by EvalRefML3Parser#eval.
	VisitEval(ctx *EvalContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#IntValue.
	VisitIntValue(ctx *IntValueContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#BoolValue.
	VisitBoolValue(ctx *BoolValueContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#LocValue.
	VisitLocValue(ctx *LocValueContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#FunValue.
	VisitFunValue(ctx *FunValueContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#RecFunValue.
	VisitRecFunValue(ctx *RecFunValueContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#env.
	VisitEnv(ctx *EnvContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#bind.
	VisitBind(ctx *BindContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#store.
	VisitStore(ctx *StoreContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#BoolExp.
	VisitBoolExp(ctx *BoolExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#RefExp.
	VisitRefExp(ctx *RefExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#DerefExp.
	VisitDerefExp(ctx *DerefExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#FunExp.
	VisitFunExp(ctx *FunExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#BinOpExp.
	VisitBinOpExp(ctx *BinOpExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#LetRecExp.
	VisitLetRecExp(ctx *LetRecExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#AssignExp.
	VisitAssignExp(ctx *AssignExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#IfExp.
	VisitIfExp(ctx *IfExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#AppExp.
	VisitAppExp(ctx *AppExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#ParenExp.
	VisitParenExp(ctx *ParenExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#LetExp.
	VisitLetExp(ctx *LetExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#VarExp.
	VisitVarExp(ctx *VarExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#IntExp.
	VisitIntExp(ctx *IntExpContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#fun.
	VisitFun(ctx *FunContext) interface{}

	// Visit a parse tree produced by EvalRefML3Parser#recFun.
	VisitRecFun(ctx *RecFunContext) interface{}
}
