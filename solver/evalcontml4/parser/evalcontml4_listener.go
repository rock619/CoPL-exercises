// Code generated from EvalContML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalContML4
import "github.com/antlr4-go/antlr/v4"

// EvalContML4Listener is a complete listener for a parse tree produced by EvalContML4Parser.
type EvalContML4Listener interface {
	antlr.ParseTreeListener

	// EnterEval is called when entering the eval production.
	EnterEval(c *EvalContext)

	// EnterBoolValue is called when entering the BoolValue production.
	EnterBoolValue(c *BoolValueContext)

	// EnterFunValue is called when entering the FunValue production.
	EnterFunValue(c *FunValueContext)

	// EnterRecFunValue is called when entering the RecFunValue production.
	EnterRecFunValue(c *RecFunValueContext)

	// EnterContValue is called when entering the ContValue production.
	EnterContValue(c *ContValueContext)

	// EnterIntValue is called when entering the IntValue production.
	EnterIntValue(c *IntValueContext)

	// EnterNilValue is called when entering the NilValue production.
	EnterNilValue(c *NilValueContext)

	// EnterConsValue is called when entering the ConsValue production.
	EnterConsValue(c *ConsValueContext)

	// EnterEnv is called when entering the env production.
	EnterEnv(c *EnvContext)

	// EnterBind is called when entering the bind production.
	EnterBind(c *BindContext)

	// EnterBoolExp is called when entering the BoolExp production.
	EnterBoolExp(c *BoolExpContext)

	// EnterConsExp is called when entering the ConsExp production.
	EnterConsExp(c *ConsExpContext)

	// EnterFunExp is called when entering the FunExp production.
	EnterFunExp(c *FunExpContext)

	// EnterBinOpExp is called when entering the BinOpExp production.
	EnterBinOpExp(c *BinOpExpContext)

	// EnterLetRecExp is called when entering the LetRecExp production.
	EnterLetRecExp(c *LetRecExpContext)

	// EnterLetCCExp is called when entering the LetCCExp production.
	EnterLetCCExp(c *LetCCExpContext)

	// EnterMatchExp is called when entering the MatchExp production.
	EnterMatchExp(c *MatchExpContext)

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

	// EnterNilExp is called when entering the NilExp production.
	EnterNilExp(c *NilExpContext)

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterRecFun is called when entering the recFun production.
	EnterRecFun(c *RecFunContext)

	// EnterTerminalCont is called when entering the TerminalCont production.
	EnterTerminalCont(c *TerminalContContext)

	// EnterExpCont is called when entering the ExpCont production.
	EnterExpCont(c *ExpContContext)

	// EnterValueCont is called when entering the ValueCont production.
	EnterValueCont(c *ValueContContext)

	// EnterIfCont is called when entering the IfCont production.
	EnterIfCont(c *IfContContext)

	// EnterLetCont is called when entering the LetCont production.
	EnterLetCont(c *LetContContext)

	// EnterAppExpCont is called when entering the AppExpCont production.
	EnterAppExpCont(c *AppExpContContext)

	// EnterAppValueCont is called when entering the AppValueCont production.
	EnterAppValueCont(c *AppValueContContext)

	// EnterConsExpCont is called when entering the ConsExpCont production.
	EnterConsExpCont(c *ConsExpContContext)

	// EnterConsValueCont is called when entering the ConsValueCont production.
	EnterConsValueCont(c *ConsValueContContext)

	// EnterMatchCont is called when entering the MatchCont production.
	EnterMatchCont(c *MatchContContext)

	// ExitEval is called when exiting the eval production.
	ExitEval(c *EvalContext)

	// ExitBoolValue is called when exiting the BoolValue production.
	ExitBoolValue(c *BoolValueContext)

	// ExitFunValue is called when exiting the FunValue production.
	ExitFunValue(c *FunValueContext)

	// ExitRecFunValue is called when exiting the RecFunValue production.
	ExitRecFunValue(c *RecFunValueContext)

	// ExitContValue is called when exiting the ContValue production.
	ExitContValue(c *ContValueContext)

	// ExitIntValue is called when exiting the IntValue production.
	ExitIntValue(c *IntValueContext)

	// ExitNilValue is called when exiting the NilValue production.
	ExitNilValue(c *NilValueContext)

	// ExitConsValue is called when exiting the ConsValue production.
	ExitConsValue(c *ConsValueContext)

	// ExitEnv is called when exiting the env production.
	ExitEnv(c *EnvContext)

	// ExitBind is called when exiting the bind production.
	ExitBind(c *BindContext)

	// ExitBoolExp is called when exiting the BoolExp production.
	ExitBoolExp(c *BoolExpContext)

	// ExitConsExp is called when exiting the ConsExp production.
	ExitConsExp(c *ConsExpContext)

	// ExitFunExp is called when exiting the FunExp production.
	ExitFunExp(c *FunExpContext)

	// ExitBinOpExp is called when exiting the BinOpExp production.
	ExitBinOpExp(c *BinOpExpContext)

	// ExitLetRecExp is called when exiting the LetRecExp production.
	ExitLetRecExp(c *LetRecExpContext)

	// ExitLetCCExp is called when exiting the LetCCExp production.
	ExitLetCCExp(c *LetCCExpContext)

	// ExitMatchExp is called when exiting the MatchExp production.
	ExitMatchExp(c *MatchExpContext)

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

	// ExitNilExp is called when exiting the NilExp production.
	ExitNilExp(c *NilExpContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitRecFun is called when exiting the recFun production.
	ExitRecFun(c *RecFunContext)

	// ExitTerminalCont is called when exiting the TerminalCont production.
	ExitTerminalCont(c *TerminalContContext)

	// ExitExpCont is called when exiting the ExpCont production.
	ExitExpCont(c *ExpContContext)

	// ExitValueCont is called when exiting the ValueCont production.
	ExitValueCont(c *ValueContContext)

	// ExitIfCont is called when exiting the IfCont production.
	ExitIfCont(c *IfContContext)

	// ExitLetCont is called when exiting the LetCont production.
	ExitLetCont(c *LetContContext)

	// ExitAppExpCont is called when exiting the AppExpCont production.
	ExitAppExpCont(c *AppExpContContext)

	// ExitAppValueCont is called when exiting the AppValueCont production.
	ExitAppValueCont(c *AppValueContContext)

	// ExitConsExpCont is called when exiting the ConsExpCont production.
	ExitConsExpCont(c *ConsExpContContext)

	// ExitConsValueCont is called when exiting the ConsValueCont production.
	ExitConsValueCont(c *ConsValueContContext)

	// ExitMatchCont is called when exiting the MatchCont production.
	ExitMatchCont(c *MatchContContext)
}
