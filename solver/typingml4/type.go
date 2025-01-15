package main

import (
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

type Environment struct {
	Scopes        []Scope
	constraints   *[]Constraint
	nextTypeVarID *uint
	l             *slog.Logger
}

func NewEnvironment(logger *slog.Logger) *Environment {
	env := &Environment{
		l:             logger,
		constraints:   &[]Constraint{},
		nextTypeVarID: new(uint),
	}
	env.AddScope()
	return env
}

func (e *Environment) Clone() *Environment {
	return &Environment{
		Scopes:        slices.Clone(e.Scopes),
		constraints:   e.constraints,
		nextTypeVarID: e.nextTypeVarID,
		l:             e.l,
	}
}

type Scope []Bind

type Bind struct {
	Name string
	Type Type
}

func (e *Environment) AddBind(name string, t Type) {
	e.Scopes[len(e.Scopes)-1] = append(e.Scopes[len(e.Scopes)-1], Bind{Name: name, Type: t})
}

func (e *Environment) GetType(name string) (Type, bool) {
	for _, s := range slices.Backward(e.Scopes) {
		for _, b := range slices.Backward(s) {
			if b.Name == name {
				return b.Type, true
			}
		}
	}
	return &NilType{}, false
}

func (e *Environment) AddConstraint(left, right Type) {
	if left.Nil() || right.Nil() {
		return
	}
	*e.constraints = append(*e.constraints, Constraint{Left: left, Right: right})
}

func (e *Environment) AddScope() {
	e.Scopes = append(e.Scopes, Scope{})
}

func (e *Environment) RemoveScope() {
	e.Scopes = e.Scopes[:len(e.Scopes)-1]
}

func (e *Environment) NewTypeVar() *TypeVar {
	v := &TypeVar{ID: *e.nextTypeVarID}
	*e.nextTypeVarID++
	return v
}

func (e *Environment) String() string {
	constraints := make([]string, len(*e.constraints))
	for i, c := range *e.constraints {
		constraints[i] = c.String()
	}
	return fmt.Sprintf(
		"Environment{Scopes: %v, constraints: (%v), nextTypeVarID: %d}",
		e.Scopes,
		strings.Join(constraints, "; "),
		*e.nextTypeVarID,
	)
}

func (e *Environment) Print(subs Substitutions) string {
	binds := make([]string, 0, *e.nextTypeVarID)
	for _, scope := range e.Scopes {
		for _, bind := range scope {
			binds = append(binds, fmt.Sprintf("%s : %s", bind.Name, Substitute(bind.Type, subs).Print()))
		}
	}
	return strings.Join(binds, ", ")
}

func (e *Environment) Unify() (Substitutions, error) {
	subs := make(Substitutions)
	if err := Unify(*e.constraints, subs); err != nil {
		return nil, err
	}
	return subs, nil
}

func Unify(constraints []Constraint, subs Substitutions) error {
	if subs == nil {
		subs = make(Substitutions)
	}
	// constraintsを消費してしまうのでcopyしておく
	cs := slices.Clone(constraints)
	// i := 0
	for len(cs) > 0 {
		// if i > 1000 {
		// 	return fmt.Errorf("Unify: too many iterations: %v", cs)
		// }
		// i++
		c := cs[0]
		cs = cs[1:]

		left := Substitute(c.Left, subs)
		right := Substitute(c.Right, subs)

		switch l := left.(type) {
		case *TypeVar:
			switch r, ok := right.(*TypeVar); {
			case ok && l.ID == r.ID:
			case ok && l.ID < r.ID:
				cs = append(cs, Constraint{r, l})
			default:
				subs[l.ID] = right
			}
		case *IntType:
			switch r := right.(type) {
			case *TypeVar:
				cs = append(cs, Constraint{r, l})
			case *IntType:
			default:
				return fmt.Errorf("Unify: type mismatch: %v (%T) != %v (%T)", l, l, right, right)
			}
		case *BoolType:
			switch r := right.(type) {
			case *TypeVar:
				cs = append(cs, Constraint{r, l})
			case *BoolType:
			default:
				return fmt.Errorf("Unify: type mismatch: %v (%T) != %v (%T)", l, l, right, right)
			}
		case *FunType:
			switch r := right.(type) {
			case *TypeVar:
				cs = append(cs, Constraint{r, l})
			case *FunType:
				cs = append(cs, Constraint{l.ParamType, r.ParamType})
				cs = append(cs, Constraint{l.ReturnType, r.ReturnType})
			default:
				return fmt.Errorf("Unify: type mismatch: %v (%T) != %v (%T)", l, l, right, right)
			}
		case *ListType:
			switch r := right.(type) {
			case *TypeVar:
				cs = append(cs, Constraint{r, l})
			case *ListType:
				cs = append(cs, Constraint{l.ElemType, r.ElemType})
			default:
				return fmt.Errorf("Unify: type mismatch: %v (%T) != %v (%T)", l, l, right, right)
			}
		default:
			return fmt.Errorf("Unify: Unknown type: %v", l)
		}
	}

	return nil
}

type Constraint struct {
	Left  Type
	Right Type
}

func (c *Constraint) Substitute(s Substitutions) {
	Substitute(c.Left, s)
	Substitute(c.Right, s)
}

func (c Constraint) String() string {
	return fmt.Sprintf("%s = %s", c.Left, c.Right)
}

type Substitutions map[uint]Type

func (s Substitutions) String() string {
	ss := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		t, ok := s[uint(i)]
		if !ok {
			t = &NilType{}
		}
		ss[i] = fmt.Sprintf("#%d = %s", i, t)
	}
	return fmt.Sprintf("Substitutions{%s}", strings.Join(ss, "; "))
}

