package main

import (
	"log/slog"
	"slices"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalml2/parser"
)

type Listener struct {
	*parser.BaseEvalML2Listener
	p           *parser.EvalML2Parser
	result      []Judgement
	environment []Bind
	l           *slog.Logger
	depth       int
}

func (l *Listener) IncrDepth() {
	l.depth++
}

func (l *Listener) DecrDepth() {
	l.depth--
}

func (l *Listener) Empty() bool {
	return len(l.result) == 0
}

func (l *Listener) PopJudgement() (Judgement, bool) {
	if l.Empty() {
		return Judgement{}, false
	}

	j := l.result[len(l.result)-1]
	l.result = l.result[:len(l.result)-1]
	return j, true
}

func (l *Listener) PushJudgement(j Judgement) {
	l.l.Debug("PushJudgement", "judgement", j, "stack", l.result)
	l.result = append(l.result, j)
}

func (l *Listener) PeekParentJudgement() (parent Judgement, siblings []Judgement, ok bool) {
	for _, prev := range slices.Backward(l.result) {
		if prev.Depth < l.depth {
			return prev, siblings, true
		}
		siblings = slices.Insert(siblings, 0, prev)
	}
	return Judgement{}, siblings, false
}

func (l *Listener) PopParentJudgement() (parent Judgement, siblings []Judgement, ok bool) {
	for {
		prev, ok := l.PopJudgement()
		if !ok {
			return Judgement{}, nil, false
		}
		if prev.Depth < l.depth {
			return prev, siblings, true
		}
		siblings = slices.Insert(siblings, 0, prev)
	}
}

func (l *Listener) EnterEveryRule(c antlr.ParserRuleContext) {
	l.l.Debug("EnterEveryRule", "literal", c.GetText())
	if _, ok := c.(*parser.DefContext); ok {
		l.l.Debug("*parser.DefContext")
	}
	if _, ok := c.(parser.IExprContext); ok {
		l.EnterExpr(c)
	}
}

func (l *Listener) EnterExpr(c antlr.ParserRuleContext) {
	defer l.IncrDepth()

	l.l.Info("EnterExpr", "literal", getLiteral(l.p, c), "depth", l.depth)

	env := slices.Clone(l.environment)
	parent, siblings, ok := l.PeekParentJudgement()
	if ok {
		env = slices.Clone(parent.Environment)
	}
	if lec, ok := parent.Ctx.(*parser.LetExprContext); ok && parent.Rule == ELet && len(siblings) == 1 {
		l.l.Info("EnterExpr let, one siblings", "depth", l.depth, "siblings[0]", siblings[0])
		env = append(env, Bind{Name: lec.GetBindName().GetText(), Val: siblings[0].EvalTo})
	}

	l.PushJudgement(Judgement{
		Ctx:         c,
		Environment: env,
		Literal:     getLiteral(l.p, c),
		Depth:       l.depth,
	})
}

func (l *Listener) EnterBoolExpr(c *parser.BoolExprContext) {
	self, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	self.Rule = EBool
	l.PushJudgement(self)
}

func (l *Listener) EnterIntExpr(c *parser.IntExprContext) {
	self, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	self.Rule = EInt
	l.PushJudgement(self)
}

func (l *Listener) EnterNegIntExpr(c *parser.NegIntExprContext) {
	self, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	self.Rule = EInt
	l.PushJudgement(self)
}

func (l *Listener) EnterVarExpr(c *parser.VarExprContext) {
	self, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	self.Rule = EVar
	l.PushJudgement(self)
}

func (l *Listener) EnterBinOpExpr(c *parser.BinOpExprContext) {
	self, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	self.Rule = EBinOp
	l.PushJudgement(self)
}

func (l *Listener) EnterIfExpr(c *parser.IfExprContext) {
	self, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	self.Rule = EIf
	l.PushJudgement(self)
}

func (l *Listener) EnterLetExpr(c *parser.LetExprContext) {
	self, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	self.Rule = ELet
	l.PushJudgement(self)
}

