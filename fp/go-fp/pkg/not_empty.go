package pkg

type NotEmpty[T any] struct {
	First T
	Left  []T
}

type ErrEmptyArray struct{}

func (e ErrEmptyArray) Error() string {
	return "array is empty"
}
func NewErrEmptyArray() ErrEmptyArray {
	return ErrEmptyArray{}
}

func FromArrayToNotEmpty[T any](arr []T) (NotEmpty[T], error) {
	if len(arr) == 0 {
		return NotEmpty[T]{}, NewErrEmptyArray()
	}
	return NotEmpty[T]{First: arr[0], Left: arr[1:]}, nil
}
