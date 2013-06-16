package diff

import (
	"fmt"
	"testing"
)

func TestDiff(t *testing.T) {
	l, r := []interface{}{"a", "b", "c"}, []interface{}{"a", "c", "d"}
	ds := []map[byte][]interface{}
	d := make(map[byte][]interface{})
	deltas := Diff(l, r)
	if deltas != d {
		t.Fail()
	}
}

func ExampleDiff() {
	l, r := []interface{}{"a", "b", "c"}, []interface{}{"a", "c", "d"}
	fmt.Println(Diff(l, r))
	// Output:
	// [map[1:[] 2:[]] map[0:[a]] map[1:[b] 2:[]] map[0:[c]] map[1:[] 2:[d]]]
}