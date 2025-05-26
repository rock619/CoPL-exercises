package main

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
	"github.com/rock619/CoPL-exercises/solver/polytypingml4/parser"
)

func main() {
	if err := Run(os.Stdin, os.Stdout, os.Stderr); err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}

func Run(r io.Reader, w, errW io.Writer) error {
	inputStream := antlr.NewIoStream(r)
	p := parser.NewPolyTypingML4Parser(antlr.NewCommonTokenStream(
		parser.NewPolyTypingML4Lexer(inputStream),
		antlr.TokenDefaultChannel,
	))

	logger := setupLogger(errW)

	result, err := NewVisitor(p, logger).Do()
	if err != nil {
		logger.Error("error", "error", err)
	}

	logger.Info("result", "result.Env", result.Env, "result.Exp", result.Exp, "result.Type", result.Type)
	gen := NewTypeVarGenerator()
	gen.SetUsedFromEnv(result.Env)
	gen.SetUsedFromType(result.Type)
	inferrer := NewTypeInferrer(gen, logger)
	ires, err := inferrer.Do(result.Env, result.Exp)
	logger.Info("inferred", "inference result", ires)
	sub, err := Unify(append(ires.Substitution.TypeEquations(), TypeEquation{
		Left:  ires.Type,
		Right: result.Type,
	}))
	logger.Info("unified", "sub", sub, "type", sub.ApplyToType(ires.Type), "result type", result.Type)
	NewPrinter(w, strings.TrimSpace(inputStream.String()), sub.ApplyToJudgement(ires.Judgement), "\t", logger).Do()
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
