package main

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalcontml1/parser"
)

type Visitor struct {
	*parser.BaseEvalContML1Visitor
	p *parser.EvalContML1Parser
	l *slog.Logger
}

func NewVisitor(p *parser.EvalContML1Parser, l *slog.Logger) *Visitor {
	return &Visitor{
		p: p,
		l: l,
	}
}

func (v *Visitor) Do() (Judgement, error) {
	tree := v.p.Eval()
	if tree == nil {
		return nil, fmt.Errorf("tree is nil: %q", v.LiteralOf(tree))
	}

	res := AssertResult[Judgement](tree.Accept(v))
	return res.Val(), res.Err()
}

// VisitEval 戻り値はResult[Judgement]
func (v *Visitor) VisitEval(c *parser.EvalContext) any {
	v.l.Debug("VisitEval", "literal", v.LiteralOf(c))

	expRes := AssertResult[Exp](c.Exp().Accept(v))
	if expRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: %w", expRes.Err()))
	}

	valRes := AssertResult[Value](c.Value().Accept(v))
	if valRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: %w", valRes.Err()))
	}

	cont := Continuation(UnaryCont{})
	if c.Cont() != nil {
		contRes := AssertResult[Continuation](c.Cont().Accept(v))
		if contRes.Err() != nil {
			return Err[Judgement](fmt.Errorf("VisitEval: %w", contRes.Err()))
		}
		cont = contRes.Val()
	}

	result := ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   cont,
			evalTo: valRes.Val(),
		},
		Exp: expRes.Val(),
	}
	return OK[Judgement](&result)
}

func (v *Visitor) VisitIntExp(c *parser.IntExpContext) any {
	v.l.Debug("VisitIntExp", "literal", v.LiteralOf(c))

	i, err := strconv.Atoi(c.INT().GetText())
	if err != nil {
		return Err[Exp](fmt.Errorf("VisitIntExp: %w", err))
	}
	return OK[Exp](IntExp(i))
}

func (v *Visitor) VisitBoolExp(c *parser.BoolExpContext) any {
	v.l.Debug("VisitBoolExp", "literal", v.LiteralOf(c))

	b, err := strconv.ParseBool(c.BOOL().GetText())
	if err != nil {
		return Err[Exp](fmt.Errorf("VisitBoolExp: %w", err))
	}
	return OK[Exp](BoolExp(b))
}

func (v *Visitor) VisitBinOpExp(c *parser.BinOpExpContext) any {
	v.l.Debug("VisitBinOpExp", "literal", v.LiteralOf(c))

	leftRes := AssertResult[Exp](c.GetLeft().Accept(v))
	if leftRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitBinOpExp: %w", leftRes.Err()))
	}

	rightRes := AssertResult[Exp](c.GetRight().Accept(v))
	if rightRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitBinOpExp: %w", rightRes.Err()))
	}

	op, err := NewOp(c.GetOp())
	if err != nil {
		return Err[Exp](fmt.Errorf("VisitBinOpExp: %w", err))
	}

	return OK[Exp](BinOpExp{
		Left:  leftRes.Val(),
		Right: rightRes.Val(),
		Op:    op,
	})
}

func (v *Visitor) VisitIfExp(c *parser.IfExpContext) any {
	v.l.Debug("VisitIfExp", "literal", v.LiteralOf(c))

	condRes := AssertResult[Exp](c.GetCond().Accept(v))
	if condRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitIfExp: %w", condRes.Err()))
	}
	thenRes := AssertResult[Exp](c.GetThen().Accept(v))
	if thenRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitIfExp: %w", thenRes.Err()))
	}
	elseRes := AssertResult[Exp](c.GetElse_().Accept(v))
	if elseRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitIfExp: %w", elseRes.Err()))
	}
	return OK[Exp](IfExp{
		Cond: condRes.Val(),
		Then: thenRes.Val(),
		Else: elseRes.Val(),
	})
}

func (v *Visitor) VisitParenExp(c *parser.ParenExpContext) any {
	v.l.Debug("VisitParenExp", "literal", v.LiteralOf(c))

	innerRes := AssertResult[Exp](c.Exp().Accept(v))
	if innerRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitParenExp: %w", innerRes.Err()))
	}
	return OK[Exp](ParenExp{
		Inner: innerRes.Val(),
	})
}

func (v *Visitor) VisitUnaryCont(c *parser.UnaryContContext) any {
	v.l.Debug("VisitUnaryCont", "literal", v.LiteralOf(c))

	return OK[Continuation](UnaryCont{})
}

func (v *Visitor) VisitValueCont(c *parser.ValueContContext) any {
	v.l.Debug("VisitValueCont", "literal", v.LiteralOf(c))

	valRes := AssertResult[Value](c.Value().Accept(v))
	if valRes.Err() != nil {
		return Err[Continuation](fmt.Errorf("VisitValueCont: %w", valRes.Err()))
	}

	op, err := NewOp(c.GetOp())
	if err != nil {
		return Err[Continuation](fmt.Errorf("VisitValueCont: %w", err))
	}

	if c.Cont() == nil {
		return OK[Continuation](ValueCont{
			Left: valRes.Val(),
			Op:   op,
		})
	}

	contRes := AssertResult[Continuation](c.Cont().Accept(v))
	if contRes.Err() != nil {
		return Err[Continuation](fmt.Errorf("VisitValueCont: %w", contRes.Err()))
	}

	return OK[Continuation](ValueCont{
		Left: valRes.Val(),
	})
}

func (v *Visitor) VisitIntValue(c *parser.IntValueContext) any {
	v.l.Debug("VisitIntValue", "literal", v.LiteralOf(c))

	i, err := strconv.Atoi(c.INT().GetText())
	if err != nil {
		return Err[Value](fmt.Errorf("VisitIntValue: %w", err))
	}
	return OK[Value](IntValue(i))
}

func (v *Visitor) VisitBoolValue(c *parser.BoolValueContext) any {
	v.l.Debug("VisitBoolValue", "literal", v.LiteralOf(c))

	b, err := strconv.ParseBool(c.BOOL().GetText())
	if err != nil {
		return Err[Value](fmt.Errorf("VisitBoolValue: %w", err))
	}
	return OK[Value](BoolValue(b))
}

func (v *Visitor) LiteralOf(st antlr.SyntaxTree) string {
	text := v.p.BaseParser.GetTokenStream().GetTextFromInterval(st.GetSourceInterval())
	return normalizeSpaces(text)
}
