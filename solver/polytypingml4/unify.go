package main

import (
	"fmt"
	"slices"
)

func Unify(eqs []TypeEquation) (Substitution, error) {
	if len(eqs) == 0 {
		return nil, nil
	}
	eq := eqs[len(eqs)-1]
	eqs = eqs[:len(eqs)-1]
	if eq.Left.Equal(eq.Right) {
		return Unify(eqs)
	}

	switch l := eq.Left.(type) {
	case TypeVar:
		return unifyTypeVar(eqs, l, eq.Right)
	case BoolType, IntType:
		if r, ok := eq.Right.(TypeVar); ok {
			return unifyTypeVar(eqs, r, l)
		}
		return nil, fmt.Errorf("unification failed: %s != %s", eq.Left, eq.Right)
	case FunType:
		return unifyFunType(eqs, l, eq.Right)
	case ListType:
		return unifyListType(eqs, l, eq.Right)
	default:
		return nil, fmt.Errorf("unification failed: %s != %s", eq.Left, eq.Right)
	}
}

func unifyTypeVar(eqs []TypeEquation, v TypeVar, t Type) (Substitution, error) {
	if slices.ContainsFunc(t.FTV(), func(v2 TypeVar) bool { return v2.Equal(v) }) {
		return nil, fmt.Errorf("unification failed: %s occurs in %s", v, t)
	}

	m := TypeMapping{TypeVar: v, Type: t}
	eqs = m.ApplyToEquations(eqs)

	sub, err := Unify(m.ApplyToEquations(eqs))
	if err != nil {
		return nil, err
	}

	m.Type = sub.ApplyToType(m.Type)
	return append(sub, m), nil
}

func unifyFunType(eqs []TypeEquation, f FunType, t Type) (Substitution, error) {
	switch t := t.(type) {
	case TypeVar:
		return unifyTypeVar(eqs, t, f)
	case FunType:
		return Unify(append(eqs,
			TypeEquation{
				Left:  f.Param,
				Right: t.Param,
			},
			TypeEquation{
				Left:  f.Return,
				Right: t.Return,
			},
		))
	default:
		return nil, fmt.Errorf("unification failed: %s != %s", f, t)
	}
}

func unifyListType(eqs []TypeEquation, l ListType, t Type) (Substitution, error) {
	switch t := t.(type) {
	case TypeVar:
		return unifyTypeVar(eqs, t, l)
	case ListType:
		return Unify(append(eqs,
			TypeEquation{
				Left:  l.Elem,
				Right: t.Elem,
			},
		))
	default:
		return nil, fmt.Errorf("unification failed: %s != %s", l, t)
	}
}

type TypeEquation struct {
	Left  Type
	Right Type
}

func (e TypeEquation) String() string {
	return fmt.Sprintf("%s = %s", e.Left, e.Right)
}

type Substitution []TypeMapping

func (s Substitution) TypeEquations() []TypeEquation {
	eqs := make([]TypeEquation, len(s))
	for i, m := range s {
		eqs[i] = TypeEquation{
			Left:  m.TypeVar,
			Right: m.Type,
		}
	}
	return eqs
}

func (s Substitution) ApplyToJudgement(j Judgement) Judgement {
	j.Env = s.ApplyToEnv(j.Env)
	j.Type = s.ApplyToType(j.Type)
	for i, p := range j.Premises {
		j.Premises[i] = s.ApplyToJudgement(p)
	}
	return j
}

func (s Substitution) ApplyToEnv(env Env) Env {
	for _, m := range s {
		env = m.ApplyToEnv(env)
	}
	return env
}

func (s Substitution) ApplyToType(t Type) Type {
	for _, m := range s {
		t = m.ApplyToType(t)
	}
	return t
}

type TypeMapping struct {
	TypeVar TypeVar
	Type    Type
}

func (m TypeMapping) String() string {
	return fmt.Sprintf("%s => %s", m.TypeVar, m.Type)
}

func (m TypeMapping) ApplyToEnv(env Env) Env {
	for i, bind := range env {
		env[i].TypeScheme = m.ApplyToTypeScheme(bind.TypeScheme)
	}
	return env
}

func (m TypeMapping) ApplyToTypeScheme(ts TypeScheme) TypeScheme {
	if slices.Contains(ts.FTV(), m.TypeVar) {
		ts.Type = m.ApplyToType(ts.Type)
	}
	// ts.Vars = slices.DeleteFunc(slices.Clone(ts.Vars), func(v TypeVar) bool {
	// 	return v.Equal(m.TypeVar)
	// })
	// ts.Type = m.ApplyToType(ts.Type)
	return ts
}

func (m TypeMapping) ApplyToEquations(eqs []TypeEquation) []TypeEquation {
	for i, eq := range eqs {
		eqs[i] = m.ApplyToEquation(eq)
	}
	return eqs
}

func (m TypeMapping) ApplyToEquation(eq TypeEquation) TypeEquation {
	return TypeEquation{
		Left:  m.ApplyToType(eq.Left),
		Right: m.ApplyToType(eq.Right),
	}
}

func (m TypeMapping) ApplyToType(t Type) Type {
	switch t := t.(type) {
	case TypeVar:
		if t.Equal(m.TypeVar) {
			return m.Type
		}
		return t
	case FunType:
		return FunType{
			Param:  m.ApplyToType(t.Param),
			Return: m.ApplyToType(t.Return),
		}
	case ListType:
		return ListType{
			Elem: m.ApplyToType(t.Elem),
		}
	default:
		return t
	}
}

func (m TypeMapping) ApplyToJudgement(j Judgement) Judgement {
	j.Env = m.ApplyToEnv(j.Env)
	j.Type = m.ApplyToType(j.Type)
	for i, p := range j.Premises {
		j.Premises[i] = m.ApplyToJudgement(p)
	}
	return j
}
