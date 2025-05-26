package main

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/rock619/CoPL-exercises/solver/polytypingml4/parser"
)

// Type represents a type in the language
type Type interface {
	// isType is a marker method to distinguish between types and other interfaces
	isType()
	// Equal return true if the type is equal to another type
	Equal(Type) bool
	// String returns the string representation of the type
	String() string
	// FTV returns the free type variables of the type
	FTV() []TypeVar
}

type TypeVar string

func (TypeVar) isType() {}

func (v TypeVar) Equal(t Type) bool {
	if t2, ok := t.(TypeVar); ok {
		return string(v) == string(t2)
	}
	return false
}

func (v TypeVar) String() string {
	return fmt.Sprintf("'%s", string(v))
}

func (v TypeVar) FTV() []TypeVar {
	return []TypeVar{v}
}

type BoolType struct{}

func (BoolType) isType() {}

func (BoolType) Equal(t Type) bool {
	if _, ok := t.(BoolType); ok {
		return true
	}
	return false
}

func (BoolType) String() string {
	return "bool"
}

func (BoolType) FTV() []TypeVar {
	return nil
}

type IntType struct{}

func (IntType) isType() {}

func (IntType) Equal(t Type) bool {
	if _, ok := t.(IntType); ok {
		return true
	}
	return false
}

func (IntType) String() string {
	return "int"
}

func (IntType) FTV() []TypeVar {
	return nil
}

type FunType struct {
	Param  Type
	Return Type
}

func (FunType) isType() {}

func (t FunType) Equal(t2 Type) bool {
	if t2, ok := t2.(FunType); ok {
		return t.Param.Equal(t2.Param) && t.Return.Equal(t2.Return)
	}
	return false
}

func (t FunType) String() string {
	if _, ok := t.Param.(FunType); ok {
		return fmt.Sprintf("(%s) -> %s", t.Param, t.Return)
	}
	return fmt.Sprintf("%s -> %s", t.Param, t.Return)
}

func (t FunType) FTV() []TypeVar {
	return slices.Concat(t.Param.FTV(), t.Return.FTV())
}

type ListType struct {
	Elem Type
}

func (ListType) isType() {}

func (t ListType) Equal(t2 Type) bool {
	if t2, ok := t2.(ListType); ok {
		return t.Elem.Equal(t2.Elem)
	}
	return false
}

func (t ListType) String() string {
	if _, ok := t.Elem.(FunType); ok {
		return fmt.Sprintf("(%s) list", t.Elem)
	}
	return fmt.Sprintf("%s list", t.Elem)
}

func (t ListType) FTV() []TypeVar {
	return t.Elem.FTV()
}

type TypeScheme struct {
	vars map[TypeVar]struct{}
	Type Type
}

func NewTypeScheme(typeVars []TypeVar, t Type) TypeScheme {
	s := TypeScheme{
		vars: make(map[TypeVar]struct{}),
		Type: t,
	}
	for _, v := range typeVars {
		s.vars[v] = struct{}{}
	}
	return s
}

func (s TypeScheme) TypeVars() []TypeVar {
	vars := make([]TypeVar, 0, len(s.vars))
	for k := range s.vars {
		vars = append(vars, k)
	}
	return vars
}

func (s TypeScheme) String() string {
	if len(s.vars) == 0 {
		return s.Type.String()
	}
	vars := make([]string, len(s.vars))
	for i, v := range slices.Sorted(maps.Keys(s.vars)) {
		vars[i] = v.String()
	}
	return fmt.Sprintf("%s . %s", strings.Join(vars, " "), s.Type.String())
}

func (s *TypeScheme) AddTypeVar(v TypeVar) TypeScheme {
	if s.vars == nil {
		s.vars = make(map[TypeVar]struct{})
	}
	s.vars[v] = struct{}{}
	return *s
}

func (s TypeScheme) FTV() []TypeVar {
	ftv := make([]TypeVar, 0)
	for _, v := range s.Type.FTV() {
		if _, ok := s.vars[v]; !ok {
			ftv = append(ftv, v)
		}
	}
	return ftv
}

type Env []Bind

func (e Env) Get(v Var) (TypeScheme, bool) {
	for _, b := range e {
		if b.Var == v {
			return b.TypeScheme, true
		}
	}
	return TypeScheme{}, false
}

func (e Env) String() string {
	var binds []string
	for _, b := range e {
		binds = append(binds, b.String())
	}
	return strings.Join(binds, ", ")
}

func (e Env) FTV() []TypeVar {
	var ftv []TypeVar
	for _, b := range e {
		ftv = append(ftv, b.TypeScheme.FTV()...)
	}
	return ftv
}

