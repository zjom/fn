package fn

// Map applies a function to each element of a list and returns a new list with the results.
func Map[T, R any](l *List[T], f func(item T) R) *List[R] {
	if l == nil {
		return nil
	}

	return &List[R]{a: f(l.a), b: Map(l.b, f)}
}

// Reduce applies an accumulator function to each element of a list and returns the final accumulator value.
func Reduce[T any](l *List[T], f func(acc T, item T) T) T {
	if l == nil {
		var v T
		return v
	}
	return f(Reduce(l.Rest(), f), l.Head())
}

// Fold applies an accumulator function to each element of a list and an accumulator and returns the final accumulator value.
func Fold[T, R any](l *List[T], z R, f func(acc R, item T) R) R {
	if l == nil {
		return z
	}
	return Fold(l.Rest(), f(z, l.Head()), f)
}

// Filter returns a new list with only the elements that satisfy a predicate function.
func Filter[T any](l *List[T], f func(item T) bool) *List[T] {
	if l == nil {
		return l
	}

	if f(l.Head()) {
		return &List[T]{l.Head(), Filter(l.Rest(), f)}
	}

	return Filter(l.Rest(), f)
}

// Zip returns a new list with the elements of two lists combined into pairs.
func Zip[A, B any](a *List[A], b *List[B]) *List[Pair[A, B]] {
	if a == nil || b == nil {
		return nil
	}
	return &List[Pair[A, B]]{NewPair(a.Head(), b.Head()), Zip(a.Rest(), b.Rest())}
}

// All returns Option[True] if all elements of a list satisfy a predicate function.
// Otherwise, it returns Option[False].
// Returns None if the list is empty.
func All[T any](l *List[T], f func(item T) bool) Option[bool] {
	if l == nil {
		return None[bool]()
	}

	if !f(l.Head()) {
		return Some(false)
	}

	return All(l.Rest(), f)
}

func Any[T any](l *List[T], f func(item T) bool) Option[bool] {
	if l == nil {
		return None[bool]()
	}

	if f(l.Head()) {
		return Some(true)
	}

	return Any(l.Rest(), f)
}
