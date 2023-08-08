package go_rx

import "sync"

// Subscription : 購読を管理するためのインターフェース
type Subscription interface {
	// IsUnsubscribed : 購読が解除済みかどうか
	IsUnsubscribed() bool

	// Unsubscribe : 購読を解除する
	Unsubscribe()
}

// callbackSubscription : 購読解除時にコールバック関数を実行する Subscription
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

// newCallbackSubscription : 購読解除時にコールバック関数を実行する Subscription を生成する
func newCallbackSubscription(callback func()) *callbackSubscription {
	return &callbackSubscription{
		isUnsubscribed: false,
		callback:       callback,
	}
}

// emptySubscription : 空の Subscription
type emptySubscription struct {
}

func (s *emptySubscription) IsUnsubscribed() bool {
	return true
}

func (s *emptySubscription) Unsubscribe() {
}

// newEmptySubscription : 空の Subscription を生成する
func newEmptySubscription() *emptySubscription {
	return &emptySubscription{}
}
