package rx

import "sync"

type Subscription interface {
	IsUnsubscribed() bool
	Unsubscribe()
}

type subscriptionImpl struct {
	mu             sync.Mutex
	isUnsubscribed bool
	callback       func()
}

func (s *subscriptionImpl) IsUnsubscribed() bool {
	s.mu.Lock()
	unsubscribed := s.isUnsubscribed
	s.mu.Unlock()

	return unsubscribed
}

func (s *subscriptionImpl) Unsubscribe() {
	s.mu.Lock()
	unsubscribed := s.isUnsubscribed

	if unsubscribed {
		s.mu.Unlock()
		return
	}

	s.isUnsubscribed = true

	s.mu.Unlock()

	s.callback()
}

func newSubscriptionImpl(callback func()) *subscriptionImpl {
	return &subscriptionImpl{
		isUnsubscribed: false,
		callback:       callback,
	}
}
