package types

import (
	"fmt"
	"sync"
	"time"
)

type Option struct {
	Label string `json:"label"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Icon  string `json:"icon,omitempty"`
	Info  string `json:"info,omitempty"`
}

type SchemaOptions struct {
	Simple  []Option            `json:"simple,omitempty"`
	Complex map[string][]Option `json:"complex,omitempty"`
}

type OptionsProvider interface {
	GetOptions(chainId int, action Action) (map[int]SchemaOptions, error)
}

// Cache error messages
var (
	errFailedGetOptions = "failed to get options: %w"
)

type cacheKey struct {
	chainId int
	action  Action
}

type cachedOptions struct {
	options map[int]SchemaOptions
	expiry  time.Time
}

type CachedOptionsProvider struct {
	provider OptionsProvider
	cache    map[cacheKey]cachedOptions
	mu       sync.RWMutex
}

func NewCachedOptionsProvider(provider OptionsProvider) *CachedOptionsProvider {
	// Pre-allocate cache map with a reasonable initial size
	return &CachedOptionsProvider{
		provider: provider,
		cache:    make(map[cacheKey]cachedOptions, 32), // Space for 32 action/chain combinations
	}
}

func (c *CachedOptionsProvider) GetOptions(chainId int, action Action) (map[int]SchemaOptions, error) {
	key := cacheKey{chainId: chainId, action: action}

	// Try to get from cache first using RLock
	c.mu.RLock()
	if cached, ok := c.cache[key]; ok && time.Now().Before(cached.expiry) {
		c.mu.RUnlock()
		return cached.options, nil
	}
	c.mu.RUnlock()

	// If not in cache or expired, get fresh options
	options, err := c.provider.GetOptions(chainId, action)
	if err != nil {
		return nil, fmt.Errorf(errFailedGetOptions, err)
	}

	// Pre-allocate the options map if needed
	if options == nil {
		options = make(map[int]SchemaOptions)
	}

	// Update cache with write lock
	c.mu.Lock()
	c.cache[key] = cachedOptions{
		options: options,
		expiry:  time.Now().Add(30 * time.Minute),
	}
	c.mu.Unlock()

	return options, nil
}

func (c *CachedOptionsProvider) PreWarmCache(chainId int, actions []Action) {
	const maxWorkers = 4
	sem := make(chan struct{}, maxWorkers)

	completed := 0
	var mu sync.Mutex

	for _, action := range actions {
		sem <- struct{}{}
		go func(action Action) {
			defer func() {
				<-sem
				mu.Lock()
				completed++
				mu.Unlock()
			}()
			if _, err := c.GetOptions(chainId, action); err != nil {
				// Silently continue on error
			}
		}(action)
	}

	for i := 0; i < maxWorkers; i++ {
		sem <- struct{}{}
	}
}
