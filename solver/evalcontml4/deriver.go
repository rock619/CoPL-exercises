package main

import (
	"fmt"
	"log/slog"
	"slices"
)

type Judgement interface {
	isJudgement()
	Literal() string
	By() Rule
	Premises() []Judgement
}

type Rule string

const (
	EInt    Rule = "E-Int"
	EBool   Rule = "E-Bool"
	EIf     Rule = "E-If"
	EBinOp  Rule = "E-BinOp"
	EVar    Rule = "E-Var"
	ELet    Rule = "E-Let"
	EFun    Rule = "E-Fun"
	EApp    Rule = "E-App"
	ELetRec Rule = "E-LetRec"
	ENil    Rule = "E-Nil"
	ECons   Rule = "E-Cons"
	EMatch  Rule = "E-Match"
	ELetCC  Rule = "E-LetCc"

	CRet       Rule = "C-Ret"
	CEvalR     Rule = "C-EvalR"
	CPlus      Rule = "C-Plus"
	CMinus     Rule = "C-Minus"
	CTimes     Rule = "C-Times"
	CLT        Rule = "C-Lt"
	CIfT       Rule = "C-IfT"
	CIfF       Rule = "C-IfF"
	CLetBody   Rule = "C-LetBody"
	CEvalArg   Rule = "C-EvalArg"
	CEvalFun   Rule = "C-EvalFun"
	CEvalFunR  Rule = "C-EvalFunR"
	CEvalFunC  Rule = "C-EvalFunC"
	CEvalConsR Rule = "C-EvalConsR"
	CCons      Rule = "C-Cons"
	CMatchNil  Rule = "C-MatchNil"
	CMatchCons Rule = "C-MatchCons"

	BPlus  Rule = "B-Plus"
	BMinus Rule = "B-Minus"
	BTimes Rule = "B-Times"
	BLt    Rule = "B-Lt"
)

type BaseJudgement struct {
	literal  string
	cont     Cont
	evalTo   Value
	by       Rule
	premises []Judgement
}

func (j *BaseJudgement) Literal() string {
	return j.literal
}

func (j *BaseJudgement) EvalTo() Value {
	return j.evalTo
}

func (j *BaseJudgement) By() Rule {
	return j.by
}

func (j *BaseJudgement) Premises() []Judgement {
	return j.premises
}

func (j *BaseJudgement) String() string {
	return fmt.Sprintf("%+v", *j)
}

type ValueJudgement struct {
	*BaseJudgement
	Value Value
}

func (ValueJudgement) isJudgement() {}

func (j *ValueJudgement) String() string {
	conts := ""
	for c := j.cont; c != nil; c = c.Next() {
		if _, ok := c.(TerminalCont); ok {
			break
		}
		if conts == "" {
			conts = c.String()
		} else {
			conts = fmt.Sprintf("%s >> %s", conts, c.String())
		}
	}
	if conts == "" {
		conts = "_"
	}
	return fmt.Sprintf("%s => %s evalto %s by %s", j.Value, conts, j.evalTo, j.By())
}

type ExpJudgement struct {
	*BaseJudgement
	Env Env
	Exp Exp
}

func (ExpJudgement) isJudgement() {}

func (j *ExpJudgement) String() string {
	conts := ""
	for c := j.cont; c != nil; c = c.Next() {
		if _, ok := c.(TerminalCont); ok {
			break
		}

		conts = fmt.Sprintf("%s >> %s", conts, c.String())
	}
	return fmt.Sprintf("%s %s%s evalto %s by %s", envPrefix(j.Env), j.Exp, conts, j.evalTo, j.By())
}

type BinOpJudgement struct {
	*BaseJudgement
	Left  Value
	Op    Op
	Right Value
}

func NewBinOpJudgement(left Value, op Op, right Value) *BinOpJudgement {
	j := &BinOpJudgement{
		BaseJudgement: &BaseJudgement{},
		Left:          left,
		Op:            op,
		Right:         right,
	}

	switch op {
	case OpPlus:
		j.by = BPlus
		j.evalTo = IntValue(int(left.(IntValue)) + int(right.(IntValue)))
	case OpMinus:
		j.by = BMinus
		j.evalTo = IntValue(int(left.(IntValue)) - int(right.(IntValue)))
	case OpTimes:
		j.by = BTimes
		j.evalTo = IntValue(int(left.(IntValue)) * int(right.(IntValue)))
	case OpLT:
		j.by = BLt
		j.evalTo = BoolValue(left.(IntValue) < right.(IntValue))
	default:
		panic("unknown operator")
	}

	return j
}

