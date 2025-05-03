package main

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalcontml1/parser"
)

type Judgement interface {
	isJudgement()
	Literal() string
	Derive()
	EvalTo() Value
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
	CIfT   Rule = "C-IfT"
	CIfF   Rule = "C-IfF"

	BPlus  Rule = "B-Plus"
	BMinus Rule = "B-Minus"
	BTimes Rule = "B-Times"
	BLT    Rule = "B-LT"
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

func (j *ValueJudgement) Literal() string {
	return fmt.Sprintf("%s by %s", j.Value, j.by)
}

func (j *ValueJudgement) Derive() {
	switch v := j.Value.(type) {
	case IntValue:
		switch c := j.cont.(type) {
		case UnaryCont:
			// v => _ evalto v by C-Ret {}
			j.by = CRet
		case ExpCont:
			// v1 => {_ op e} >> k evalto v2 by C-EvalR
			//     e >> {v1 op _} >> k evalto v2
			j.by = CEvalR

			cont2 := c.Next
			cont1 := &ValueCont{
				Left: v,
				Op:   c.Op,
				Next: cont2,
			}
			child := &ExpJudgement{
				BaseJudgement: &BaseJudgement{
					cont:   cont1,
					evalTo: j.evalTo,
				},
				Exp: c.Right,
			}
			child.Derive()
			j.premises = append(j.premises, child)
		case ValueCont:
			// i2 => {i1 op _} >> k evalto v by C-*
			// TODO:
			j.by = CPlus
		default:
			panic("unknown continuation type")
		}
	case BoolValue:
		if bool(v) {
			j.by = CIfT
		} else {
			j.by = CIfF
		}
	default:
		panic("unknown value type")
	}
}

type ExpJudgement struct {
	*BaseJudgement
	Exp Exp
}

func (ExpJudgement) isJudgement() {}

func (j *ExpJudgement) String() string {
	return fmt.Sprintf("%+v", *j)
}

func (j *ExpJudgement) Derive() {
	switch exp := j.Exp.(type) {
	case IntExp:
		j.by = EInt
		child := &ValueJudgement{
			BaseJudgement: &BaseJudgement{
				cont:   j.cont,
				evalTo: j.evalTo,
			},
			Value: IntValue(exp),
		}
		child.Derive()
		j.premises = append(j.premises, child)
	case BoolExp:
		j.by = EBool
	case BinOpExp:
		j.by = EBinOp
	case IfExp:
		j.by = EIf
	default:
		panic("unknown expression type")
	}
}

type Value interface {
	isValue()
}

type IntValue int

func (IntValue) isValue() {}

type BoolValue bool

func (BoolValue) isValue() {}

type Exp interface {
	isExpression()
}

type IntExp int

func (IntExp) isExpression() {}

type BoolExp bool

func (BoolExp) isExpression() {}

type BinOpExp struct {
	Left  Exp
	Op    Op
	Right Exp
}

func (BinOpExp) isExpression() {}

type IfExp struct {
	Cond Exp
	Then Exp
	Else Exp
}

func (IfExp) isExpression() {}

type Continuation interface {
	isContinuation()
	HasNext() bool
}

type UnaryCont struct{}

func (UnaryCont) isContinuation() {}

func (c UnaryCont) HasNext() bool {
	return false
}

type ExpCont struct {
	Op    Op
	Right Exp
	Next  Continuation
}

func (ExpCont) isContinuation() {}

func (c ExpCont) HasNext() bool {
	return c.Right != nil
}

type ValueCont struct {
	Left Value
	Op   Op
	Next Continuation
}

func (ValueCont) isContinuation() {}

func (c ValueCont) HasNext() bool {
	return c.Next != nil
}

type IfCont struct {
	Then Exp
	Else Exp
	Next Continuation
}

func (IfCont) isContinuation() {}

func (c IfCont) HasNext() bool {
	return c.Next != nil
}

type Op string

const (
	OpPlus  Op = "+"
	OpMinus Op = "-"
	OpTimes Op = "*"
	OpLT    Op = "<"
)

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
