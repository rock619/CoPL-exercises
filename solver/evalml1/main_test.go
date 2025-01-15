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
			name: "Q25",
			args: args{
				input: "3 + 5 evalto 8",
			},
			wantW: readTestdata(t, 25),
		},
		{
			name: "Q26",
			args: args{
				input: "8 - 2 - 3 evalto 3",
			},
			wantW: readTestdata(t, 26),
		},
		{
			name: "Q27",
			args: args{
				input: "(4 + 5) * (1 - 10) evalto -81",
			},
			wantW: readTestdata(t, 27),
		},
		{
			name: "Q28",
			args: args{
				input: "if 4 < 5 then 2 + 3 else 8 * 8 evalto 5",
			},
			wantW: readTestdata(t, 28),
		},
		{
			name: "Q29",
			args: args{
				input: "3 + if -23 < -2 * 8 then 8 else 2 + 4 evalto 11",
			},
			wantW: readTestdata(t, 29),
		},
		{
			name: "Q30",
			args: args{
				input: "3 + (if -23 < -2 * 8 then 8 else 2) + 4 evalto 15",
			},
			wantW: readTestdata(t, 30),
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