func (BinOpJudgement) isJudgement() {}

func (j *BinOpJudgement) String() string {
	o := ""
	switch j.Op {
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
	return fmt.Sprintf("%s %s %s is %s by %s", j.Left, o, j.Right, j.EvalTo(), j.By())
}

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
	switch c := j.cont.(type) {
	case TerminalCont, nil:
		// v => _ evalto v by C-Ret {}
		j.by = CRet
	case BinOpExpCont:
		d.deriveBinOpExpCont(j, c)
	case BinOpValueCont:
		d.deriveBinOpValueCont(j, c)
	case IfCont:
		d.deriveIfCont(j, c)
	case LetCont:
		d.deriveLetCont(j, c)
	case AppExpCont:
		d.deriveAppExpCont(j, c)
	case AppValueCont:
		d.deriveAppValueCont(j, c)
	case ConsExpCont:
		d.deriveConsExpCont(j, c)
	case ConsValueCont:
		d.deriveConsValueCont(j, c)
	case MatchCont:
		d.deriveMatchCont(j, c)
	default:
		d.l.Error("deriveValueJudgement: unknown continuation type", "type", fmt.Sprintf("%T", c), "judgement", j)
	}
}

// deriveBinOpExpCont
//
//	v1 => {𝓔 |- _ op e} >> k evalto v2 by C-EvalR {
//		𝓔 |- e >> {v1 op _} >> k evalto v2
//	}
func (d *Deriver) deriveBinOpExpCont(j *ValueJudgement, c BinOpExpCont) {
	d.l.Info("deriveBinOpExpCont", "judgement", j, "exp", c)

	j.by = CEvalR
	cont2 := c.Next()
	cont1 := BinOpValueCont{
		Left: j.Value,
		Op:   c.Op,
		next: cont2,
	}
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   cont1,
			evalTo: j.evalTo,
		},
		Env: slices.Clone(c.Env),
		Exp: c.Right,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveBinOpValueCont
//
//	i2 => {i1 op _} >> k evalto v by C-? {
//		i1 op i2 is i3
//		i3 => k evalto v
//	}
func (d *Deriver) deriveBinOpValueCont(j *ValueJudgement, c BinOpValueCont) {
	d.l.Info("deriveBinOpValueCont", "judgement", j, "value", c)

	left, ok := c.Left.(IntValue)
	if !ok {
		d.l.Error("deriveBinOpValueCont: left value is not an IntValue", "left", c.Left)
		return
	}
	right, ok := j.Value.(IntValue)
	if !ok {
		d.l.Error("deriveBinOpValueCont: right value is not an IntValue", "right", j.Value)
		return
	}

	bj := NewBinOpJudgement(left, c.Op, right)
	d.l.Info("deriveValueJudgement: binop judgement", "judgement", bj, "bj.EvalTo()", bj.EvalTo())
	vj := &ValueJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   c.Next(),
			evalTo: j.EvalTo(),
		},
		Value: bj.EvalTo(),
	}
	if vj.cont == nil {
		vj.cont = TerminalCont{}
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
		j.by = CLT
	default:
		panic("unknown operator")
	}
}

