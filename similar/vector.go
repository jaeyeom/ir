package similar

import "github.com/jaeyeom/math/vector"

// wordVector generates word frequency vectors.
type wordVector map[string]int

// strings returns a vector from words.
func (w wordVector) strings(words []string) vector.Vector {
	v := vector.New(len(w))
	for _, word := range words {
		i := len(v)
		if j, ok := w[word]; ok {
			i = j
		} else {
			w[word] = i
		}
		if i >= len(v) {
			v = append(v, make([]int, i+1-len(v))...)
		}
		v[i]++
	}
	return v
}
