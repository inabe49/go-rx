package rx

type Observable[T any] interface {
	Filter(predicate func(T) bool) Observable[T]
	Subscribe(onNext func(T), onError func(err error), onCompleted func()) Subscription
}
