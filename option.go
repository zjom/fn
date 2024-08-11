package fn

type Option[T any] struct {
	value        T
	instantiated bool
}

func Some[T any](v T) Option[T] {
	return Option[T]{value: v, instantiated: true}
}

func None[T any]() Option[T] {
	return Option[T]{}
}

func (o Option[T]) IsSome() bool {
	return o.instantiated
}

func (o Option[T]) IsNone() bool {
	return !o.instantiated
}

func (o Option[T]) Unwrap() T {
	return o.value
}

func (o Option[T]) UnwrapOr(v T) T {
	if o.instantiated {
		return o.value
	}

	return v
}

func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.instantiated {
		return o.value
	}

	return f()
}
