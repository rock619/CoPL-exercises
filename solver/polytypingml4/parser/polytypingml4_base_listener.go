// Code generated from PolyTypingML4.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // PolyTypingML4
import "github.com/antlr4-go/antlr/v4"

// BasePolyTypingML4Listener is a complete listener for a parse tree produced by PolyTypingML4Parser.
type BasePolyTypingML4Listener struct{}

var _ PolyTypingML4Listener = &BasePolyTypingML4Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePolyTypingML4Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePolyTypingML4Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePolyTypingML4Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePolyTypingML4Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterEval is called when production eval is entered.
func (s *BasePolyTypingML4Listener) EnterEval(ctx *EvalContext) {}

// ExitEval is called when production eval is exited.
func (s *BasePolyTypingML4Listener) ExitEval(ctx *EvalContext) {}

// EnterVar is called when production var is entered.
func (s *BasePolyTypingML4Listener) EnterVar(ctx *VarContext) {}

// ExitVar is called when production var is exited.
func (s *BasePolyTypingML4Listener) ExitVar(ctx *VarContext) {}

// EnterTVar is called when production tVar is entered.
func (s *BasePolyTypingML4Listener) EnterTVar(ctx *TVarContext) {}

// ExitTVar is called when production tVar is exited.
func (s *BasePolyTypingML4Listener) ExitTVar(ctx *TVarContext) {}

// EnterBoolType is called when production BoolType is entered.
func (s *BasePolyTypingML4Listener) EnterBoolType(ctx *BoolTypeContext) {}

// ExitBoolType is called when production BoolType is exited.
func (s *BasePolyTypingML4Listener) ExitBoolType(ctx *BoolTypeContext) {}

// EnterTypeVar is called when production TypeVar is entered.
func (s *BasePolyTypingML4Listener) EnterTypeVar(ctx *TypeVarContext) {}

// ExitTypeVar is called when production TypeVar is exited.
func (s *BasePolyTypingML4Listener) ExitTypeVar(ctx *TypeVarContext) {}

