package ratelimit

import (
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type Bucket struct {
	// Next reset.
	Next time.Time
	// Current number of tokens, starts at limit.
	Tokens int
	// The limit given in the header without any modifications.
	BaseLimit int
	// Maximum amount of tokens, modified by the LimitUsageFactor.
	Limit int
	// Time interval in seconds.
	Interval         time.Duration
	IntervalOverhead time.Duration
	mutex            sync.Mutex
}

func (b *Bucket) MarshalZerologObject(encoder *zerolog.Event) {
	encoder.Int("t", b.Tokens).Int("l", b.Limit).Dur("i", b.Interval).Time("n", b.Next)
}

func NewBucket(interval time.Duration, intervalOverhead time.Duration, baseLimit int, limit int, tokens int) *Bucket {
	return &Bucket{
		Next:             time.Now().Add(interval + intervalOverhead),
		Tokens:           tokens,
		BaseLimit:        baseLimit,
		Limit:            limit,
		Interval:         interval,
		IntervalOverhead: intervalOverhead,
		mutex:            sync.Mutex{},
	}
}

// Checks if the 'next' reset is in the past, and if so, resets the bucket tokens and sets the next reset.
func (b *Bucket) Check() {
	now := time.Now()
	if b.Next.Before(now) {
		b.Tokens = 0
		b.Next = now.Add(b.Interval + b.IntervalOverhead)
	}
}

// Increments the number of tokens in the bucket and returns if the bucket is rate limited.
func (b *Bucket) IsRateLimited() bool {
	b.Check()
	if b.BaseLimit == 0 {
		return false
	}
	b.Tokens++
	return b.Tokens >= b.Limit
}
