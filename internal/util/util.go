package util

func Zero[T interface{}]() T {
	return *new(T)
}
