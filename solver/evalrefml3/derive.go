package main

import (
	"fmt"
	"log/slog"
	"strings"
)

type Judgement interface {
	isJudgement()
	EvalTo() Value
	By() Rule
	Premises() []Judgement
	String() string
}

type Rule string

const (
	EInt    Rule = "E-Int"
	EBool   Rule = "E-Bool"
	EIfT    Rule = "E-IfT"
	EIfF    Rule = "E-IfF"
	EPlus   Rule = "E-Plus"
	EMinus  Rule = "E-Minus"
	EMult   Rule = "E-Mult"
	ELT     Rule = "E-Lt"
	EVar    Rule = "E-Var"
	ELet    Rule = "E-Let"
	EFun    Rule = "E-Fun"
	EApp    Rule = "E-App"
	ELetRec Rule = "E-LetRec"
	EAppRec Rule = "E-AppRec"
	ERef    Rule = "E-Ref"
	EDeref  Rule = "E-Deref"
	EAssign Rule = "E-Assign"

	BPlus  Rule = "B-Plus"
	BMinus Rule = "B-Minus"
	BMult  Rule = "B-Mult"
	BLT    Rule = "B-Lt"
)

type BinOpJudgement struct {
	left   Value
	op     Op
	right  Value
	evalTo Value
	by     Rule
}

func NewBinOpJudgement(left Value, op Op, right Value) (*BinOpJudgement, error) {
	l, ok := left.(IntValue)
	if !ok {
		return nil, fmt.Errorf("NewBinOpJudgement: left operand is not an IntValue: got %v %T", left, left)
	}
	r, ok := right.(IntValue)
	if !ok {
		return nil, fmt.Errorf("NewBinOpJudgement: right operand is not an IntValue: got %v %T", right, right)
	}

	j := &BinOpJudgement{
		left:  left,
		op:    op,
		right: right,
	}

	switch op {
	case OpPlus:
		j.by = BPlus
		j.evalTo = l + r
	case OpMinus:
		j.by = BMinus
		j.evalTo = l - r
	case OpTimes:
		j.by = BMult
		j.evalTo = l * r
	case OpLT:
		j.by = BLT
		j.evalTo = BoolValue(l < r)
	default:
		return nil, fmt.Errorf("NewBinOpJudgement: unknown operator: %s", op)
	}
	return j, nil
}

func (BinOpJudgement) isJudgement() {}

func (j *BinOpJudgement) By() Rule {
	return j.by
}

func (j *BinOpJudgement) Premises() []Judgement {
	return nil
}

func (j *BinOpJudgement) EvalTo() Value {
	return j.evalTo
}

func (j *BinOpJudgement) String() string {
	o := ""
	switch j.op {
	case OpPlus:
		o = "plus"
	case OpMinus:
		o = "minus"
	case OpTimes:
		o = "times"
	case OpLT:
		o = "less than"
	default:
		panic("unknown operator")
	}
	return fmt.Sprintf("%s %s %s is %s by %s", j.left, o, j.right, j.EvalTo(), j.By())
}

type ExpJudgement struct {
	storeIn  Store
	env      Env
	exp      Exp
	evalTo   Value
	storeOut Store
	by       Rule
	premises []Judgement
}

func NewExpJudgement(storeIn Store, env Env, exp Exp) *ExpJudgement {
	return &ExpJudgement{
		storeIn: storeIn,
		env:     env,
		exp:     exp,
	}
}

func NewEvaledExpJudgement(storeIn Store, env Env, exp Exp, evalTo Value, storeOut Store) *ExpJudgement {
	return &ExpJudgement{
		storeIn:  storeIn,
		env:      env,
		exp:      exp,
		evalTo:   evalTo,
		storeOut: storeOut,
	}
}

func (ExpJudgement) isJudgement() {}

func (j *ExpJudgement) EvalTo() Value {
	return j.evalTo
}

func (j *ExpJudgement) By() Rule {
	return j.by
}

func (j *ExpJudgement) Premises() []Judgement {
	return j.premises
}

func (j *ExpJudgement) String() string {
	parts := make([]string, 0, 11)

	storeIn := j.storeIn.String()
	if len(storeIn) > 0 {
		parts = append(parts, storeIn, "/")
	}

	env := j.env.String()
	if len(env) > 0 {
		parts = append(parts, env)
	}

	parts = append(parts, "|-", j.exp.String(), "evalto")
	if j.EvalTo() == nil {
		parts = append(parts, "?")
	} else {
		parts = append(parts, j.EvalTo().String())
	}

	storeOut := j.storeOut.String()
	if len(storeOut) > 0 {
		parts = append(parts, "/", storeOut)
	}

	parts = append(parts, "by", string(j.By()))
	if strings.Contains(strings.Join(parts, " "), "error") {
		panic("ExpJudgement: error in string" + strings.Join(parts, " "))
	}

	return strings.Join(parts, " ")
}

