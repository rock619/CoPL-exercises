// Code generated from EvalContML1.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalContML1
import "github.com/antlr4-go/antlr/v4"

// BaseEvalContML1Listener is a complete listener for a parse tree produced by EvalContML1Parser.
type BaseEvalContML1Listener struct{}

var _ EvalContML1Listener = &BaseEvalContML1Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEvalContML1Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEvalContML1Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEvalContML1Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEvalContML1Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEval is called when production eval is entered.
func (s *BaseEvalContML1Listener) EnterEval(ctx *EvalContext) {}

// ExitEval is called when production eval is exited.
func (s *BaseEvalContML1Listener) ExitEval(ctx *EvalContext) {}

// EnterIntValue is called when production IntValue is entered.
func (s *BaseEvalContML1Listener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production IntValue is exited.
func (s *BaseEvalContML1Listener) ExitIntValue(ctx *IntValueContext) {}

// EnterBoolValue is called when production BoolValue is entered.
func (s *BaseEvalContML1Listener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production BoolValue is exited.
func (s *BaseEvalContML1Listener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterBoolExp is called when production BoolExp is entered.
func (s *BaseEvalContML1Listener) EnterBoolExp(ctx *BoolExpContext) {}

// ExitBoolExp is called when production BoolExp is exited.
func (s *BaseEvalContML1Listener) ExitBoolExp(ctx *BoolExpContext) {}

// EnterIfExp is called when production IfExp is entered.
func (s *BaseEvalContML1Listener) EnterIfExp(ctx *IfExpContext) {}

// ExitIfExp is called when production IfExp is exited.
func (s *BaseEvalContML1Listener) ExitIfExp(ctx *IfExpContext) {}

// EnterParenExp is called when production ParenExp is entered.
func (s *BaseEvalContML1Listener) EnterParenExp(ctx *ParenExpContext) {}

// ExitParenExp is called when production ParenExp is exited.
func (s *BaseEvalContML1Listener) ExitParenExp(ctx *ParenExpContext) {}

// EnterIntExp is called when production IntExp is entered.
func (s *BaseEvalContML1Listener) EnterIntExp(ctx *IntExpContext) {}

// ExitIntExp is called when production IntExp is exited.
func (s *BaseEvalContML1Listener) ExitIntExp(ctx *IntExpContext) {}

// EnterBinOpExp is called when production BinOpExp is entered.
func (s *BaseEvalContML1Listener) EnterBinOpExp(ctx *BinOpExpContext) {}

// ExitBinOpExp is called when production BinOpExp is exited.
func (s *BaseEvalContML1Listener) ExitBinOpExp(ctx *BinOpExpContext) {}

// EnterUnaryCont is called when production UnaryCont is entered.
func (s *BaseEvalContML1Listener) EnterUnaryCont(ctx *UnaryContContext) {}

// ExitUnaryCont is called when production UnaryCont is exited.
func (s *BaseEvalContML1Listener) ExitUnaryCont(ctx *UnaryContContext) {}

// EnterExpCont is called when production ExpCont is entered.
func (s *BaseEvalContML1Listener) EnterExpCont(ctx *ExpContContext) {}

// ExitExpCont is called when production ExpCont is exited.
func (s *BaseEvalContML1Listener) ExitExpCont(ctx *ExpContContext) {}

// EnterValueCont is called when production ValueCont is entered.
func (s *BaseEvalContML1Listener) EnterValueCont(ctx *ValueContContext) {}

// ExitValueCont is called when production ValueCont is exited.
func (s *BaseEvalContML1Listener) ExitValueCont(ctx *ValueContContext) {}

// EnterIfCont is called when production IfCont is entered.
func (s *BaseEvalContML1Listener) EnterIfCont(ctx *IfContContext) {}

// ExitIfCont is called when production IfCont is exited.
func (s *BaseEvalContML1Listener) ExitIfCont(ctx *IfContContext) {}
