// Code generated from EvalML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML4
import "github.com/antlr4-go/antlr/v4"

// EvalML4Listener is a complete listener for a parse tree produced by EvalML4Parser.
type EvalML4Listener interface {
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

	// EnterAppExpr is called when entering the AppExpr production.
	EnterAppExpr(c *AppExprContext)

	// EnterEmptyListExpr is called when entering the EmptyListExpr production.
	EnterEmptyListExpr(c *EmptyListExprContext)

	// EnterConsExpr is called when entering the ConsExpr production.
	EnterConsExpr(c *ConsExprContext)

	// EnterVarExpr is called when entering the VarExpr production.
	EnterVarExpr(c *VarExprContext)

	// EnterBinOpExpr is called when entering the BinOpExpr production.
	EnterBinOpExpr(c *BinOpExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterFunExpr is called when entering the FunExpr production.
	EnterFunExpr(c *FunExprContext)

	// EnterMatchExpr is called when entering the MatchExpr production.
	EnterMatchExpr(c *MatchExprContext)

	// EnterDefList is called when entering the defList production.
	EnterDefList(c *DefListContext)

	// EnterDef is called when entering the def production.
	EnterDef(c *DefContext)

	// EnterEmptyPattern is called when entering the emptyPattern production.
	EnterEmptyPattern(c *EmptyPatternContext)

	// EnterConsPattern is called when entering the consPattern production.
	EnterConsPattern(c *ConsPatternContext)

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterRecFun is called when entering the recFun production.
	EnterRecFun(c *RecFunContext)

	// EnterEmptyListValue is called when entering the EmptyListValue production.
	EnterEmptyListValue(c *EmptyListValueContext)

	// EnterBoolValue is called when entering the BoolValue production.
	EnterBoolValue(c *BoolValueContext)

	// EnterFunValue is called when entering the FunValue production.
	EnterFunValue(c *FunValueContext)

	// EnterRecFunValue is called when entering the RecFunValue production.
	EnterRecFunValue(c *RecFunValueContext)

	// EnterIntValue is called when entering the IntValue production.
	EnterIntValue(c *IntValueContext)

	// EnterConsValue is called when entering the ConsValue production.
	EnterConsValue(c *ConsValueContext)

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

	// ExitAppExpr is called when exiting the AppExpr production.
	ExitAppExpr(c *AppExprContext)

	// ExitEmptyListExpr is called when exiting the EmptyListExpr production.
	ExitEmptyListExpr(c *EmptyListExprContext)

	// ExitConsExpr is called when exiting the ConsExpr production.
	ExitConsExpr(c *ConsExprContext)

	// ExitVarExpr is called when exiting the VarExpr production.
	ExitVarExpr(c *VarExprContext)

	// ExitBinOpExpr is called when exiting the BinOpExpr production.
	ExitBinOpExpr(c *BinOpExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitFunExpr is called when exiting the FunExpr production.
	ExitFunExpr(c *FunExprContext)

	// ExitMatchExpr is called when exiting the MatchExpr production.
	ExitMatchExpr(c *MatchExprContext)

	// ExitDefList is called when exiting the defList production.
	ExitDefList(c *DefListContext)

	// ExitDef is called when exiting the def production.
	ExitDef(c *DefContext)

	// ExitEmptyPattern is called when exiting the emptyPattern production.
	ExitEmptyPattern(c *EmptyPatternContext)

	// ExitConsPattern is called when exiting the consPattern production.
	ExitConsPattern(c *ConsPatternContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitRecFun is called when exiting the recFun production.
	ExitRecFun(c *RecFunContext)

	// ExitEmptyListValue is called when exiting the EmptyListValue production.
	ExitEmptyListValue(c *EmptyListValueContext)

	// ExitBoolValue is called when exiting the BoolValue production.
	ExitBoolValue(c *BoolValueContext)

	// ExitFunValue is called when exiting the FunValue production.
	ExitFunValue(c *FunValueContext)

	// ExitRecFunValue is called when exiting the RecFunValue production.
	ExitRecFunValue(c *RecFunValueContext)

	// ExitIntValue is called when exiting the IntValue production.
	ExitIntValue(c *IntValueContext)

	// ExitConsValue is called when exiting the ConsValue production.
	ExitConsValue(c *ConsValueContext)
}
