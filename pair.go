package fn

type Pair[T any] struct {
	a T
	b T
}

func NewPair[T any](a, b T) Pair[T] {
	return Pair[T]{a: a, b: b}
}

func (p Pair[T]) First() T {
	return p.a
}

func (p Pair[T]) Second() T {
	return p.b
}
