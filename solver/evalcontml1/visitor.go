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
	contRes := AssertResult[Continuation](c.Cont().Accept(v))
	if contRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: %w", contRes.Err()))
	}
	valRes := AssertResult[Value](c.Value().Accept(v))
	if valRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: %w", valRes.Err()))
	}

	result := ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   contRes.val,
			evalTo: valRes.val,
		},
		Exp: expRes.val,
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
			Left: valRes.val,
			Op:   op,
		})
	}

	contRes := AssertResult[Continuation](c.Cont().Accept(v))
	if contRes.Err() != nil {
		return Err[Continuation](fmt.Errorf("VisitValueCont: %w", contRes.Err()))
	}

	return OK[Continuation](ValueCont{
		Left: valRes.val,
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
	// v.p.GetTokenStream().GetAllText()
	text := v.p.BaseParser.GetTokenStream().GetTextFromInterval(st.GetSourceInterval())
	return normalizeSpaces(text)
}
