package go_rx

import "time"

type Observable[T any] interface {
	// Delay :
	Delay(delay time.Duration) Observable[T]

	// Filter :
	Filter(predicate func(T) bool) Observable[T]

	// Subscribe :
	Subscribe(onNext func(T), onError func(err error), onCompleted func()) Subscription
}
