package main

import (
	"fmt"
	"io"
	"strings"
)

type Printer struct {
	w      io.Writer
	j      Judgement
	indent string
}

func (p *Printer) Do() {
	p.do(p.j, 0)
}

func (p *Printer) do(j Judgement, depth int) {
	prefix := strings.Repeat(p.indent, depth)
	if j.Rule.B() {
		fmt.Fprintf(p.w, "%s%s is %s by %s {", prefix, j.Literal, j.EvalTo, j.Rule)
	} else {
		binds := make([]string, len(j.Environment))
		for i, bind := range j.Environment {
			binds[i] = fmt.Sprintf("%s = %s", bind.Name, bind.Val)
		}
		env := strings.Join(binds, ", ")
		if env != "" {
			env += " "
		}
		fmt.Fprintf(p.w, "%s%s|- %s evalto %s by %s {", prefix, env, j.Literal, j.EvalTo, j.Rule)
	}

	for i, premise := range j.Premises {
		fmt.Fprintln(p.w)
		p.do(premise, depth+1)
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
