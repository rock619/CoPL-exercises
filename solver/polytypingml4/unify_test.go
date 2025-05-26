package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func validSubstitution(t *testing.T, sub Substitution) {
	t.Helper()

	for _, s := range sub {
		if s.TypeVar == "" {
			t.Fatal("expected a non-empty TypeVar in substitution")
		}

		if typeContainsTypeVar(s.Type) {
			t.Fatalf("expected Type %s in substitution to not contain TypeVar %s", s.Type, s.TypeVar)
		}
	}
}

func typeContainsTypeVar(t Type) bool {
	switch t := t.(type) {
	case TypeVar:
		return true
	case FunType:
		return typeContainsTypeVar(t.Param) || typeContainsTypeVar(t.Return)
	case ListType:
		return typeContainsTypeVar(t.Elem)
	default:
		return false
	}
}

func TestUnify(t *testing.T) {
	type args struct {
		eqs []TypeEquation
	}
	tests := []struct {
		name    string
		args    args
		want    Substitution
		wantErr bool
	}{
		{
			name: "empty equations",
			args: args{
				eqs: nil,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "simple unification",
			args: args{
				eqs: []TypeEquation{
					{Left: TypeVar("'a"), Right: IntType{}},
				},
			},
			want: Substitution{
				{TypeVar: "'a", Type: IntType{}},
			},
			wantErr: false,
		},
		{
			name: "unification with a function type",
			args: args{
				eqs: []TypeEquation{
					{Left: TypeVar("'a"), Right: FunType{Param: IntType{}, Return: BoolType{}}},
				},
			},
			want: Substitution{
				{TypeVar: "'a", Type: FunType{Param: IntType{}, Return: BoolType{}}},
			},
			wantErr: false,
		},
		{
			name: "unification with a function type",
			args: args{
				eqs: []TypeEquation{
					{Left: TypeVar("1"), Right: IntType{}},
					{Left: TypeVar("0"), Right: FunType{Param: TypeVar("1"), Return: BoolType{}}},
				},
			},
			want: Substitution{
				{TypeVar: "1", Type: IntType{}},
				{TypeVar: "0", Type: FunType{Param: IntType{}, Return: BoolType{}}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unify(tt.args.eqs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Unify() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_unifyTypeVar(t *testing.T) {
	type args struct {
		eqs []TypeEquation
		v   TypeVar
		t   Type
	}
	tests := []struct {
		name    string
		args    args
		want    Substitution
		wantErr bool
	}{
		{
			name: "FTV of t contains v",
			args: args{
				eqs: nil,
				v:   TypeVar("'a"),
				t:   FunType{Param: TypeVar("'a"), Return: TypeVar("'b")},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "t is a TypeVar",
			args: args{
				eqs: nil,
				v:   TypeVar("'a"),
				t:   TypeVar("'b"),
			},
			want: Substitution{
				{
					TypeVar: "'a",
					Type:    TypeVar("'b"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unifyTypeVar(tt.args.eqs, tt.args.v, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("unifyTypeVar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("unifyTypeVar() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_unifyFunType(t *testing.T) {
	type args struct {
		eqs []TypeEquation
		f   FunType
		t   Type
	}
	tests := []struct {
		name    string
		args    args
		want    Substitution
		wantErr bool
	}{
		{
			name: "t is a TypeVar",
			args: args{
				eqs: nil,
				f:   FunType{Param: IntType{}, Return: BoolType{}},
				t:   TypeVar("'a"),
			},
			want: Substitution{
				{
					TypeVar: "'a",
					Type:    FunType{Param: IntType{}, Return: BoolType{}},
				},
			},
			wantErr: false,
		},
		{
			name: "t is a FunType",
			args: args{
				eqs: nil,
				f: FunType{
					Param:  TypeVar("'a"),
					Return: TypeVar("'b"),
				},
				t: FunType{
					Param:  TypeVar("'c"),
					Return: TypeVar("'d"),
				},
			},
			want: Substitution{
				{
					TypeVar: "'a",
					Type:    TypeVar("'c"),
				},
				{
					TypeVar: "'b",
					Type:    TypeVar("'d"),
				},
			},
			wantErr: false,
		},
		{
			name: "t is another type",
			args: args{
				eqs: nil,
				f: FunType{
					Param:  TypeVar("'a"),
					Return: BoolType{},
				},
				t: IntType{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unifyFunType(tt.args.eqs, tt.args.f, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("unifyFunType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("unifyFunType() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func Test_unifyListType(t *testing.T) {
	type args struct {
		eqs []TypeEquation
		l   ListType
		t   Type
	}
	tests := []struct {
		name    string
		args    args
		want    Substitution
		wantErr bool
	}{
		{
			name: "t is a TypeVar",
			args: args{
				eqs: nil,
				l: ListType{
					Elem: IntType{},
				},
				t: TypeVar("'a"),
			},
			want: Substitution{
				{
					TypeVar: "'a",
					Type:    ListType{Elem: IntType{}},
				},
			},
			wantErr: false,
		},
		{
			name: "t is a ListType",
			args: args{
				eqs: nil,
				l: ListType{
					Elem: TypeVar("'a"),
				},
				t: ListType{
					Elem: TypeVar("'b"),
				},
			},
			want: Substitution{
				{
					TypeVar: "'a",
					Type:    TypeVar("'b"),
				},
			},
			wantErr: false,
		},
		{
			name: "t is another type",
			args: args{
				eqs: nil,
				l: ListType{
					Elem: TypeVar("'a"),
				},
				t: IntType{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unifyListType(tt.args.eqs, tt.args.l, tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("unifyListType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("unifyListType() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSubstitution_ApplyEquation(t *testing.T) {
	type fields struct {
		TypeVar TypeVar
		Type    Type
	}
	type args struct {
		eq TypeEquation
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   TypeEquation
	}{
		{
			name: "the substitution is applied to the left side of the equation",
			fields: fields{
				TypeVar: TypeVar("'a"),
				Type:    IntType{},
			},
			args: args{
				eq: TypeEquation{
					Left:  TypeVar("'a"),
					Right: TypeVar("'b"),
				},
			},
			want: TypeEquation{
				Left:  IntType{},
				Right: TypeVar("'b"),
			},
		},
		{
			name: "the substitution is applied to the right side of the equation",
			fields: fields{
				TypeVar: TypeVar("'a"),
				Type:    IntType{},
			},
			args: args{
				eq: TypeEquation{
					Left:  TypeVar("'b"),
					Right: TypeVar("'a"),
				},
			},
			want: TypeEquation{
				Left:  TypeVar("'b"),
				Right: IntType{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TypeMapping{
				TypeVar: tt.fields.TypeVar,
				Type:    tt.fields.Type,
			}
			got := s.ApplyToEquation(tt.args.eq)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Substitution.ApplyEquation() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestSubstitution_ApplyType(t *testing.T) {
	type fields struct {
		TypeVar TypeVar
		Type    Type
	}
	type args struct {
		t Type
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Type
	}{
		{
			name: "t is a TypeVar; t is equal to s.TypeVar",
			fields: fields{
				TypeVar: TypeVar("'a"),
				Type:    TypeVar("'b"),
			},
			args: args{
				t: TypeVar("'a"),
			},
			want: TypeVar("'b"),
		},
		{
			name: "t is a TypeVar; t is not equal to s.TypeVar",
			fields: fields{
				TypeVar: TypeVar("'a"),
				Type:    TypeVar("'b"),
			},
			args: args{
				t: TypeVar("'c"),
			},
			want: TypeVar("'c"),
		},
		{
			name: "t is a FunType",
			fields: fields{
				TypeVar: TypeVar("'b"),
				Type:    IntType{},
			},
			args: args{
				t: FunType{Param: TypeVar("'a"), Return: TypeVar("'b")},
			},
			want: FunType{Param: TypeVar("'a"), Return: IntType{}},
		},
		{
			name: "t is a ListType",
			fields: fields{
				TypeVar: TypeVar("'a"),
				Type:    IntType{},
			},
			args: args{
				t: ListType{Elem: TypeVar("'a")},
			},
			want: ListType{Elem: IntType{}},
		},
		{
			name: "t is another type",
			fields: fields{
				TypeVar: TypeVar("'a"),
				Type:    IntType{},
			},
			args: args{
				t: BoolType{},
			},
			want: BoolType{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := TypeMapping{
				TypeVar: tt.fields.TypeVar,
				Type:    tt.fields.Type,
			}
			got := s.ApplyToType(tt.args.t)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Substitution.ApplyType() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
