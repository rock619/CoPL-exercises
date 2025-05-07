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
	"github.com/rock619/CoPL-exercises/solver/evalcontml1/parser"
)

func main() {
	if err := Run(os.Stdin, os.Stdout, os.Stderr); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func Run(r io.Reader, w, errW io.Writer) error {
	inputStream := antlr.NewIoStream(r)
	p := parser.NewEvalContML1Parser(antlr.NewCommonTokenStream(
		parser.NewEvalContML1Lexer(inputStream),
		antlr.TokenDefaultChannel,
	))

	logger := setupLogger(errW)

	result, err := NewVisitor(p, logger).Do()
	if err != nil {
		logger.Error("error", "error", err)
	}
	logger.Info("result", "result", result)
	result = Derive(result, logger)
	logger.Info("result", "result", fmt.Sprintf("%#v", result))
	NewPrinter(w, strings.TrimSpace(inputStream.String()), result, "\t", logger).Do()
	return nil
}

func setupLogger(w io.Writer) *slog.Logger {
	f, ok := w.(*os.File)
	return slog.New(tint.NewHandler(w, &tint.Options{
		AddSource: true,
		Level:     slog.LevelDebug,
		NoColor:   !ok || !isatty.IsTerminal(f.Fd()),
	}))
}
