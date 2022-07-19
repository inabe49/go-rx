package rx

type Subject[T any] interface {
	Observable[T]
	onCompleted()
	onError(err error)
	onNext(value T)
}
