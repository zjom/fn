// Package of convenience functions for working with *fn.List
package list

import (
	"github.com/zjom/fn"
)

// Map applies a function to each element of a list and returns a new list with the results.
func Map[T, R any](l *fn.List[T], f func(item T) R) *fn.List[R] {
	if l == nil {
		return nil
	}
	return &fn.List[R]{A: f(l.A), B: Map(l.B, f)}
}

// Reduce applies an accumulator function to each element of a list and returns the final accumulator value.
func Reduce[T any](l *fn.List[T], f func(acc T, item T) T) T {
	if l == nil {
		var v T
		return v
	}
	return f(Reduce(l.Rest(), f), l.Head())
}

// Fold applies an accumulator function to each element of a list and an accumulator and returns the final accumulator value.
func Fold[T, R any](l *fn.List[T], z R, f func(acc R, item T) R) R {
	if l == nil {
		return z
	}
	return Fold(l.Rest(), f(z, l.Head()), f)
}

// Filter returns a new list with only the elements that satisfy a predicate function.
func Filter[T any](l *fn.List[T], f func(item T) bool) *fn.List[T] {
	if l == nil {
		return l
	}

	if f(l.Head()) {
		return &fn.List[T]{A: l.Head(), B: Filter(l.Rest(), f)}
	}

	return Filter(l.Rest(), f)
}

// Zip returns a new list with the elements of two lists combined into pairs.
func Zip[A, B any](a *fn.List[A], b *fn.List[B]) *fn.List[fn.Pair[A, B]] {
	if a == nil || b == nil {
		return nil
	}
	return &fn.List[fn.Pair[A, B]]{
		A: fn.NewPair(a.Head(), b.Head()),
		B: Zip(a.Rest(), b.Rest()),
	}
}

// All returns Option[True] if all elements of a list satisfy a predicate function.
// Otherwise, it returns Option[False].
// Returns None if the list is empty.
func All[T any](l *fn.List[T], f func(item T) bool) fn.Option[bool] {
	if l == nil {
		return fn.None[bool]()
	}

	if !f(l.Head()) {
		return fn.Some(false)
	}

	return All(l.Rest(), f)
}

// Any returns Option[True] if any elements of a list satisfies the predicate
// Returns None if the list is empty
func Any[T any](l *fn.List[T], f func(item T) bool) fn.Option[bool] {
	if l == nil {
		return fn.None[bool]()
	}

	if f(l.Head()) {
		return fn.Some(true)
	}

	return Any(l.Rest(), f)
}
