package similar

import (
	"reflect"
	"testing"

	"github.com/jaeyeom/math/vector"
)

func TestWordVector_Strings(t *testing.T) {
	w := wordVector{}
	v := w.strings([]string{"a", "brown", "brown", "fox"})
	if !reflect.DeepEqual(vector.Vector{1, 2, 1}, v) {
		t.Errorf("%v mismatches with [1 2 1]", v)
	}
	if !reflect.DeepEqual(wordVector{
		"a":     0,
		"brown": 1,
		"fox":   2,
	}, w) {
		t.Errorf("%v mismatches with map[a:0 brown:1 fox:2]", w)
	}
}
