package similar

import "github.com/jaeyeom/math/vector"

// Returns a cosine similarity of a and b.
func Cosine(a, b []string) float64 {
	w := wordVector{}
	va, vb := w.strings(a), w.strings(b)
	dot := vector.Dot(va, vb)
	norm := va.Norm() * vb.Norm()
	if norm == 0.0 {
		return 0.0
	}
	return float64(dot) / norm
}
