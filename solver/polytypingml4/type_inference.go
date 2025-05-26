package main

import (
	"fmt"
	"log/slog"
	"slices"
)

type TypeInferrer struct {
	typeVarGen *TypeVarGenerator
	l          *slog.Logger
}

func NewTypeInferrer(gen *TypeVarGenerator, logger *slog.Logger) *TypeInferrer {
	return &TypeInferrer{
		typeVarGen: gen,
		l:          logger,
	}
}

func (i *TypeInferrer) NewTypeVar() TypeVar {
	return i.typeVarGen.NewTypeVar()
}

func (i *TypeInferrer) Do(env Env, exp Exp) (InferenceResult, error) {
	switch exp := exp.(type) {
	case IntExp:
		return i.intExp(env, exp)
	case BoolExp:
		return i.boolExp(env, exp)
	case VarExp:
		return i.varExp(env, exp)
	case BinOpExp:
		return i.binOpExp(env, exp)
	case IfExp:
		return i.ifExp(env, exp)
	case LetExp:
		return i.letExp(env, exp)
	case FunExp:
		return i.funExp(env, exp)
	case AppExp:
		return i.appExp(env, exp)
	case LetRecExp:
		return i.letRecExp(env, exp)
	case NilExp:
		return i.nilExp(env, exp)
	case ConsExp:
		return i.consExp(env, exp)
	case MatchExp:
		return i.matchExp(env, exp)
	default:
		return InferenceResult{}, fmt.Errorf("unknown expression type: %T", exp)
	}
}

func (i *TypeInferrer) intExp(env Env, exp IntExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.intExp", "env", env, "exp", exp)

	return NewInferenceResult(env, exp, IntType{}, nil, TInt, nil), nil
}

func (i *TypeInferrer) boolExp(env Env, exp BoolExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.boolExp", "env", env, "exp", exp)

	return NewInferenceResult(env, exp, BoolType{}, nil, TBool, nil), nil
}

func (i *TypeInferrer) varExp(env Env, exp VarExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.varExp", "env", env, "exp", exp)

	tyScheme, ok := env.Get(Var(exp))
	if !ok {
		return InferenceResult{}, fmt.Errorf("variable %s not found in environment", exp)
	}
	t := tyScheme.Type
	for _, v := range tyScheme.TypeVars() {
		tv := i.NewTypeVar()
		mapping := TypeMapping{
			TypeVar: v,
			Type:    tv,
		}
		t = mapping.ApplyToType(t)
	}
	return NewInferenceResult(env, exp, t, nil, TVar, nil), nil
}

func (i *TypeInferrer) binOpExp(env Env, exp BinOpExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.binOpExp", "env", env, "exp", exp)

	left, err := i.Do(env, exp.Left)
	if err != nil {
		return InferenceResult{}, err
	}
	right, err := i.Do(env, exp.Right)
	if err != nil {
		return InferenceResult{}, err
	}
	sub, err := Unify(append(slices.Concat(left.Substitution, right.Substitution).TypeEquations(),
		TypeEquation{left.Type, IntType{}},
		TypeEquation{right.Type, IntType{}},
	))
	if err != nil {
		return InferenceResult{}, fmt.Errorf("unification failed: %w", err)
	}
	if exp.Op == OpLT {
		return NewInferenceResult(env, exp, BoolType{}, sub, TLT, []Judgement{left.Judgement, right.Judgement}), nil
	}
	return NewInferenceResult(env, exp, IntType{}, sub, exp.Op.Rule(), []Judgement{left.Judgement, right.Judgement}), nil
}

func (i *TypeInferrer) ifExp(env Env, exp IfExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.ifExp", "env", env, "exp", exp)

	cond, err := i.Do(env, exp.Cond)
	if err != nil {
		return InferenceResult{}, err
	}
	then, err := i.Do(env, exp.Then)
	if err != nil {
		return InferenceResult{}, err
	}
	els, err := i.Do(env, exp.Else)
	if err != nil {
		return InferenceResult{}, err
	}
	sub, err := Unify(append(slices.Concat(cond.Substitution, then.Substitution, els.Substitution).TypeEquations(),
		TypeEquation{cond.Type, BoolType{}},
		TypeEquation{then.Type, els.Type},
	))
	if err != nil {
		return InferenceResult{}, fmt.Errorf("unification failed: %w", err)
	}
	return NewInferenceResult(env, exp, sub.ApplyToType(then.Type), sub, TIf, []Judgement{
		cond.Judgement, then.Judgement, els.Judgement,
	}), nil
}

