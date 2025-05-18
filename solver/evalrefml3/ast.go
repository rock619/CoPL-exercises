package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalrefml3/parser"
)

type Value interface {
	isValue()
	String() string
	Clone() Value
}

type IntValue int

func (IntValue) isValue() {}

func (v IntValue) String() string {
	return strconv.Itoa(int(v))
}

func (v IntValue) Clone() Value {
	return v
}

type BoolValue bool

func (BoolValue) isValue() {}

func (v BoolValue) String() string {
	return strconv.FormatBool(bool(v))
}

func (v BoolValue) Clone() Value {
	return v
}

type LocValue string

func (LocValue) isValue() {}

func (v LocValue) String() string {
	return string(v)
}

func (v LocValue) Clone() Value {
	return v
}

type FunValue struct {
	Env Env
	Fun Fun
}

func (FunValue) isValue() {}

func (v FunValue) String() string {
	return fmt.Sprintf("(%s)[%s]", v.Env, v.Fun)
}

func (v FunValue) Clone() Value {
	return FunValue{
		Env: v.Env.Clone(),
		Fun: v.Fun.Clone(),
	}
}

type Fun struct {
	Param Var
	Body  Exp
}

func (f Fun) String() string {
	return fmt.Sprintf("fun %s -> %s", f.Param, f.Body)
}

func (f Fun) Clone() Fun {
	return Fun{
		Param: f.Param,
		Body:  f.Body.Clone(),
	}
}

type RecFunValue struct {
	Env Env
	Fun RecFun
}

func (RecFunValue) isValue() {}

func (v RecFunValue) String() string {
	return fmt.Sprintf("(%s)[%s]", v.Env, v.Fun)
}

func (v RecFunValue) Clone() Value {
	return RecFunValue{
		Env: v.Env.Clone(),
		Fun: v.Fun.Clone(),
	}
}

type RecFun struct {
	Name Var
	Fun  Fun
}

func (f RecFun) String() string {
	return fmt.Sprintf("rec %s = %s", f.Name, f.Fun)
}

func (f RecFun) Clone() RecFun {
	return RecFun{
		Name: f.Name,
		Fun:  f.Fun.Clone(),
	}
}

type Env []Bind

func (e Env) String() string {
	var binds []string
	for _, b := range e {
		binds = append(binds, b.String())
	}
	return strings.Join(binds, ", ")
}

func (e Env) Clone() Env {
	env := make(Env, len(e))
	for i, b := range e {
		env[i] = b.Clone()
	}
	return env
}

func (e Env) Get(v Var) (got Value, ok bool) {
	for _, bind := range slices.Backward(e) {
		if bind.Var == v {
			return bind.Value, true
		}
	}
	return nil, false
}

func (e *Env) Set(b Bind) {
	if b.Value == nil {
		panic("b.Value is nil: Var = " + string(b.Var))
	}
	*e = append(*e, b)
}

type Bind struct {
	Var   Var
	Value Value
}

func (b Bind) String() string {
	return fmt.Sprintf("%s = %s", b.Var, b.Value)
}

func (b Bind) Clone() Bind {
	return Bind{
		Var:   b.Var,
		Value: b.Value.Clone(),
	}
}

type Var string

type Store []Assign

func (s Store) Empty() bool {
	return len(s) == 0
}

func (s Store) String() string {
	var assigns []string
	for _, a := range s {
		assigns = append(assigns, a.String())
	}
	return strings.Join(assigns, ", ")
}

func (s Store) Clone() Store {
	store := make(Store, len(s))
	for i, a := range s {
		store[i] = a.Clone()
	}
	return store
}

func (s Store) Get(loc LocValue) (v Value, ok bool) {
	for _, assign := range s {
		if assign.Loc == loc {
			return assign.Value, true
		}
	}
	return nil, false
}

func (s *Store) Set(a Assign) {
	for i, assign := range *s {
		if assign.Loc == a.Loc {
			(*s)[i] = a
			return
		}
	}
	*s = append(*s, a)
}

type Assign struct {
	Loc   LocValue
	Value Value
}

func (a Assign) String() string {
	return fmt.Sprintf("%s = %s", a.Loc, a.Value)
}

func (a Assign) Clone() Assign {
	return Assign{
		Loc:   a.Loc,
		Value: a.Value.Clone(),
	}
}

type Exp interface {
	isExpression()
	String() string
	Clone() Exp
}

type IntExp int

func (IntExp) isExpression() {}

func (e IntExp) String() string {
	return strconv.Itoa(int(e))
}

func (e IntExp) Clone() Exp {
	return e
}

type BoolExp bool

func (BoolExp) isExpression() {}

func (e BoolExp) String() string {
	return strconv.FormatBool(bool(e))
}

func (e BoolExp) Clone() Exp {
	return e
}

type VarExp Var

func (VarExp) isExpression() {}

func (e VarExp) String() string {
	return string(e)
}

func (e VarExp) Clone() Exp {
	return e
}

type BinOpExp struct {
	Left  Exp
	Op    Op
	Right Exp
}

func (BinOpExp) isExpression() {}