func Derive(j Judgement, l *slog.Logger) Judgement {
	l.Info("Derive", "judgement", j)

	d := &Deriver{
		l: l,
	}
	return d.Do(j)
}

type Deriver struct {
	l     *slog.Logger
	store Store
}

func (d *Deriver) NextLoc() LocValue {
	if len(d.store) == 0 {
		return LocValue("@INVALID")
	}
	l := d.store[0].Loc
	d.store = d.store[1:]
	return l
}

func (d *Deriver) Do(j Judgement) Judgement {
	if ej, ok := j.(*ExpJudgement); ok {
		d.store = ej.storeOut
	}
	d.deriveExpJudgement(j)
	return j
}

func (d *Deriver) deriveExpJudgement(j Judgement) {
	ej, ok := j.(*ExpJudgement)
	if !ok {
		return
	}

	switch e := ej.exp.(type) {
	case IntExp:
		d.deriveIntExpJudgement(ej, e)
	case BoolExp:
		d.deriveBoolExpJudgement(ej, e)
	case IfExp:
		d.deriveIfExpJudgement(ej, e)
	case BinOpExp:
		d.deriveBinOpExpJudgement(ej, e)
	case VarExp:
		d.deriveVarExpJudgement(ej, e)
	case LetExp:
		d.deriveLetExpJudgement(ej, e)
	case FunExp:
		d.deriveFunExpJudgement(ej, e)
	case AppExp:
		d.deriveAppExpJudgement(ej, e)
	case LetRecExp:
		d.deriveLetRecExpJudgement(ej, e)
	case RefExp:
		d.deriveRefExpJudgement(ej, e)
	case DerefExp:
		d.deriveDerefExpJudgement(ej, e)
	case AssignExp:
		d.deriveAssignExpJudgement(ej, e)
	case ParenExp:
		d.deriveParenExpJudgement(ej, e)
	default:
		d.l.Error("deriveExp: unknown expression type", "type", fmt.Sprintf("%T", e))
	}
}

func (d *Deriver) deriveIntExpJudgement(ej *ExpJudgement, e IntExp) {
	d.l.Debug("deriveIntExpJudgement", "judgement", ej, "expression", e)

	ej.by = EInt
	ej.premises = nil
	ej.evalTo = IntValue(e)
	ej.storeOut = ej.storeIn.Clone()
}

func (d *Deriver) deriveBoolExpJudgement(ej *ExpJudgement, e BoolExp) {
	d.l.Debug("deriveBoolExpJudgement", "judgement", ej, "expression", e)

	ej.by = EBool
	ej.premises = nil
	ej.evalTo = BoolValue(e)
	ej.storeOut = ej.storeIn.Clone()
}

func (d *Deriver) deriveBinOpExpJudgement(ej *ExpJudgement, e BinOpExp) {
	d.l.Debug("deriveBinOpExpJudgement", "judgement", ej, "expression", e)
	p1 := NewExpJudgement(ej.storeIn.Clone(), ej.env.Clone(), e.Left)
	d.deriveExpJudgement(p1)
	p2 := NewExpJudgement(p1.storeOut.Clone(), ej.env.Clone(), e.Right)
	d.deriveExpJudgement(p2)
	p3, err := NewBinOpJudgement(p1.EvalTo(), e.Op, p2.EvalTo())
	if err != nil {
		d.l.Error("deriveBinOpExpJudgement: NewBinOpJudgement error", "error", err)
		return
	}
	ej.premises = []Judgement{p1, p2, p3}
	ej.evalTo = p3.EvalTo()
	switch e.Op {
	case OpPlus:
		ej.by = EPlus
	case OpMinus:
		ej.by = EMinus
	case OpTimes:
		ej.by = EMult
	case OpLT:
		ej.by = ELT
	default:
		d.l.Error("deriveBinOpExpJudgement: unknown operator", "operator", e.Op)
	}
	ej.storeOut = p2.storeOut.Clone()
}

func (d *Deriver) deriveVarExpJudgement(ej *ExpJudgement, e VarExp) {
	d.l.Debug("deriveVarExpJudgement: begin", "judgement", ej, "expression", e)

	v, ok := ej.env.Get(Var(e))
	if !ok {
		d.l.Error("deriveVarExpJudgement: variable not found in environment", "env", ej.env.Clone(), "var", e)
		return
	}
	ej.evalTo = v
	ej.by = EVar
	ej.storeOut = ej.storeIn.Clone()
	d.l.Debug("deriveVarExpJudgement: end", "judgement", ej, "expression", e)
}