type Type interface {
	typ()
	Print() string
	// -
	Nil() bool
	TypeVar() bool
}

func Substitute(t Type, subs Substitutions) Type {
	switch t := t.(type) {
	case *TypeVar:
		if sub, ok := subs[t.ID]; ok {
			return Substitute(sub, subs)
		}
		return t
	case *FunType:
		t.ParamType = Substitute(t.ParamType, subs)
		t.ReturnType = Substitute(t.ReturnType, subs)
		return t
	case *ListType:
		t.ElemType = Substitute(t.ElemType, subs)
		return t
	default:
		return t
	}
}

// anyType 型を決定できない場合に使う適当な具体型
var anyType = &IntType{}

type TypeVar struct {
	ID uint
}

func (*TypeVar) typ() {}

func (v *TypeVar) Print() string {
	return anyType.Print()
}

func (v *TypeVar) Nil() bool {
	return false
}

func (v *TypeVar) String() string {
	return fmt.Sprintf("Type#%d", v.ID)
}

func (*TypeVar) TypeVar() bool {
	return true
}

type NilType struct{}

func (*NilType) typ() {}

func (*NilType) Print() string {
	return "?"
}

func (t *NilType) String() string {
	return t.Print()
}

func (*NilType) Nil() bool {
	return true
}

func (*NilType) TypeVar() bool {
	return false
}

type BoolType struct{}

func (*BoolType) typ() {}

func (*BoolType) Print() string {
	return "bool"
}

func (t *BoolType) String() string {
	return t.Print()
}

func (*BoolType) Nil() bool {
	return false
}

func (*BoolType) TypeVar() bool {
	return false
}

type IntType struct{}

func (*IntType) typ() {}

func (*IntType) Print() string {
	return "int"
}

func (t *IntType) String() string {
	return t.Print()
}

func (*IntType) Nil() bool {
	return false
}

func (*IntType) TypeVar() bool {
	return false
}

type FunType struct {
	ParamType  Type
	ReturnType Type
}

func (*FunType) typ() {}

func (t *FunType) Print() string {
	if _, ok := t.ParamType.(*FunType); ok {
		return fmt.Sprintf("(%s) -> %s", t.ParamType.Print(), t.ReturnType.Print())
	}
	return fmt.Sprintf("%s -> %s", t.ParamType.Print(), t.ReturnType.Print())
}

func (t *FunType) String() string {
	if _, ok := t.ParamType.(*FunType); ok {
		return fmt.Sprintf("(%s) -> %s", t.ParamType, t.ReturnType)
	}
	return fmt.Sprintf("%s -> %s", t.ParamType, t.ReturnType)
}

func (t *FunType) Nil() bool {
	return t.ParamType.Nil() || t.ReturnType.Nil()
}

func (t *FunType) TypeVar() bool {
	return false
}

type ListType struct {
	ElemType Type
}

func (*ListType) typ() {}

func (t *ListType) String() string {
	if _, ok := t.ElemType.(*FunType); ok {
		return fmt.Sprintf("(%s) list", t.ElemType)
	}
	return fmt.Sprintf("%s list", t.ElemType)
}

func (t *ListType) Print() string {
	if _, ok := t.ElemType.(*FunType); ok {
		return fmt.Sprintf("(%s) list", t.ElemType.Print())
	}
	return fmt.Sprintf("%s list", t.ElemType.Print())
}

func (t *ListType) Nil() bool {
	return t.ElemType.Nil()
}

func (t *ListType) TypeVar() bool {
	return t.ElemType.TypeVar()
}
