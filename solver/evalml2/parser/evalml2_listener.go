// Code generated from EvalML2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML2
import "github.com/antlr4-go/antlr/v4"

// EvalML2Listener is a complete listener for a parse tree produced by EvalML2Parser.
type EvalML2Listener interface {
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

	// EnterVarExpr is called when entering the VarExpr production.
	EnterVarExpr(c *VarExprContext)

	// EnterBinOpExpr is called when entering the BinOpExpr production.
	EnterBinOpExpr(c *BinOpExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterNegIntExpr is called when entering the NegIntExpr production.
	EnterNegIntExpr(c *NegIntExprContext)

	// EnterDefList is called when entering the defList production.
	EnterDefList(c *DefListContext)

	// EnterDef is called when entering the def production.
	EnterDef(c *DefContext)

	// EnterIntValue is called when entering the IntValue production.
	EnterIntValue(c *IntValueContext)

	// EnterBoolValue is called when entering the BoolValue production.
	EnterBoolValue(c *BoolValueContext)

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

	// ExitVarExpr is called when exiting the VarExpr production.
	ExitVarExpr(c *VarExprContext)

	// ExitBinOpExpr is called when exiting the BinOpExpr production.
	ExitBinOpExpr(c *BinOpExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitNegIntExpr is called when exiting the NegIntExpr production.
	ExitNegIntExpr(c *NegIntExprContext)

	// ExitDefList is called when exiting the defList production.
	ExitDefList(c *DefListContext)

	// ExitDef is called when exiting the def production.
	ExitDef(c *DefContext)

	// ExitIntValue is called when exiting the IntValue production.
	ExitIntValue(c *IntValueContext)

	// ExitBoolValue is called when exiting the BoolValue production.
	ExitBoolValue(c *BoolValueContext)
}