func (d *Deriver) deriveLetExpJudgement(ej *ExpJudgement, e LetExp) {
	d.l.Debug("deriveLetExpJudgement", "judgement", ej, "expression", e, "bind", e.Bind)

	p1 := NewExpJudgement(ej.storeIn.Clone(), ej.env.Clone(), e.Bind)
	d.deriveExpJudgement(p1)

	p2Env := ej.env.Clone()
	if p1.EvalTo() == nil {
		d.l.Error("deriveLetExpJudgement: p1.EvalTo() is nil", "p1", p1)
		return
	}
	p2Env.Set(Bind{Var: e.Var, Value: p1.EvalTo()})
	p2 := NewExpJudgement(p1.storeOut.Clone(), p2Env, e.Body)
	d.deriveExpJudgement(p2)
	ej.evalTo = p2.EvalTo()
	ej.premises = []Judgement{p1, p2}
	ej.by = ELet
	ej.storeOut = p2.storeOut.Clone()
}

func (d *Deriver) deriveFunExpJudgement(ej *ExpJudgement, e FunExp) {
	d.l.Debug("deriveFunExpJudgement", "judgement", ej, "expression", e)

	ej.evalTo = FunValue{
		Env: ej.env.Clone(),
		Fun: e.Fun.Clone(),
	}
	ej.by = EFun
	ej.storeOut = ej.storeIn.Clone()
}

func (d *Deriver) deriveAppExpJudgement(ej *ExpJudgement, e AppExp) {
	d.l.Debug("deriveAppExpJudgement", "judgement", ej, "expression", e)

	p1 := NewExpJudgement(ej.storeIn.Clone(), ej.env.Clone(), e.Fun)
	d.deriveExpJudgement(p1)
	ej.premises = []Judgement{p1}
	switch f := p1.EvalTo().(type) {
	case FunValue:
		d.deriveAppExpFunJudgement(ej, e, f)
	case RecFunValue:
		d.deriveAppExpRecFunJudgement(ej, e, f)
	default:
		d.l.Error("deriveAppExpJudgement: function not found", "value", p1.EvalTo())
	}
	return
}

func (d *Deriver) deriveAppExpFunJudgement(ej *ExpJudgement, e AppExp, f FunValue) {
	d.l.Debug("deriveAppExpFunJudgement", "judgement", ej, "expression", e, "function", f)

	p1 := ej.premises[0].(*ExpJudgement)
	p2 := NewExpJudgement(p1.storeOut.Clone(), ej.env.Clone(), e.Arg)
	d.deriveExpJudgement(p2)
	p3Env := f.Env.Clone()
	p3Env.Set(Bind{Var: f.Fun.Param, Value: p2.EvalTo()})
	p3 := NewExpJudgement(p2.storeOut.Clone(), p3Env, f.Fun.Body)
	d.deriveExpJudgement(p3)
	ej.evalTo = p3.EvalTo()
	ej.premises = append(ej.premises, p2, p3)
	ej.by = EApp
	ej.storeOut = p3.storeOut.Clone()
}

func (d *Deriver) deriveAppExpRecFunJudgement(ej *ExpJudgement, e AppExp, f RecFunValue) {
	d.l.Debug("deriveAppExpRecFunJudgement", "judgement", ej, "expression", e, "function", f)

	p1 := ej.premises[0].(*ExpJudgement)
	p2 := NewExpJudgement(p1.storeOut.Clone(), ej.env.Clone(), e.Arg)
	d.deriveExpJudgement(p2)
	p3Env := f.Env.Clone()
	p3Env.Set(Bind{Var: f.Fun.Name, Value: f})
	p3Env.Set(Bind{Var: f.Fun.Fun.Param, Value: p2.EvalTo()})

	p3 := NewExpJudgement(p2.storeOut.Clone(), p3Env, f.Fun.Fun.Body)
	d.deriveExpJudgement(p3)
	ej.evalTo = p3.EvalTo()
	ej.premises = append(ej.premises, p2, p3)
	ej.by = EAppRec
	ej.storeOut = p3.storeOut.Clone()
}

