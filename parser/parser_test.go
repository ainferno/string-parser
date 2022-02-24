package parser

import "testing"

// func TestStrings1(t *testing.T) {

// 	got := Parse("abcd")
// 	want := "abcd"

// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}
// }
// func TestStrings2(t *testing.T) {

// 	got := Parse("2{a}")
// 	want := "aa"

// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}
// }
// func TestStrings3(t *testing.T) {

// 	got := Parse("2{a}2{b}")
// 	want := "aabb"

// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}
// }
// func TestStrings4(t *testing.T) {

// 	got := Parse("2{ab}")
// 	want := "abab"

// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}
// }
// func TestStrings5(t *testing.T) {

// 	got := Parse("2{a2{b}}")
// 	want := "abbabb"

// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}
// }
// func TestStrings6(t *testing.T) {

// 	got := Parse("2{b}3{fg}")
// 	want := "bbfgfgfg"

// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}
// }
// func TestStrings7(t *testing.T) {

// 	got := Parse("4{b3{a}}")
// 	want := "baaabaaabaaabaaa"

// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}
// }
// func TestStrings8(t *testing.T) {

// 	got := Parse("0{a}1{b}")
// 	want := "b"

// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}
// }
// func TestStrings9(t *testing.T) {

// 	got := Parse("2{2{a}}")
// 	want := "aaaa"

// 	if got != want {
// 		t.Errorf("got %s, wanted %s", got, want)
// 	}
// }

func TestParse(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				str: "abcd",
			},
			want: "abcd",
		},
		{
			name: "test case 2",
			args: args{
				str: "2{a}",
			},
			want: "aa",
		},
		{
			name: "test case 3",
			args: args{
				str: "2{a}2{b}",
			},
			want: "aabb",
		},
		{
			name: "test case 4",
			args: args{
				str: "2{ab}",
			},
			want: "abab",
		},
		{
			name: "test case 5",
			args: args{
				str: "2{a2{b}}",
			},
			want: "abbabb",
		},
		{
			name: "test case 6",
			args: args{
				str: "2{b}3{fg}",
			},
			want: "bbfgfgfg",
		},
		{
			name: "test case 7",
			args: args{
				str: "4{b3{a}}",
			},
			want: "baaabaaabaaabaaa",
		},
		{
			name: "test case 8",
			args: args{
				str: "0{a}1{b}",
			},
			want: "b",
		},
		{
			name: "test case 9",
			args: args{
				str: "2{2{a}}",
			},
			want: "aaaa",
		},
		{
			name: "test case 10",
			args: args{
				str: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := Parse(tt.args.str); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
