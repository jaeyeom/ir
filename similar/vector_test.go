package similar

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jaeyeom/math/vector"
)

func ExampleWordWeights_Get() {
	w := MapWordWeights{
		"a": 0.7,
		"b": 2.0,
	}
	fmt.Printf("%.2f\n", w.Get("a"))
	fmt.Printf("%.2f\n", w.Get("b"))
	fmt.Printf("%.2f\n", w.Get("c"))
	// Output:
	// 0.70
	// 2.00
	// 1.00
}

func TestWordVector_Strings_NilWeights(t *testing.T) {
	wv := newWordVector(nil)
	v := wv.strings([]string{"a", "brown", "brown", "fox"})
	if !reflect.DeepEqual(vector.Vector{1, 2, 1}, v) {
		t.Errorf("%v mismatches with [1 2 1]", v)
	}
	if !reflect.DeepEqual(wordVector{m: map[string]int{
		"a":     0,
		"brown": 1,
		"fox":   2,
	}, w: nil}, wv) {
		t.Errorf("%v mismatches with map[a:0 brown:1 fox:2]", wv)
	}
}

func TestWordVector_Strings_WithWeights(t *testing.T) {
	w := MapWordWeights{"a": 0.1}
	wv := newWordVector(w)
	v := wv.strings([]string{"a", "brown", "brown", "fox"})
	if !reflect.DeepEqual(vector.Vector{0.1, 2, 1}, v) {
		t.Errorf("%v mismatches with [0.1 2 1]", v)
	}
	if !reflect.DeepEqual(wordVector{m: map[string]int{
		"a":     0,
		"brown": 1,
		"fox":   2,
	}, w: w}, wv) {
		t.Errorf("%v mismatches with map[a:0 brown:1 fox:2]", wv)
	}
}
