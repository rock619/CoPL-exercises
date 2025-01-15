// Code generated from EvalML2.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML2
import "github.com/antlr4-go/antlr/v4"

// BaseEvalML2Listener is a complete listener for a parse tree produced by EvalML2Parser.
type BaseEvalML2Listener struct{}

var _ EvalML2Listener = &BaseEvalML2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEvalML2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEvalML2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEvalML2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEvalML2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuestion is called when production question is entered.
func (s *BaseEvalML2Listener) EnterQuestion(ctx *QuestionContext) {}

// ExitQuestion is called when production question is exited.
func (s *BaseEvalML2Listener) ExitQuestion(ctx *QuestionContext) {}

// EnterEval is called when production eval is entered.
func (s *BaseEvalML2Listener) EnterEval(ctx *EvalContext) {}

// ExitEval is called when production eval is exited.
func (s *BaseEvalML2Listener) ExitEval(ctx *EvalContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseEvalML2Listener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseEvalML2Listener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterIfExpr is called when production IfExpr is entered.
func (s *BaseEvalML2Listener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production IfExpr is exited.
func (s *BaseEvalML2Listener) ExitIfExpr(ctx *IfExprContext) {}

// EnterLetExpr is called when production LetExpr is entered.
func (s *BaseEvalML2Listener) EnterLetExpr(ctx *LetExprContext) {}

// ExitLetExpr is called when production LetExpr is exited.
func (s *BaseEvalML2Listener) ExitLetExpr(ctx *LetExprContext) {}

// EnterVarExpr is called when production VarExpr is entered.
func (s *BaseEvalML2Listener) EnterVarExpr(ctx *VarExprContext) {}

// ExitVarExpr is called when production VarExpr is exited.
func (s *BaseEvalML2Listener) ExitVarExpr(ctx *VarExprContext) {}

// EnterBinOpExpr is called when production BinOpExpr is entered.
func (s *BaseEvalML2Listener) EnterBinOpExpr(ctx *BinOpExprContext) {}

// ExitBinOpExpr is called when production BinOpExpr is exited.
func (s *BaseEvalML2Listener) ExitBinOpExpr(ctx *BinOpExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseEvalML2Listener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseEvalML2Listener) ExitIntExpr(ctx *IntExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseEvalML2Listener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseEvalML2Listener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterNegIntExpr is called when production NegIntExpr is entered.
func (s *BaseEvalML2Listener) EnterNegIntExpr(ctx *NegIntExprContext) {}

// ExitNegIntExpr is called when production NegIntExpr is exited.
func (s *BaseEvalML2Listener) ExitNegIntExpr(ctx *NegIntExprContext) {}

// EnterDefList is called when production defList is entered.
func (s *BaseEvalML2Listener) EnterDefList(ctx *DefListContext) {}

// ExitDefList is called when production defList is exited.
func (s *BaseEvalML2Listener) ExitDefList(ctx *DefListContext) {}

// EnterDef is called when production def is entered.
func (s *BaseEvalML2Listener) EnterDef(ctx *DefContext) {}

// ExitDef is called when production def is exited.
func (s *BaseEvalML2Listener) ExitDef(ctx *DefContext) {}

// EnterIntValue is called when production IntValue is entered.
func (s *BaseEvalML2Listener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production IntValue is exited.
func (s *BaseEvalML2Listener) ExitIntValue(ctx *IntValueContext) {}

// EnterBoolValue is called when production BoolValue is entered.
func (s *BaseEvalML2Listener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production BoolValue is exited.
func (s *BaseEvalML2Listener) ExitBoolValue(ctx *BoolValueContext) {}
