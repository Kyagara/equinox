package ratelimit

import (
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type Bucket struct {
	// Next reset
	next time.Time
	// Current number of tokens, starts at limit
	tokens int
	// Maximum number of tokens
	limit int
	// Time interval in seconds
	interval time.Duration
	mutex    sync.Mutex
}

func (b *Bucket) MarshalZerologObject(encoder *zerolog.Event) {
	encoder.Int("tokens", b.tokens).Int("limit", b.limit).Dur("interval", b.interval).Time("next", b.next)
}

func NewBucket(interval time.Duration, limit int, tokens int) *Bucket {
	return &Bucket{
		interval: interval * time.Second,
		limit:    limit,
		tokens:   tokens,
		next:     time.Now().Add(interval * time.Second),
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
