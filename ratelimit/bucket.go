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
	// The limit given in the header without any modification
	baseLimit int
	// Maximum number of tokens
	limit int
	// Time interval in seconds
	interval         time.Duration
	intervalOverhead time.Duration
	mutex            sync.Mutex
}

func (b *Bucket) MarshalZerologObject(encoder *zerolog.Event) {
	encoder.Int("tokens", b.tokens).Int("limit", b.limit).Dur("interval", b.interval).Time("next", b.next)
}

func NewBucket(interval time.Duration, intervalOverhead time.Duration, baseLimit int, limit int, tokens int) *Bucket {
	return &Bucket{
		interval:         interval,
		intervalOverhead: intervalOverhead,
		baseLimit:        baseLimit,
		limit:            limit,
		tokens:           tokens,
		next:             time.Now().Add(interval + intervalOverhead),
		mutex:            sync.Mutex{},
	}
}

// Checks if the 'next' reset is in the past, and if so, resets the bucket tokens and sets the next reset.
func (b *Bucket) check() {
	now := time.Now()
	if b.next.Before(now) {
		b.tokens = 0
		b.next = now.Add(b.interval + b.intervalOverhead)
	}
}

// Increments the number of tokens in the bucket and returns if the bucket is rate limited.
func (b *Bucket) isRateLimited() bool {
	b.check()
	if b.limit == 0 {
		return false
	}
	b.tokens++
	return b.tokens >= b.limit
}
