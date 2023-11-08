package pkg

type NotEmpty[T any] struct {
	First T
	Left  []T
}

func FromArrayToNotEmpty[T any](arr []T) (NotEmpty[T], bool) {
	if len(arr) == 0 {
		return NotEmpty[T]{}, false
	}
	return NotEmpty[T]{First: arr[0], Left: arr[1:]}, true
}
func ToArray[T any](ne NotEmpty[T]) []T {
	return append([]T{ne.First}, ne.Left...)
}