func (i *TypeInferrer) letExp(env Env, exp LetExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.letExp", "env", env, "exp", exp)

	bind, err := i.Do(env, exp.Bind)
	if err != nil {
		return InferenceResult{}, err
	}
	tyScheme := Closure(bind.Type, bind.Substitution.ApplyToEnv(env))
	body, err := i.Do(
		append(env, Bind{Var: exp.Var, TypeScheme: tyScheme}),
		exp.Body,
	)
	if err != nil {
		return InferenceResult{}, err
	}
	sub, err := Unify(slices.Concat(bind.Substitution, body.Substitution).TypeEquations())
	if err != nil {
		return InferenceResult{}, fmt.Errorf("unification failed: %w", err)
	}
	return NewInferenceResult(env, exp, sub.ApplyToType(body.Type), sub, TLet, []Judgement{
		bind.Judgement, body.Judgement,
	},
	), nil
}

func (i *TypeInferrer) funExp(env Env, exp FunExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.funExp", "env", env, "exp", exp)

	paramType := i.NewTypeVar()
	result, err := i.Do(append(env, Bind{Var: exp.Param, TypeScheme: TypeScheme{Type: paramType}}), exp.Body)
	if err != nil {
		return InferenceResult{}, err
	}
	return NewInferenceResult(env, exp,
		FunType{Param: result.Substitution.ApplyToType(paramType), Return: result.Type},
		result.Substitution,
		TAbs,
		[]Judgement{result.Judgement},
	), nil
}

func (i *TypeInferrer) appExp(env Env, exp AppExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.appExp", "env", env, "exp", exp)

	fun, err := i.Do(env, exp.Fun)
	if err != nil {
		return InferenceResult{}, err
	}
	arg, err := i.Do(env, exp.Arg)
	if err != nil {
		return InferenceResult{}, err
	}
	returnType := i.NewTypeVar()
	sub, err := Unify(append(slices.Concat(fun.Substitution, arg.Substitution).TypeEquations(),
		TypeEquation{
			Left:  fun.Type,
			Right: FunType{Param: arg.Type, Return: returnType},
		},
	))
	if err != nil {
		return InferenceResult{}, fmt.Errorf("unification failed: %w", err)
	}
	return NewInferenceResult(env, exp,
		sub.ApplyToType(returnType),
		sub,
		TApp,
		[]Judgement{fun.Judgement, arg.Judgement},
	), nil
}

func (i *TypeInferrer) letRecExp(env Env, exp LetRecExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.letRecExp", "env", env, "exp", exp)

	varType := i.NewTypeVar()
	paramType := i.NewTypeVar()
	bind, err := i.Do(
		append(env,
			Bind{Var: exp.Var, TypeScheme: TypeScheme{Type: varType}},
			Bind{Var: exp.Bind.Param, TypeScheme: TypeScheme{Type: paramType}},
		),
		exp.Bind.Body,
	)
	if err != nil {
		return InferenceResult{}, err
	}
	unifiedBindSub, err := Unify(append(bind.Substitution.TypeEquations(),
		TypeEquation{Left: varType, Right: FunType{Param: paramType, Return: bind.Type}},
	))
	tyScheme := Closure(unifiedBindSub.ApplyToType(varType), unifiedBindSub.ApplyToEnv(env))
	body, err := i.Do(append(env, Bind{Var: exp.Var, TypeScheme: tyScheme}), exp.Body)
	if err != nil {
		return InferenceResult{}, err
	}
	resultSub, err := Unify(slices.Concat(unifiedBindSub.TypeEquations(), body.Substitution.TypeEquations()))
	if err != nil {
		return InferenceResult{}, fmt.Errorf("unification failed: %w", err)
	}
	return NewInferenceResult(env, exp, resultSub.ApplyToType(body.Type), resultSub, TLetRec,
		[]Judgement{bind.Judgement, body.Judgement},
	), nil
}

func (i *TypeInferrer) nilExp(env Env, exp NilExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.nilExp", "env", env, "exp", exp)

	return NewInferenceResult(env, exp, ListType{Elem: i.NewTypeVar()}, nil, TNil, nil), nil
}

