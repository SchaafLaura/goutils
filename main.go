package utils

func Map[T any](arr []T, fn func(T) T) []T {

	var output []T = make([]T, len(arr))
	for i, val := range arr {
		output[i] = fn(val)
	}
	return output
}

func Mapper[T any](fn func(T) T) func(arr []T, fn func(T) T) []T {
	return func(arr []T, fn func(T) T) []T {
		var output []T = make([]T, len(arr))
		for i, val := range arr {
			output[i] = fn(val)
		}
		return output
	}
}