package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func readSolution(t *testing.T, number int) string {
	t.Helper()
	b, err := os.ReadFile(fmt.Sprintf("../../solutions/q%03d.txt", number))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func readQuestion(t *testing.T, number int) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("testdata", fmt.Sprintf("q%03d.txt", number)))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func TestRun(t *testing.T) {
	for i := 124; i <= 129; i++ {
		t.Run(fmt.Sprintf("q%d", i), func(t *testing.T) {
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}
			r := strings.NewReader(readQuestion(t, i))
			want := readSolution(t, i)
			if err := Run(r, w, errW); err != nil {
				t.Errorf("Run() error = %v", err)
				return
			}
			got := w.String()
			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("Run() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
