package rx

import (
	"context"
	"sync"
)

func Map[P any, T any](parent Observable[P], operator func(P) T) Observable[T] {
	return &transformedObservable[P, T]{
		parent: parent,
		operator: func(value P, observer Observer[T]) {
			observer.OnNext(operator(value))
		},
	}
}

func FlatMap[P, T any](parent Observable[P], operator func(P) Observable[T]) Observable[T] {
	return &transformedObservable[P, T]{
		parent: parent,
		operator: func(value P, observer Observer[T]) {
			var subscription Subscription
			subscription = operator(value).Subscribe(func(value T) {
				observer.OnNext(value)
			}, func(err error) {
				observer.OnError(err)
			}, func() {
				subscription.Unsubscribe()
			})
		},
	}
}

func WithContext[T any](ctx context.Context, observable Observable[T]) Observable[T] {
	var mu sync.RWMutex
	subject := newPublishSubject[T]()
	closed := false

	var subscription Subscription
	subscription = observable.Subscribe(func(value T) {
		mu.Lock()
		defer mu.Unlock()

		if closed {
			return
		}

		select {
		case <-ctx.Done():
		default:
			subject.OnNext(value)
		}
	}, func(err error) {
		mu.Lock()
		defer mu.Unlock()

		if closed {
			return
		}

		closed = true

		if !subscription.IsUnsubscribed() {
			subscription.Unsubscribe()
		}

		subject.OnError(err)
	}, func() {
		mu.Lock()
		defer mu.Unlock()

		if closed {
			return
		}

		closed = true

		if !subscription.IsUnsubscribed() {
			subscription.Unsubscribe()
		}

		subject.OnCompleted()
	})

	go func() {
		select {
		case <-ctx.Done():
			mu.Lock()
			defer mu.Unlock()

			if closed {
				return
			}

			closed = true

			if !subscription.IsUnsubscribed() {
				subscription.Unsubscribe()
			}

			subject.OnCompleted()
		}
	}()

	return subject
}

func WaitFirst[T any](ctx context.Context, observable Observable[T]) <-chan Item[T] {
	var mu sync.Mutex
	received := make(chan Item[T])
	done := false

	go (func() {
		var subscription Subscription

		subscription = observable.Subscribe(func(value T) {
			mu.Lock()
			defer mu.Unlock()

			if done {
				return
			}

			done = true
			received <- NewValueItem[T](value)
			subscription.Unsubscribe()
		}, func(err error) {
			mu.Lock()
			defer mu.Unlock()

			if done {
				return
			}

			done = true
			received <- NewErrorItem[T](err)
			subscription.Unsubscribe()
		}, func() {
			mu.Lock()
			defer mu.Unlock()

			if done {
				return
			}

			done = true
			close(received)
			subscription.Unsubscribe()
		})
	})()

	return received
}
