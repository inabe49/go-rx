package rx

type Observer[T any] interface {
	OnCompleted()
	OnError(err error)
	OnNext(value T)
}
