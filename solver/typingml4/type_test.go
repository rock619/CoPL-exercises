package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConstraint_Substitute(t *testing.T) {
	type fields struct {
		Left  Type
		Right Type
	}
	type args struct {
		s Substitutions
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Constraint{
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			c.Substitute(tt.args.s)
		})
	}
}

func TestSubstitute(t *testing.T) {
	type args struct {
		t    Type
		subs Substitutions
	}
	tests := []struct {
		name string
		args args
		want Type
	}{
		{
			name: "if TypeVar#0 = int; then TypeVar#0 = int",
			args: args{
				t: &TypeVar{
					ID: 0,
				},
				subs: Substitutions{
					0: &IntType{},
				},
			},
			want: &IntType{},
		},
		{
			name: "if TypeVar#0 = TypeVar#1, TypeVar#1 = TypeVar#2, TypeVar#2 = bool; then TypeVar#0 = bool",
			args: args{
				t: &TypeVar{
					ID: 0,
				},
				subs: Substitutions{
					0: &TypeVar{ID: 1},
					1: &TypeVar{ID: 2},
					2: &BoolType{},
				},
			},
			want: &BoolType{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Substitute(tt.args.t, tt.args.subs)
			if got != tt.want {
				t.Errorf("Substitute() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnify(t *testing.T) {
	type args struct {
		constraints []Constraint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    Substitutions
	}{
		{
			name: "all typeVars can be unified",
			args: args{
				constraints: []Constraint{
					{Left: &TypeVar{ID: 0}, Right: &IntType{}},
					{Left: &TypeVar{ID: 1}, Right: &TypeVar{ID: 0}},
					{Left: &TypeVar{ID: 2}, Right: &FunType{ParamType: &TypeVar{ID: 1}, ReturnType: &BoolType{}}},
					{Left: &TypeVar{ID: 3}, Right: &ListType{ElemType: &TypeVar{ID: 2}}},
				},
			},
			wantErr: false,
			want: Substitutions{
				0: &IntType{},
				1: &IntType{},
				2: &FunType{ParamType: &IntType{}, ReturnType: &BoolType{}},
				3: &ListType{ElemType: &FunType{ParamType: &IntType{}, ReturnType: &BoolType{}}},
			},
		},
		{
			name: "all typeVars cannot be contrete types",
			args: args{
				constraints: []Constraint{
					{Left: &TypeVar{ID: 0}, Right: &TypeVar{ID: 2}},
					{Left: &TypeVar{ID: 1}, Right: &TypeVar{ID: 0}},
					{Left: &TypeVar{ID: 2}, Right: &FunType{ParamType: &TypeVar{ID: 3}, ReturnType: &BoolType{}}},
					{Left: &TypeVar{ID: 3}, Right: &ListType{ElemType: &TypeVar{ID: 4}}},
				},
			},
			wantErr: false,
			want: Substitutions{
				0: &FunType{ParamType: &ListType{ElemType: &TypeVar{ID: 4}}, ReturnType: &BoolType{}},
				1: &TypeVar{ID: 0},
				2: &FunType{ParamType: &ListType{ElemType: &TypeVar{ID: 4}}, ReturnType: &BoolType{}},
				3: &ListType{ElemType: &TypeVar{ID: 4}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			subs := make(Substitutions)
			if err := Unify(tt.args.constraints, subs); (err != nil) != tt.wantErr {
				t.Errorf("Unify() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(subs, tt.want) {
				t.Errorf("Unify() = %v, want %v", subs, tt.want)
			}
		})
	}
}
