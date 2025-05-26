// Code generated from PolyTypingML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PolyTypingML4
import "github.com/antlr4-go/antlr/v4"

// PolyTypingML4Listener is a complete listener for a parse tree produced by PolyTypingML4Parser.
type PolyTypingML4Listener interface {
	antlr.ParseTreeListener

	// EnterEval is called when entering the eval production.
	EnterEval(c *EvalContext)

	// EnterVar is called when entering the var production.
	EnterVar(c *VarContext)

	// EnterTVar is called when entering the tVar production.
	EnterTVar(c *TVarContext)

	// EnterBoolType is called when entering the BoolType production.
	EnterBoolType(c *BoolTypeContext)

	// EnterTypeVar is called when entering the TypeVar production.
	EnterTypeVar(c *TypeVarContext)

	// EnterListType is called when entering the ListType production.
	EnterListType(c *ListTypeContext)

	// EnterFunType is called when entering the FunType production.
	EnterFunType(c *FunTypeContext)

	// EnterParenType is called when entering the ParenType production.
	EnterParenType(c *ParenTypeContext)

	// EnterIntType is called when entering the IntType production.
	EnterIntType(c *IntTypeContext)

	// EnterTyScheme is called when entering the tyScheme production.
	EnterTyScheme(c *TySchemeContext)

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

	// EnterIfExp is called when entering the IfExp production.
	EnterIfExp(c *IfExpContext)

	// EnterMatchExp is called when entering the MatchExp production.
	EnterMatchExp(c *MatchExpContext)

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

	// ExitEval is called when exiting the eval production.
	ExitEval(c *EvalContext)

	// ExitVar is called when exiting the var production.
	ExitVar(c *VarContext)

	// ExitTVar is called when exiting the tVar production.
	ExitTVar(c *TVarContext)

	// ExitBoolType is called when exiting the BoolType production.
	ExitBoolType(c *BoolTypeContext)

	// ExitTypeVar is called when exiting the TypeVar production.
	ExitTypeVar(c *TypeVarContext)

	// ExitListType is called when exiting the ListType production.
	ExitListType(c *ListTypeContext)

	// ExitFunType is called when exiting the FunType production.
	ExitFunType(c *FunTypeContext)

	// ExitParenType is called when exiting the ParenType production.
	ExitParenType(c *ParenTypeContext)

	// ExitIntType is called when exiting the IntType production.
	ExitIntType(c *IntTypeContext)

	// ExitTyScheme is called when exiting the tyScheme production.
	ExitTyScheme(c *TySchemeContext)

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

	// ExitIfExp is called when exiting the IfExp production.
	ExitIfExp(c *IfExpContext)

	// ExitMatchExp is called when exiting the MatchExp production.
	ExitMatchExp(c *MatchExpContext)

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
}
