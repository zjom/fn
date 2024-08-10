package fn

// Map applies a function to each element of a list and returns a new list with the results.
func Map[T, R any](l *List[T], f func(item T) R) *List[R] {
	if l == nil {
		return nil
	}
	return &List[R]{f(l.Head()), Map(l.Rest(), f)}
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