type Bind struct {
	Var        Var
	TypeScheme TypeScheme
}

func (b Bind) String() string {
	return fmt.Sprintf("%s : %s", b.Var, b.TypeScheme)
}

type Var string

type Exp interface {
	isExp()
	String() string
}

type IntExp int

func (IntExp) isExp() {}

func (e IntExp) String() string {
	return strconv.Itoa(int(e))
}

type BoolExp bool

func (BoolExp) isExp() {}

func (e BoolExp) String() string {
	return strconv.FormatBool(bool(e))
}

type VarExp Var

func (VarExp) isExp() {}

func (e VarExp) String() string {
	return string(e)
}

type BinOpExp struct {
	Left  Exp
	Op    Op
	Right Exp
}

func (BinOpExp) isExp() {}

func (e BinOpExp) String() string {
	return fmt.Sprintf("%s %s %s", e.Left, e.Op, e.Right)
}

type IfExp struct {
	Cond Exp
	Then Exp
	Else Exp
}

func (IfExp) isExp() {}

func (e IfExp) String() string {
	return fmt.Sprintf("if %s then %s else %s", e.Cond, e.Then, e.Else)
}

type LetExp struct {
	Var  Var
	Bind Exp
	Body Exp
}

func (LetExp) isExp() {}

func (e LetExp) String() string {
	return fmt.Sprintf("let %s = %s in %s", e.Var, e.Bind, e.Body)
}

type FunExp Fun

func (FunExp) isExp() {}

func (e FunExp) String() string {
	return Fun(e).String()
}

type Fun struct {
	Param Var
	Body  Exp
}

func (f Fun) String() string {
	return fmt.Sprintf("fun %s -> %s", f.Param, f.Body)
}

type AppExp struct {
	Fun Exp
	Arg Exp
}

func (AppExp) isExp() {}

func (e AppExp) String() string {
	switch e.Arg.(type) {
	case VarExp, IntExp, BoolExp:
		return fmt.Sprintf("%s %s", e.Fun, e.Arg)
	default:
		return fmt.Sprintf("%s (%s)", e.Fun, e.Arg)
	}
}

type LetRecExp struct {
	Var  Var
	Bind Fun
	Body Exp
}

func (LetRecExp) isExp() {}

func (e LetRecExp) String() string {
	return fmt.Sprintf("let rec %s = %s in %s", e.Var, e.Bind, e.Body)
}

type NilExp struct{}

func (NilExp) isExp() {}

func (NilExp) String() string {
	return "[]"
}

type ConsExp struct {
	Head Exp
	Tail Exp
}

func (ConsExp) isExp() {}

func (e ConsExp) String() string {
	switch h := e.Head.(type) {
	case AppExp, VarExp, IntExp, BoolExp, NilExp:
		return fmt.Sprintf("%s :: %s", e.Head, e.Tail)
	case BinOpExp:
		if h.Op != OpLT {
			return fmt.Sprintf("%s :: %s", e.Head, e.Tail)
		}
		return fmt.Sprintf("(%s) :: %s", e.Head, e.Tail)
	default:
		return fmt.Sprintf("(%s) :: %s", e.Head, e.Tail)
	}
}

type MatchExp struct {
	Arg      Exp
	NilCase  Exp
	HeadVar  Var
	TailVar  Var
	ConsCase Exp
}

func (MatchExp) isExp() {}

func (e MatchExp) String() string {
	return fmt.Sprintf(
		"match %s with [] -> %s | %s :: %s -> %s",
		e.Arg, e.NilCase, e.HeadVar, e.TailVar, e.ConsCase,
	)
}

type ParenExp struct {
	Inner Exp
}

func (ParenExp) isExp() {}

func (e ParenExp) String() string {
	return fmt.Sprintf("(%s)", e.Inner.String())
}

type Op string

const (
	OpPlus  Op = "+"
	OpMinus Op = "-"
	OpTimes Op = "*"
	OpLT    Op = "<"
)

func NewOp(token antlr.Token) (Op, error) {
	switch token.GetTokenType() {
	case parser.PolyTypingML4ParserPLUS:
		return OpPlus, nil
	case parser.PolyTypingML4LexerMINUS:
		return OpMinus, nil
	case parser.PolyTypingML4LexerTIMES:
		return OpTimes, nil
	case parser.PolyTypingML4LexerLT:
		return OpLT, nil
	default:
		return "", fmt.Errorf("NewOp: unknown operator: %s", token.GetText())
	}
}

func (o Op) Rule() Rule {
	return map[Op]Rule{
		OpPlus:  TPlus,
		OpMinus: TMinus,
		OpTimes: TMult,
		OpLT:    TLT,
	}[o]
}
