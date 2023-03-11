package rx

import (
	"sync"
	"time"
)

type publishSubject[T any] struct {
	mu sync.RWMutex
}

func (s *publishSubject[T]) OnCompleted() {

}

func (s *publishSubject[T]) OnError(err error) {

}

func (s *publishSubject[T]) OnNext(value T) {

}

func (s *publishSubject[T]) Delay(delay time.Duration) Observable[T] {
	panic("???")
}

func (s *publishSubject[T]) Filter(predicate func(T) bool) Observable[T] {
	panic("???")
}

func (s *publishSubject[T]) Subscribe(onNext func(T), onError func(err error), onCompleted func()) Subscription {
	panic("???")
}

func newPublishSubject[T any]() *publishSubject[T] {
	return &publishSubject[T]{}
}
