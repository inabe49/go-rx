package rx

type Item[T any] interface {
	IsError() bool
	E() error
	V() T
}

type itemImpl[T any] struct {
	isError bool
	err     error
	value   T
}

func (i *itemImpl[T]) IsError() bool {
	return i.isError
}

func (i *itemImpl[T]) E() error {
	return i.err
}

func (i *itemImpl[T]) V() T {
	return i.value
}

func NewValueItem[T any](value T) Item[T] {
	return &itemImpl[T]{
		isError: false,
		value:   value,
	}
}

func NewErrorItem[T any](err error) Item[T] {
	return &itemImpl[T]{
		isError: true,
		err:     err,
	}
}
