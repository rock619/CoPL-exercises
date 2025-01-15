package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
	"github.com/rock619/CoPL-exercises/solver/evalml3/parser"
)

func main() {
	if err := Run(os.Stdin, os.Stdout, os.Stderr); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func Run(r io.Reader, w, errW io.Writer) error {
	inputStream := antlr.NewIoStream(r)
	p := parser.NewEvalML3Parser(antlr.NewCommonTokenStream(
		parser.NewEvalML3Lexer(inputStream),
		antlr.TokenDefaultChannel,
	))
	tree := p.Question().Eval()
	logger := setupLogger(errW)

	l := &Listener{p: p, l: logger}

	antlr.NewParseTreeWalker().Walk(l, tree)
	// dump.P(l.result)
	root, ok := l.Pop()
	if !ok {
		return fmt.Errorf("no root expr: %v", l.result)
	}
	printer := NewPrinter(w, strings.TrimSpace(inputStream.String()), root, "\t")
	printer.Do()
	return nil
}

func setupLogger(w io.Writer) *slog.Logger {
	f, ok := w.(*os.File)
	return slog.New(tint.NewHandler(w, &tint.Options{
		Level:   slog.LevelDebug,
		NoColor: !ok || !isatty.IsTerminal(f.Fd()),
	}))
}