// deriveIfCont
//
//	true => {if _ then e1 else e2} >> k evalto v by C-IfT {
//		e1 >> k evalto v
//	}
//
//	false => {if e1 then e2 else e3} >> k evalto v by C-IfF {
//		e2 >> k evalto v
//	}
func (d *Deriver) deriveIfCont(j *ValueJudgement, c IfCont) {
	d.l.Info("deriveIfCont", "judgement", j, "if", c)

	j.by = CIfF
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   c.Next(),
			evalTo: j.evalTo,
		},
		Env: slices.Clone(c.Env),
		Exp: c.Else,
	}
	if j.Value.(BoolValue) {
		j.by = CIfT
		child.Exp = c.Then
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveLetCont
//
//	v1 => {𝓔 |- let x = _ in e} >> k evalto v2 by C-LetBody
//		𝓔, x = v1 |- e >> k evalto v2
func (d *Deriver) deriveLetCont(j *ValueJudgement, c LetCont) {
	d.l.Info("deriveLetCont", "judgement", j, "let", c)

	j.by = CLetBody
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   c.Next(),
			evalTo: j.evalTo,
		},
		Env: append(slices.Clone(c.Env), Bind{Var: c.Var, Value: j.Value}),
		Exp: c.Body,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveAppExpCont
//
//	v1 => {𝓔 |- _ e} >> k evalto v by C-EvalFun {
//		𝓔 |- e >> {v1 _} >> k evalto v
//	}
func (d *Deriver) deriveAppExpCont(j *ValueJudgement, c AppExpCont) {
	d.l.Info("deriveAppExpCont", "judgement", j, "exp", c)

	j.by = CEvalArg
	cont2 := c.Next()
	cont1 := AppValueCont{
		Fun:  j.Value,
		next: cont2,
	}
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   cont1,
			evalTo: j.evalTo,
		},
		Env: slices.Clone(c.Env),
		Exp: c.Arg,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveAppValueCont
func (d *Deriver) deriveAppValueCont(j *ValueJudgement, c AppValueCont) {
	switch f := c.Fun.(type) {
	case FunValue:
		d.deriveAppFunValueCont(j, c, f)
	case RecFunValue:
		d.deriveAppRecFunValueCont(j, c, f)
	case ContValue:
		d.deriveAppContValueCont(j, f)
	default:
		d.l.Error("deriveAppValueCont: unknown function type", "type", fmt.Sprintf("%T", f), "fun", f)
	}
}

// deriveAppFunValueCont
//
//	v1 => {(𝓔)[fun x -> e] _} >> k evalto v2 by C-EvalFun {
//		𝓔, x = v1 |- e >> k evalto v2
//	}
func (d *Deriver) deriveAppFunValueCont(j *ValueJudgement, c AppValueCont, f FunValue) {
	d.l.Info("deriveAppFunValueCont", "judgement", j, "fun", f)

	j.by = CEvalFun
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   c.Next(),
			evalTo: j.evalTo,
		},
		Env: append(slices.Clone(f.Env), Bind{Var: f.Fun.Param, Value: j.Value}),
		Exp: f.Fun.Body,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveAppRecFunValueCont
//
//	v1 => {(𝓔)[rec x = fun y -> e] _} >> k evalto v2 by C-EvalFunR {
//		𝓔, x = (𝓔)[rec x = fun y -> e], y = v1 |- e >> k evalto v2
//	}
func (d *Deriver) deriveAppRecFunValueCont(j *ValueJudgement, c AppValueCont, f RecFunValue) {
	d.l.Info("deriveAppRecFunValueCont", "judgement", j, "recfun", f)

	j.by = CEvalFunR
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   c.Next(),
			evalTo: j.evalTo,
		},
		Env: append(
			slices.Clone(f.Env),
			Bind{Var: f.Fun.Name, Value: f},
			Bind{Var: f.Fun.Fun.Param, Value: j.Value},
		),
		Exp: f.Fun.Fun.Body,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveAppContValueCont
//
//	v1 => {[k1] _} >> k2 evalto v2 by C-EvalFunC {
//		v1 => k1 evalto v2
//	}
func (d *Deriver) deriveAppContValueCont(j *ValueJudgement, f ContValue) {
	d.l.Info("deriveAppContValueCont", "judgement", j, "cont", f)

	j.by = CEvalFunC
	child := &ValueJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   f.Cont,
			evalTo: j.evalTo,
		},
		Value: j.Value,
	}
	d.deriveValueJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveConsExpCont
//
//	v1 => {𝓔 |- _ :: e} >> k evalto v2 by C-EvalConsR {
//		𝓔 |- e >> {v1 _} >> k evalto v2
//	}
func (d *Deriver) deriveConsExpCont(j *ValueJudgement, c ConsExpCont) {
	d.l.Info("deriveConsExpCont", "judgement", j, "exp", c)

	j.by = CEvalConsR
	cont2 := c.Next()
	cont1 := ConsValueCont{
		Head: j.Value,
		next: cont2,
	}
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   cont1,
			evalTo: j.evalTo,
		},
		Env: slices.Clone(c.Env),
		Exp: c.Tail,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveConsValueCont
