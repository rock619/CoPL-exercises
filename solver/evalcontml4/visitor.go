package main

import (
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalcontml4/parser"
)

type Visitor struct {
	*parser.BaseEvalContML4Visitor
	p *parser.EvalContML4Parser
	l *slog.Logger
}

func NewVisitor(p *parser.EvalContML4Parser, l *slog.Logger) *Visitor {
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

	env := make(Env, 0)
	if c.Env() != nil {
		envRes := AssertResult[Env](c.Env().Accept(v))
		if envRes.Err() != nil {
			return Err[Judgement](fmt.Errorf("VisitEval: %w", envRes.Err()))
		}
		env = envRes.Val()
	}

	expRes := AssertResult[Exp](c.Exp().Accept(v))
	if expRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: %w", expRes.Err()))
	}

	valRes := AssertResult[Value](c.Value().Accept(v))
	if valRes.Err() != nil {
		return Err[Judgement](fmt.Errorf("VisitEval: %w", valRes.Err()))
	}

	cont := Cont(TerminalCont{})
	if c.Cont() != nil {
		contRes := AssertResult[Cont](c.Cont().Accept(v))
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
		Env: env,
		Exp: expRes.Val(),
	}
	return OK[Judgement](&result)
}

func (v *Visitor) VisitEnv(c *parser.EnvContext) any {
	v.l.Debug("VisitEnv", "literal", v.LiteralOf(c))
	env := make(Env, 0)
	for _, bind := range c.AllBind() {
		bindRes := AssertResult[Bind](bind.Accept(v))
		if bindRes.Err() != nil {
			return Err[Env](fmt.Errorf("VisitEnv: %w", bindRes.Err()))
		}
		env = append(env, bindRes.Val())
	}

	return OK(env)
}

func (v *Visitor) VisitBind(c *parser.BindContext) any {
	v.l.Debug("VisitBind", "literal", v.LiteralOf(c))

	valRes := AssertResult[Value](c.Value().Accept(v))
	if valRes.Err() != nil {
		return Err[Bind](fmt.Errorf("VisitBind: %w", valRes.Err()))
	}
	return OK(Bind{
		Var:   c.IDENTIFIER().GetText(),
		Value: valRes.Val(),
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

func (v *Visitor) VisitFunValue(c *parser.FunValueContext) any {
	v.l.Debug("VisitFunValue", "literal", v.LiteralOf(c))

	envRes := AssertResult[Env](c.Env().Accept(v))
	if envRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitFunValue: %w", envRes.Err()))
	}

	funRes := AssertResult[Fun](c.Fun().Accept(v))
	if funRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitFunValue: %w", funRes.Err()))
	}
	return OK[Value](FunValue{
		Env: envRes.Val(),
		Fun: funRes.Val(),
	})
}

func (v *Visitor) VisitFun(c *parser.FunContext) any {
	v.l.Debug("VisitFun", "literal", v.LiteralOf(c))

	bodyRes := AssertResult[Exp](c.Exp().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Fun](fmt.Errorf("VisitFun: %w", bodyRes.Err()))
	}
	return OK(Fun{
		Param: c.GetParam().GetText(),
		Body:  bodyRes.Val(),
	})
}

func (v *Visitor) VisitRecFunValue(c *parser.RecFunValueContext) any {
	v.l.Debug("VisitRecFunValue", "literal", v.LiteralOf(c))
	envRes := AssertResult[Env](c.Env().Accept(v))
	if envRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitRecFunValue: %w", envRes.Err()))
	}
	funRes := AssertResult[RecFun](c.RecFun().Accept(v))
	if funRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitRecFunValue: %w", funRes.Err()))
	}
	return OK[Value](RecFunValue{
		Env: envRes.Val(),
		Fun: funRes.Val(),
	})
}

