package fn

import "fmt"

func Ok[T any](v T) Result[T] {
	return Result[T]{value: v}
}

func Err[T any](e error) Result[T] {
	return Result[T]{err: e}
}

type Result[T any] struct {
	value T
	err   error
}

func (r Result[T]) IsOk() bool {
	return r.err == nil
}

func (r Result[T]) IsErr() bool {
	return r.err != nil
}

// Error returns the error of the result.
func (r Result[T]) Error() error {
	return r.err
}

// Unwrap returns the value of the result.
// If the result is an error, it returns the nil value for type T.
func (r Result[T]) Unwrap() T {
	return r.value
}

// UnwrapOrPanic returns the value of the result.
// If the result is an error, it panics.
func (r Result[T]) UnwrapOrPanic() T {
	if r.err != nil {
		panic(r.err)
	}
	return r.value
}

// UnwrapOr returns the value of the result.
// If the result is an error, it returns the provided value.
func (r Result[T]) UnwrapOr(v T) T {
	if r.err != nil {
		return v
	}
	return r.value
}

// UnwrapOrElse returns the value of the result.
// If the result is an error, it returns the result of the provided function.
func (r Result[T]) UnwrapOrElse(f func() T) T {
	if r.err != nil {
		return f()
	}
	return r.value
}

// Expect calls the provided function if the result is an error.
// Useful for logging or error handling.
// If the result is not an error, the function is not called and has no effect.
// Useful for side effects.
func (r Result[T]) Expect(f func(r Result[T])) {
	if r.err != nil {
		f(r)
	}
}

func (r Result[T]) String() string {
	if r.err != nil {
		return r.err.Error()
	}
	return fmt.Sprintf("%v", r.value)
}
