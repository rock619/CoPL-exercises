package main

import (
	"fmt"
	"io"
	"strings"
)

type Printer struct {
	w      io.Writer
	input  string
	root   Expr
	subs   Substitutions
	indent string
}

func NewPrinter(w io.Writer, input string, root Expr, indent string) (*Printer, error) {
	subs, err := root.Env().Unify()
	if err != nil {
		return nil, err
	}

	return &Printer{
		w:      w,
		input:  input,
		root:   root,
		subs:   subs,
		indent: indent,
	}, nil
}

func (p *Printer) Do() {
	p.do(p.root, 0)
}

func (p *Printer) do(e Expr, depth int) {
	prefix := strings.Repeat(p.indent, depth)
	literal := normalizeSpaces(e.Literal())
	if depth == 0 {
		fmt.Fprintf(p.w, "%s by %s {", p.input, e.Rule())
	} else {
		envStr := e.Env().Print(p.subs)
		if envStr != "" {
			envStr += " "
		}
		fmt.Fprintf(
			p.w,
			"%s%s|- %s : %s by %s {",
			prefix, envStr, literal, Substitute(e.Type(), p.subs).Print(), e.Rule(),
		)
	}

	for i, c := range e.Children() {
		fmt.Fprintln(p.w)
		p.do(c, depth+1)
		if i == len(e.Children())-1 {
			fmt.Fprintln(p.w)
		} else {
			fmt.Fprint(p.w, ";")
		}
	}
	if len(e.Children()) == 0 {
		fmt.Fprint(p.w, "}")
	} else {
		fmt.Fprint(p.w, prefix+"}")
	}

	if depth == 0 {
		fmt.Fprintln(p.w)
	}
}
