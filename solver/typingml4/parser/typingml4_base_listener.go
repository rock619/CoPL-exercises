// Code generated from TypingML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TypingML4
import "github.com/antlr4-go/antlr/v4"

// BaseTypingML4Listener is a complete listener for a parse tree produced by TypingML4Parser.
type BaseTypingML4Listener struct{}

var _ TypingML4Listener = &BaseTypingML4Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTypingML4Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTypingML4Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTypingML4Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTypingML4Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterJudgement is called when production judgement is entered.
func (s *BaseTypingML4Listener) EnterJudgement(ctx *JudgementContext) {}

// ExitJudgement is called when production judgement is exited.
func (s *BaseTypingML4Listener) ExitJudgement(ctx *JudgementContext) {}

// EnterBoolType is called when production BoolType is entered.
func (s *BaseTypingML4Listener) EnterBoolType(ctx *BoolTypeContext) {}

// ExitBoolType is called when production BoolType is exited.
func (s *BaseTypingML4Listener) ExitBoolType(ctx *BoolTypeContext) {}

// EnterListType is called when production ListType is entered.
func (s *BaseTypingML4Listener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production ListType is exited.
func (s *BaseTypingML4Listener) ExitListType(ctx *ListTypeContext) {}

// EnterFunType is called when production FunType is entered.
func (s *BaseTypingML4Listener) EnterFunType(ctx *FunTypeContext) {}

// ExitFunType is called when production FunType is exited.
func (s *BaseTypingML4Listener) ExitFunType(ctx *FunTypeContext) {}

// EnterParenType is called when production ParenType is entered.
func (s *BaseTypingML4Listener) EnterParenType(ctx *ParenTypeContext) {}

// ExitParenType is called when production ParenType is exited.
func (s *BaseTypingML4Listener) ExitParenType(ctx *ParenTypeContext) {}

// EnterIntType is called when production IntType is entered.
func (s *BaseTypingML4Listener) EnterIntType(ctx *IntTypeContext) {}

// ExitIntType is called when production IntType is exited.
func (s *BaseTypingML4Listener) ExitIntType(ctx *IntTypeContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseTypingML4Listener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseTypingML4Listener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterIfExpr is called when production IfExpr is entered.
func (s *BaseTypingML4Listener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production IfExpr is exited.
func (s *BaseTypingML4Listener) ExitIfExpr(ctx *IfExprContext) {}

// EnterLetExpr is called when production LetExpr is entered.
func (s *BaseTypingML4Listener) EnterLetExpr(ctx *LetExprContext) {}

// ExitLetExpr is called when production LetExpr is exited.
func (s *BaseTypingML4Listener) ExitLetExpr(ctx *LetExprContext) {}

// EnterLetRecExpr is called when production LetRecExpr is entered.
func (s *BaseTypingML4Listener) EnterLetRecExpr(ctx *LetRecExprContext) {}

// ExitLetRecExpr is called when production LetRecExpr is exited.
func (s *BaseTypingML4Listener) ExitLetRecExpr(ctx *LetRecExprContext) {}

// EnterAppExpr is called when production AppExpr is entered.
func (s *BaseTypingML4Listener) EnterAppExpr(ctx *AppExprContext) {}

// ExitAppExpr is called when production AppExpr is exited.
func (s *BaseTypingML4Listener) ExitAppExpr(ctx *AppExprContext) {}

// EnterEmptyListExpr is called when production EmptyListExpr is entered.
func (s *BaseTypingML4Listener) EnterEmptyListExpr(ctx *EmptyListExprContext) {}

// ExitEmptyListExpr is called when production EmptyListExpr is exited.
func (s *BaseTypingML4Listener) ExitEmptyListExpr(ctx *EmptyListExprContext) {}

// EnterConsExpr is called when production ConsExpr is entered.
func (s *BaseTypingML4Listener) EnterConsExpr(ctx *ConsExprContext) {}

// ExitConsExpr is called when production ConsExpr is exited.
func (s *BaseTypingML4Listener) ExitConsExpr(ctx *ConsExprContext) {}

// EnterVarExpr is called when production VarExpr is entered.
func (s *BaseTypingML4Listener) EnterVarExpr(ctx *VarExprContext) {}

// ExitVarExpr is called when production VarExpr is exited.
func (s *BaseTypingML4Listener) ExitVarExpr(ctx *VarExprContext) {}

// EnterBinOpExpr is called when production BinOpExpr is entered.
func (s *BaseTypingML4Listener) EnterBinOpExpr(ctx *BinOpExprContext) {}

// ExitBinOpExpr is called when production BinOpExpr is exited.
func (s *BaseTypingML4Listener) ExitBinOpExpr(ctx *BinOpExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseTypingML4Listener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseTypingML4Listener) ExitIntExpr(ctx *IntExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseTypingML4Listener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseTypingML4Listener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterFunExpr is called when production FunExpr is entered.
func (s *BaseTypingML4Listener) EnterFunExpr(ctx *FunExprContext) {}

// ExitFunExpr is called when production FunExpr is exited.
func (s *BaseTypingML4Listener) ExitFunExpr(ctx *FunExprContext) {}

// EnterMatchExpr is called when production MatchExpr is entered.
func (s *BaseTypingML4Listener) EnterMatchExpr(ctx *MatchExprContext) {}

// ExitMatchExpr is called when production MatchExpr is exited.
func (s *BaseTypingML4Listener) ExitMatchExpr(ctx *MatchExprContext) {}

// EnterEnv is called when production env is entered.
func (s *BaseTypingML4Listener) EnterEnv(ctx *EnvContext) {}

// ExitEnv is called when production env is exited.
func (s *BaseTypingML4Listener) ExitEnv(ctx *EnvContext) {}

// EnterBind is called when production bind is entered.
func (s *BaseTypingML4Listener) EnterBind(ctx *BindContext) {}

// ExitBind is called when production bind is exited.
func (s *BaseTypingML4Listener) ExitBind(ctx *BindContext) {}

// EnterEmptyPattern is called when production emptyPattern is entered.
func (s *BaseTypingML4Listener) EnterEmptyPattern(ctx *EmptyPatternContext) {}

// ExitEmptyPattern is called when production emptyPattern is exited.
func (s *BaseTypingML4Listener) ExitEmptyPattern(ctx *EmptyPatternContext) {}

// EnterConsPattern is called when production consPattern is entered.
func (s *BaseTypingML4Listener) EnterConsPattern(ctx *ConsPatternContext) {}

// ExitConsPattern is called when production consPattern is exited.
func (s *BaseTypingML4Listener) ExitConsPattern(ctx *ConsPatternContext) {}

// EnterFun is called when production fun is entered.
func (s *BaseTypingML4Listener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BaseTypingML4Listener) ExitFun(ctx *FunContext) {}

// EnterRecFun is called when production recFun is entered.
func (s *BaseTypingML4Listener) EnterRecFun(ctx *RecFunContext) {}

// ExitRecFun is called when production recFun is exited.
func (s *BaseTypingML4Listener) ExitRecFun(ctx *RecFunContext) {}
