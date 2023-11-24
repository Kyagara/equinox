package ratelimit

import (
	"context"
	"sync"
	"time"

	"go.uber.org/zap/zapcore"
)

type Bucket struct {
	// Time interval in seconds
	interval time.Duration
	// Maximum number of tokens
	limit int
	// Current number of tokens
	tokens int
	// Updates every time its checked
	updated time.Time
	// Next reset
	next  time.Time
	mutex sync.Mutex
}

func (b *Bucket) MarshalLogObject(encoder zapcore.ObjectEncoder) error {
	encoder.AddDuration("interval", b.interval)
	encoder.AddInt("tokens", b.tokens)
	encoder.AddInt("limit", b.limit)
	return nil
}

func NewBucket(interval time.Duration, limit int, tokens int) *Bucket {
	now := time.Now()
	return &Bucket{
		interval: interval * time.Second,
		limit:    limit,
		tokens:   tokens,
		updated:  now,
		next:     now.Add(interval * time.Second),
		mutex:    sync.Mutex{},
	}
}

// Responsible for updating the bucket, resets the tokens if necessary.
func (b *Bucket) check() {
	now := time.Now()
	if b.next.Before(now) {
		b.tokens = b.limit
		b.next = now.Add(b.interval)
	}
	b.updated = now
}

func (b *Bucket) IsRateLimited(ctx context.Context) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.check()
	if b.limit == 0 {
		return nil
	}
	if b.tokens-1 <= 0 {
		deadline, ok := ctx.Deadline()
		if ok && deadline.Before(b.next) {
			return ErrContextDeadlineExceeded
		}
		return ErrRateLimited
	}
	b.tokens--
	return nil
}

// wait should block if the rate limit is reached and wait until the bucket resets.
func (b *Bucket) wait(ctx context.Context) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	timer := time.NewTimer(time.Until(b.next))
	defer timer.Stop()
	select {
	case <-timer.C:
		b.check()
		if b.tokens-1 <= 0 {
			return ErrRateLimited
		}
		b.tokens--
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
