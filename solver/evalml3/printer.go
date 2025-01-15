package main

import (
	"fmt"
	"io"
	"strings"
)

type Printer struct {
	w      io.Writer
	input  string
	e      Expr
	indent string
}

func NewPrinter(w io.Writer, input string, root Expr, indent string) *Printer {
	return &Printer{
		w:      w,
		input:  input,
		e:      root,
		indent: indent,
	}
}

func (p *Printer) Do() {
	p.do(p.e, 0)
}

func (p *Printer) do(e Expr, depth int) {
	prefix := strings.Repeat(p.indent, depth)
	literal := strings.Join(strings.Fields(e.Literal()), " ")
	switch {
	case depth == 0:
		fmt.Fprintf(p.w, "%s by %s {", p.input, e.Rule())
	case e.Rule().B():
		fmt.Fprintf(p.w, "%s%s by %s {", prefix, literal, e.Rule())
	default:
		binds := make([]string, len(e.Env()))
		for i, bind := range e.Env() {
			binds[i] = fmt.Sprintf("%s = %s", bind.Name, bind.Val)
		}
		env := strings.Join(binds, ", ")
		if env != "" {
			env += " "
		}
		fmt.Fprintf(p.w, "%s%s|- %s evalto %s by %s {", prefix, env, literal, e.EvalTo(), e.Rule())
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
