// Code generated from EvalRefML3.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalRefML3
import "github.com/antlr4-go/antlr/v4"

// BaseEvalRefML3Listener is a complete listener for a parse tree produced by EvalRefML3Parser.
type BaseEvalRefML3Listener struct{}

var _ EvalRefML3Listener = &BaseEvalRefML3Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEvalRefML3Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEvalRefML3Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEvalRefML3Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEvalRefML3Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEval is called when production eval is entered.
func (s *BaseEvalRefML3Listener) EnterEval(ctx *EvalContext) {}

// ExitEval is called when production eval is exited.
func (s *BaseEvalRefML3Listener) ExitEval(ctx *EvalContext) {}

// EnterIntValue is called when production IntValue is entered.
func (s *BaseEvalRefML3Listener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production IntValue is exited.
func (s *BaseEvalRefML3Listener) ExitIntValue(ctx *IntValueContext) {}

// EnterBoolValue is called when production BoolValue is entered.
func (s *BaseEvalRefML3Listener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production BoolValue is exited.
func (s *BaseEvalRefML3Listener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterLocValue is called when production LocValue is entered.
func (s *BaseEvalRefML3Listener) EnterLocValue(ctx *LocValueContext) {}

// ExitLocValue is called when production LocValue is exited.
func (s *BaseEvalRefML3Listener) ExitLocValue(ctx *LocValueContext) {}

// EnterFunValue is called when production FunValue is entered.
func (s *BaseEvalRefML3Listener) EnterFunValue(ctx *FunValueContext) {}

// ExitFunValue is called when production FunValue is exited.
func (s *BaseEvalRefML3Listener) ExitFunValue(ctx *FunValueContext) {}

// EnterRecFunValue is called when production RecFunValue is entered.
func (s *BaseEvalRefML3Listener) EnterRecFunValue(ctx *RecFunValueContext) {}

// ExitRecFunValue is called when production RecFunValue is exited.
func (s *BaseEvalRefML3Listener) ExitRecFunValue(ctx *RecFunValueContext) {}

// EnterEnv is called when production env is entered.
func (s *BaseEvalRefML3Listener) EnterEnv(ctx *EnvContext) {}

// ExitEnv is called when production env is exited.
func (s *BaseEvalRefML3Listener) ExitEnv(ctx *EnvContext) {}

// EnterBind is called when production bind is entered.
func (s *BaseEvalRefML3Listener) EnterBind(ctx *BindContext) {}

// ExitBind is called when production bind is exited.
func (s *BaseEvalRefML3Listener) ExitBind(ctx *BindContext) {}

// EnterStore is called when production store is entered.
func (s *BaseEvalRefML3Listener) EnterStore(ctx *StoreContext) {}

// ExitStore is called when production store is exited.
func (s *BaseEvalRefML3Listener) ExitStore(ctx *StoreContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseEvalRefML3Listener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseEvalRefML3Listener) ExitAssign(ctx *AssignContext) {}

// EnterBoolExp is called when production BoolExp is entered.
func (s *BaseEvalRefML3Listener) EnterBoolExp(ctx *BoolExpContext) {}

// ExitBoolExp is called when production BoolExp is exited.
func (s *BaseEvalRefML3Listener) ExitBoolExp(ctx *BoolExpContext) {}

// EnterRefExp is called when production RefExp is entered.
func (s *BaseEvalRefML3Listener) EnterRefExp(ctx *RefExpContext) {}

// ExitRefExp is called when production RefExp is exited.
func (s *BaseEvalRefML3Listener) ExitRefExp(ctx *RefExpContext) {}

// EnterDerefExp is called when production DerefExp is entered.
func (s *BaseEvalRefML3Listener) EnterDerefExp(ctx *DerefExpContext) {}

// ExitDerefExp is called when production DerefExp is exited.
func (s *BaseEvalRefML3Listener) ExitDerefExp(ctx *DerefExpContext) {}

// EnterFunExp is called when production FunExp is entered.
func (s *BaseEvalRefML3Listener) EnterFunExp(ctx *FunExpContext) {}

// ExitFunExp is called when production FunExp is exited.
func (s *BaseEvalRefML3Listener) ExitFunExp(ctx *FunExpContext) {}

// EnterBinOpExp is called when production BinOpExp is entered.
func (s *BaseEvalRefML3Listener) EnterBinOpExp(ctx *BinOpExpContext) {}

// ExitBinOpExp is called when production BinOpExp is exited.
func (s *BaseEvalRefML3Listener) ExitBinOpExp(ctx *BinOpExpContext) {}

// EnterLetRecExp is called when production LetRecExp is entered.
func (s *BaseEvalRefML3Listener) EnterLetRecExp(ctx *LetRecExpContext) {}

// ExitLetRecExp is called when production LetRecExp is exited.
func (s *BaseEvalRefML3Listener) ExitLetRecExp(ctx *LetRecExpContext) {}

// EnterAssignExp is called when production AssignExp is entered.
func (s *BaseEvalRefML3Listener) EnterAssignExp(ctx *AssignExpContext) {}

// ExitAssignExp is called when production AssignExp is exited.
func (s *BaseEvalRefML3Listener) ExitAssignExp(ctx *AssignExpContext) {}

// EnterIfExp is called when production IfExp is entered.
func (s *BaseEvalRefML3Listener) EnterIfExp(ctx *IfExpContext) {}

// ExitIfExp is called when production IfExp is exited.
func (s *BaseEvalRefML3Listener) ExitIfExp(ctx *IfExpContext) {}

// EnterAppExp is called when production AppExp is entered.
func (s *BaseEvalRefML3Listener) EnterAppExp(ctx *AppExpContext) {}

// ExitAppExp is called when production AppExp is exited.
func (s *BaseEvalRefML3Listener) ExitAppExp(ctx *AppExpContext) {}

// EnterParenExp is called when production ParenExp is entered.
func (s *BaseEvalRefML3Listener) EnterParenExp(ctx *ParenExpContext) {}

// ExitParenExp is called when production ParenExp is exited.
func (s *BaseEvalRefML3Listener) ExitParenExp(ctx *ParenExpContext) {}

// EnterLetExp is called when production LetExp is entered.
func (s *BaseEvalRefML3Listener) EnterLetExp(ctx *LetExpContext) {}

// ExitLetExp is called when production LetExp is exited.
func (s *BaseEvalRefML3Listener) ExitLetExp(ctx *LetExpContext) {}

// EnterVarExp is called when production VarExp is entered.
func (s *BaseEvalRefML3Listener) EnterVarExp(ctx *VarExpContext) {}

// ExitVarExp is called when production VarExp is exited.
func (s *BaseEvalRefML3Listener) ExitVarExp(ctx *VarExpContext) {}

// EnterIntExp is called when production IntExp is entered.
func (s *BaseEvalRefML3Listener) EnterIntExp(ctx *IntExpContext) {}

// ExitIntExp is called when production IntExp is exited.
func (s *BaseEvalRefML3Listener) ExitIntExp(ctx *IntExpContext) {}

// EnterFun is called when production fun is entered.
func (s *BaseEvalRefML3Listener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BaseEvalRefML3Listener) ExitFun(ctx *FunContext) {}

// EnterRecFun is called when production recFun is entered.
func (s *BaseEvalRefML3Listener) EnterRecFun(ctx *RecFunContext) {}

// ExitRecFun is called when production recFun is exited.
func (s *BaseEvalRefML3Listener) ExitRecFun(ctx *RecFunContext) {}
