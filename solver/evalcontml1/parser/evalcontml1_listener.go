// Code generated from EvalContML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalContML1
import "github.com/antlr4-go/antlr/v4"

// EvalContML1Listener is a complete listener for a parse tree produced by EvalContML1Parser.
type EvalContML1Listener interface {
	antlr.ParseTreeListener

	// EnterEval is called when entering the eval production.
	EnterEval(c *EvalContext)

	// EnterIntValue is called when entering the IntValue production.
	EnterIntValue(c *IntValueContext)

	// EnterBoolValue is called when entering the BoolValue production.
	EnterBoolValue(c *BoolValueContext)

	// EnterBoolExp is called when entering the BoolExp production.
	EnterBoolExp(c *BoolExpContext)

	// EnterIfExp is called when entering the IfExp production.
	EnterIfExp(c *IfExpContext)

	// EnterParenExp is called when entering the ParenExp production.
	EnterParenExp(c *ParenExpContext)

	// EnterIntExp is called when entering the IntExp production.
	EnterIntExp(c *IntExpContext)

	// EnterBinOpExp is called when entering the BinOpExp production.
	EnterBinOpExp(c *BinOpExpContext)

	// EnterUnaryCont is called when entering the UnaryCont production.
	EnterUnaryCont(c *UnaryContContext)

	// EnterExpCont is called when entering the ExpCont production.
	EnterExpCont(c *ExpContContext)

	// EnterValueCont is called when entering the ValueCont production.
	EnterValueCont(c *ValueContContext)

	// EnterIfCont is called when entering the IfCont production.
	EnterIfCont(c *IfContContext)

	// ExitEval is called when exiting the eval production.
	ExitEval(c *EvalContext)

	// ExitIntValue is called when exiting the IntValue production.
	ExitIntValue(c *IntValueContext)

	// ExitBoolValue is called when exiting the BoolValue production.
	ExitBoolValue(c *BoolValueContext)

	// ExitBoolExp is called when exiting the BoolExp production.
	ExitBoolExp(c *BoolExpContext)

	// ExitIfExp is called when exiting the IfExp production.
	ExitIfExp(c *IfExpContext)

	// ExitParenExp is called when exiting the ParenExp production.
	ExitParenExp(c *ParenExpContext)

	// ExitIntExp is called when exiting the IntExp production.
	ExitIntExp(c *IntExpContext)

	// ExitBinOpExp is called when exiting the BinOpExp production.
	ExitBinOpExp(c *BinOpExpContext)

	// ExitUnaryCont is called when exiting the UnaryCont production.
	ExitUnaryCont(c *UnaryContContext)

	// ExitExpCont is called when exiting the ExpCont production.
	ExitExpCont(c *ExpContContext)

	// ExitValueCont is called when exiting the ValueCont production.
	ExitValueCont(c *ValueContContext)

	// ExitIfCont is called when exiting the IfCont production.
	ExitIfCont(c *IfContContext)
}
