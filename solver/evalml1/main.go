package main

import (
	"io"
	"log/slog"
	"os"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/evalml1/parser"
)

type Listener struct {
	*parser.BaseEvalML1Listener
	p      *parser.EvalML1Parser
	result []Judgement
	l      *slog.Logger
}

func (l *Listener) PopJudgement() Judgement {
	if len(l.result) == 0 {
		l.l.Error("PopJudgement: empty stack")
	}
	j := l.result[len(l.result)-1]
	l.result = l.result[:len(l.result)-1]
	return j
}

func (l *Listener) PushJudgement(j Judgement) {
	l.l.Debug("PushJudgement", "judgement", j, "stack", l.result)
	l.result = append(l.result, j)
}

func (l *Listener) ExitNegIntExpr(c *parser.NegIntExprContext) {
	l.l.Info("ExitNegIntExpr", "literal", getLiteral(l.p, c))
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err)
	}
	l.PushJudgement(Judgement{
		Literal: getLiteral(l.p, c),
		EvalTo:  IntValue(i),
		Rule:    EInt,
	})
}

func (l *Listener) ExitBinOpExpr(c *parser.BinOpExprContext) {
	right, left := l.PopJudgement(), l.PopJudgement()
	l.l.Info("ExitBinOpExpr", "literal", getLiteral(l.p, c), "left", left.Literal, "right", right.Literal)
	l.PushJudgement(NewEJudgementFromBinOpExpr(getLiteral(l.p, c), NewOp(c.GetOp().GetText()), left, right))
}

func (l *Listener) ExitIfExpr(c *parser.IfExprContext) {
	l.l.Info("ExitIfExpr", "literal", getLiteral(l.p, c), "last", l.result[len(l.result)-1], "last2", l.result[len(l.result)-2], "last3", l.result[len(l.result)-3])
	els, then, cond := l.PopJudgement(), l.PopJudgement(), l.PopJudgement()
	if cond.EvalTo.(BoolValue) {
		l.PushJudgement(Judgement{
			Literal:  getLiteral(l.p, c),
			EvalTo:   then.EvalTo,
			Rule:     EIfT,
			Premises: []Judgement{cond, then},
		})
	} else {
		l.PushJudgement(Judgement{
			Literal:  getLiteral(l.p, c),
			EvalTo:   els.EvalTo,
			Rule:     EIfF,
			Premises: []Judgement{cond, els},
		})
	}
}

func (l *Listener) ExitIntExpr(c *parser.IntExprContext) {
	l.l.Info("ExitIntExpr", "literal", getLiteral(l.p, c))
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err)
	}
	l.result = append(l.result, Judgement{
		Literal: getLiteral(l.p, c),
		EvalTo:  IntValue(i),
		Rule:    EInt,
	})
}

func (l *Listener) ExitBoolExpr(c *parser.BoolExprContext) {
	l.result = append(l.result, Judgement{
		Literal: getLiteral(l.p, c),
		Rule:    EBool,
	})
}

func getLiteral(p antlr.Parser, st antlr.SyntaxTree) string {
	ts := p.GetTokenStream()
	tx := ts.GetTextFromInterval(st.GetSourceInterval())
	return tx
}

func main() {
	if err := Run(os.Args[1], os.Stdout, os.Stderr); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func Run(input string, w, errW io.Writer) error {
	p := parser.NewEvalML1Parser(antlr.NewCommonTokenStream(
		parser.NewEvalML1Lexer(antlr.NewInputStream(input)),
		antlr.TokenDefaultChannel,
	))
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Question()
	logger := slog.New(slog.NewTextHandler(errW, &slog.HandlerOptions{}))
	l := &Listener{p: p, l: logger}

	antlr.NewParseTreeWalker().Walk(l, tree)
	printer := &Printer{w: w, j: l.PopJudgement(), indent: "\t"}
	printer.Do()
	return nil
}
