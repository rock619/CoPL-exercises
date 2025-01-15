package main

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
	"github.com/rock619/CoPL-exercises/solver/evalml4/parser"
)

func main() {
	if err := Run(os.Stdin, os.Stdout, os.Stderr); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func Run(r io.Reader, w, errW io.Writer) error {
	inputStream := antlr.NewIoStream(r)
	p := parser.NewEvalML4Parser(antlr.NewCommonTokenStream(
		parser.NewEvalML4Lexer(inputStream),
		antlr.TokenDefaultChannel,
	))
	logger := setupLogger(errW)

	v := NewVisitor(p, logger)
	result, err := v.Do(p.Question())
	if err != nil {
		return err
	}

	printer := NewPrinter(w, strings.TrimSpace(inputStream.String()), result, "\t")
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
