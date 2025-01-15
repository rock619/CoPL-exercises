// Code generated from EvalML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalML1
import "github.com/antlr4-go/antlr/v4"

// BaseEvalML1Listener is a complete listener for a parse tree produced by EvalML1Parser.
type BaseEvalML1Listener struct{}

var _ EvalML1Listener = &BaseEvalML1Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEvalML1Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEvalML1Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEvalML1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEvalML1Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuestion is called when production question is entered.
func (s *BaseEvalML1Listener) EnterQuestion(ctx *QuestionContext) {}

// ExitQuestion is called when production question is exited.
func (s *BaseEvalML1Listener) ExitQuestion(ctx *QuestionContext) {}

// EnterEval is called when production eval is entered.
func (s *BaseEvalML1Listener) EnterEval(ctx *EvalContext) {}

// ExitEval is called when production eval is exited.
func (s *BaseEvalML1Listener) ExitEval(ctx *EvalContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseEvalML1Listener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseEvalML1Listener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterIfExpr is called when production IfExpr is entered.
func (s *BaseEvalML1Listener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production IfExpr is exited.
func (s *BaseEvalML1Listener) ExitIfExpr(ctx *IfExprContext) {}

// EnterBinOpExpr is called when production BinOpExpr is entered.
func (s *BaseEvalML1Listener) EnterBinOpExpr(ctx *BinOpExprContext) {}

// ExitBinOpExpr is called when production BinOpExpr is exited.
func (s *BaseEvalML1Listener) ExitBinOpExpr(ctx *BinOpExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseEvalML1Listener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseEvalML1Listener) ExitIntExpr(ctx *IntExprContext) {}

// EnterParenExpr is called when production ParenExpr is entered.
func (s *BaseEvalML1Listener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production ParenExpr is exited.
func (s *BaseEvalML1Listener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterNegIntExpr is called when production NegIntExpr is entered.
func (s *BaseEvalML1Listener) EnterNegIntExpr(ctx *NegIntExprContext) {}

// ExitNegIntExpr is called when production NegIntExpr is exited.
func (s *BaseEvalML1Listener) ExitNegIntExpr(ctx *NegIntExprContext) {}

// EnterIntValue is called when production IntValue is entered.
func (s *BaseEvalML1Listener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production IntValue is exited.
func (s *BaseEvalML1Listener) ExitIntValue(ctx *IntValueContext) {}

// EnterBoolValue is called when production BoolValue is entered.
func (s *BaseEvalML1Listener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production BoolValue is exited.
func (s *BaseEvalML1Listener) ExitBoolValue(ctx *BoolValueContext) {}
