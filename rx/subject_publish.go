package rx

import (
	"sync"
	"time"
)

type publishSubject[T any] struct {
	mu        sync.RWMutex
	observers []Observer[T]
}

func (s *publishSubject[T]) OnCompleted() {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, observer := range s.observers {
		observer.OnCompleted()
	}

	s.observers = nil
}

func (s *publishSubject[T]) OnError(err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, observer := range s.observers {
		observer.OnError(err)
	}

	s.observers = nil
}

func (s *publishSubject[T]) OnNext(value T) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, observer := range s.observers {
		observer.OnNext(value)
	}
}

func (s *publishSubject[T]) Delay(delay time.Duration) Observable[T] {
	return &transformedObservable[T, T]{
		parent: s,
		operator: func(value T, observer Observer[T]) {
			<-time.After(delay)
			observer.OnNext(value)
		},
	}
}

func (s *publishSubject[T]) Filter(predicate func(T) bool) Observable[T] {
	return &transformedObservable[T, T]{
		parent: s,
		operator: func(value T, observer Observer[T]) {
			if predicate(value) {
				observer.OnNext(value)
			}
		},
	}
}

func (s *publishSubject[T]) Subscribe(onNext func(T), onError func(err error), onCompleted func()) Subscription {
	observer := newObserver[T](onNext, onError, onCompleted)

	subscription := newCallbackSubscription(func() {
		s.mu.Lock()
		defer s.mu.Unlock()

		var observers []Observer[T]

		for _, o := range s.observers {
			if o != observer {
				observers = append(observers, o)
			}
		}

		s.observers = observers
	})

	s.mu.Lock()
	defer s.mu.Unlock()

	s.observers = append(s.observers, observer)

	return subscription
}

func newPublishSubject[T any]() *publishSubject[T] {
	return &publishSubject[T]{}
}
