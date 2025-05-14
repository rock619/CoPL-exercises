// Code generated from EvalContML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // EvalContML4
import "github.com/antlr4-go/antlr/v4"

// BaseEvalContML4Listener is a complete listener for a parse tree produced by EvalContML4Parser.
type BaseEvalContML4Listener struct{}

var _ EvalContML4Listener = &BaseEvalContML4Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseEvalContML4Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseEvalContML4Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseEvalContML4Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseEvalContML4Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEval is called when production eval is entered.
func (s *BaseEvalContML4Listener) EnterEval(ctx *EvalContext) {}

// ExitEval is called when production eval is exited.
func (s *BaseEvalContML4Listener) ExitEval(ctx *EvalContext) {}

// EnterBoolValue is called when production BoolValue is entered.
func (s *BaseEvalContML4Listener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production BoolValue is exited.
func (s *BaseEvalContML4Listener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterFunValue is called when production FunValue is entered.
func (s *BaseEvalContML4Listener) EnterFunValue(ctx *FunValueContext) {}

// ExitFunValue is called when production FunValue is exited.
func (s *BaseEvalContML4Listener) ExitFunValue(ctx *FunValueContext) {}

// EnterRecFunValue is called when production RecFunValue is entered.
func (s *BaseEvalContML4Listener) EnterRecFunValue(ctx *RecFunValueContext) {}

// ExitRecFunValue is called when production RecFunValue is exited.
func (s *BaseEvalContML4Listener) ExitRecFunValue(ctx *RecFunValueContext) {}

// EnterContValue is called when production ContValue is entered.
func (s *BaseEvalContML4Listener) EnterContValue(ctx *ContValueContext) {}

// ExitContValue is called when production ContValue is exited.
func (s *BaseEvalContML4Listener) ExitContValue(ctx *ContValueContext) {}

// EnterIntValue is called when production IntValue is entered.
func (s *BaseEvalContML4Listener) EnterIntValue(ctx *IntValueContext) {}

// ExitIntValue is called when production IntValue is exited.
func (s *BaseEvalContML4Listener) ExitIntValue(ctx *IntValueContext) {}

// EnterNilValue is called when production NilValue is entered.
func (s *BaseEvalContML4Listener) EnterNilValue(ctx *NilValueContext) {}

// ExitNilValue is called when production NilValue is exited.
func (s *BaseEvalContML4Listener) ExitNilValue(ctx *NilValueContext) {}

// EnterConsValue is called when production ConsValue is entered.
func (s *BaseEvalContML4Listener) EnterConsValue(ctx *ConsValueContext) {}

// ExitConsValue is called when production ConsValue is exited.
func (s *BaseEvalContML4Listener) ExitConsValue(ctx *ConsValueContext) {}

// EnterEnv is called when production env is entered.
func (s *BaseEvalContML4Listener) EnterEnv(ctx *EnvContext) {}

// ExitEnv is called when production env is exited.
func (s *BaseEvalContML4Listener) ExitEnv(ctx *EnvContext) {}

// EnterBind is called when production bind is entered.
func (s *BaseEvalContML4Listener) EnterBind(ctx *BindContext) {}

// ExitBind is called when production bind is exited.
func (s *BaseEvalContML4Listener) ExitBind(ctx *BindContext) {}

// EnterBoolExp is called when production BoolExp is entered.
func (s *BaseEvalContML4Listener) EnterBoolExp(ctx *BoolExpContext) {}

// ExitBoolExp is called when production BoolExp is exited.
func (s *BaseEvalContML4Listener) ExitBoolExp(ctx *BoolExpContext) {}

// EnterConsExp is called when production ConsExp is entered.
func (s *BaseEvalContML4Listener) EnterConsExp(ctx *ConsExpContext) {}

// ExitConsExp is called when production ConsExp is exited.
func (s *BaseEvalContML4Listener) ExitConsExp(ctx *ConsExpContext) {}

// EnterFunExp is called when production FunExp is entered.
func (s *BaseEvalContML4Listener) EnterFunExp(ctx *FunExpContext) {}

// ExitFunExp is called when production FunExp is exited.
func (s *BaseEvalContML4Listener) ExitFunExp(ctx *FunExpContext) {}

// EnterBinOpExp is called when production BinOpExp is entered.
func (s *BaseEvalContML4Listener) EnterBinOpExp(ctx *BinOpExpContext) {}

// ExitBinOpExp is called when production BinOpExp is exited.
func (s *BaseEvalContML4Listener) ExitBinOpExp(ctx *BinOpExpContext) {}

// EnterLetRecExp is called when production LetRecExp is entered.
func (s *BaseEvalContML4Listener) EnterLetRecExp(ctx *LetRecExpContext) {}

// ExitLetRecExp is called when production LetRecExp is exited.
func (s *BaseEvalContML4Listener) ExitLetRecExp(ctx *LetRecExpContext) {}

// EnterLetCCExp is called when production LetCCExp is entered.
func (s *BaseEvalContML4Listener) EnterLetCCExp(ctx *LetCCExpContext) {}

// ExitLetCCExp is called when production LetCCExp is exited.
func (s *BaseEvalContML4Listener) ExitLetCCExp(ctx *LetCCExpContext) {}

// EnterMatchExp is called when production MatchExp is entered.
func (s *BaseEvalContML4Listener) EnterMatchExp(ctx *MatchExpContext) {}

// ExitMatchExp is called when production MatchExp is exited.
func (s *BaseEvalContML4Listener) ExitMatchExp(ctx *MatchExpContext) {}

// EnterIfExp is called when production IfExp is entered.
func (s *BaseEvalContML4Listener) EnterIfExp(ctx *IfExpContext) {}

// ExitIfExp is called when production IfExp is exited.
func (s *BaseEvalContML4Listener) ExitIfExp(ctx *IfExpContext) {}

// EnterAppExp is called when production AppExp is entered.
func (s *BaseEvalContML4Listener) EnterAppExp(ctx *AppExpContext) {}

// ExitAppExp is called when production AppExp is exited.
func (s *BaseEvalContML4Listener) ExitAppExp(ctx *AppExpContext) {}

// EnterParenExp is called when production ParenExp is entered.
func (s *BaseEvalContML4Listener) EnterParenExp(ctx *ParenExpContext) {}

// ExitParenExp is called when production ParenExp is exited.
func (s *BaseEvalContML4Listener) ExitParenExp(ctx *ParenExpContext) {}

// EnterLetExp is called when production LetExp is entered.
func (s *BaseEvalContML4Listener) EnterLetExp(ctx *LetExpContext) {}

// ExitLetExp is called when production LetExp is exited.
func (s *BaseEvalContML4Listener) ExitLetExp(ctx *LetExpContext) {}

// EnterVarExp is called when production VarExp is entered.
func (s *BaseEvalContML4Listener) EnterVarExp(ctx *VarExpContext) {}

// ExitVarExp is called when production VarExp is exited.
func (s *BaseEvalContML4Listener) ExitVarExp(ctx *VarExpContext) {}

// EnterIntExp is called when production IntExp is entered.
func (s *BaseEvalContML4Listener) EnterIntExp(ctx *IntExpContext) {}

// ExitIntExp is called when production IntExp is exited.
func (s *BaseEvalContML4Listener) ExitIntExp(ctx *IntExpContext) {}

// EnterNilExp is called when production NilExp is entered.
func (s *BaseEvalContML4Listener) EnterNilExp(ctx *NilExpContext) {}

// ExitNilExp is called when production NilExp is exited.
func (s *BaseEvalContML4Listener) ExitNilExp(ctx *NilExpContext) {}

// EnterFun is called when production fun is entered.
func (s *BaseEvalContML4Listener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BaseEvalContML4Listener) ExitFun(ctx *FunContext) {}

// EnterRecFun is called when production recFun is entered.
func (s *BaseEvalContML4Listener) EnterRecFun(ctx *RecFunContext) {}

// ExitRecFun is called when production recFun is exited.
func (s *BaseEvalContML4Listener) ExitRecFun(ctx *RecFunContext) {}

// EnterTerminalCont is called when production TerminalCont is entered.
func (s *BaseEvalContML4Listener) EnterTerminalCont(ctx *TerminalContContext) {}

// ExitTerminalCont is called when production TerminalCont is exited.
func (s *BaseEvalContML4Listener) ExitTerminalCont(ctx *TerminalContContext) {}

// EnterExpCont is called when production ExpCont is entered.
func (s *BaseEvalContML4Listener) EnterExpCont(ctx *ExpContContext) {}

// ExitExpCont is called when production ExpCont is exited.
func (s *BaseEvalContML4Listener) ExitExpCont(ctx *ExpContContext) {}

// EnterValueCont is called when production ValueCont is entered.
func (s *BaseEvalContML4Listener) EnterValueCont(ctx *ValueContContext) {}

// ExitValueCont is called when production ValueCont is exited.
func (s *BaseEvalContML4Listener) ExitValueCont(ctx *ValueContContext) {}

// EnterIfCont is called when production IfCont is entered.
func (s *BaseEvalContML4Listener) EnterIfCont(ctx *IfContContext) {}

// ExitIfCont is called when production IfCont is exited.
func (s *BaseEvalContML4Listener) ExitIfCont(ctx *IfContContext) {}

// EnterLetCont is called when production LetCont is entered.
func (s *BaseEvalContML4Listener) EnterLetCont(ctx *LetContContext) {}

// ExitLetCont is called when production LetCont is exited.
func (s *BaseEvalContML4Listener) ExitLetCont(ctx *LetContContext) {}

// EnterAppExpCont is called when production AppExpCont is entered.
func (s *BaseEvalContML4Listener) EnterAppExpCont(ctx *AppExpContContext) {}

// ExitAppExpCont is called when production AppExpCont is exited.
func (s *BaseEvalContML4Listener) ExitAppExpCont(ctx *AppExpContContext) {}

// EnterAppValueCont is called when production AppValueCont is entered.
func (s *BaseEvalContML4Listener) EnterAppValueCont(ctx *AppValueContContext) {}

// ExitAppValueCont is called when production AppValueCont is exited.
func (s *BaseEvalContML4Listener) ExitAppValueCont(ctx *AppValueContContext) {}

// EnterConsExpCont is called when production ConsExpCont is entered.
func (s *BaseEvalContML4Listener) EnterConsExpCont(ctx *ConsExpContContext) {}

// ExitConsExpCont is called when production ConsExpCont is exited.
func (s *BaseEvalContML4Listener) ExitConsExpCont(ctx *ConsExpContContext) {}

// EnterConsValueCont is called when production ConsValueCont is entered.
func (s *BaseEvalContML4Listener) EnterConsValueCont(ctx *ConsValueContContext) {}

// ExitConsValueCont is called when production ConsValueCont is exited.
func (s *BaseEvalContML4Listener) ExitConsValueCont(ctx *ConsValueContContext) {}

// EnterMatchCont is called when production MatchCont is entered.
func (s *BaseEvalContML4Listener) EnterMatchCont(ctx *MatchContContext) {}

// ExitMatchCont is called when production MatchCont is exited.
func (s *BaseEvalContML4Listener) ExitMatchCont(ctx *MatchContContext) {}
