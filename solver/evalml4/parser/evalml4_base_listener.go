// Code generated from EvalML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML4
import "github.com/antlr4-go/antlr/v4"

// BaseEvalML4Listener is a complete listener for a parse tree produced by EvalML4Parser.
type BaseEvalML4Listener struct{}

var _ EvalML4Listener = &BaseEvalML4Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEvalML4Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEvalML4Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEvalML4Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEvalML4Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuestion is called when production question is entered.
func (s *BaseEvalML4Listener) EnterQuestion(ctx *QuestionContext) {}

// ExitQuestion is called when production question is exited.
func (s *BaseEvalML4Listener) ExitQuestion(ctx *QuestionContext) {}

// EnterEval is called when production eval is entered.
func (s *BaseEvalML4Listener) EnterEval(ctx *EvalContext) {}

// ExitEval is called when production eval is exited.
func (s *BaseEvalML4Listener) ExitEval(ctx *EvalContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseEvalML4Listener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseEvalML4Listener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterIfExpr is called when production IfExpr is entered.
func (s *BaseEvalML4Listener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production IfExpr is exited.
func (s *BaseEvalML4Listener) ExitIfExpr(ctx *IfExprContext) {}

// EnterLetExpr is called when production LetExpr is entered.
func (s *BaseEvalML4Listener) EnterLetExpr(ctx *LetExprContext) {}

// ExitLetExpr is called when production LetExpr is exited.
func (s *BaseEvalML4Listener) ExitLetExpr(ctx *LetExprContext) {}

// EnterLetRecExpr is called when production LetRecExpr is entered.
func (s *BaseEvalML4Listener) EnterLetRecExpr(ctx *LetRecExprContext) {}

// ExitLetRecExpr is called when production LetRecExpr is exited.
func (s *BaseEvalML4Listener) ExitLetRecExpr(ctx *LetRecExprContext) {}

// EnterAppExpr is called when production AppExpr is entered.
func (s *BaseEvalML4Listener) EnterAppExpr(ctx *AppExprContext) {}

// ExitAppExpr is called when production AppExpr is exited.
func (s *BaseEvalML4Listener) ExitAppExpr(ctx *AppExprContext) {}

// EnterEmptyListExpr is called when production EmptyListExpr is entered.
func (s *BaseEvalML4Listener) EnterEmptyListExpr(ctx *EmptyListExprContext) {}

// ExitEmptyListExpr is called when production EmptyListExpr is exited.
func (s *BaseEvalML4Listener) ExitEmptyListExpr(ctx *EmptyListExprContext) {}

// EnterConsExpr is called when production ConsExpr is entered.
func (s *BaseEvalML4Listener) EnterConsExpr(ctx *ConsExprContext) {}

// ExitConsExpr is called when production ConsExpr is exited.
func (s *BaseEvalML4Listener) ExitConsExpr(ctx *ConsExprContext) {}

// EnterVarExpr is called when production VarExpr is entered.
func (s *BaseEvalML4Listener) EnterVarExpr(ctx *VarExprContext) {}

// ExitVarExpr is called when production VarExpr is exited.
func (s *BaseEvalML4Listener) ExitVarExpr(ctx *VarExprContext) {}

// EnterBinOpExpr is called when production BinOpExpr is entered.
func (s *BaseEvalML4Listener) EnterBinOpExpr(ctx *BinOpExprContext) {}

// ExitBinOpExpr is called when production BinOpExpr is exited.
func (s *BaseEvalML4Listener) ExitBinOpExpr(ctx *BinOpExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseEvalML4Listener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseEvalML4Listener) ExitIntExpr(ctx *IntExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseEvalML4Listener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseEvalML4Listener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterFunExpr is called when production FunExpr is entered.
func (s *BaseEvalML4Listener) EnterFunExpr(ctx *FunExprContext) {}

// ExitFunExpr is called when production FunExpr is exited.
func (s *BaseEvalML4Listener) ExitFunExpr(ctx *FunExprContext) {}

// EnterMatchExpr is called when production MatchExpr is entered.
func (s *BaseEvalML4Listener) EnterMatchExpr(ctx *MatchExprContext) {}

// ExitMatchExpr is called when production MatchExpr is exited.
func (s *BaseEvalML4Listener) ExitMatchExpr(ctx *MatchExprContext) {}

// EnterDefList is called when production defList is entered.
func (s *BaseEvalML4Listener) EnterDefList(ctx *DefListContext) {}

// ExitDefList is called when production defList is exited.
func (s *BaseEvalML4Listener) ExitDefList(ctx *DefListContext) {}

// EnterDef is called when production def is entered.
func (s *BaseEvalML4Listener) EnterDef(ctx *DefContext) {}

// ExitDef is called when production def is exited.
func (s *BaseEvalML4Listener) ExitDef(ctx *DefContext) {}

// EnterEmptyPattern is called when production emptyPattern is entered.
func (s *BaseEvalML4Listener) EnterEmptyPattern(ctx *EmptyPatternContext) {}

// ExitEmptyPattern is called when production emptyPattern is exited.
func (s *BaseEvalML4Listener) ExitEmptyPattern(ctx *EmptyPatternContext) {}

// EnterConsPattern is called when production consPattern is entered.
func (s *BaseEvalML4Listener) EnterConsPattern(ctx *ConsPatternContext) {}

// ExitConsPattern is called when production consPattern is exited.
func (s *BaseEvalML4Listener) ExitConsPattern(ctx *ConsPatternContext) {}

// EnterFun is called when production fun is entered.
func (s *BaseEvalML4Listener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BaseEvalML4Listener) ExitFun(ctx *FunContext) {}

// EnterRecFun is called when production recFun is entered.
func (s *BaseEvalML4Listener) EnterRecFun(ctx *RecFunContext) {}

// ExitRecFun is called when production recFun is exited.
func (s *BaseEvalML4Listener) ExitRecFun(ctx *RecFunContext) {}

// EnterEmptyListValue is called when production EmptyListValue is entered.
func (s *BaseEvalML4Listener) EnterEmptyListValue(ctx *EmptyListValueContext) {}

// ExitEmptyListValue is called when production EmptyListValue is exited.
func (s *BaseEvalML4Listener) ExitEmptyListValue(ctx *EmptyListValueContext) {}

// EnterBoolValue is called when production BoolValue is entered.
func (s *BaseEvalML4Listener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production BoolValue is exited.
func (s *BaseEvalML4Listener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterFunValue is called when production FunValue is entered.
func (s *BaseEvalML4Listener) EnterFunValue(ctx *FunValueContext) {}

// ExitFunValue is called when production FunValue is exited.
func (s *BaseEvalML4Listener) ExitFunValue(ctx *FunValueContext) {}

// EnterRecFunValue is called when production RecFunValue is entered.
func (s *BaseEvalML4Listener) EnterRecFunValue(ctx *RecFunValueContext) {}

// ExitRecFunValue is called when production RecFunValue is exited.
func (s *BaseEvalML4Listener) ExitRecFunValue(ctx *RecFunValueContext) {}

// EnterIntValue is called when production IntValue is entered.
func (s *BaseEvalML4Listener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production IntValue is exited.
func (s *BaseEvalML4Listener) ExitIntValue(ctx *IntValueContext) {}

// EnterConsValue is called when production ConsValue is entered.
func (s *BaseEvalML4Listener) EnterConsValue(ctx *ConsValueContext) {}

// ExitConsValue is called when production ConsValue is exited.
func (s *BaseEvalML4Listener) ExitConsValue(ctx *ConsValueContext) {}
