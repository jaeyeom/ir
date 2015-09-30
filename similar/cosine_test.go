package similar

import "fmt"

func ExampleCosine() {
	testCases := []struct {
		a, b []string
	}{{
		[]string{"a", "brown", "fox"},
		[]string{"a", "brown", "fox"},
	}, {
		[]string{"a", "brown", "fox"},
		[]string{"the", "red", "fox"},
	}, {
		[]string{"a", "brown", "fox"},
		[]string{"a", "a", "brown", "fox"},
	}, {
		[]string{"a", "brown", "fox"},
		[]string{"a", "b", "brown", "fox"},
	}}
	for _, testCase := range testCases {
		a, b := testCase.a, testCase.b
		fmt.Printf("%v, %v: %.3f\n", a, b, Cosine(a, b))
	}
	// Output:
	// [a brown fox], [a brown fox]: 1.000
	// [a brown fox], [the red fox]: 0.333
	// [a brown fox], [a a brown fox]: 0.943
	// [a brown fox], [a b brown fox]: 0.866
}
