// Code generated from EvalML3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML3
import "github.com/antlr4-go/antlr/v4"

// BaseEvalML3Listener is a complete listener for a parse tree produced by EvalML3Parser.
type BaseEvalML3Listener struct{}

var _ EvalML3Listener = &BaseEvalML3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEvalML3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEvalML3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEvalML3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEvalML3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuestion is called when production question is entered.
func (s *BaseEvalML3Listener) EnterQuestion(ctx *QuestionContext) {}

// ExitQuestion is called when production question is exited.
func (s *BaseEvalML3Listener) ExitQuestion(ctx *QuestionContext) {}

// EnterEval is called when production eval is entered.
func (s *BaseEvalML3Listener) EnterEval(ctx *EvalContext) {}

// ExitEval is called when production eval is exited.
func (s *BaseEvalML3Listener) ExitEval(ctx *EvalContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseEvalML3Listener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseEvalML3Listener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterIfExpr is called when production IfExpr is entered.
func (s *BaseEvalML3Listener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production IfExpr is exited.
func (s *BaseEvalML3Listener) ExitIfExpr(ctx *IfExprContext) {}

// EnterLetExpr is called when production LetExpr is entered.
func (s *BaseEvalML3Listener) EnterLetExpr(ctx *LetExprContext) {}

// ExitLetExpr is called when production LetExpr is exited.
func (s *BaseEvalML3Listener) ExitLetExpr(ctx *LetExprContext) {}

// EnterLetRecExpr is called when production LetRecExpr is entered.
func (s *BaseEvalML3Listener) EnterLetRecExpr(ctx *LetRecExprContext) {}

// ExitLetRecExpr is called when production LetRecExpr is exited.
func (s *BaseEvalML3Listener) ExitLetRecExpr(ctx *LetRecExprContext) {}

// EnterVarExpr is called when production VarExpr is entered.
func (s *BaseEvalML3Listener) EnterVarExpr(ctx *VarExprContext) {}

// ExitVarExpr is called when production VarExpr is exited.
func (s *BaseEvalML3Listener) ExitVarExpr(ctx *VarExprContext) {}

// EnterAppExpr is called when production AppExpr is entered.
func (s *BaseEvalML3Listener) EnterAppExpr(ctx *AppExprContext) {}

// ExitAppExpr is called when production AppExpr is exited.
func (s *BaseEvalML3Listener) ExitAppExpr(ctx *AppExprContext) {}

// EnterBinOpExpr is called when production BinOpExpr is entered.
func (s *BaseEvalML3Listener) EnterBinOpExpr(ctx *BinOpExprContext) {}

// ExitBinOpExpr is called when production BinOpExpr is exited.
func (s *BaseEvalML3Listener) ExitBinOpExpr(ctx *BinOpExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseEvalML3Listener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseEvalML3Listener) ExitIntExpr(ctx *IntExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseEvalML3Listener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseEvalML3Listener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterNegIntExpr is called when production NegIntExpr is entered.
func (s *BaseEvalML3Listener) EnterNegIntExpr(ctx *NegIntExprContext) {}

// ExitNegIntExpr is called when production NegIntExpr is exited.
func (s *BaseEvalML3Listener) ExitNegIntExpr(ctx *NegIntExprContext) {}

// EnterFunExpr is called when production FunExpr is entered.
func (s *BaseEvalML3Listener) EnterFunExpr(ctx *FunExprContext) {}

// ExitFunExpr is called when production FunExpr is exited.
func (s *BaseEvalML3Listener) ExitFunExpr(ctx *FunExprContext) {}

// EnterDefList is called when production defList is entered.
func (s *BaseEvalML3Listener) EnterDefList(ctx *DefListContext) {}

// ExitDefList is called when production defList is exited.
func (s *BaseEvalML3Listener) ExitDefList(ctx *DefListContext) {}

// EnterDef is called when production def is entered.
func (s *BaseEvalML3Listener) EnterDef(ctx *DefContext) {}

// ExitDef is called when production def is exited.
func (s *BaseEvalML3Listener) ExitDef(ctx *DefContext) {}

// EnterFun is called when production fun is entered.
func (s *BaseEvalML3Listener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BaseEvalML3Listener) ExitFun(ctx *FunContext) {}

// EnterRecFun is called when production recFun is entered.
func (s *BaseEvalML3Listener) EnterRecFun(ctx *RecFunContext) {}

// ExitRecFun is called when production recFun is exited.
func (s *BaseEvalML3Listener) ExitRecFun(ctx *RecFunContext) {}

// EnterIntValue is called when production IntValue is entered.
func (s *BaseEvalML3Listener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production IntValue is exited.
func (s *BaseEvalML3Listener) ExitIntValue(ctx *IntValueContext) {}

// EnterBoolValue is called when production BoolValue is entered.
func (s *BaseEvalML3Listener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production BoolValue is exited.
func (s *BaseEvalML3Listener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterFunValue is called when production FunValue is entered.
func (s *BaseEvalML3Listener) EnterFunValue(ctx *FunValueContext) {}

// ExitFunValue is called when production FunValue is exited.
func (s *BaseEvalML3Listener) ExitFunValue(ctx *FunValueContext) {}

// EnterRecFunValue is called when production RecFunValue is entered.
func (s *BaseEvalML3Listener) EnterRecFunValue(ctx *RecFunValueContext) {}

// ExitRecFunValue is called when production RecFunValue is exited.
func (s *BaseEvalML3Listener) ExitRecFunValue(ctx *RecFunValueContext) {}
