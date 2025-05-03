package main

import (
	"log/slog"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalcontml1/parser"
)

type Listener struct {
	*parser.BaseEvalContML1Listener
	p      *parser.EvalContML1Parser
	l      *slog.Logger
	result Judgement
}

func NewListener(p *parser.EvalContML1Parser, l *slog.Logger) *Listener {
	return &Listener{
		p: p,
		l: l,
	}
}

func (l *Listener) EnterEval(c *parser.EvalContext) {
	l.l.Debug("EnterEval", "literal", l.LiteralOf(c))
}

func (l *Listener) ExitEval(c *parser.EvalContext) {
	l.l.Debug("ExitEval", "literal", l.LiteralOf(c))
	l.result = &ExpJudgement{
		BaseJudgement: &BaseJudgement{},
	}
}

func (l *Listener) EnterIntExp(c *parser.IntExpContext) {
	l.l.Debug("EnterExp", "literal", l.LiteralOf(c))
}

func (l *Listener) EnterCont(c *parser.ContContext) {
	l.l.Debug("EnterCont", "literal", l.LiteralOf(c))
}

func (l *Listener) LiteralOf(st antlr.SyntaxTree) string {
	text := l.p.GetTokenStream().GetTextFromInterval(st.GetSourceInterval())
	return normalizeSpaces(text)
}

func normalizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
