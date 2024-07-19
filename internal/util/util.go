// Package util houses utility functions
package util

// Zero returns an empty T
func Zero[T any]() T {
	return *new(T)
}