//
//	v2 => {v1 :: _} >> k evalto v3 by C-Cons {
//		v1 :: v2 => k evalto v3
//	}
func (d *Deriver) deriveConsValueCont(j *ValueJudgement, c ConsValueCont) {
	d.l.Info("deriveConsValueCont", "judgement", j, "value", c)

	j.by = CCons
	child := &ValueJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   c.Next(),
			evalTo: j.evalTo,
		},
		Value: ConsValue{
			Head: c.Head,
			Tail: j.Value,
		},
	}
	d.deriveValueJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveMatchCont
func (d *Deriver) deriveMatchCont(j *ValueJudgement, c MatchCont) {
	switch v := j.Value.(type) {
	case NilValue:
		d.deriveMatchNilCont(j, c)
	case ConsValue:
		d.deriveMatchConsCont(j, c, v)
	default:
		d.l.Error("deriveMatchCont: unknown value type", "type", fmt.Sprintf("%T", j.Value), "value", j.Value)
	}
}

// deriveMatchNilCont
//
//	[] => {𝓔 |- match _ with [] -> e1 | x :: y -> e2} >> k evalto v by C-MatchNil {
//		𝓔 |- e2 >> k evalto v
//	}
func (d *Deriver) deriveMatchNilCont(j *ValueJudgement, c MatchCont) {
	d.l.Info("deriveMatchNilCont", "judgement", j, "match", c)

	j.by = CMatchNil
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   c.Next(),
			evalTo: j.evalTo,
		},
		Env: slices.Clone(c.Env),
		Exp: c.NilCase,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveMatchConsCont
//
//	v1 :: v2 => {𝓔 |- match _ with [] -> e1 | x :: y -> e2} >> k evalto v by C-MatchCons {
//		𝓔, x = v1, y = v2 |- e2 >> k evalto v
//	}
func (d *Deriver) deriveMatchConsCont(j *ValueJudgement, c MatchCont, v ConsValue) {
	d.l.Info("deriveMatchConsCont", "judgement", j, "match", c, "cons", v)

	j.by = CMatchCons
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   c.Next(),
			evalTo: j.evalTo,
		},
		Env: append(
			slices.Clone(c.Env),
			Bind{Var: c.HeadVar, Value: v.Head},
			Bind{Var: c.TailVar, Value: v.Tail},
		),
		Exp: c.ConsCase,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

func (d *Deriver) deriveExpJudgement(j *ExpJudgement) {
	switch e := j.Exp.(type) {
	case IntExp:
		d.deriveIntExpJudgement(j, e)
	case BoolExp:
		d.deriveBoolExpJudgement(j, e)
	case IfExp:
		d.deriveIfExpJudgement(j, e)
	case BinOpExp:
		d.deriveBinOpExpJudgement(j, e)
	case VarExp:
		d.deriveVarExpJudgement(j, e)
	case LetExp:
		d.deriveLetExpJudgement(j, e)
	case FunExp:
		d.deriveFunExpJudgement(j, e)
	case AppExp:
		d.deriveAppExpJudgement(j, e)
	case LetRecExp:
		d.deriveLetRecExpJudgement(j, e)
	case NilExp:
		d.deriveNilExpJudgement(j, e)
	case ConsExp:
		d.deriveConsExpJudgement(j, e)
	case MatchExp:
		d.deriveMatchExpJudgement(j, e)
	case LetCCExp:
		d.deriveLetCCExpJudgement(j, e)
	case ParenExp:
		d.deriveParenExpJudgement(j, e)
	default:
		d.l.Error("deriveExpJudgement: unknown expression type", "type", fmt.Sprintf("%T", j.Exp), "exp", j.Exp)
	}
}

// deriveIntExpJudgement
//
//	𝓔 |- i >> k evalto v by E-Int {
//		i => k evalto v
//	}
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

// deriveBoolExpJudgement
//
//	𝓔 |- b >> k evalto v by E-Bool {
//		b => k evalto v
//	}
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

