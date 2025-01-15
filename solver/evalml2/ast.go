package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalml2/parser"
)

type Bind struct {
	Name string
	Val  Value
}

func getLiteral(p antlr.Parser, st antlr.SyntaxTree) string {
	ts := p.GetTokenStream()
	tx := ts.GetTextFromInterval(st.GetSourceInterval())
	return tx
}

type Judgement struct {
	Ctx         antlr.ParserRuleContext
	Environment []Bind
	Literal     string
	EvalTo      Value
	Rule        Rule
	Premises    []Judgement
	Depth       int
}

func NewBJudgement(left, right IntValue, op antlr.Token) Judgement {
	switch op.GetTokenType() {
	case parser.EvalML2ParserPLUS:
		return Judgement{
			Literal: fmt.Sprintf("%d plus %d", left, right),
			EvalTo:  IntValue(left + right),
			Rule:    BPlus,
		}
	case parser.EvalML2ParserMINUS:
		return Judgement{
			Literal: fmt.Sprintf("%d minus %d", left, right),
			EvalTo:  IntValue(left - right),
			Rule:    BMinus,
		}
	case parser.EvalML2ParserTIMES:
		return Judgement{
			Literal: fmt.Sprintf("%d times %d", left, right),
			EvalTo:  IntValue(left * right),
			Rule:    BTimes,
		}
	case parser.EvalML2ParserLT:
		return Judgement{
			Literal: fmt.Sprintf("%d less than %d", left, right),
			EvalTo:  BoolValue(left < right),
			Rule:    BLt,
		}
	default:
		panic("Unknown op: " + op.GetText())
	}
}

func NewEJudgementFromBinOpExpr(
	ctx antlr.ParserRuleContext,
	env []Bind,
	literal string,
	op antlr.Token,
	leftPremise, rightPremise Judgement,
	depth int,
) Judgement {
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
		Ctx:         ctx,
		Environment: env,
		Literal:     literal,
		Rule:        NewERuleFromOp(op),
		EvalTo:      bJudgement.EvalTo,
		Premises:    append([]Judgement{leftPremise, rightPremise}, bJudgement),
		Depth:       depth,
	}
}

func NewERuleFromOp(op antlr.Token) Rule {
	switch op.GetTokenType() {
	case parser.EvalML2ParserPLUS:
		return EPlus
	case parser.EvalML2ParserMINUS:
		return EMinus
	case parser.EvalML2ParserTIMES:
		return ETimes
	case parser.EvalML2ParserLT:
		return ELt
	default:
		panic("Unknown op: " + op.GetText())
	}
}

type Rule string

const (
	EInt   Rule = "E-Int"
	EBool  Rule = "E-Bool"
	EVar1  Rule = "E-Var1"
	EVar2  Rule = "E-Var2"
	EPlus  Rule = "E-Plus"
	EMinus Rule = "E-Minus"
	ETimes Rule = "E-Times"
	ELt    Rule = "E-Lt"
	EIfT   Rule = "E-IfT"
	EIfF   Rule = "E-IfF"
	ELet   Rule = "E-Let"
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

type Value interface {
	value()
	String() string
}

func ValueFromLiteral(literal string) Value {
	i, err := strconv.Atoi(literal)
	if err == nil {
		return IntValue(i)
	}
	b, err := strconv.ParseBool(literal)
	if err == nil {
		return BoolValue(b)
	}
	panic("Unknown value literal: " + literal)
}

func ValueFromToken(tok antlr.Token) Value {
	switch tok.GetTokenType() {
	case parser.EvalML2ParserINT:
		i, err := strconv.Atoi(tok.GetText())
		if err != nil {
			panic(err)
		}
		return IntValue(i)
	case parser.EvalML2ParserBOOL:
		b, err := strconv.ParseBool(tok.GetText())
		if err != nil {
			panic(err)
		}
		return BoolValue(b)
	default:
		panic("Unknown token type: " + tok.GetText())
	}
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

func newEVarJudgement(ctx antlr.ParserRuleContext, env []Bind, literal string, varName string, depth int) Judgement {
	if len(env) == 0 {
		panic("Variable not found: " + varName)
	}

	j := Judgement{
		Ctx:         ctx,
		Environment: env,
		Literal:     literal,
		Depth:       depth,
	}
	last := env[len(env)-1]
	if last.Name == varName {
		j.Rule = EVar1
		j.EvalTo = last.Val
		return j
	}
	j.Rule = EVar2
	j.Premises = []Judgement{newEVarJudgement(ctx, env[:len(env)-1], literal, varName, depth+1)}
	j.EvalTo = j.Premises[0].EvalTo
	return j
}
