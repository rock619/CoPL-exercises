package main

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/typingml4/parser"
)

type Expr interface {
	Children() []Expr
	AddChildren(...Expr)
	Env() *Environment
	SetEnv(env *Environment)
	Literal() string
	Type() Type
	SetType(Type)
	Rule() Rule
}

type BaseExpr struct {
	children []Expr
	env      *Environment
	literal  string
	typ      Type
	rule     Rule
}

func (e BaseExpr) Children() []Expr {
	return e.children
}

func (e BaseExpr) Env() *Environment {
	return e.env
}

func (e *BaseExpr) SetEnv(env *Environment) {
	e.env = env
}

func (e BaseExpr) Literal() string {
	return e.literal
}

func (e BaseExpr) Type() Type {
	return e.typ
}

func (e *BaseExpr) SetType(t Type) {
	e.typ = t
}

func (e BaseExpr) Rule() Rule {
	return e.rule
}

func (e *BaseExpr) AddChildren(children ...Expr) {
	e.children = append(e.children, children...)
}

// ParenExpr (e)
// e = Inner
type ParenExpr struct {
	*BaseExpr
	Inner Expr
}

// FunExpr fun x -> e
// x = Param, e = Body
type FunExpr struct {
	*BaseExpr
	Fun
}

type Fun struct {
	Param       string
	Body        Expr
	BodyLiteral string
	BodyCtx     parser.IExprContext
}

// AppExpr e1 e2
// e1 = Fun, e2 = Arg
type AppExpr struct {
	*BaseExpr
	Fun     Expr
	Arg     Expr
	Applied Expr
}

// MatchExpr match e1 with [] -> e2 | x :: y -> e3
// e1 = Matched, e2 = Empty, x = HeadVar, y = TailVar, e3 = Cons
type MatchExpr struct {
	*BaseExpr
	Matched Expr
	Empty   Expr
	HeadVar string
	TailVar string
	Cons    Expr
}

// EmptyListExpr []
type EmptyListExpr struct {
	*BaseExpr
}

// ConsExpr e1 :: e2
// e1 = Head, e2 = Tail
type ConsExpr struct {
	*BaseExpr
	Head Expr
	Tail Expr
}

// BinOpExpr e1 op e2
// e1 = Left, op = Op, e2 = Right
type BinOpExpr struct {
	*BaseExpr
	Left  Expr
	Op    antlr.Token
	Right Expr
}

// IfExpr if e1 then e2 else e3
// e1 = Cond, e2 = Then, e3 = Else
type IfExpr struct {
	*BaseExpr
	Cond Expr
	Then Expr
	Else Expr
}

// LetExpr let x = e1 in e2
// x = Var, e1 = Bind, e2 = Body
type LetExpr struct {
	*BaseExpr
	Var  string
	Bind Expr
	Body Expr
}

// LetRecExpr let rec x = fun y -> e1 in e2
// x = Name, y = Param, e1 = Fun, e2 = Body
type LetRecExpr struct {
	*BaseExpr
	RecFun
}

type RecFun struct {
	Name string
	Fun
}

type BoolExpr struct {
	*BaseExpr
}

type IntExpr struct {
	*BaseExpr
}

type VarExpr struct {
	*BaseExpr
	Name string
}

type Rule string

const (
	TInt    Rule = "T-Int"
	TBool   Rule = "T-Bool"
	TIf     Rule = "T-If"
	TPlus   Rule = "T-Plus"
	TMinus  Rule = "T-Minus"
	TTimes  Rule = "T-Times"
	TLt     Rule = "T-Lt"
	TVar    Rule = "T-Var"
	TLet    Rule = "T-Let"
	TFun    Rule = "T-Fun"
	TApp    Rule = "T-App"
	TLetRec Rule = "T-LetRec"
	TNil    Rule = "T-Nil"
	TCons   Rule = "T-Cons"
	TMatch  Rule = "T-Match"
)
