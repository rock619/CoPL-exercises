// Code generated from EvalML3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML3
import "github.com/antlr4-go/antlr/v4"

// EvalML3Listener is a complete listener for a parse tree produced by EvalML3Parser.
type EvalML3Listener interface {
	antlr.ParseTreeListener

	// EnterQuestion is called when entering the question production.
	EnterQuestion(c *QuestionContext)

	// EnterEval is called when entering the eval production.
	EnterEval(c *EvalContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterIfExpr is called when entering the IfExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterLetExpr is called when entering the LetExpr production.
	EnterLetExpr(c *LetExprContext)

	// EnterLetRecExpr is called when entering the LetRecExpr production.
	EnterLetRecExpr(c *LetRecExprContext)

	// EnterVarExpr is called when entering the VarExpr production.
	EnterVarExpr(c *VarExprContext)

	// EnterAppExpr is called when entering the AppExpr production.
	EnterAppExpr(c *AppExprContext)

	// EnterBinOpExpr is called when entering the BinOpExpr production.
	EnterBinOpExpr(c *BinOpExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterNegIntExpr is called when entering the NegIntExpr production.
	EnterNegIntExpr(c *NegIntExprContext)

	// EnterFunExpr is called when entering the FunExpr production.
	EnterFunExpr(c *FunExprContext)

	// EnterDefList is called when entering the defList production.
	EnterDefList(c *DefListContext)

	// EnterDef is called when entering the def production.
	EnterDef(c *DefContext)

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterRecFun is called when entering the recFun production.
	EnterRecFun(c *RecFunContext)

	// EnterIntValue is called when entering the IntValue production.
	EnterIntValue(c *IntValueContext)

	// EnterBoolValue is called when entering the BoolValue production.
	EnterBoolValue(c *BoolValueContext)

	// EnterFunValue is called when entering the FunValue production.
	EnterFunValue(c *FunValueContext)

	// EnterRecFunValue is called when entering the RecFunValue production.
	EnterRecFunValue(c *RecFunValueContext)

	// ExitQuestion is called when exiting the question production.
	ExitQuestion(c *QuestionContext)

	// ExitEval is called when exiting the eval production.
	ExitEval(c *EvalContext)

	// ExitBoolExpr is called when exiting the BoolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitIfExpr is called when exiting the IfExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitLetExpr is called when exiting the LetExpr production.
	ExitLetExpr(c *LetExprContext)

	// ExitLetRecExpr is called when exiting the LetRecExpr production.
	ExitLetRecExpr(c *LetRecExprContext)

	// ExitVarExpr is called when exiting the VarExpr production.
	ExitVarExpr(c *VarExprContext)

	// ExitAppExpr is called when exiting the AppExpr production.
	ExitAppExpr(c *AppExprContext)

	// ExitBinOpExpr is called when exiting the BinOpExpr production.
	ExitBinOpExpr(c *BinOpExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitNegIntExpr is called when exiting the NegIntExpr production.
	ExitNegIntExpr(c *NegIntExprContext)

	// ExitFunExpr is called when exiting the FunExpr production.
	ExitFunExpr(c *FunExprContext)

	// ExitDefList is called when exiting the defList production.
	ExitDefList(c *DefListContext)

	// ExitDef is called when exiting the def production.
	ExitDef(c *DefContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitRecFun is called when exiting the recFun production.
	ExitRecFun(c *RecFunContext)

	// ExitIntValue is called when exiting the IntValue production.
	ExitIntValue(c *IntValueContext)

	// ExitBoolValue is called when exiting the BoolValue production.
	ExitBoolValue(c *BoolValueContext)

	// ExitFunValue is called when exiting the FunValue production.
	ExitFunValue(c *FunValueContext)

	// ExitRecFunValue is called when exiting the RecFunValue production.
	ExitRecFunValue(c *RecFunValueContext)
}
