package ratelimit

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type Bucket struct {
	// Current number of tokens, starts at limit
	tokens int
	// Maximum number of tokens
	limit int
	// Time interval in seconds
	interval time.Duration
	// Next reset
	next  time.Time
	mutex sync.Mutex
}

func (b *Bucket) MarshalZerologObject(encoder *zerolog.Event) {
	encoder.Int("tokens", b.tokens).Int("limit", b.limit).Dur("interval", b.interval).Time("next", b.next)
}

func NewBucket(interval time.Duration, limit int, tokens int) *Bucket {
	now := time.Now()
	return &Bucket{
		interval: interval * time.Second,
		limit:    limit,
		tokens:   tokens,
		next:     now.Add(interval * time.Second),
		mutex:    sync.Mutex{},
	}
}

func (b *Bucket) check() {
	now := time.Now()
	if b.next.Before(now) {
		b.tokens = b.limit
		b.next = now.Add(b.interval)
	}
}

func (b *Bucket) isRateLimited() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.check()
	if b.limit == 0 {
		return false
	}
	if b.tokens-1 <= 0 {
		return true
	}
	b.tokens--
	return false
}

// wait() should block if the rate limit is reached and wait until the bucket resets.
func (b *Bucket) wait(ctx context.Context) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.check()
	deadline, ok := ctx.Deadline()
	if ok && deadline.Before(b.next) {
		return ErrContextDeadlineExceeded
	}
	select {
	case <-time.After(time.Until(b.next)):
		b.check()
		b.tokens--
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
