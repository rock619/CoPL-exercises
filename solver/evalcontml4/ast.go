package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalcontml4/parser"
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
	Param string
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
	Name string
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

type NilValue struct{}

func (NilValue) isValue() {}

func (NilValue) String() string {
	return "[]"
}

func (v NilValue) Clone() Value {
	return v
}

type ConsValue struct {
	Head Value
	Tail Value
}

func (ConsValue) isValue() {}

func (v ConsValue) String() string {
	return fmt.Sprintf("%s :: %s", v.Head, v.Tail)
}

func (v ConsValue) Clone() Value {
	return ConsValue{
		Head: v.Head.Clone(),
		Tail: v.Tail.Clone(),
	}
}

type ContValue struct {
	Cont Cont
}

func (ContValue) isValue() {}

func (v ContValue) String() string {
	if IsTerminal(v.Cont) {
		return "[_]"
	}
	var ss []string
	for c := v.Cont; c != nil && !IsTerminal(c); c = c.Next() {
		ss = append(ss, c.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(ss, " >> "))
}

func (v ContValue) Clone() Value {
	return ContValue{
		Cont: v.Cont.Clone(),
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

type Bind struct {
	Var   string
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

type VarExp string

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
	Var  string
	Bind Exp
	Body Exp
}

func (LetExp) isExpression() {}

func (e LetExp) String() string {
	return fmt.Sprintf("let %s in %s", e.Bind, e.Body)
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

type NilExp struct{}

func (NilExp) isExpression() {}

func (e NilExp) String() string {
	return "[]"
}

func (e NilExp) Clone() Exp {
	return e
}

type ConsExp struct {
	Head Exp
	Tail Exp
}

func (ConsExp) isExpression() {}

func (e ConsExp) String() string {
	return fmt.Sprintf("%s :: %s", e.Head, e.Tail)
}

func (e ConsExp) Clone() Exp {
	return ConsExp{
		Head: e.Head.Clone(),
		Tail: e.Tail.Clone(),
	}
}

type MatchExp struct {
	Matched  Exp
	NilCase  Exp
	HeadVar  string
	TailVar  string
	ConsCase Exp
}

func (MatchExp) isExpression() {}

func (e MatchExp) String() string {
	return fmt.Sprintf("match %s with [] -> %s | %s :: %s -> %s", e.Matched, e.NilCase, e.HeadVar, e.TailVar, e.ConsCase)
}

func (e MatchExp) Clone() Exp {
	return MatchExp{
		Matched:  e.Matched.Clone(),
		NilCase:  e.NilCase.Clone(),
		HeadVar:  e.HeadVar,
		TailVar:  e.TailVar,
		ConsCase: e.ConsCase.Clone(),
	}
}

type LetCCExp struct {
	Var  string
	Body Exp
}

func (LetCCExp) isExpression() {}

func (e LetCCExp) String() string {
	return fmt.Sprintf("letcc %s in %s", e.Var, e.Body)
}

func (e LetCCExp) Clone() Exp {
	return LetCCExp{
		Var:  e.Var,
		Body: e.Body.Clone(),
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

type Cont interface {
	isContinuation()
	Next() Cont
	String() string
	Clone() Cont
}

func IsTerminal(c Cont) bool {
	_, ok := c.(TerminalCont)
	return ok
}

type TerminalCont struct{}

func (TerminalCont) isContinuation() {}

func (c TerminalCont) Next() Cont {
	return nil
}

func (c TerminalCont) String() string {
	return "_"
}

func (c TerminalCont) Clone() Cont {
	return c
}

type BinOpExpCont struct {
	Env   Env
	Op    Op
	Right Exp
	next  Cont
}

func (BinOpExpCont) isContinuation() {}

func (c BinOpExpCont) Next() Cont {
	return c.next
}

func (c BinOpExpCont) String() string {
	right := c.Right.String()
	if parenExp, ok := c.Right.(ParenExp); ok {
		if inner, ok := parenExp.Inner.(BinOpExp); ok {
			if inner.Op.LT(c.Op) {
				right = fmt.Sprintf("(%s)", right)
			}
		}
	}
	env := c.Env.String()
	if env != "" {
		env += " "
	}
	return fmt.Sprintf("{%s _ %s %s}", envPrefix(c.Env), c.Op, right)
}

func (c BinOpExpCont) Clone() Cont {
	return BinOpExpCont{
		Env:   c.Env.Clone(),
		Op:    c.Op,
		Right: c.Right.Clone(),
		next:  c.next.Clone(),
	}
}

func envPrefix(env Env) string {
	e := env.String()
	if e == "" {
		return "|-"
	}
	return fmt.Sprintf("%s |-", e)
}

type BinOpValueCont struct {
	Left Value
	Op   Op
	next Cont
}

func (BinOpValueCont) isContinuation() {}

func (c BinOpValueCont) Next() Cont {
	return c.next
}

func (c BinOpValueCont) String() string {
	return fmt.Sprintf("{%s %s _}", c.Left, c.Op)
}

func (c BinOpValueCont) Clone() Cont {
	return BinOpValueCont{
		Left: c.Left.Clone(),
		Op:   c.Op,
		next: c.next.Clone(),
	}
}

type IfCont struct {
	Env  Env
	Then Exp
	Else Exp
	next Cont
}

func (IfCont) isContinuation() {}

func (c IfCont) Next() Cont {
	return c.next
}

func (c IfCont) String() string {
	return fmt.Sprintf("{%s if _ then %s else %s}", envPrefix(c.Env), c.Then, c.Else)
}

func (c IfCont) Clone() Cont {
	return IfCont{
		Env:  c.Env.Clone(),
		Then: c.Then.Clone(),
		Else: c.Else.Clone(),
		next: c.next.Clone(),
	}
}

type LetCont struct {
	Env  Env
	Var  string
	Body Exp
	next Cont
}

func (LetCont) isContinuation() {}

func (c LetCont) Next() Cont {
	return c.next
}

func (c LetCont) String() string {
	return fmt.Sprintf("{%s let %s = _ in %s}", envPrefix(c.Env), c.Var, c.Body)
}

func (c LetCont) Clone() Cont {
	return LetCont{
		Env:  c.Env.Clone(),
		Var:  c.Var,
		Body: c.Body.Clone(),
		next: c.next.Clone(),
	}
}

type AppExpCont struct {
	Env  Env
	Arg  Exp
	next Cont
}

func (AppExpCont) isContinuation() {}

func (c AppExpCont) Next() Cont {
	return c.next
}

func (c AppExpCont) String() string {
	return fmt.Sprintf("{%s _ %s}", envPrefix(c.Env), c.Arg)
}

func (c AppExpCont) Clone() Cont {
	return AppExpCont{
		Env:  c.Env.Clone(),
		Arg:  c.Arg.Clone(),
		next: c.next.Clone(),
	}
}

type AppValueCont struct {
	Fun  Value
	next Cont
}

func (AppValueCont) isContinuation() {}

func (c AppValueCont) Next() Cont {
	return c.next
}

func (c AppValueCont) String() string {
	return fmt.Sprintf("{%s _}", c.Fun)
}

func (c AppValueCont) Clone() Cont {
	return AppValueCont{
		Fun:  c.Fun.Clone(),
		next: c.next.Clone(),
	}
}

type ConsExpCont struct {
	Env  Env
	Tail Exp
	next Cont
}

func (ConsExpCont) isContinuation() {}

func (c ConsExpCont) Next() Cont {
	return c.next
}

func (c ConsExpCont) String() string {
	return fmt.Sprintf("{%s _ :: %s}", envPrefix(c.Env), c.Tail)
}

func (c ConsExpCont) Clone() Cont {
	return ConsExpCont{
		Env:  c.Env.Clone(),
		Tail: c.Tail.Clone(),
		next: c.next.Clone(),
	}
}

type ConsValueCont struct {
	Head Value
	next Cont
}

func (ConsValueCont) isContinuation() {}

func (c ConsValueCont) Next() Cont {
	return c.next
}

func (c ConsValueCont) String() string {
	return fmt.Sprintf("{%s :: _}", c.Head)
}

func (c ConsValueCont) Clone() Cont {
	return ConsValueCont{
		Head: c.Head.Clone(),
		next: c.next.Clone(),
	}
}

type MatchCont struct {
	Env      Env
	NilCase  Exp
	HeadVar  string
	TailVar  string
	ConsCase Exp
	next     Cont
}

func (MatchCont) isContinuation() {}

func (c MatchCont) Next() Cont {
	return c.next
}

func (c MatchCont) String() string {
	return fmt.Sprintf(
		"{%s match _ with [] -> %s | %s :: %s -> %s}",
		envPrefix(c.Env), c.NilCase, c.HeadVar, c.TailVar, c.ConsCase,
	)
}

func (c MatchCont) Clone() Cont {
	return MatchCont{
		Env:      c.Env.Clone(),
		NilCase:  c.NilCase.Clone(),
		HeadVar:  c.HeadVar,
		TailVar:  c.TailVar,
		ConsCase: c.ConsCase.Clone(),
		next:     c.next.Clone(),
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
	case parser.EvalContML4ParserPLUS:
		return OpPlus, nil
	case parser.EvalContML4LexerMINUS:
		return OpMinus, nil
	case parser.EvalContML4LexerTIMES:
		return OpTimes, nil
	case parser.EvalContML4LexerLT:
		return OpLT, nil
	default:
		return "", fmt.Errorf("unknown operator: %s", token.GetText())
	}
}