func (e BinOpExp) String() string {
	left := e.Left.String()
	if parenExp, ok := e.Left.(ParenExp); ok {
		if inner, ok := parenExp.Inner.(BinOpExp); ok {
			if inner.Op.LT(e.Op) {
				left = fmt.Sprintf("(%s)", left)
			}
		}
	}
	right := e.Right.String()
	if parenExp, ok := e.Right.(ParenExp); ok {
		if inner, ok := parenExp.Inner.(BinOpExp); ok {
			if e.Op.LT(inner.Op) {
				right = fmt.Sprintf("(%s)", right)
			}
		}
	}
	return fmt.Sprintf("%s %s %s", left, e.Op, right)
}

func (e BinOpExp) Clone() Exp {
	return BinOpExp{
		Left:  e.Left.Clone(),
		Op:    e.Op,
		Right: e.Right.Clone(),
	}
}

type IfExp struct {
	Cond Exp
	Then Exp
	Else Exp
}

func (IfExp) isExpression() {}

func (e IfExp) String() string {
	return fmt.Sprintf("if %s then %s else %s", e.Cond, e.Then, e.Else)
}

func (e IfExp) Clone() Exp {
	return IfExp{
		Cond: e.Cond.Clone(),
		Then: e.Then.Clone(),
		Else: e.Else.Clone(),
	}
}

type LetExp struct {
	Var  Var
	Bind Exp
	Body Exp
}

func (LetExp) isExpression() {}

func (e LetExp) String() string {
	return fmt.Sprintf("let %s = %s in %s", e.Var, e.Bind, e.Body)
}

func (e LetExp) Clone() Exp {
	return LetExp{
		Var:  e.Var,
		Bind: e.Bind.Clone(),
		Body: e.Body.Clone(),
	}
}

type FunExp struct {
	Fun Fun
}

func (FunExp) isExpression() {}

func (e FunExp) String() string {
	return e.Fun.String()
}

func (e FunExp) Clone() Exp {
	return FunExp{
		Fun: e.Fun.Clone(),
	}
}

type AppExp struct {
	Fun Exp
	Arg Exp
}

func (AppExp) isExpression() {}

func (e AppExp) String() string {
	if _, ok := e.Arg.(BinOpExp); ok {
		return fmt.Sprintf("%s (%s)", e.Fun, e.Arg)
	}
	return fmt.Sprintf("%s %s", e.Fun, e.Arg)
}

func (e AppExp) Clone() Exp {
	return AppExp{
		Fun: e.Fun.Clone(),
		Arg: e.Arg.Clone(),
	}
}

type LetRecExp struct {
	RecFun RecFun
	Body   Exp
}

func (LetRecExp) isExpression() {}

func (e LetRecExp) String() string {
	return fmt.Sprintf("let %s in %s", e.RecFun, e.Body)
}

func (e LetRecExp) Clone() Exp {
	return LetRecExp{
		RecFun: e.RecFun.Clone(),
		Body:   e.Body.Clone(),
	}
}

type RefExp struct {
	Exp Exp
}

func (RefExp) isExpression() {}

func (e RefExp) String() string {
	return fmt.Sprintf("ref %s", e.Exp)
}

func (e RefExp) Clone() Exp {
	return RefExp{
		Exp: e.Exp.Clone(),
	}
}

type DerefExp struct {
	Exp Exp
}

func (DerefExp) isExpression() {}

func (e DerefExp) String() string {
	return fmt.Sprintf("!%s", e.Exp)
}

func (e DerefExp) Clone() Exp {
	return DerefExp{
		Exp: e.Exp.Clone(),
	}
}

type AssignExp struct {
	Left  Exp
	Right Exp
}

func (AssignExp) isExpression() {}

func (e AssignExp) String() string {
	return fmt.Sprintf("%s := %s", e.Left, e.Right)
}

func (e AssignExp) Clone() Exp {
	return AssignExp{
		Left:  e.Left.Clone(),
		Right: e.Right.Clone(),
	}
}

type ParenExp struct {
	Inner Exp
}

func (ParenExp) isExpression() {}

func (e ParenExp) String() string {
	return fmt.Sprintf("(%s)", e.Inner.String())
}

func (e ParenExp) Clone() Exp {
	return ParenExp{
		Inner: e.Inner.Clone(),
	}
}

type Op string

const (
	OpPlus  Op = "+"
	OpMinus Op = "-"
	OpTimes Op = "*"
	OpLT    Op = "<"
)

func (o Op) LT(other Op) bool {
	switch o {
	case OpLT:
		return slices.Contains([]Op{OpPlus, OpMinus, OpTimes}, other)
	case OpPlus, OpMinus:
		return other == OpTimes
	default:
		return false
	}
}

func NewOp(token antlr.Token) (Op, error) {
	switch token.GetTokenType() {
	case parser.EvalRefML3ParserPLUS:
		return OpPlus, nil
	case parser.EvalRefML3LexerMINUS:
		return OpMinus, nil
	case parser.EvalRefML3LexerTIMES:
		return OpTimes, nil
	case parser.EvalRefML3LexerLT:
		return OpLT, nil
	default:
		return "", fmt.Errorf("unknown operator: %s", token.GetText())
	}
}
