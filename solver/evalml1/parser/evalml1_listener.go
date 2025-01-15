// Code generated from EvalML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML1
import "github.com/antlr4-go/antlr/v4"

// EvalML1Listener is a complete listener for a parse tree produced by EvalML1Parser.
type EvalML1Listener interface {
	antlr.ParseTreeListener

	// EnterQuestion is called when entering the question production.
	EnterQuestion(c *QuestionContext)

	// EnterEval is called when entering the eval production.
	EnterEval(c *EvalContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterIfExpr is called when entering the IfExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterBinOpExpr is called when entering the BinOpExpr production.
	EnterBinOpExpr(c *BinOpExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterParenExpr is called when entering the ParenExpr production.
	EnterParenExpr(c *ParenExprContext)

	// EnterNegIntExpr is called when entering the NegIntExpr production.
	EnterNegIntExpr(c *NegIntExprContext)

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

	// ExitBinOpExpr is called when exiting the BinOpExpr production.
	ExitBinOpExpr(c *BinOpExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitParenExpr is called when exiting the ParenExpr production.
	ExitParenExpr(c *ParenExprContext)

	// ExitNegIntExpr is called when exiting the NegIntExpr production.
	ExitNegIntExpr(c *NegIntExprContext)

	// ExitIntValue is called when exiting the IntValue production.
	ExitIntValue(c *IntValueContext)

	// ExitBoolValue is called when exiting the BoolValue production.
	ExitBoolValue(c *BoolValueContext)
}
