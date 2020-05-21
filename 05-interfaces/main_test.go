package main

import (
	"strings"
	"testing"
)

func TestUpper(t *testing.T) {
	tests := map[string]struct {
		input      string
		want       string
		uppercaser uppercaser
	}{
		"basic":    {input: "hello", want: "HELLO", uppercaser: mockCaser{}},
		"accented": {input: "café", want: "CAFÉ", uppercaser: mockCaser{}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := upcase(tc.uppercaser, tc.input)
			if tc.want != got {
				t.Fatalf("wanted %s but got %s", tc.want, got)
			}
		})
	}
}

func TestLower(t *testing.T) {
	tests := map[string]struct {
		input      string
		want       string
		lowercaser lowercaser
	}{
		"basic":    {input: "HELLO", want: "hello", lowercaser: mockCaser{}},
		"accented": {input: "CAFÉ", want: "café", lowercaser: mockCaser{}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := lowcase(tc.lowercaser, tc.input)
			if tc.want != got {
				t.Fatalf("wanted %s but got %s", tc.want, got)
			}
		})
	}
}

type mockCaser struct{}

func (mockCaser) Upper(s string) string {
	return strings.ToUpper(s)
}

func (mockCaser) Lower(s string) string {
	return strings.ToLower(s)
}
