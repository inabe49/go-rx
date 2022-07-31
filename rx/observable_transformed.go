package rx

import "time"

type transformedOperator[P any, T any] func(value P, onNext func(T), onError func(err error), onCompleted func())

type transformedObservable[P any, T any] struct {
	parent   Observable[P]
	operator transformedOperator[P, T]
}

func (o *transformedObservable[P, T]) Delay(delay time.Duration) Observable[T] {
	panic("???")
}

func (o *transformedObservable[P, T]) Filter(predicate func(T) bool) Observable[T] {
	panic("???")
}

func (o *transformedObservable[P, T]) Subscribe(onNext func(T), onError func(err error), onCompleted func()) Subscription {
	subscription := o.parent.Subscribe(func(value P) {
		o.operator(value, onNext, onError, onCompleted)
	}, func(err error) {
		onError(err)
	}, func() {
		onCompleted()
	})

	return subscription
}
