# Cache

This package provides an interface to interact with different cache `stores`, like BigCache and Redis. It allows you to perform common cache operations like `Get`/`Set`/`Delete` and `Clear` the entire cache.

Cache is optional, you can pass nil when building the equinox client and it will disable caching. The cache is enabled by default and uses BigCache.

The idea is to keep this package small and simple, providing only one in-memory cache and an external database, preferably Redis.

You can interact with the cache from the equinox client:

```go
client, err := equinox.NewClient("RIOT_API_KEY")
body, err := client.Cache.Get("https://euw1.api.riotgames.com/...") // []byte, error
```

## Stores

- [BigCache](https://github.com/allegro/bigcache) - In-memory cache
- [Redis](https://github.com/redis/go-redis) - Redis cache
