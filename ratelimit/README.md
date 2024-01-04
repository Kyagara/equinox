# Rate Limit

Rate limiting is enabled by default in a default equinox client. For now the only `store` available is internal, though I want to add Redis support in the future, maybe using a lua script.

Info on rate limiting:

- [Hextechdocs](https://hextechdocs.dev/rate-limiting/)
- [Riot developer portal](https://developer.riotgames.com/docs/portal#web-apis_rate-limiting)

You can create an InternalRateLimit with `NewInternalRateLimit()`. RateLimit includes the following:

```go
type RateLimit struct {
	Region  map[string]*Limits
	Enabled bool
	// Factor to be applied to the limit. E.g. if set to 0.5, the limit will be reduced by 50%.
	LimitUsageFactor float32
	// Delay in milliseconds to be add to reset intervals.
	IntervalOverhead time.Duration
	mutex            sync.Mutex
}

func NewInternalRateLimit(limitUsageFactor float32, intervalOverhead time.Duration) *RateLimit {
	if limitUsageFactor < 0.0 || limitUsageFactor > 1.0 {
		limitUsageFactor = 0.99
	}
	if intervalOverhead < 0 {
		intervalOverhead = 1 * time.Second
	}
	return &RateLimit{
		Region:           make(map[string]*Limits, 1),
		LimitUsageFactor: limitUsageFactor,
		IntervalOverhead: intervalOverhead,
		Enabled:          true,
	}
}
```

### Bucket

A `Limit` for the App/Method is a `Bucket`:

```go
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
```

When creating a bucket, `interval` is the time in seconds between resets, `limit` is the maximum number of tokens taking into account the `LimitUsageFactor` from the `RateLimit`, and `tokens` is the current number of tokens.

```go
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
```

- When a bucket is full, the amount of tokens will be the same as the limit - Not Rate limited, able to `Take()` a token from the bucket.
- When a bucket is empty, the amount of tokens will be 0 - Rate limited, not able to `Take()` a token without waiting.

When initializing a bucket, the current amount of tokens will be the same as the limit minus the current count provided from the `X-App-Rate-Limit-Count` or `X-Method-Rate-Limit-Count` headers.

```go
func (r *RateLimit) parseHeaders(limitHeader string, countHeader string, limitType string) *Limit {
	if limitHeader == "" || countHeader == "" {
		return NewLimit(limitType)
	}

	limits := strings.Split(limitHeader, ",")
	counts := strings.Split(countHeader, ",")

	if len(limits) == 0 {
		return NewLimit(limitType)
	}

	limit := &Limit{
		buckets:    make([]*Bucket, len(limits)),
		limitType:  limitType,
		retryAfter: 0,
		mutex:      sync.Mutex{},
	}

	for i, limitString := range limits {
		baseLimit, interval := getNumbersFromPair(limitString)
		count, _ := getNumbersFromPair(counts[i])
		newLimit := int(math.Max(1, float64(baseLimit)*r.LimitUsageFactor))
		currentTokens := newLimit - count
		limit.buckets[i] = NewBucket(interval, r.IntervalOverhead, baseLimit, newLimit, currentTokens)
	}

	return limit
}
```

### Take

After creating a request and checking if it was cached, the client will use `Take()`, initializing the App and Method buckets in a **route** AND the **MethodID** if not initialized.

If rate limited, `Take()` will block until the next bucket reset. A `context` can be passed, allowing for cancellation, a check will be done to see if waiting would exceed the deadline set in a context, returning an error if it would.

`Take()` will then decrease tokens for the App and Method in the buckets involved by one.

### Update

After receiving a response, `Update()` will verify that the current buckets in memory match the ones received by the Riot API, if they don't, it will force an update in all buckets.

> By 'matching', I mean that the current **baseLimit** and **interval** in the buckets already in memory match the ones received by the Riot API.
