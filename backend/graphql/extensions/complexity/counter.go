package complexity

import (
	"context"
	"sync/atomic"
)

// Counter is a simple counter that can be used to track the number of
// operations that have been performed.
type Counter struct {
	ready chan struct{}
	limit int32
	value int32
}

type ctxKey struct{}

// NewCounter returns a new counter with the given limit.
func NewCounter(limit int) *Counter {
	return &Counter{
		limit: int32(limit),
		value: 0,
		ready: make(chan struct{}, limit),
	}
}

// Add increments the counter by the given amount.
func (c *Counter) Add(delta int) error {
	val := atomic.AddInt32(&c.value, int32(delta))
	if val > c.limit {
		return ErrLimitExceeded
	}
	return nil
}

// MakeReady makes the counter ready for use.
// As long as calculated is not called,
// the c.Value() will not return.
// It is there to ensure that we do not get the value
// before the calculation is done.
func (c *Counter) MakeReady() {
	close(c.ready)
}

// Value returns the current value of the counter.
// It will block until the counter is calculated.
func (c *Counter) Value() int {
	<-c.ready
	return int(atomic.LoadInt32(&c.value))
}

// Limit returns the limit of the counter.
func (c *Counter) Limit() int {
	return int(c.limit)
}

// NewCounterCtx adds a counter to the context.
func NewCounterCtx(ctx context.Context, limit int) context.Context {
	if _, err := CounterFromCtx(ctx); err == nil {
		panic(ErrCounterExists)
	}
	return context.WithValue(ctx, ctxKey{}, NewCounter(limit))
}

// CounterFromCtx returns the counter from the context.
func CounterFromCtx(ctx context.Context) (*Counter, error) {
	counter, ok := ctx.Value(ctxKey{}).(*Counter)
	if !ok {
		return nil, ErrNoCounterInCtx
	}
	if counter == nil {
		return nil, ErrCounterIsNil
	}
	return counter, nil
}
