package go_rx

import "time"

type transformedOperator[P any, T any] func(value P, observer Observer[T])

// transformedObservable : 要素の型が変換される演算を含めた Observable
type transformedObservable[P any, T any] struct {
	// parent : 親の Observable
	parent Observable[P]

	// operator : 要素の型が変換される演算
	operator transformedOperator[P, T]
}

func (o *transformedObservable[P, T]) Delay(delay time.Duration) Observable[T] {
	return &transformedObservable[T, T]{
		parent: o,
		operator: func(value T, observer Observer[T]) {
			<-time.After(delay)
			observer.OnNext(value)
		},
	}
}

func (o *transformedObservable[P, T]) Filter(predicate func(T) bool) Observable[T] {
	return &transformedObservable[T, T]{
		parent: o,
		operator: func(value T, observer Observer[T]) {
			if predicate(value) {
				observer.OnNext(value)
			}
		},
	}
}

func (o *transformedObservable[P, T]) Subscribe(onNext func(T), onError func(err error), onCompleted func()) Subscription {
	observer := newObserver[T](onNext, onError, onCompleted)

	var subscription Subscription
	subscription = o.parent.Subscribe(func(value P) {
		o.operator(value, observer)
	}, func(err error) {
		onError(err)
	}, func() {
		onCompleted()
		subscription.Unsubscribe()
	})

	return subscription
}
