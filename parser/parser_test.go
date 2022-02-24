package parser

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		in   string
		out  string
	}{
		{
			name: "test abcd",
			in:   "abcd",
			out:  "abcd",
		},
		{
			name: "test case 2{a}",
			in:   "2{a}",
			out:  "aa",
		},
		{
			name: "test 2{a}2{b}",
			in:   "2{a}2{b}",
			out:  "aabb",
		},
		{
			name: "test 2{ab}",
			in:   "2{ab}",
			out:  "abab",
		},
		{
			name: "test 2{a2{b}}",
			in:   "2{a2{b}}",
			out:  "abbabb",
		},
		{
			name: "test 2{b}3{fg}",
			in:   "2{b}3{fg}",
			out:  "bbfgfgfg",
		},
		{
			name: "test 4{b3{a}}",
			in:   "4{b3{a}}",
			out:  "baaabaaabaaabaaa",
		},
		{
			name: "test 0{a}1{b}",
			in:   "0{a}1{b}",
			out:  "b",
		},
		{
			name: "test 2{2{a}}",
			in:   "2{2{a}}",
			out:  "aaaa",
		},
		{
			name: "test \"\"",
			in:   "",
			out:  "",
		},
		{
			name: "test 2{}2{a}",
			in:   "2{}2{a}",
			out:  "aa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.in); got != tt.out {
				t.Errorf("Parse() = %v, want %v", got, tt.out)
			}
		})
	}
}
