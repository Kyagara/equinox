# Rate Limit

> Warning: this is still a work in progress. Tests succeeding != production ready.

Rate limiting is enabled by default in a default Equinox client. For now the only `store` available is in-memory, though I want to add Redis support in the future, maybe using a lua script.

Info on rate limiting:

- [Hextechdocs](https://hextechdocs.dev/rate-limiting/)
- [Riot developer portal](https://developer.riotgames.com/docs/portal#web-apis_rate-limiting)

## How does it work

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

When creating a bucket, `interval` is the time in seconds between resets, `limit` is the maximum number of tokens, and `tokens` is the current number of tokens.

```go
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
```

When initializing a bucket, the current amount of tokens will be the same as the limit minus the current count provided from the `X-App-Rate-Limit-Count` or `X-Method-Rate-Limit-Count` headers.

```go
func parseHeaders(limitHeader string, countHeader string) []*Bucket {
	if limitHeader == "" || countHeader == "" {
		return []*Bucket{}
	}
	limits := strings.Split(limitHeader, ",")
	counts := strings.Split(countHeader, ",")
	rates := make([]*Bucket, len(limits))
	for i := range limits {
		limit, seconds := getNumbersFromPair(limits[i])
		count, _ := getNumbersFromPair(counts[i])
		rates[i] = NewBucket(time.Duration(seconds), limit, limit-count)
	}
	return rates
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