// EnterListType is called when production ListType is entered.
func (s *BasePolyTypingML4Listener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production ListType is exited.
func (s *BasePolyTypingML4Listener) ExitListType(ctx *ListTypeContext) {}

// EnterFunType is called when production FunType is entered.
func (s *BasePolyTypingML4Listener) EnterFunType(ctx *FunTypeContext) {}

// ExitFunType is called when production FunType is exited.
func (s *BasePolyTypingML4Listener) ExitFunType(ctx *FunTypeContext) {}

// EnterParenType is called when production ParenType is entered.
func (s *BasePolyTypingML4Listener) EnterParenType(ctx *ParenTypeContext) {}

// ExitParenType is called when production ParenType is exited.
func (s *BasePolyTypingML4Listener) ExitParenType(ctx *ParenTypeContext) {}

// EnterIntType is called when production IntType is entered.
func (s *BasePolyTypingML4Listener) EnterIntType(ctx *IntTypeContext) {}

// ExitIntType is called when production IntType is exited.
func (s *BasePolyTypingML4Listener) ExitIntType(ctx *IntTypeContext) {}

// EnterTyScheme is called when production tyScheme is entered.
func (s *BasePolyTypingML4Listener) EnterTyScheme(ctx *TySchemeContext) {}

// ExitTyScheme is called when production tyScheme is exited.
func (s *BasePolyTypingML4Listener) ExitTyScheme(ctx *TySchemeContext) {}

// EnterEnv is called when production env is entered.
func (s *BasePolyTypingML4Listener) EnterEnv(ctx *EnvContext) {}

// ExitEnv is called when production env is exited.
func (s *BasePolyTypingML4Listener) ExitEnv(ctx *EnvContext) {}

// EnterBind is called when production bind is entered.
func (s *BasePolyTypingML4Listener) EnterBind(ctx *BindContext) {}

// ExitBind is called when production bind is exited.
func (s *BasePolyTypingML4Listener) ExitBind(ctx *BindContext) {}

// EnterBoolExp is called when production BoolExp is entered.
func (s *BasePolyTypingML4Listener) EnterBoolExp(ctx *BoolExpContext) {}

// ExitBoolExp is called when production BoolExp is exited.
func (s *BasePolyTypingML4Listener) ExitBoolExp(ctx *BoolExpContext) {}

// EnterConsExp is called when production ConsExp is entered.
func (s *BasePolyTypingML4Listener) EnterConsExp(ctx *ConsExpContext) {}

// ExitConsExp is called when production ConsExp is exited.
func (s *BasePolyTypingML4Listener) ExitConsExp(ctx *ConsExpContext) {}

// EnterFunExp is called when production FunExp is entered.
func (s *BasePolyTypingML4Listener) EnterFunExp(ctx *FunExpContext) {}

// ExitFunExp is called when production FunExp is exited.
func (s *BasePolyTypingML4Listener) ExitFunExp(ctx *FunExpContext) {}

// EnterBinOpExp is called when production BinOpExp is entered.
func (s *BasePolyTypingML4Listener) EnterBinOpExp(ctx *BinOpExpContext) {}

// ExitBinOpExp is called when production BinOpExp is exited.
func (s *BasePolyTypingML4Listener) ExitBinOpExp(ctx *BinOpExpContext) {}

// EnterLetRecExp is called when production LetRecExp is entered.
func (s *BasePolyTypingML4Listener) EnterLetRecExp(ctx *LetRecExpContext) {}

// ExitLetRecExp is called when production LetRecExp is exited.
func (s *BasePolyTypingML4Listener) ExitLetRecExp(ctx *LetRecExpContext) {}

// EnterIfExp is called when production IfExp is entered.
func (s *BasePolyTypingML4Listener) EnterIfExp(ctx *IfExpContext) {}

// ExitIfExp is called when production IfExp is exited.
func (s *BasePolyTypingML4Listener) ExitIfExp(ctx *IfExpContext) {}

// EnterMatchExp is called when production MatchExp is entered.
func (s *BasePolyTypingML4Listener) EnterMatchExp(ctx *MatchExpContext) {}

// ExitMatchExp is called when production MatchExp is exited.
func (s *BasePolyTypingML4Listener) ExitMatchExp(ctx *MatchExpContext) {}

// EnterAppExp is called when production AppExp is entered.
func (s *BasePolyTypingML4Listener) EnterAppExp(ctx *AppExpContext) {}

// ExitAppExp is called when production AppExp is exited.
func (s *BasePolyTypingML4Listener) ExitAppExp(ctx *AppExpContext) {}

// EnterParenExp is called when production ParenExp is entered.
func (s *BasePolyTypingML4Listener) EnterParenExp(ctx *ParenExpContext) {}

// ExitParenExp is called when production ParenExp is exited.
func (s *BasePolyTypingML4Listener) ExitParenExp(ctx *ParenExpContext) {}

// EnterLetExp is called when production LetExp is entered.
func (s *BasePolyTypingML4Listener) EnterLetExp(ctx *LetExpContext) {}

// ExitLetExp is called when production LetExp is exited.
func (s *BasePolyTypingML4Listener) ExitLetExp(ctx *LetExpContext) {}

// EnterVarExp is called when production VarExp is entered.
func (s *BasePolyTypingML4Listener) EnterVarExp(ctx *VarExpContext) {}

// ExitVarExp is called when production VarExp is exited.
func (s *BasePolyTypingML4Listener) ExitVarExp(ctx *VarExpContext) {}

// EnterIntExp is called when production IntExp is entered.
func (s *BasePolyTypingML4Listener) EnterIntExp(ctx *IntExpContext) {}

// ExitIntExp is called when production IntExp is exited.
func (s *BasePolyTypingML4Listener) ExitIntExp(ctx *IntExpContext) {}

// EnterNilExp is called when production NilExp is entered.
func (s *BasePolyTypingML4Listener) EnterNilExp(ctx *NilExpContext) {}

// ExitNilExp is called when production NilExp is exited.
func (s *BasePolyTypingML4Listener) ExitNilExp(ctx *NilExpContext) {}

// EnterFun is called when production fun is entered.
func (s *BasePolyTypingML4Listener) EnterFun(ctx *FunContext) {}

// ExitFun is called when production fun is exited.
func (s *BasePolyTypingML4Listener) ExitFun(ctx *FunContext) {}
