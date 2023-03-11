package rx

import "sync"

type Subscription interface {
	// IsUnsubscribed :
	IsUnsubscribed() bool

	// Unsubscribe :
	Unsubscribe()
}

type callbackSubscription struct {
	mu             sync.Mutex
	isUnsubscribed bool
	callback       func()
}

func (s *callbackSubscription) IsUnsubscribed() bool {
	s.mu.Lock()
	unsubscribed := s.isUnsubscribed
	s.mu.Unlock()

	return unsubscribed
}

func (s *callbackSubscription) Unsubscribe() {
	s.mu.Lock()
	defer s.mu.Unlock()

	unsubscribed := s.isUnsubscribed

	if unsubscribed {
		return
	}

	s.isUnsubscribed = true

	s.callback()
}

func newCallbackSubscription(callback func()) *callbackSubscription {
	return &callbackSubscription{
		isUnsubscribed: false,
		callback:       callback,
	}
}

type emptySubscription struct {
}

func (s *emptySubscription) IsUnsubscribed() bool {
	return true
}

func (s *emptySubscription) Unsubscribe() {
}

func newEmptySubscription() *emptySubscription {
	return &emptySubscription{}
}
