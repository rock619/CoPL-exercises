package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalml3/parser"
)

type Expr interface {
	ID() int
	Children() []Expr
	AddChild(Expr)
	Ctx() parser.IExprContext
	Env() Env
	Literal() string
	Depth() int
	EvalTo() Value
	Rule() Rule
	Skip() bool
}

type BaseExpr struct {
	id       int
	children []Expr
	ctx      parser.IExprContext
	env      Env
	literal  string
	depth    int
	evalTo   Value
	rule     Rule
	skip     bool
}

func (e BaseExpr) Clone() *BaseExpr {
	return &BaseExpr{
		id:      e.id,
		ctx:     e.ctx,
		env:     slices.Clone(e.env),
		literal: e.literal,
		depth:   e.depth,
		evalTo:  e.evalTo,
		rule:    e.rule,
	}
}

func (e BaseExpr) ID() int {
	return e.id
}

func (e BaseExpr) Children() []Expr {
	return e.children
}

func (e BaseExpr) Ctx() parser.IExprContext {
	return e.ctx
}

func (e BaseExpr) Env() Env {
	return slices.Clone(e.env)
}

func (e BaseExpr) Literal() string {
	return e.literal
}

func (e BaseExpr) Depth() int {
	return e.depth
}

func (e BaseExpr) EvalTo() Value {
	return e.evalTo
}

func (e BaseExpr) Rule() Rule {
	return e.rule
}

func (e *BaseExpr) AddChild(child Expr) {
	e.children = append(e.children, child)
}

func (e *BaseExpr) Skip() bool {
	return e.skip
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
	Fun  Expr
	Arg  Expr
	Eval Expr
}

// FunExpr fun x -> e
// x = Param, e = Body
type FunExpr struct {
	*BaseExpr
	Fun
}

type Fun struct {
	Param string
	Body  Expr
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

func NewBoolValue(tok antlr.Token) (BoolValue, error) {
	if tok.GetTokenType() != parser.EvalML3ParserBOOL {
		return false, fmt.Errorf("expected BOOL token, but got %s", tok.GetText())
	}
	b, err := strconv.ParseBool(tok.GetText())
	if err != nil {
		return false, fmt.Errorf("parse BOOL token: %s", tok.GetText())
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
	return fmt.Sprintf("(%s)[fun %s -> %s]", strings.Join(binds, ", "), f.Param, f.Body.Literal())
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
	return fmt.Sprintf("(%s)[rec %s = fun %s -> %s]", strings.Join(binds, ", "), f.Name, f.Param, f.Body.Literal())
}

type VarValue struct {
	Name string
}

func (VarValue) value() {}

func (v VarValue) String() string {
	return v.Name
}

type BExpr struct {
	*BaseExpr
}

func NewBExpr(left, right IntValue, op antlr.Token) BExpr {
	switch op.GetTokenType() {
	case parser.EvalML3ParserPLUS:
		return BExpr{
			BaseExpr: &BaseExpr{
				literal: fmt.Sprintf("%d plus %d is %d", left, right, left+right),
				rule:    BPlus,
			},
		}
	case parser.EvalML3ParserMINUS:
		return BExpr{
			BaseExpr: &BaseExpr{
				literal: fmt.Sprintf("%d minus %d is %d", left, right, left-right),
				rule:    BMinus,
			},
		}
	case parser.EvalML3ParserTIMES:
		return BExpr{
			BaseExpr: &BaseExpr{
				literal: fmt.Sprintf("%d times %d is %d", left, right, left*right),
				rule:    BTimes,
			},
		}
	case parser.EvalML3ParserLT:
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
	case parser.EvalML3ParserPLUS:
		return EPlus
	case parser.EvalML3ParserMINUS:
		return EMinus
	case parser.EvalML3ParserTIMES:
		return ETimes
	case parser.EvalML3ParserLT:
		return ELt
	default:
		panic("Unknown op: " + op.GetText())
	}
}

type Rule string

const (
	EInt    Rule = "E-Int"
	EBool   Rule = "E-Bool"
	EVar1   Rule = "E-Var1"
	EVar2   Rule = "E-Var2"
	EPlus   Rule = "E-Plus"
	EMinus  Rule = "E-Minus"
	ETimes  Rule = "E-Times"
	ELt     Rule = "E-Lt"
	EIfT    Rule = "E-IfT"
	EIfF    Rule = "E-IfF"
	ELet    Rule = "E-Let"
	EFun    Rule = "E-Fun"
	EApp    Rule = "E-App"
	ELetRec Rule = "E-LetRec"
	EAppRec Rule = "E-AppRec"

	BPlus  Rule = "B-Plus"
	BMinus Rule = "B-Minus"
	BTimes Rule = "B-Times"
	BLt    Rule = "B-Lt"

	// 評価するまで判断できないため仮で設定するRule
	EIf    Rule = "E-If?"
	EVar   Rule = "E-Var?"
	EBinOp Rule = "E-BinOp?"
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
