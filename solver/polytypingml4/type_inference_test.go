package main

import "testing"

func TestTypeVarGenerator_NewTypeVar(t *testing.T) {
	type fields struct {
		nextID int
		used   map[TypeVar]struct{}
	}
	tests := []struct {
		name   string
		fields fields
		want   TypeVar
	}{
		{
			name: "initial case",
			fields: fields{
				nextID: 0,
				used:   make(map[TypeVar]struct{}),
			},
			want: TypeVar("a"),
		},
		{
			name: "26th case",
			fields: fields{
				nextID: 25,
				used:   make(map[TypeVar]struct{}),
			},
			want: TypeVar("z"),
		},
		{
			name: "27th case",
			fields: fields{
				nextID: 26,
				used:   make(map[TypeVar]struct{}),
			},
			want: TypeVar("a1"),
		},
		{
			name: "28th case",
			fields: fields{
				nextID: 27,
				used:   make(map[TypeVar]struct{}),
			},
			want: TypeVar("b1"),
		},
		{
			name: "if a is used, it should return b",
			fields: fields{
				nextID: 0,
				used: map[TypeVar]struct{}{
					TypeVar("a"): {},
				},
			},
			want: TypeVar("b"),
		},
		{
			name: "if nextID is 1 and b, c, and d is used, it should return e",
			fields: fields{
				nextID: 1,
				used: map[TypeVar]struct{}{
					TypeVar("b"): {},
					TypeVar("c"): {},
					TypeVar("d"): {},
				},
			},
			want: TypeVar("e"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &TypeVarGenerator{
				nextID: tt.fields.nextID,
				used:   tt.fields.used,
			}
			if got := g.NewTypeVar(); got != tt.want {
				t.Errorf("TypeVarGenerator.NewTypeVar() = %v, want %v", got, tt.want)
			}
		})
	}
}
