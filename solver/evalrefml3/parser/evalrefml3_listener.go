// Code generated from EvalRefML3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalRefML3
import "github.com/antlr4-go/antlr/v4"

// EvalRefML3Listener is a complete listener for a parse tree produced by EvalRefML3Parser.
type EvalRefML3Listener interface {
	antlr.ParseTreeListener

	// EnterEval is called when entering the eval production.
	EnterEval(c *EvalContext)

	// EnterIntValue is called when entering the IntValue production.
	EnterIntValue(c *IntValueContext)

	// EnterBoolValue is called when entering the BoolValue production.
	EnterBoolValue(c *BoolValueContext)

	// EnterLocValue is called when entering the LocValue production.
	EnterLocValue(c *LocValueContext)

	// EnterFunValue is called when entering the FunValue production.
	EnterFunValue(c *FunValueContext)

	// EnterRecFunValue is called when entering the RecFunValue production.
	EnterRecFunValue(c *RecFunValueContext)

	// EnterEnv is called when entering the env production.
	EnterEnv(c *EnvContext)

	// EnterBind is called when entering the bind production.
	EnterBind(c *BindContext)

	// EnterStore is called when entering the store production.
	EnterStore(c *StoreContext)

	// EnterAssign is called when entering the assign production.
	EnterAssign(c *AssignContext)

	// EnterBoolExp is called when entering the BoolExp production.
	EnterBoolExp(c *BoolExpContext)

	// EnterRefExp is called when entering the RefExp production.
	EnterRefExp(c *RefExpContext)

	// EnterDerefExp is called when entering the DerefExp production.
	EnterDerefExp(c *DerefExpContext)

	// EnterFunExp is called when entering the FunExp production.
	EnterFunExp(c *FunExpContext)

	// EnterBinOpExp is called when entering the BinOpExp production.
	EnterBinOpExp(c *BinOpExpContext)

	// EnterLetRecExp is called when entering the LetRecExp production.
	EnterLetRecExp(c *LetRecExpContext)

	// EnterAssignExp is called when entering the AssignExp production.
	EnterAssignExp(c *AssignExpContext)

	// EnterIfExp is called when entering the IfExp production.
	EnterIfExp(c *IfExpContext)

	// EnterAppExp is called when entering the AppExp production.
	EnterAppExp(c *AppExpContext)

	// EnterParenExp is called when entering the ParenExp production.
	EnterParenExp(c *ParenExpContext)

	// EnterLetExp is called when entering the LetExp production.
	EnterLetExp(c *LetExpContext)

	// EnterVarExp is called when entering the VarExp production.
	EnterVarExp(c *VarExpContext)

	// EnterIntExp is called when entering the IntExp production.
	EnterIntExp(c *IntExpContext)

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterRecFun is called when entering the recFun production.
	EnterRecFun(c *RecFunContext)

	// ExitEval is called when exiting the eval production.
	ExitEval(c *EvalContext)

	// ExitIntValue is called when exiting the IntValue production.
	ExitIntValue(c *IntValueContext)

	// ExitBoolValue is called when exiting the BoolValue production.
	ExitBoolValue(c *BoolValueContext)

	// ExitLocValue is called when exiting the LocValue production.
	ExitLocValue(c *LocValueContext)

	// ExitFunValue is called when exiting the FunValue production.
	ExitFunValue(c *FunValueContext)

	// ExitRecFunValue is called when exiting the RecFunValue production.
	ExitRecFunValue(c *RecFunValueContext)

	// ExitEnv is called when exiting the env production.
	ExitEnv(c *EnvContext)

	// ExitBind is called when exiting the bind production.
	ExitBind(c *BindContext)

	// ExitStore is called when exiting the store production.
	ExitStore(c *StoreContext)

	// ExitAssign is called when exiting the assign production.
	ExitAssign(c *AssignContext)

	// ExitBoolExp is called when exiting the BoolExp production.
	ExitBoolExp(c *BoolExpContext)

	// ExitRefExp is called when exiting the RefExp production.
	ExitRefExp(c *RefExpContext)

	// ExitDerefExp is called when exiting the DerefExp production.
	ExitDerefExp(c *DerefExpContext)

	// ExitFunExp is called when exiting the FunExp production.
	ExitFunExp(c *FunExpContext)

	// ExitBinOpExp is called when exiting the BinOpExp production.
	ExitBinOpExp(c *BinOpExpContext)

	// ExitLetRecExp is called when exiting the LetRecExp production.
	ExitLetRecExp(c *LetRecExpContext)

	// ExitAssignExp is called when exiting the AssignExp production.
	ExitAssignExp(c *AssignExpContext)

	// ExitIfExp is called when exiting the IfExp production.
	ExitIfExp(c *IfExpContext)

	// ExitAppExp is called when exiting the AppExp production.
	ExitAppExp(c *AppExpContext)

	// ExitParenExp is called when exiting the ParenExp production.
	ExitParenExp(c *ParenExpContext)

	// ExitLetExp is called when exiting the LetExp production.
	ExitLetExp(c *LetExpContext)

	// ExitVarExp is called when exiting the VarExp production.
	ExitVarExp(c *VarExpContext)

	// ExitIntExp is called when exiting the IntExp production.
	ExitIntExp(c *IntExpContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitRecFun is called when exiting the recFun production.
	ExitRecFun(c *RecFunContext)
}
