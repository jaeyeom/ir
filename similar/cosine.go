package similar

import "github.com/jaeyeom/math/vector"

type WeightedCosine struct {
	weights WordWeights
}

func NewWeightedCosine(weights WordWeights) WeightedCosine {
	return WeightedCosine{weights}
}

func (w WeightedCosine) Cosine(a, b []string) float64 {
	wv := newWordVector(w.weights)
	va, vb := wv.strings(a), wv.strings(b)
	dot := vector.Dot(va, vb)
	norm := va.Norm() * vb.Norm()
	if norm == 0.0 {
		return 0.0
	}
	return float64(dot) / norm
}

// Returns a cosine similarity of a and b.
func Cosine(a, b []string) float64 {
	return WeightedCosine{}.Cosine(a, b)
}
