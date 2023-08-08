package go_rx

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestWithContext_Complete(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	subject := newPublishSubject[int32]()
	done := false
	var mu sync.Mutex

	subscription := WithContext[int32](ctx, subject).Subscribe(func(value int32) {
		if value > 2 {
			assert.Fail(t, "receive value after cancel")
		}
	}, func(err error) {
		assert.Fail(t, "receive error")
	}, func() {
		mu.Lock()
		defer mu.Unlock()

		fmt.Println("complete")

		assert.False(t, done)
		done = true
	})

	subject.OnNext(1)
	subject.OnNext(2)

	cancel()

	subject.OnNext(3)
	subject.OnNext(4)

	subscription.Unsubscribe()
}

func TestWaitFirst_Success(t *testing.T) {
	ctx := context.Background()
	subject := newPublishSubject[int32]()

	go (func() {
		<-time.After(1 * time.Second)
		subject.OnNext(1)
	})()

	<-WaitFirst[int32](ctx, subject)

	assert.True(t, true)
}
