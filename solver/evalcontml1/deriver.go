package main

import (
	"fmt"
	"log/slog"
)

func Derive(j Judgement, l *slog.Logger) Judgement {
	l.Info("Derive", "judgement", j)

	d := &Deriver{l: l}
	return d.Do(j)
}

type Deriver struct {
	l *slog.Logger
}

func (d *Deriver) Do(j Judgement) Judgement {
	switch j := j.(type) {
	case *ValueJudgement:
		d.deriveValueJudgement(j)
	case *ExpJudgement:
		d.deriveExpJudgement(j)
	default:
		panic(fmt.Sprintf("unknown judgement type; got %T", j))
	}

	return j
}

func (d *Deriver) deriveValueJudgement(j *ValueJudgement) {
	switch v := j.Value.(type) {
	case IntValue:
		d.deriveIntValueJudgement(j, v)
	case BoolValue:
		d.deriveBoolValueJudgement(j, v)
	default:
		panic(fmt.Sprintf("unknown value type; got %T", v))
	}
}

func (d *Deriver) deriveIntValueJudgement(j *ValueJudgement, v IntValue) {
	d.l.Info("deriveIntValueJudgement", "judgement", j, "value", v)

	switch c := j.cont.(type) {
	case UnaryCont, nil:
		// v => _ evalto v by C-Ret {}
		j.by = CRet
	case ExpCont:
		// v1 => {_ op e} >> k evalto v2 by C-EvalR
		//     e >> {v1 op _} >> k evalto v2
		j.by = CEvalR

		cont2 := c.Next()
		cont1 := ValueCont{
			Left: v,
			Op:   c.Op,
			next: cont2,
		}
		child := &ExpJudgement{
			BaseJudgement: &BaseJudgement{
				cont:   cont1,
				evalTo: j.evalTo,
			},
			Exp: c.Right,
		}
		d.deriveExpJudgement(child)
		j.premises = append(j.premises, child)
	case ValueCont:
		// i2 => {i1 op _} >> k evalto v by C-? {
		//     i1 op i2 is i3 by B-? {};
		//     i3 => k evalto v by ?
		// }

		bj := NewBinOpJudgement(c.Left, c.Op, v)
		d.l.Info("deriveValueJudgement: binop judgement", "judgement", bj, "bj.EvalTo()", bj.EvalTo())
		vj := &ValueJudgement{
			BaseJudgement: &BaseJudgement{
				cont:   c.Next(),
				evalTo: j.EvalTo(),
			},
			Value: bj.EvalTo(),
		}
		if vj.cont == nil {
			vj.cont = UnaryCont{}
		}
		d.deriveValueJudgement(vj)
		j.premises = append(j.premises, bj, vj)
		switch c.Op {
		case OpPlus:
			j.by = CPlus
		case OpMinus:
			j.by = CMinus
		case OpTimes:
			j.by = CTimes
		case OpLT:
			j.by = CLt
		default:
			panic("unknown operator")
		}
	default:
		d.l.Error("deriveIntValueJudgement: unknown continuation type", "type", fmt.Sprintf("%T", c), "judgement", j)
	}
}

func (d *Deriver) deriveBoolValueJudgement(j *ValueJudgement, v BoolValue) {
	d.l.Info("deriveBoolValueJudgement", "judgement", j, "value", v)

	switch c := j.cont.(type) {
	case UnaryCont, nil:
		// v => _ evalto v by C-Ret {}
		j.by = CRet
	case IfCont:
		// true => {if _ then e1 else e2} >> k evalto v by C-IfT {
		//     e1 >> k evalto v by ? {}
		// }
		// false => {if e1 then e2 else e3} >> k evalto v by C-IfF {
		//     e2 >> k evalto v by ? {}
		// }
		j.by = CIfF
		child := &ExpJudgement{
			BaseJudgement: &BaseJudgement{
				cont:   c.Next(),
				evalTo: j.evalTo,
			},
			Exp: c.Else,
		}
		if v {
			j.by = CIfT
			child.Exp = c.Then
		}
		d.deriveExpJudgement(child)
		j.premises = append(j.premises, child)
	default:
		d.l.Error("deriveBoolValueJudgement: unknown continuation type", "type", fmt.Sprintf("%T", c), "judgement", j)
	}
}

func (d *Deriver) deriveExpJudgement(j *ExpJudgement) {
	switch e := j.Exp.(type) {
	case IntExp:
		d.deriveIntExpJudgement(j, e)
	case BoolExp:
		d.deriveBoolExpJudgement(j, e)
	case BinOpExp:
		d.deriveBinOpExpJudgement(j, e)
	case IfExp:
		d.deriveIfExpJudgement(j, e)
	case ParenExp:
		d.deriveParenExpJudgement(j, e)
	default:
		d.l.Error("deriveExpJudgement: unknown expression type", "type", fmt.Sprintf("%T", j.Exp), "exp", j.Exp)
	}
}

func (d *Deriver) deriveIntExpJudgement(j *ExpJudgement, e IntExp) {
	d.l.Info("deriveIntExpJudgement", "judgement", j, "exp", e)

	j.by = EInt
	child := &ValueJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   j.cont,
			evalTo: j.evalTo,
		},
		Value: IntValue(e),
	}
	d.deriveValueJudgement(child)
	j.premises = append(j.premises, child)
}

func (d *Deriver) deriveBoolExpJudgement(j *ExpJudgement, e BoolExp) {
	d.l.Info("deriveBoolExpJudgement", "judgement", j, "exp", e)

	j.by = EBool
	child := &ValueJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   j.cont,
			evalTo: j.evalTo,
		},
		Value: BoolValue(e),
	}
	d.deriveValueJudgement(child)
	j.premises = append(j.premises, child)
}

func (d *Deriver) deriveBinOpExpJudgement(j *ExpJudgement, e BinOpExp) {
	d.l.Info("deriveBinOpExpJudgement", "judgement", j, "exp", e)

	j.by = EBinOp
	ej := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont: ExpCont{
				Op:    e.Op,
				Right: e.Right,
				next:  j.cont,
			},
			evalTo: j.evalTo,
		},
		Exp: e.Left,
	}
	d.deriveExpJudgement(ej)
	j.premises = append(j.premises, ej)
}

func (d *Deriver) deriveIfExpJudgement(j *ExpJudgement, e IfExp) {
	d.l.Info("deriveIfExpJudgement", "judgement", j, "exp", e)

	j.by = EIf
	cont := IfCont{
		Then: e.Then,
		Else: e.Else,
		next: j.cont,
	}
	ej := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   cont,
			evalTo: j.evalTo,
		},
		Exp: e.Cond,
	}
	d.deriveExpJudgement(ej)
	j.premises = append(j.premises, ej)
}

func (d *Deriver) deriveParenExpJudgement(j *ExpJudgement, e ParenExp) {
	d.l.Info("deriveParenExpJudgement", "judgement", j, "exp", e)

	j.Exp = e.Inner
	d.deriveExpJudgement(j)
}
