package go_rx

import "time"

type justObservable[T any] struct {
	items []T
}

func (o *justObservable[T]) Delay(delay time.Duration) Observable[T] {
	panic("???")
}

func (o *justObservable[T]) Filter(predicate func(T) bool) Observable[T] {
	var filtered []T = nil

	for _, item := range o.items {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}

	return newJustObservable[T](filtered)
}

func (o *justObservable[T]) Subscribe(onNext func(T), onError func(err error), onCompleted func()) Subscription {
	for _, item := range o.items {
		onNext(item)
	}

	onCompleted()

	return newEmptySubscription()
}

func newJustObservable[T any](items []T) Observable[T] {
	return &justObservable[T]{
		items: items,
	}
}
