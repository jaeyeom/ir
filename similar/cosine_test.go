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
		[]string{"the", "brown", "fox"},
	}, {
		[]string{"a", "brown", "fox"},
		[]string{"a", "red", "fox"},
	}, {
		[]string{"a", "brown", "fox"},
		[]string{"a", "a", "brown", "fox"},
	}, {
		[]string{"a", "brown", "fox"},
		[]string{"a", "b", "brown", "fox"},
	}}
	fmt.Println("Unweighted:")
	for _, testCase := range testCases {
		a, b := testCase.a, testCase.b
		fmt.Printf("%v, %v: %.3f\n", a, b, Cosine(a, b))
	}
	weighted := NewWeightedCosine(MapWordWeights{
		"a":   0.2,
		"b":   0.2,
		"the": 0.3,
	})
	fmt.Println("Weighted:")
	for _, testCase := range testCases {
		a, b := testCase.a, testCase.b
		fmt.Printf("%v, %v: %.3f\n", a, b, weighted.Cosine(a, b))
	}
	// Output:
	// Unweighted:
	// [a brown fox], [a brown fox]: 1.000
	// [a brown fox], [the brown fox]: 0.667
	// [a brown fox], [a red fox]: 0.667
	// [a brown fox], [a a brown fox]: 0.943
	// [a brown fox], [a b brown fox]: 0.866
	// Weighted:
	// [a brown fox], [a brown fox]: 1.000
	// [a brown fox], [the brown fox]: 0.969
	// [a brown fox], [a red fox]: 0.510
	// [a brown fox], [a a brown fox]: 0.991
	// [a brown fox], [a b brown fox]: 0.990
}
