// Code generated from TypingML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TypingML4
import "github.com/antlr4-go/antlr/v4"

// TypingML4Listener is a complete listener for a parse tree produced by TypingML4Parser.
type TypingML4Listener interface {
	antlr.ParseTreeListener

	// EnterJudgement is called when entering the judgement production.
	EnterJudgement(c *JudgementContext)

	// EnterBoolType is called when entering the BoolType production.
	EnterBoolType(c *BoolTypeContext)

	// EnterListType is called when entering the ListType production.
	EnterListType(c *ListTypeContext)

	// EnterFunType is called when entering the FunType production.
	EnterFunType(c *FunTypeContext)

	// EnterParenType is called when entering the ParenType production.
	EnterParenType(c *ParenTypeContext)

	// EnterIntType is called when entering the IntType production.
	EnterIntType(c *IntTypeContext)

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

	// EnterEnv is called when entering the env production.
	EnterEnv(c *EnvContext)

	// EnterBind is called when entering the bind production.
	EnterBind(c *BindContext)

	// EnterEmptyPattern is called when entering the emptyPattern production.
	EnterEmptyPattern(c *EmptyPatternContext)

	// EnterConsPattern is called when entering the consPattern production.
	EnterConsPattern(c *ConsPatternContext)

	// EnterFun is called when entering the fun production.
	EnterFun(c *FunContext)

	// EnterRecFun is called when entering the recFun production.
	EnterRecFun(c *RecFunContext)

	// ExitJudgement is called when exiting the judgement production.
	ExitJudgement(c *JudgementContext)

	// ExitBoolType is called when exiting the BoolType production.
	ExitBoolType(c *BoolTypeContext)

	// ExitListType is called when exiting the ListType production.
	ExitListType(c *ListTypeContext)

	// ExitFunType is called when exiting the FunType production.
	ExitFunType(c *FunTypeContext)

	// ExitParenType is called when exiting the ParenType production.
	ExitParenType(c *ParenTypeContext)

	// ExitIntType is called when exiting the IntType production.
	ExitIntType(c *IntTypeContext)

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

	// ExitEnv is called when exiting the env production.
	ExitEnv(c *EnvContext)

	// ExitBind is called when exiting the bind production.
	ExitBind(c *BindContext)

	// ExitEmptyPattern is called when exiting the emptyPattern production.
	ExitEmptyPattern(c *EmptyPatternContext)

	// ExitConsPattern is called when exiting the consPattern production.
	ExitConsPattern(c *ConsPatternContext)

	// ExitFun is called when exiting the fun production.
	ExitFun(c *FunContext)

	// ExitRecFun is called when exiting the recFun production.
	ExitRecFun(c *RecFunContext)
}