func (l *Listener) ExitEveryRule(c antlr.ParserRuleContext) {
	l.l.Debug("ExitEveryRule", "literal", c.GetText())
	if ec, ok := c.(parser.IExprContext); ok {
		l.ExitExpr(ec)
	}
}

func (l *Listener) ExitExpr(c parser.IExprContext) {
	l.DecrDepth()

	l.l.Info("ExitExpr", "literal", getLiteral(l.p, c), "depth", l.depth)
}

func (l *Listener) ExitDef(c *parser.DefContext) {
	l.l.Info("ExitDef", "literal", getLiteral(l.p, c), "VARNAME", c.VARNAME().GetText(), "value", c.Value().GetText())
	l.environment = append(l.environment, Bind{
		Name: c.VARNAME().GetText(),
		Val:  ValueFromLiteral(c.Value().GetText()),
	})
}

func (l *Listener) ExitLetExpr(c *parser.LetExprContext) {
	l.l.Info("ExitLetExpr", "literal", getLiteral(l.p, c))
	body, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No body judgement")
		return
	}
	bind, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No bind judgement")
		return
	}
	self, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	self.EvalTo = body.EvalTo
	self.Rule = ELet
	self.Premises = []Judgement{bind, body}
	l.PushJudgement(self)
}

func (l *Listener) ExitBoolExpr(c *parser.BoolExprContext) {
	j, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	b := ValueFromLiteral(j.Literal)
	j.EvalTo = b
	j.Rule = EBool
	l.PushJudgement(j)
}

func (l *Listener) ExitIntExpr(c *parser.IntExprContext) {
	l.l.Info("ExitIntExpr", "literal", getLiteral(l.p, c))
	j, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	i := ValueFromLiteral(j.Literal)
	j.EvalTo = i
	j.Rule = EInt
	l.PushJudgement(j)
}

func (l *Listener) ExitNegIntExpr(c *parser.NegIntExprContext) {
	l.l.Info("ExitNegIntExpr", "literal", getLiteral(l.p, c))
	j, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	i := ValueFromLiteral(j.Literal)
	j.EvalTo = i
	j.Rule = EInt
	l.PushJudgement(j)
}

func (l *Listener) ExitVarExpr(c *parser.VarExprContext) {
	l.l.Info("ExitVarExpr", "literal", getLiteral(l.p, c))
	varName := c.VARNAME().GetText()
	j, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}
	l.PushJudgement(newEVarJudgement(c, j.Environment, j.Literal, varName, l.depth))
}

func (l *Listener) ExitBinOpExpr(c *parser.BinOpExprContext) {
	l.l.Info("ExitBinOpExpr", "literal", getLiteral(l.p, c))
	right, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No right judgement")
		return
	}
	left, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No left judgement")
		return
	}
	self, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}

	l.l.Info("ExitBinOpExpr", "literal", getLiteral(l.p, c), "left", left.Literal, "right", right.Literal, "self", self.Literal)
	l.PushJudgement(NewEJudgementFromBinOpExpr(c, self.Environment, self.Literal, c.GetOp(), left, right, self.Depth))
}

func (l *Listener) ExitIfExpr(c *parser.IfExprContext) {
	l.l.Info("ExitIfExpr", "literal", getLiteral(l.p, c), "last", l.result[len(l.result)-1], "last2", l.result[len(l.result)-2], "last3", l.result[len(l.result)-3])
	els, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No else judgement")
		return
	}
	then, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No then judgement")
		return
	}
	cond, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No condition judgement")
		return
	}
	self, ok := l.PopJudgement()
	if !ok {
		l.l.Error("No self judgement")
		return
	}

	if _, ok := cond.EvalTo.(BoolValue); ok {
		self.EvalTo = then.EvalTo
		self.Rule = EIfT
		self.Premises = []Judgement{cond, then}
	} else {
		self.EvalTo = els.EvalTo
		self.Rule = EIfF
		self.Premises = []Judgement{cond, els}
	}
	l.PushJudgement(self)
}
