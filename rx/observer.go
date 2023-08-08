package rx

type Observer[T any] interface {
	OnNext(value T)
	OnError(err error)
	OnCompleted()
}

type observerImpl[T any] struct {
	onNext      func(value T)
	onError     func(err error)
	onCompleted func()
}

func (o *observerImpl[T]) OnCompleted() {
	o.onCompleted()
}

func (o *observerImpl[T]) OnError(err error) {
	o.onError(err)
}

func (o *observerImpl[T]) OnNext(value T) {
	o.onNext(value)
}

func newObserver[T any](onNext func(T), onError func(error), onCompleted func()) Observer[T] {
	return &observerImpl[T]{
		onNext:      onNext,
		onError:     onError,
		onCompleted: onCompleted,
	}
}