// deriveIfExpJudgement
//
//	𝓔 |- if e1 then e2 else e3 >> k evalto v by E-If {
//		𝓔 |- e1 >> {𝓔 |- if _ then e2 else e3} >> k evalto v
//	}
func (d *Deriver) deriveIfExpJudgement(j *ExpJudgement, e IfExp) {
	d.l.Info("deriveIfExpJudgement", "judgement", j, "exp", e)

	j.by = EIf
	ej := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont: IfCont{
				Env:  slices.Clone(j.Env),
				Then: e.Then,
				Else: e.Else,
				next: j.cont,
			},
			evalTo: j.evalTo,
		},
		Env: slices.Clone(j.Env),
		Exp: e.Cond,
	}
	d.deriveExpJudgement(ej)
	j.premises = append(j.premises, ej)
}

// deriveBinOpExpJudgement
//
//	𝓔 |- e1 op e2 >> k evalto v by E-BinOp {
//		𝓔 |- e1 >> {𝓔 |- _ op e2} >> k evalto v
//	}
func (d *Deriver) deriveBinOpExpJudgement(j *ExpJudgement, e BinOpExp) {
	d.l.Info("deriveBinOpExpJudgement", "judgement", j, "exp", e)

	j.by = EBinOp
	ej := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont: BinOpExpCont{
				Env:   slices.Clone(j.Env),
				Op:    e.Op,
				Right: e.Right,
				next:  j.cont,
			},
			evalTo: j.evalTo,
		},
		Env: slices.Clone(j.Env),
		Exp: e.Left,
	}
	d.deriveExpJudgement(ej)
	j.premises = append(j.premises, ej)
}

// deriveVarExpJudgement
//
//	𝓔 |- x >> k evalto v2 by E-Var {
//		(𝓔(x) = v1) v1 => k evalto v2
//	}
func (d *Deriver) deriveVarExpJudgement(j *ExpJudgement, e VarExp) {
	d.l.Info("deriveVarExpJudgement", "judgement", j, "exp", e)

	j.by = EVar

	var v Value
	for _, b := range slices.Backward(j.Env) {
		if b.Var == string(e) {
			v = b.Value
			break
		}
	}

	child := &ValueJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   j.cont,
			evalTo: j.evalTo,
		},
		Value: v,
	}
	d.deriveValueJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveLetExpJudgement
//
//	𝓔 |- let x = e1 in e2 >> k evalto v by E-Let {
//		𝓔 |- e1 >> {𝓔 |- let x = _ in e2} >> k evalto v
//	}
func (d *Deriver) deriveLetExpJudgement(j *ExpJudgement, e LetExp) {
	d.l.Info("deriveLetExpJudgement", "judgement", j, "exp", e)

	j.by = ELet
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont: LetCont{
				Env:  slices.Clone(j.Env),
				Var:  e.Var,
				Body: e.Body,
				next: j.cont,
			},
			evalTo: j.evalTo,
		},
		Env: slices.Clone(j.Env),
		Exp: e.Bind,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveFunExpJudgement
//
//	𝓔 |- fun x -> e >> k evalto v by E-Fun {
//		(𝓔)[fun x -> e] => k evalto v
//	}
func (d *Deriver) deriveFunExpJudgement(j *ExpJudgement, e FunExp) {
	d.l.Info("deriveFunExpJudgement", "judgement", j, "exp", e)

	j.by = EFun
	child := &ValueJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   j.cont,
			evalTo: j.evalTo,
		},
		Value: FunValue{
			Env: slices.Clone(j.Env),
			Fun: e.Fun,
		},
	}
	d.deriveValueJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveAppExpJudgement
//
//	𝓔 |- e1 e2 >> k evalto v by E-App {
//		𝓔 |- e1 >> {𝓔 |- _ e2} >> k evalto v
//	}
func (d *Deriver) deriveAppExpJudgement(j *ExpJudgement, e AppExp) {
	d.l.Info("deriveAppExpJudgement", "judgement", j, "exp", e)

	j.by = EApp
	ej := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont: AppExpCont{
				Env:  slices.Clone(j.Env),
				Arg:  e.Arg,
				next: j.cont,
			},
			evalTo: j.evalTo,
		},
		Env: slices.Clone(j.Env),
		Exp: e.Fun,
	}
	d.deriveExpJudgement(ej)
	j.premises = append(j.premises, ej)
}

