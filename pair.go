package fn

import "fmt"

type Pair[A, B any] struct {
	a A
	b B
}

func NewPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{a: a, b: b}
}

func (p Pair[A, _]) First() A {
	return p.a
}

func (p Pair[_, B]) Second() B {
	return p.b
}

func (p Pair[A, B]) String() string {
	return fmt.Sprintf("(%v %v)", p.a, p.b)
}