func (v *Visitor) VisitRecFun(c *parser.RecFunContext) any {
	v.l.Debug("VisitRecFun", "literal", v.LiteralOf(c))

	bodyRes := AssertResult[Exp](c.Fun().Exp().Accept(v))
	if bodyRes.Err() != nil {
		return Err[RecFun](fmt.Errorf("VisitRecFun: %w", bodyRes.Err()))
	}
	return OK(RecFun{
		Name: c.GetFunName().GetText(),
		Fun: Fun{
			Param: c.Fun().GetParam().GetText(),
			Body:  bodyRes.Val(),
		},
	})
}

func (v *Visitor) VisitNilValue(c *parser.NilValueContext) any {
	v.l.Debug("VisitNilValue", "literal", v.LiteralOf(c))

	return OK[Value](NilValue{})
}

func (v *Visitor) VisitConsValue(c *parser.ConsValueContext) any {
	v.l.Debug("VisitConsValue", "literal", v.LiteralOf(c))

	headRes := AssertResult[Value](c.GetHead().Accept(v))
	if headRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitConsValue: %w", headRes.Err()))
	}
	tailRes := AssertResult[Value](c.GetTail().Accept(v))
	if tailRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitConsValue: %w", tailRes.Err()))
	}
	return OK[Value](ConsValue{
		Head: headRes.Val(),
		Tail: tailRes.Val(),
	})
}

func (v *Visitor) VisitContValue(c *parser.ContValueContext) any {
	v.l.Debug("VisitContValue", "literal", v.LiteralOf(c), "cont", c.Cont().GetChildCount())

	contRes := AssertResult[Cont](c.Cont().Accept(v))
	if contRes.Err() != nil {
		return Err[Value](fmt.Errorf("VisitContValue: %w", contRes.Err()))
	}
	return OK[Value](ContValue{
		Cont: contRes.Val(),
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

func (v *Visitor) VisitFunExp(c *parser.FunExpContext) any {
	v.l.Debug("VisitFunExp", "literal", v.LiteralOf(c))

	bodyRes := AssertResult[Exp](c.Fun().Exp().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitFunExp: %w", bodyRes.Err()))
	}
	return OK[Exp](FunExp{
		Fun: Fun{
			Param: c.Fun().GetParam().GetText(),
			Body:  bodyRes.Val(),
		},
	})
}

func (v *Visitor) VisitMatchExp(c *parser.MatchExpContext) any {
	v.l.Debug("VisitMatchExp", "literal", v.LiteralOf(c))

	matchedRes := AssertResult[Exp](c.GetMatchedExp().Accept(v))
	if matchedRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitMatchExp: %w", matchedRes.Err()))
	}

	nilRes := AssertResult[Exp](c.GetNilExp().Accept(v))
	if nilRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitMatchExp: %w", nilRes.Err()))
	}

	consRes := AssertResult[Exp](c.GetConsExp().Accept(v))
	if consRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitMatchExp: %w", consRes.Err()))
	}
	return OK[Exp](MatchExp{
		Matched:  matchedRes.Val(),
		NilCase:  nilRes.Val(),
		HeadVar:  c.GetHeadVar().GetText(),
		TailVar:  c.GetTailVar().GetText(),
		ConsCase: consRes.Val(),
	})
}

func (v *Visitor) VisitNilExp(c *parser.NilExpContext) any {
	v.l.Debug("VisitNilExp", "literal", v.LiteralOf(c))

	return OK[Exp](NilExp{})
}

func (v *Visitor) VisitConsExp(c *parser.ConsExpContext) any {
	v.l.Debug("VisitConsExp", "literal", v.LiteralOf(c))

	headRes := AssertResult[Exp](c.GetHead().Accept(v))
	if headRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitConsExp: %w", headRes.Err()))
	}
	tailRes := AssertResult[Exp](c.GetTail().Accept(v))
	if tailRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitConsExp: %w", tailRes.Err()))
	}
	return OK[Exp](ConsExp{
		Head: headRes.Val(),
		Tail: tailRes.Val(),
	})
}

