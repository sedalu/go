package nqueens

import (
	"reflect"
	"testing"
)

var tsprint = []struct {
	s   []int
	str string
}{
	{nil, "++\n++\n"},
	{make([]int, 0), "++\n++\n"},
	{[]int{}, "++\n++\n"},
	{[]int{0}, "+---+\n| Q |\n+---+\n"},
	{[]int{0, 1}, "+---+---+\n| Q |   |\n+---+---+\n|   | Q |\n+---+---+\n"},
}

func TestSprint(t *testing.T) {
	for _, tt := range tsprint {
		// TODO: use 1.7 subtests
		str := Sprint(tt.s)

		if str != tt.str {
			t.Errorf("Solution(%v).String() => %q, want %q", []int(tt.s), str, tt.str)
		}
	}
}

var tSolve = []struct {
	n   uint
	len int
	cmp bool // true => compare solition set; false => compare len()
	s   [][]int
}{
	{0, 0, true, nil},
	{1, 1, true, [][]int{{0}}},
	{2, 0, true, nil},
	{3, 0, true, nil},
	{4, 2, true, [][]int{{1, 3, 0, 2}, {2, 0, 3, 1}}},
	{5, 10, false, nil},
	{6, 4, true, [][]int{{1, 3, 5, 0, 2, 4}, {2, 5, 1, 4, 0, 3}, {3, 0, 4, 1, 5, 2}, {4, 2, 0, 5, 3, 1}}},
	{7, 40, false, nil},
	{8, 92, false, nil},
}

func TestSolve(t *testing.T) {
	for _, tt := range tSolve {
		s := Solve(tt.n)

		if tt.cmp && !reflect.DeepEqual(s, tt.s) {
			t.Errorf("Solve(%v) => %v, want %v", tt.n, s, tt.s)
		} else if len(s) != tt.len {
			t.Errorf("len(Solve(%v)) => %v, want %v", tt.n, len(s), tt.len)
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solve(8)
	}
}

var tvalidate = []struct {
	x, y  int
	s     []int
	valid bool
}{
	{0, 0, nil, true},
	{0, 0, []int{}, true},
	{0, 0, []int{0}, false},
	{1, 0, []int{0}, false},
	{0, 1, []int{0}, false},
	{1, 1, []int{0}, false},
	{2, 1, []int{0}, true},
}

func Test_validate(t *testing.T) {
	for _, tt := range tvalidate {
		valid := validate(tt.x, tt.y, tt.s)

		if valid != tt.valid {
			t.Errorf("validate(%v, %v, %v) => %v, want %v", tt.x, tt.y, tt.s, valid, tt.valid)
		}
	}
}

var tdiagonal = []struct {
	x1, y1, x2, y2 int
	s              bool
}{
	{-1, -1, -1, -1, false},
	{-1, -1, 0, 0, true},
	{0, 0, -1, -1, true},
	{0, 0, 0, 0, false},
	{0, 0, 1, 1, true},
	{1, 1, 0, 0, true},
	{1, 1, 1, 1, false},
	{2, 2, 0, 0, true},
	{2, 0, 0, 0, false},
	{2, 0, 0, 1, false},
}

func Test_isDiagonal(t *testing.T) {
	for _, tt := range tdiagonal {
		s := isDiagonal(tt.x1, tt.y1, tt.x2, tt.y2)

		if s != tt.s {
			t.Errorf("slope(%v, %v, %v, %v) => %v, want %v", tt.x1, tt.y1, tt.x2, tt.y2, s, tt.s)
		}
	}
}
