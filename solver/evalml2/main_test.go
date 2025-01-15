package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func readTestdata(t *testing.T, number int) string {
	t.Helper()
	b, err := os.ReadFile(fmt.Sprintf("../../solutions/q%03d.txt", number))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestRun(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		{
			name: "Q34",
			args: args{
				input: "x = 3, y = 2 |- x evalto 3",
			},
			wantW: readTestdata(t, 34),
		},
		{
			name: "Q35",
			args: args{
				input: "x = true, y = 4 |- if x then y + 1 else y - 1 evalto 5",
			},
			wantW: readTestdata(t, 35),
		},
		{
			name: "Q36",
			args: args{
				input: "|- let x = 1 + 2 in x * 4 evalto 12",
			},
			wantW: readTestdata(t, 36),
		},
		{
			name: "Q37",
			args: args{
				input: "|- let x = 3 * 3 in let y = 4 * x in x + y evalto 45",
			},
			wantW: readTestdata(t, 37),
		},
		{
			name: "Q38",
			args: args{
				input: "x = 3 |- let x = x * 2 in x + x evalto 12",
			},
			wantW: readTestdata(t, 38),
		},
		{
			name: "Q39",
			args: args{
				input: "|- let x = let y = 3 - 2 in y * y in let y = 4 in x + y evalto 5",
			},
			wantW: readTestdata(t, 39),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}
			if err := Run(tt.args.input, w, errW); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotW := w.String()
			if diff := cmp.Diff(tt.wantW, gotW); diff != "" {
				t.Errorf("Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
