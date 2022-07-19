package rx

type Subscription interface {
	IsUnsubscribed() bool
	Unsubscribe()
}
