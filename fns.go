// Convenience functions for working with iterables in go
package fn

func Map[T, R any](xs []T, fn func(x T, i int) R) []R {
	var ys = make([]R, 0, len(xs))

	for i, x := range xs {
		ys = append(ys, fn(x, i))
	}

	return ys
}