// deriveLetRecExpJudgement
//
//	𝓔 |- let rec x = fun y -> e1 in e2 >> k evalto v by E-LetRec {
//		𝓔, x = (𝓔)[rec x = fun y -> e1] |- e2 >> k evalto v
//	}
func (d *Deriver) deriveLetRecExpJudgement(j *ExpJudgement, e LetRecExp) {
	d.l.Info("deriveLetRecExpJudgement", "judgement", j, "exp", e)

	j.by = ELetRec
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   j.cont,
			evalTo: j.evalTo,
		},
		Env: append(
			slices.Clone(j.Env),
			Bind{
				Var: e.RecFun.Name,
				Value: RecFunValue{
					Env: slices.Clone(j.Env),
					Fun: e.RecFun,
				},
			},
		),
		Exp: e.Body,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveNilExpJudgement
//
//	𝓔 |- [] >> k evalto v by E-Nil {
//		[] => k evalto v
//	}
func (d *Deriver) deriveNilExpJudgement(j *ExpJudgement, e NilExp) {
	d.l.Info("deriveNilExpJudgement", "judgement", j, "exp", e)

	j.by = ENil
	child := &ValueJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   j.cont,
			evalTo: j.EvalTo(),
		},
		Value: NilValue{},
	}
	d.deriveValueJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveConsExpJudgement
//
//	𝓔 |- e1 :: e2 >> k evalto v by E-Cons {
//		𝓔 |- e1 >> {𝓔 |- _ :: e2} >> k evalto v
//	}
func (d *Deriver) deriveConsExpJudgement(j *ExpJudgement, e ConsExp) {
	d.l.Info("deriveConsExpJudgement", "judgement", j, "exp", e)

	j.by = ECons
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont: ConsExpCont{
				Env:  slices.Clone(j.Env),
				Tail: e.Tail,
				next: j.cont,
			},
			evalTo: j.evalTo,
		},
		Env: slices.Clone(j.Env),
		Exp: e.Head,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveMatchExpJudgement
//
//	𝓔 |- match e1 with [] -> e2 | x :: y -> e3 >> k evalto v by E-Match {
//		𝓔 |- e1 >> {𝓔 |- match _ with [] -> e2 | x :: y -> e3} >> k evalto v
//	}
func (d *Deriver) deriveMatchExpJudgement(j *ExpJudgement, e MatchExp) {
	d.l.Info("deriveMatchExpJudgement", "judgement", j, "exp", e)

	j.by = EMatch
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont: MatchCont{
				Env:      slices.Clone(j.Env),
				NilCase:  e.NilCase,
				HeadVar:  e.HeadVar,
				TailVar:  e.TailVar,
				ConsCase: e.ConsCase,
				next:     j.cont,
			},
			evalTo: j.evalTo,
		},
		Env: slices.Clone(j.Env),
		Exp: e.Matched,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

// deriveLetCCExpJudgement
//
//	𝓔 |- letcc x in e >> k evalto v by E-LetCC {
//		𝓔, x = [k] |- e >> k evalto v
//	}
func (d *Deriver) deriveLetCCExpJudgement(j *ExpJudgement, e LetCCExp) {
	d.l.Info("deriveLetCCExpJudgement", "judgement", j, "exp", e, "j.cont.Clone()", j.cont.Clone())

	j.by = ELetCC
	child := &ExpJudgement{
		BaseJudgement: &BaseJudgement{
			cont:   j.cont,
			evalTo: j.evalTo,
		},
		Env: append(slices.Clone(j.Env), Bind{Var: e.Var, Value: ContValue{Cont: j.cont.Clone()}}),
		Exp: e.Body,
	}
	d.deriveExpJudgement(child)
	j.premises = append(j.premises, child)
}

func (d *Deriver) deriveParenExpJudgement(j *ExpJudgement, e ParenExp) {
	d.l.Info("deriveParenExpJudgement", "judgement", j, "exp", e)

	j.Exp = e.Inner
	d.deriveExpJudgement(j)
}
