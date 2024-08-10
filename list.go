package fn

import "fmt"

type List[T any] Pair[T, *List[T]]

func NewList[T any](s ...T) *List[T] {
	if len(s) == 0 {
		return nil
	}
	return &List[T]{s[0], NewList(s[1:]...)}
}

func NewListFromSlice[T any](s []T) *List[T] {
	if len(s) == 0 {
		return nil
	}
	return &List[T]{s[0], NewListFromSlice(s[1:])}
}

func (l *List[T]) Head() T {
	return l.a
}

func (l *List[T]) Rest() *List[T] {
	return l.b
}

func (l *List[T]) Cons(a T) *List[T] {
	return &List[T]{a, l}
}

func (l *List[T]) Destructure() (T, *List[T]) {
	return l.Head(), l.Rest()
}

func (l *List[T]) ToSlice() []T {
	a, b := l.Destructure()
	if b == nil {
		return []T{a}
	}

	return append([]T{a}, b.ToSlice()...)
}

func (l *List[T]) String() string {
	if l == nil {
		return "()"
	}

	return Fold(l, "(", func(acc string, item T) string {
		if acc == "(" {
			return fmt.Sprintf("%s%v", acc, item)
		}
		return fmt.Sprintf("%s %v", acc, item)
	}) + ")"
}

// Iter calls a function for each element of a list.
// Useful for side effects.
func (l *List[T]) Iter(f func(T)) {
	if l == nil {
		return
	}
	f(l.Head())
	l.Rest().Iter(f)
}
