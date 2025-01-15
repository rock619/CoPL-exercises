package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalml4/parser"
)

type Expr interface {
	Children() []Expr
	AddChildren(...Expr)
	Env() Env
	Literal() string
	EvalTo() Value
	Rule() Rule
}

type BaseExpr struct {
	children []Expr
	env      Env
	literal  string
	evalTo   Value
	rule     Rule
}

func (e BaseExpr) Clone() *BaseExpr {
	return &BaseExpr{
		env:     slices.Clone(e.env),
		literal: e.literal,
		evalTo:  e.evalTo,
		rule:    e.rule,
	}
}

func (e BaseExpr) Children() []Expr {
	return e.children
}

func (e BaseExpr) Env() Env {
	return slices.Clone(e.env)
}

func (e BaseExpr) Literal() string {
	return e.literal
}

func (e BaseExpr) EvalTo() Value {
	return e.evalTo
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

// AppExpr e1 e2
// e1 = Fun, e2 = Arg
type AppExpr struct {
	*BaseExpr
	Fun     Expr
	Arg     Expr
	Applied Expr
}

// FunExpr fun x -> e
// x = Param, e = Body
type FunExpr struct {
	*BaseExpr
	Fun
}

type Fun struct {
	Param       string
	BodyLiteral string
	BodyCtx     parser.IExprContext
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
	Val BoolValue
}

type IntExpr struct {
	*BaseExpr
	Val IntValue
}

type VarExpr struct {
	*BaseExpr
	Name string
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

type Env []Bind

func (e Env) String() string {
	binds := make([]string, len(e))
	for i, b := range e {
		binds[i] = b.String()
	}
	return strings.Join(binds, ", ")
}

// Bind x = v
// x = Name, v = Val
type Bind struct {
	Name string
	Val  Value
}

func (b Bind) String() string {
	return fmt.Sprintf("%s = %s", b.Name, b.Val)
}

type Value interface {
	value()
	String() string
}

type IntValue int

func (IntValue) value() {}

func (i IntValue) String() string {
	return strconv.Itoa(int(i))
}

func NewIntValue(literal string) (IntValue, error) {
	i, err := strconv.Atoi(literal)
	if err != nil {
		return 0, fmt.Errorf("parse INT literal: %s", literal)
	}
	return IntValue(i), nil
}

type BoolValue bool

func (BoolValue) value() {}

func (b BoolValue) String() string {
	return strconv.FormatBool(bool(b))
}

func NewBoolValue(literal string) (BoolValue, error) {
	b, err := strconv.ParseBool(literal)
	if err != nil {
		return false, fmt.Errorf("parse BOOL literal: %s", literal)
	}
	return BoolValue(b), nil
}

// FunValue (ℰ)[fun x -> e]
// ℰ = Env, x = Param, e = Body
type FunValue struct {
	Fun
	Env Env
}

func (FunValue) value() {}

func (f FunValue) String() string {
	binds := make([]string, len(f.Env))
	for i, b := range f.Env {
		binds[i] = b.String()
	}
	return fmt.Sprintf("(%s)[fun %s -> %s]", strings.Join(binds, ", "), f.Param, f.BodyLiteral)
}

// RecFunValue (ℰ)[rec fun x = y -> e]
// ℰ = Env, x = Name, y = Param, e = Body
type RecFunValue struct {
	RecFun
	Env Env
}

func (RecFunValue) value() {}

func (f RecFunValue) String() string {
	binds := make([]string, len(f.Env))
	for i, b := range f.Env {
		binds[i] = fmt.Sprintf("%s = %s", b.Name, b.Val)
	}
	return fmt.Sprintf("(%s)[rec %s = fun %s -> %s]", strings.Join(binds, ", "), f.Name, f.Param, f.BodyLiteral)
}

type VarValue struct {
	Name string
}

func (VarValue) value() {}

func (v VarValue) String() string {
	return v.Name
}

type EmptyListValue struct{}

func (EmptyListValue) value() {}

func (EmptyListValue) String() string {
	return "[]"
}

type ConsValue struct {
	Head Value
	Tail Value
}

func (ConsValue) value() {}

func (c ConsValue) String() string {
	// Cons :: は右結合のため、HeadがConsValueの場合括弧で囲まないと評価順が変わってしまう
	if _, ok := c.Head.(ConsValue); ok {
		return fmt.Sprintf("(%s) :: %s", c.Head, c.Tail)
	}
	return fmt.Sprintf("%s :: %s", c.Head, c.Tail)
}

type BExpr struct {
	*BaseExpr
}

func NewBExpr(left, right IntValue, op antlr.Token) BExpr {
	switch op.GetTokenType() {
	case parser.EvalML4ParserPLUS:
		return BExpr{
			BaseExpr: &BaseExpr{
				literal: fmt.Sprintf("%d plus %d is %d", left, right, left+right),
				rule:    BPlus,
			},
		}
	case parser.EvalML4ParserMINUS:
		return BExpr{
			BaseExpr: &BaseExpr{
				literal: fmt.Sprintf("%d minus %d is %d", left, right, left-right),
				rule:    BMinus,
			},
		}
	case parser.EvalML4ParserTIMES:
		return BExpr{
			BaseExpr: &BaseExpr{
				literal: fmt.Sprintf("%d times %d is %d", left, right, left*right),
				rule:    BTimes,
			},
		}
	case parser.EvalML4ParserLT:
		return BExpr{
			BaseExpr: &BaseExpr{
				literal: fmt.Sprintf("%d less than %d is %t", left, right, left < right),
				rule:    BLt,
			},
		}
	default:
		panic("Unknown op: " + op.GetText())
	}
}

func NewERuleFromOp(op antlr.Token) Rule {
	switch op.GetTokenType() {
	case parser.EvalML4ParserPLUS:
		return EPlus
	case parser.EvalML4ParserMINUS:
		return EMinus
	case parser.EvalML4ParserTIMES:
		return ETimes
	case parser.EvalML4ParserLT:
		return ELt
	default:
		panic("Unknown op: " + op.GetText())
	}
}

type Rule string

const (
	EInt       Rule = "E-Int"
	EBool      Rule = "E-Bool"
	EVar       Rule = "E-Var"
	EPlus      Rule = "E-Plus"
	EMinus     Rule = "E-Minus"
	ETimes     Rule = "E-Times"
	ELt        Rule = "E-Lt"
	EIfT       Rule = "E-IfT"
	EIfF       Rule = "E-IfF"
	ELet       Rule = "E-Let"
	EFun       Rule = "E-Fun"
	EApp       Rule = "E-App"
	ELetRec    Rule = "E-LetRec"
	EAppRec    Rule = "E-AppRec"
	ENil       Rule = "E-Nil"
	ECons      Rule = "E-Cons"
	EMatchNil  Rule = "E-MatchNil"
	EMatchCons Rule = "E-MatchCons"

	BPlus  Rule = "B-Plus"
	BMinus Rule = "B-Minus"
	BTimes Rule = "B-Times"
	BLt    Rule = "B-Lt"
)

func (r Rule) B() bool {
	return slices.Contains([]Rule{BPlus, BMinus, BTimes, BLt}, r)
}

func getLiteral(p antlr.Parser, st antlr.SyntaxTree) string {
	ts := p.GetTokenStream()
	tx := ts.GetTextFromInterval(st.GetSourceInterval())
	return normalizeSpaces(tx)
}

func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
