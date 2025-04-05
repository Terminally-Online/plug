package cache

import (
	"time"
)

var (
	CachePeriod = 5 * time.Minute
)

type CacheOptions struct {
	duration    time.Duration
	useStale    bool
	staleBuffer time.Duration
}

type CacheOption func(*CacheOptions)

func defaultCacheOptions() *CacheOptions {
	return &CacheOptions{
		duration:    CachePeriod,
		useStale:    false,
		staleBuffer: 10 * time.Minute,
	}
}

func WithOptions(options ...CacheOption) []CacheOption {
	return options
}

func WithDuration(duration time.Duration) CacheOption {
	return func(o *CacheOptions) {
		o.duration = duration
	}
}

func WithStaleData(useStale bool) CacheOption {
	return func(o *CacheOptions) {
		o.useStale = useStale
	}
}

func WithStaleBuffer(buffer time.Duration) CacheOption {
	return func(o *CacheOptions) {
		o.staleBuffer = buffer
	}
}

