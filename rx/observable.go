package rx

import "time"

type Observable[T any] interface {
	Delay(delay time.Duration) Observable[T]
	Filter(predicate func(T) bool) Observable[T]
	Subscribe(onNext func(T), onError func(err error), onCompleted func()) Subscription
}
