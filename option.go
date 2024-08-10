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
