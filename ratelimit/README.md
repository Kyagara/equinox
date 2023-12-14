# Rate Limit

> Warning: this is still a work in progress. Tests succeeding != production ready.

Rate limiting is enabled by default in a default equinox client. For now the only `store` available is in-memory, though I want to add Redis support in the future, maybe using a lua script.

Info on rate limiting:

- [Hextechdocs](https://hextechdocs.dev/rate-limiting/)
- [Riot developer portal](https://developer.riotgames.com/docs/portal#web-apis_rate-limiting)

You can create an InternalRateLimit with `NewInternalRateLimit()`. RateLimit includes the following:

```go
type RateLimit struct {
	// 'any' is used here because routes can be PlatformRoute, RegionalRoute...
	Region  map[any]*Limits
	Enabled bool
	// Factor to be applied to the limit. E.g. if set to 0.5, the limit will be reduced by 50%.
	LimitUsageFactor float32
	// Delay in milliseconds to be add to reset intervals.
	IntervalOverhead time.Duration
	mutex            sync.Mutex
}
```

### Bucket

A limit for the App or Method is a bucket of tokens:

```go
type Bucket struct {
	// Time interval in seconds
	interval time.Duration
	// Maximum number of tokens
	limit int
	// Current number of tokens, starts at limit
	tokens int
	// Next reset
	next  time.Time
	mutex sync.Mutex
}
```

When creating a bucket, `interval` is the time in seconds between resets, `limit` is the maximum number of tokens taking into account the 'LimitUsageFactor', and `tokens` is the current number of tokens.

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

When initializing a bucket, the current amount of tokens will be the same as the limit minus the current count provided from the `X-App-Rate-Limit-Count` or `X-Method-Rate-Limit-Count` headers.

```go
func parseHeaders(limitHeader string, countHeader string) []*Bucket {
	if limitHeader == "" || countHeader == "" {
		return NewLimit(limitType)
	}

	limit := NewLimit(limitType)

	limits := strings.Split(limitHeader, ",")
	counts := strings.Split(countHeader, ",")
	rates := make([]*Bucket, len(limits))

	for i := range limits {
		baseLimit, interval := getNumbersFromPair(limits[i])
		count, _ := getNumbersFromPair(counts[i])
		limit := int(math.Floor(math.Max(1, float64(baseLimit)*float64(r.LimitUsageFactor))))
		rates[i] = NewBucket(interval, r.IntervalOverhead, baseLimit, limit, limit-count)
	}

	limit.buckets = rates
	return limit
}
```

- When a bucket is full, the amount of tokens will be the same as the limit - Not Rate limited, able to `Take()` a token from the bucket.
- When a bucket is empty, the amount of tokens will be 0 - Rate limited, not able to `Take()` a token without waiting.

### Take

After creating a request and checking if it was cached, the client will use `Take()`, initializing the App and Method buckets in a **route** AND the **MethodID** if not initialized.

If rate limited, `Take()` will block until the next bucket reset. A `context` can be passed, allowing for cancellation and checking if the deadline will be exceeded, cancelling the block if it will.

`Take()` will then decrease tokens for the App and Method in the buckets by one.

### Update

After receiving a response, `Update()` will verify that the current buckets in-memory match the ones received by the Riot API, if they don't, it will force an update in all buckets.

> By 'matching', I mean that the current **limit** and **interval** in these buckets match the ones received by the Riot API.
