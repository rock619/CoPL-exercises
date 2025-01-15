package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/lmittmann/tint"
	"github.com/rock619/CoPL-exercises/solver/evalml2/parser"
)

func main() {
	if err := Run(os.Args[1], os.Stdout, os.Stderr); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func Run(input string, w, errW io.Writer) error {
	inputStream := antlr.NewInputStream(input)
	p := parser.NewEvalML2Parser(antlr.NewCommonTokenStream(
		parser.NewEvalML2Lexer(inputStream),
		antlr.TokenDefaultChannel,
	))
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Question()
	logger := slog.New(tint.NewHandler(errW, &tint.Options{
		Level: slog.LevelDebug,
	}))
	l := &Listener{p: p, l: logger}

	antlr.NewParseTreeWalker().Walk(l, tree)
	logger.Info("Result", "judgements", l.result)
	root, ok := l.PopJudgement()
	if !ok {
		return fmt.Errorf("no root judgement: %v", l.result)
	}
	printer := &Printer{w: w, j: root, indent: "\t"}
	printer.Do()
	return nil
}
