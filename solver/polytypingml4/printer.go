package main

import (
	"fmt"
	"io"
	"log/slog"
	"strings"
)

type Printer struct {
	w      io.Writer
	input  string
	root   Judgement
	indent string
	l      *slog.Logger
}

func NewPrinter(w io.Writer, input string, root Judgement, indent string, l *slog.Logger) *Printer {
	return &Printer{
		w:      w,
		input:  input,
		root:   root,
		indent: indent,
		l:      l,
	}
}

func (p *Printer) Do() {
	// fmt.Fprintln(p.w, "// Derivation System PolyTypingML4")
	p.do(p.root, 0)
}

func (p *Printer) do(j Judgement, depth int) {
	prefix := strings.Repeat(p.indent, depth)
	if depth == 0 {
		fmt.Fprintf(p.w, "%s by %s {", p.input, j.By)
	} else {
		fmt.Fprintf(p.w, "%s%s {", prefix, j)
	}

	for i, v := range j.Premises {
		fmt.Fprintln(p.w)
		p.do(v, depth+1)
		if i == len(j.Premises)-1 {
			fmt.Fprintln(p.w)
		} else {
			fmt.Fprint(p.w, ";")
		}
	}

	if len(j.Premises) == 0 {
		fmt.Fprint(p.w, "}")
	} else {
		fmt.Fprint(p.w, prefix+"}")
	}

	if depth == 0 {
		fmt.Fprintln(p.w)
	}
}
