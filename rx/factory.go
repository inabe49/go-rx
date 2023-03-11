package rx

import "time"

func Amb[T any](sources ...Observable[T]) Observable[T] {
	panic("???")
}

func Empty[T any]() Observable[T] {
	panic("???")
}

func Error[T any](err error) Observable[T] {
	panic("???")
}

func Interval(initialDelay time.Duration, period time.Duration) Observable[int] {
	panic("???")
}

func Just[T any](items ...T) Observable[T] {
	return newJustObservable[T](items)
}