func (v *Visitor) VisitAppExp(c *parser.AppExpContext) any {
	v.l.Debug("VisitAppExp", "literal", v.LiteralOf(c))

	funRes := AssertResult[Exp](c.GetFn().Accept(v))
	if funRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitAppExp: %w", funRes.Err()))
	}
	argRes := AssertResult[Exp](c.GetArg().Accept(v))
	if argRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitAppExp: %w", argRes.Err()))
	}
	return OK[Exp](AppExp{
		Fun: funRes.Val(),
		Arg: argRes.Val(),
	})
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

func (v *Visitor) VisitLetExp(c *parser.LetExpContext) any {
	v.l.Debug("VisitLetExp", "literal", v.LiteralOf(c))
	bindRes := AssertResult[Exp](c.GetBindExp().Accept(v))
	if bindRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetExp: %w", bindRes.Err()))
	}
	bodyRes := AssertResult[Exp](c.GetBody().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetExp: %w", bodyRes.Err()))
	}
	return OK[Exp](LetExp{
		Var:  c.GetVar_().GetText(),
		Bind: bindRes.Val(),
		Body: bodyRes.Val(),
	})
}

func (v *Visitor) VisitLetRecExp(c *parser.LetRecExpContext) any {
	v.l.Debug("VisitLetRecExp", "literal", v.LiteralOf(c))

	funBodyRes := AssertResult[Exp](c.RecFun().Fun().GetBody().Accept(v))
	if funBodyRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetRecExp: %w", funBodyRes.Err()))
	}
	bodyRes := AssertResult[Exp](c.GetBody().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetRecExp: %w", bodyRes.Err()))
	}
	return OK[Exp](LetRecExp{
		RecFun: RecFun{
			Name: c.RecFun().GetFunName().GetText(),
			Fun: Fun{
				Param: c.RecFun().Fun().GetParam().GetText(),
				Body:  funBodyRes.Val(),
			},
		},
		Body: bodyRes.Val(),
	})
}

func (v *Visitor) VisitLetCCExp(c *parser.LetCCExpContext) any {
	v.l.Debug("VisitLetCCExp", "literal", v.LiteralOf(c))

	bodyRes := AssertResult[Exp](c.GetBody().Accept(v))
	if bodyRes.Err() != nil {
		return Err[Exp](fmt.Errorf("VisitLetCCExp: %w", bodyRes.Err()))
	}
	return OK[Exp](LetCCExp{
		Var:  c.GetVar_().GetText(),
		Body: bodyRes.Val(),
	})
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

func (v *Visitor) VisitVarExp(c *parser.VarExpContext) any {
	v.l.Debug("VisitVarExp", "literal", v.LiteralOf(c))

	return OK[Exp](VarExp(c.IDENTIFIER().GetText()))
}

func (v *Visitor) VisitTerminalCont(c *parser.TerminalContContext) any {
	v.l.Debug("VisitTerminalCont", "literal", v.LiteralOf(c))

	return OK[Cont](TerminalCont{})
}

func (v *Visitor) VisitValueCont(c *parser.ValueContContext) any {
	v.l.Debug("VisitValueCont", "literal", v.LiteralOf(c))

	valRes := AssertResult[Value](c.Value().Accept(v))
	if valRes.Err() != nil {
		return Err[Cont](fmt.Errorf("VisitValueCont: %w", valRes.Err()))
	}

	op, err := NewOp(c.GetOp())
	if err != nil {
		return Err[Cont](fmt.Errorf("VisitValueCont: %w", err))
	}

	if c.GetNext() == nil {
		return OK[Cont](BinOpValueCont{
			Left: valRes.Val(),
			Op:   op,
		})
	}

	nextRes := AssertResult[Cont](c.GetNext().Accept(v))
	if nextRes.Err() != nil {
		return Err[Cont](fmt.Errorf("VisitValueCont: %w", nextRes.Err()))
	}

	return OK[Cont](BinOpValueCont{
		Left: valRes.Val(),
		Op:   op,
		next: nextRes.Val(),
	})
}

func (v *Visitor) LiteralOf(st antlr.SyntaxTree) string {
	text := v.p.BaseParser.GetTokenStream().GetTextFromInterval(st.GetSourceInterval())
	return normalizeSpaces(text)
}

func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
