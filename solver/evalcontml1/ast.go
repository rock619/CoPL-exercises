package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalcontml1/parser"
)

type Judgement interface {
	isJudgement()
	Literal() string
	By() Rule
	Premises() []Judgement
}

type Rule string

const (
	EInt   Rule = "E-Int"
	EBool  Rule = "E-Bool"
	EBinOp Rule = "E-BinOp"
	EIf    Rule = "E-If"

	CRet   Rule = "C-Ret"
	CEvalR Rule = "C-EvalR"
	CPlus  Rule = "C-Plus"
	CMinus Rule = "C-Minus"
	CTimes Rule = "C-Times"
	CLt    Rule = "C-Lt"
	CIfT   Rule = "C-IfT"
	CIfF   Rule = "C-IfF"

	BPlus  Rule = "B-Plus"
	BMinus Rule = "B-Minus"
	BTimes Rule = "B-Times"
	BLt    Rule = "B-Lt"
)

type BaseJudgement struct {
	literal  string
	cont     Continuation
	evalTo   Value
	by       Rule
	premises []Judgement
}

func (j *BaseJudgement) Literal() string {
	return j.literal
}

func (j *BaseJudgement) EvalTo() Value {
	return j.evalTo
}

func (j *BaseJudgement) By() Rule {
	return j.by
}

func (j *BaseJudgement) Premises() []Judgement {
	return j.premises
}

func (j *BaseJudgement) String() string {
	return fmt.Sprintf("%+v", *j)
}

type ValueJudgement struct {
	*BaseJudgement
	Value Value
}

func (ValueJudgement) isJudgement() {}

func (j *ValueJudgement) String() string {
	conts := ""
	for c := j.cont; c != nil; c = c.Next() {
		if _, ok := c.(UnaryCont); ok {
			break
		}
		if conts == "" {
			conts = c.String()
		} else {
			conts = fmt.Sprintf("%s >> %s", conts, c.String())
		}
	}
	if conts == "" {
		conts = "_"
	}
	return fmt.Sprintf("%s => %s evalto %s by %s", j.Value, conts, j.evalTo, j.By())
}

type ExpJudgement struct {
	*BaseJudgement
	Exp Exp
}

func (ExpJudgement) isJudgement() {}

func (j *ExpJudgement) String() string {
	conts := ""
	for c := j.cont; c != nil; c = c.Next() {
		if _, ok := c.(UnaryCont); ok {
			break
		}

		conts = fmt.Sprintf("%s >> %s", conts, c.String())
	}
	return fmt.Sprintf("%s%s evalto %s by %s", j.Exp, conts, j.evalTo, j.By())
}

type BinOpJudgement struct {
	*BaseJudgement
	Left  Value
	Op    Op
	Right Value
}

func NewBinOpJudgement(left Value, op Op, right Value) *BinOpJudgement {
	j := &BinOpJudgement{
		BaseJudgement: &BaseJudgement{},
		Left:          left,
		Op:            op,
		Right:         right,
	}

	switch op {
	case OpPlus:
		j.by = BPlus
		j.evalTo = IntValue(int(left.(IntValue)) + int(right.(IntValue)))
	case OpMinus:
		j.by = BMinus
		j.evalTo = IntValue(int(left.(IntValue)) - int(right.(IntValue)))
	case OpTimes:
		j.by = BTimes
		j.evalTo = IntValue(int(left.(IntValue)) * int(right.(IntValue)))
	case OpLT:
		j.by = BLt
		j.evalTo = BoolValue(left.(IntValue) < right.(IntValue))
	default:
		panic("unknown operator")
	}

	return j
}

func (BinOpJudgement) isJudgement() {}

func (j *BinOpJudgement) String() string {
	o := ""
	switch j.Op {
	case OpPlus:
		o = "plus"
	case OpMinus:
		o = "minus"
	case OpTimes:
		o = "times"
	case OpLT:
		o = "less than"
	default:
		panic("unknown operator")
	}
	return fmt.Sprintf("%s %s %s is %s by %s", j.Left, o, j.Right, j.EvalTo(), j.By())
}

type Value interface {
	isValue()
	String() string
}

type IntValue int

func (IntValue) isValue() {}

func (v IntValue) String() string {
	return strconv.Itoa(int(v))
}

type BoolValue bool

func (BoolValue) isValue() {}

func (v BoolValue) String() string {
	return strconv.FormatBool(bool(v))
}

type Exp interface {
	isExpression()
	String() string
}

type IntExp int

func (IntExp) isExpression() {}

func (e IntExp) String() string {
	return strconv.Itoa(int(e))
}

type BoolExp bool

func (BoolExp) isExpression() {}

func (e BoolExp) String() string {
	return strconv.FormatBool(bool(e))
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

type IfExp struct {
	Cond Exp
	Then Exp
	Else Exp
}

func (IfExp) isExpression() {}

func (e IfExp) String() string {
	return fmt.Sprintf("if %s then %s else %s", e.Cond, e.Then, e.Else)
}

type ParenExp struct {
	Inner Exp
}

func (ParenExp) isExpression() {}

func (e ParenExp) String() string {
	return e.Inner.String()
}

type Continuation interface {
	isContinuation()
	Next() Continuation
	String() string
}

type UnaryCont struct{}

func (UnaryCont) isContinuation() {}

func (c UnaryCont) Next() Continuation {
	return nil
}

func (c UnaryCont) String() string {
	return "_"
}

type ExpCont struct {
	Op    Op
	Right Exp
	next  Continuation
}

func (ExpCont) isContinuation() {}

func (c ExpCont) Next() Continuation {
	return c.next
}

func (c ExpCont) String() string {
	right := c.Right.String()
	if parenExp, ok := c.Right.(ParenExp); ok {
		if inner, ok := parenExp.Inner.(BinOpExp); ok {
			if inner.Op.LT(c.Op) {
				right = fmt.Sprintf("(%s)", right)
			}
		}
	}
	return fmt.Sprintf("{_ %s %s}", c.Op, right)
}

type ValueCont struct {
	Left Value
	Op   Op
	next Continuation
}

func (ValueCont) isContinuation() {}

func (c ValueCont) Next() Continuation {
	return c.next
}

func (c ValueCont) String() string {
	return fmt.Sprintf("{%s %s _}", c.Left, c.Op)
}

type IfCont struct {
	Then Exp
	Else Exp
	next Continuation
}

func (IfCont) isContinuation() {}

func (c IfCont) Next() Continuation {
	return c.next
}

func (c IfCont) String() string {
	return fmt.Sprintf("{if _ then %s else %s}", c.Then, c.Else)
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
	case parser.EvalContML1ParserPLUS:
		return OpPlus, nil
	case parser.EvalContML1LexerMINUS:
		return OpMinus, nil
	case parser.EvalContML1LexerTIMES:
		return OpTimes, nil
	case parser.EvalContML1LexerLT:
		return OpLT, nil
	default:
		return "", fmt.Errorf("unknown operator: %s", token.GetText())
	}
}
