package similar

import "github.com/jaeyeom/math/vector"

type WordWeights interface {
	Get(word string) float64
}

// MapWordWeights implements a map based WordWeights with the default
// value 1.0.
type MapWordWeights map[string]float64

// Get returns the weight of the given word.
func (w MapWordWeights) Get(word string) float64 {
	v, ok := w[word]
	if !ok {
		return 1.0
	}
	return v
}

// wordVector generates word frequency vectors.
type wordVector struct {
	m map[string]int
	w WordWeights
}

func newWordVector(w WordWeights) wordVector {
	return wordVector{m: map[string]int{}, w: w}
}

// strings returns a vector from words.
func (w wordVector) strings(words []string) vector.Vector {
	v := vector.New(len(w.m))
	for _, word := range words {
		i := len(v)
		if j, ok := w.m[word]; ok {
			i = j
		} else {
			w.m[word] = i
		}
		if i >= len(v) {
			v = append(v, vector.New(i+1-len(v))...)
		}
		if w.w == nil {
			v[i]++
		} else {
			v[i] += w.w.Get(word)
		}
	}
	return v
}