func (d *Deriver) deriveIfExpJudgement(ej *ExpJudgement, e IfExp) {
	d.l.Debug("deriveIfExpJudgement", "judgement", ej, "expression", e)
	p1 := NewExpJudgement(ej.storeIn.Clone(), ej.env.Clone(), e.Cond)
	d.deriveExpJudgement(p1)
	condEvalTo, ok := p1.EvalTo().(BoolValue)
	if !ok {
		d.l.Error("deriveIfExpJudgement: condition eval to not a BoolValue", "value", p1.EvalTo())
		return
	}
	var p2 *ExpJudgement
	if condEvalTo {
		p2 = NewExpJudgement(p1.storeOut.Clone(), ej.env.Clone(), e.Then)
		ej.by = EIfT
	} else {
		p2 = NewExpJudgement(p1.storeOut.Clone(), ej.env.Clone(), e.Else)
		ej.by = EIfF
	}
	d.deriveExpJudgement(p2)
	ej.evalTo = p2.EvalTo()
	ej.premises = []Judgement{p1, p2}
	ej.storeOut = p2.storeOut.Clone()
}

func (d *Deriver) deriveLetRecExpJudgement(ej *ExpJudgement, e LetRecExp) {
	d.l.Debug("deriveLetRecExpJudgement", "judgement", ej, "expression", e)

	p1Env := ej.env.Clone()
	p1Env.Set(Bind{
		Var: e.RecFun.Name,
		Value: RecFunValue{
			Env: ej.env.Clone(),
			Fun: e.RecFun,
		},
	})
	p1 := NewExpJudgement(ej.storeIn.Clone(), p1Env, e.Body)
	d.deriveExpJudgement(p1)

	ej.premises = []Judgement{p1}
	ej.by = ELetRec
	ej.storeOut = p1.storeOut.Clone()
}

func (d *Deriver) deriveRefExpJudgement(ej *ExpJudgement, e RefExp) {
	d.l.Debug("deriveRefExpJudgement", "judgement", ej, "expression", e, "storeOut", ej.storeOut)

	p := NewExpJudgement(ej.storeIn.Clone(), ej.env.Clone(), e.Exp)
	d.deriveExpJudgement(p)

	loc := d.NextLoc()
	ej.storeOut = p.storeOut.Clone()
	ej.storeOut.Set(Assign{Loc: loc, Value: p.EvalTo()})
	ej.evalTo = loc
	ej.premises = []Judgement{p}
	ej.by = ERef
}

func (d *Deriver) deriveDerefExpJudgement(ej *ExpJudgement, e DerefExp) {
	d.l.Debug("deriveDerefExpJudgement", "judgement", ej, "expression", e)

	p := NewExpJudgement(ej.storeIn.Clone(), ej.env.Clone(), e.Exp)
	d.deriveExpJudgement(p)
	loc, ok := p.EvalTo().(LocValue)
	if !ok {
		d.l.Error("deriveDerefExpJudgement: deref eval to not a LocValue", "judgement", ej, "expression", e, "value", p.EvalTo())
		return
	}
	v, ok := p.storeOut.Get(loc)
	if !ok {
		d.l.Error("deriveDerefExpJudgement: deref location not found in store", "location", loc)
		return
	}
	ej.evalTo = v
	ej.premises = []Judgement{p}
	ej.by = EDeref
	ej.storeOut = p.storeOut.Clone()
}

func (d *Deriver) deriveAssignExpJudgement(ej *ExpJudgement, e AssignExp) {
	d.l.Debug("deriveAssignExpJudgement: begin", "judgement", ej, "expression", e)
	defer d.l.Debug("deriveAssignExpJudgement: end", "judgement", ej, "expression", e)

	p1 := NewExpJudgement(ej.storeIn.Clone(), ej.env.Clone(), e.Left)
	d.deriveExpJudgement(p1)
	loc, ok := p1.EvalTo().(LocValue)
	if !ok {
		d.l.Error("deriveAssignExpJudgement: AssignExp.Left eval to not a LocValue", "value", p1.EvalTo(), "judgement", ej, "expression", e)
		return
	}
	p2 := NewExpJudgement(p1.storeOut.Clone(), ej.env.Clone(), e.Right)
	d.deriveExpJudgement(p2)

	ej.evalTo = p2.EvalTo()
	ej.premises = []Judgement{p1, p2}
	ej.by = EAssign
	ej.storeOut = p2.storeOut.Clone()
	ej.storeOut.Set(Assign{Loc: loc, Value: p2.EvalTo()})
}

func (d *Deriver) deriveParenExpJudgement(ej *ExpJudgement, e ParenExp) {
	d.l.Debug("deriveParenExpJudgement", "judgement", ej, "expression", e)

	ej.exp = e.Inner
	d.deriveExpJudgement(ej)
}
