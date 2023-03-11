package rx

type Subject[T any] interface {
	Observable[T]
	Observer[T]
}
