package similar

import "fmt"

func ExampleCosine() {
	fmt.Printf("%.3f", Cosine(
		[]string{"a", "brown", "fox"},
		[]string{"the", "red", "fox"},
	))
	// Output: 0.333
}