func (i *TypeInferrer) consExp(env Env, exp ConsExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.consExp", "env", env, "exp", exp)

	head, err := i.Do(env, exp.Head)
	if err != nil {
		return InferenceResult{}, err
	}
	tail, err := i.Do(env, exp.Tail)
	if err != nil {
		return InferenceResult{}, err
	}
	sub, err := Unify(append(slices.Concat(head.Substitution, tail.Substitution).TypeEquations(),
		TypeEquation{tail.Type, ListType{Elem: head.Type}},
	))
	if err != nil {
		return InferenceResult{}, fmt.Errorf("unification failed: %w", err)
	}
	return NewInferenceResult(env, exp, sub.ApplyToType(tail.Type), sub, TCons, []Judgement{head.Judgement, tail.Judgement}), nil
}

func (i *TypeInferrer) matchExp(env Env, exp MatchExp) (InferenceResult, error) {
	i.l.Debug("TypeInferrer.matchExp", "env", env, "exp", exp)

	arg, err := i.Do(env, exp.Arg)
	if err != nil {
		return InferenceResult{}, err
	}
	nl, err := i.Do(env, exp.NilCase)
	if err != nil {
		return InferenceResult{}, err
	}
	elemType := i.NewTypeVar()
	cons, err := i.Do(
		append(env,
			Bind{Var: exp.HeadVar, TypeScheme: TypeScheme{Type: elemType}},
			Bind{Var: exp.TailVar, TypeScheme: TypeScheme{Type: ListType{Elem: elemType}}},
		),
		exp.ConsCase,
	)
	if err != nil {
		return InferenceResult{}, err
	}
	sub, err := Unify(append(slices.Concat(arg.Substitution, nl.Substitution, cons.Substitution).TypeEquations(),
		TypeEquation{arg.Type, ListType{Elem: elemType}},
		TypeEquation{nl.Type, cons.Type},
	))
	if err != nil {
		return InferenceResult{}, fmt.Errorf("unification failed: %w", err)
	}
	nilType := sub.ApplyToType(nl.Type)
	return NewInferenceResult(env, exp, nilType, sub, TMatch, []Judgement{
		arg.Judgement, nl.Judgement, cons.Judgement,
	}), nil
}

func Closure(t Type, env Env) TypeScheme {
	envFTV := env.FTV()
	vars := slices.DeleteFunc(t.FTV(), func(v TypeVar) bool {
		return slices.Contains(envFTV, v)
	})
	return NewTypeScheme(vars, t)
}

type InferenceResult struct {
	Substitution Substitution
	Type         Type
	Judgement    Judgement
}

func NewInferenceResult(env Env, exp Exp, t Type, sub Substitution, by Rule, premises []Judgement) InferenceResult {
	return InferenceResult{
		Type:         t,
		Substitution: sub,
		Judgement: Judgement{
			Env:      env,
			Exp:      exp,
			Type:     t,
			Sub:      sub,
			By:       by,
			Premises: premises,
		},
	}
}

type TypeVarGenerator struct {
	nextID int
	used   map[TypeVar]struct{}
}

func NewTypeVarGenerator() *TypeVarGenerator {
	return &TypeVarGenerator{
		nextID: 0,
		used:   make(map[TypeVar]struct{}),
	}
}

func (g *TypeVarGenerator) NewTypeVar() TypeVar {
	for {
		tv := TypeVar(g.typeVarName(g.nextID))
		if _, exists := g.used[tv]; !exists {
			g.used[tv] = struct{}{}
			return tv
		}
		g.nextID++
	}
}

func (g *TypeVarGenerator) SetUsed(tv TypeVar) {
	g.used[tv] = struct{}{}
}

func (g *TypeVarGenerator) typeVarName(id int) string {
	alpha := string(rune('a' + id%26))
	if id < 26 {
		return alpha
	}
	return fmt.Sprintf("%s%d", alpha, id/26)
}

func (g *TypeVarGenerator) SetUsedFromEnv(env Env) {
	for _, bind := range env {
		for _, tv := range bind.TypeScheme.TypeVars() {
			g.SetUsed(tv)
		}
	}
}

func (g *TypeVarGenerator) SetUsedFromType(t Type) {
	for _, tv := range t.FTV() {
		g.SetUsed(tv)
	}
}
