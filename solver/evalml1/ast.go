package main

import (
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
)

type Printer struct {
	w      io.Writer
	j      Judgement
	indent string
}

func (p *Printer) Do() {
	p.do(p.j, 0)
}

func (p *Printer) do(j Judgement, depth int) {
	prefix := strings.Repeat(p.indent, depth)
	if j.Rule.B() {
		fmt.Fprintf(p.w, "%s%s is %s by %s {", prefix, j.Literal, j.EvalTo, j.Rule)
	} else {
		fmt.Fprintf(p.w, "%s%s evalto %s by %s {", prefix, j.Literal, j.EvalTo, j.Rule)
	}

	for i, premise := range j.Premises {
		fmt.Fprintln(p.w)
		p.do(premise, depth+1)
		if i == len(j.Premises)-1 {
			fmt.Fprintln(p.w)
		} else {
			fmt.Fprint(p.w, ";")
		}
	}
	if len(j.Premises) == 0 {
		fmt.Fprint(p.w, "}")
	} else {
		fmt.Fprint(p.w, prefix+"}")
	}

	if depth == 0 {
		fmt.Fprintln(p.w)
	}
}

type Judgement struct {
	Literal  string
	EvalTo   Value
	Rule     Rule
	Premises []Judgement
}

func NewBJudgement(left, right IntValue, op Op) Judgement {
	switch op {
	case Plus:
		return Judgement{
			Literal: fmt.Sprintf("%d plus %d", left, right),
			EvalTo:  IntValue(left + right),
			Rule:    BPlus,
		}
	case Minus:
		return Judgement{
			Literal: fmt.Sprintf("%d minus %d", left, right),
			EvalTo:  IntValue(left - right),
			Rule:    BMinus,
		}
	case Times:
		return Judgement{
			Literal: fmt.Sprintf("%d times %d", left, right),
			EvalTo:  IntValue(left * right),
			Rule:    BTimes,
		}
	case Lt:
		return Judgement{
			Literal: fmt.Sprintf("%d less than %d", left, right),
			EvalTo:  BoolValue(left < right),
			Rule:    BLt,
		}
	default:
		panic("Unknown op: " + string(op))
	}
}

func NewEJudgementFromBinOpExpr(literal string, op Op, leftPremise, rightPremise Judgement) Judgement {
	leftValue, ok := leftPremise.EvalTo.(IntValue)
	if !ok {
		panic("leftPremise.EvalTo is not an IntValue: " + leftPremise.Literal)
	}
	rightValue, ok := rightPremise.EvalTo.(IntValue)
	if !ok {
		panic("rightPremise.EvalTo is not an IntValue: " + rightPremise.Literal)
	}
	bJudgement := NewBJudgement(leftValue, rightValue, op)

	return Judgement{
		Literal:  literal,
		Rule:     NewERuleFromOp(op),
		EvalTo:   bJudgement.EvalTo,
		Premises: append([]Judgement{leftPremise, rightPremise}, bJudgement),
	}
}

func NewERuleFromOp(op Op) Rule {
	switch op {
	case Plus:
		return EPlus
	case Minus:
		return EMinus
	case Times:
		return ETimes
	case Lt:
		return ELt
	default:
		panic("Unknown op: " + string(op))
	}
}

type Rule string

const (
	EInt   Rule = "E-Int"
	EBool  Rule = "E-Bool"
	EIfT   Rule = "E-IfT"
	EIfF   Rule = "E-IfF"
	EPlus  Rule = "E-Plus"
	EMinus Rule = "E-Minus"
	ETimes Rule = "E-Times"
	ELt    Rule = "E-Lt"
	BPlus  Rule = "B-Plus"
	BMinus Rule = "B-Minus"
	BTimes Rule = "B-Times"
	BLt    Rule = "B-Lt"
)

func (r Rule) B() bool {
	return slices.Contains([]Rule{BPlus, BMinus, BTimes, BLt}, r)
}

type Op rune

const (
	Plus  Op = '+'
	Minus Op = '-'
	Times Op = '*'
	Lt    Op = '<'
)

func NewOp(op string) Op {
	switch op {
	case "+":
		return Plus
	case "-":
		return Minus
	case "*":
		return Times
	case "<":
		return Lt
	default:
		panic("Unknown operator: " + op)
	}
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

type BoolValue bool

func (BoolValue) value() {}

func (b BoolValue) String() string {
	return strconv.FormatBool(bool(b))
}
