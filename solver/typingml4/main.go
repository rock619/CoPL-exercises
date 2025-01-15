package main

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
	"github.com/rock619/CoPL-exercises/solver/typingml4/parser"
)

func main() {
	if err := Run(os.Stdin, os.Stdout, os.Stderr); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func Run(r io.Reader, w, errW io.Writer) error {
	inputStream := antlr.NewIoStream(r)
	p := parser.NewTypingML4Parser(antlr.NewCommonTokenStream(
		parser.NewTypingML4Lexer(inputStream),
		antlr.TokenDefaultChannel,
	))
	logger := setupLogger(errW)

	v := NewVisitor(p, logger)
	result, err := v.Do(p.Judgement())
	if err != nil {
		logger.Error("*Visitor.Do", "error", err)
	}
	logger.Info("*Visitor.Do", "result", result, "env", result.Env())

	printer, err := NewPrinter(w, strings.TrimSpace(inputStream.String()), result, "\t")
	if err != nil {
		logger.Error("*NewPrinter", "error", err)
	}
	logger.Info("unified", "subs", printer.subs)
	printer.Do()
	return nil
}

func setupLogger(w io.Writer) *slog.Logger {
	f, ok := w.(*os.File)
	return slog.New(tint.NewHandler(w, &tint.Options{
		AddSource: true,
		Level:     slog.LevelInfo,
		NoColor:   !ok || !isatty.IsTerminal(f.Fd()),
	}))
}
