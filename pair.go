package fn

import "fmt"

type Pair[A, B any] struct {
	A A
	B B
}

func NewPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{A: a, B: b}
}

func (p Pair[A, _]) First() A {
	return p.A
}

func (p Pair[_, B]) Second() B {
	return p.B
}

func (p Pair[A, B]) String() string {
	return fmt.Sprintf("(%v %v)", p.A, p.B)
}
