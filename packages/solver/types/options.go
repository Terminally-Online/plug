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

type cachedOptions struct {
	options map[int]SchemaOptions
	expiry  time.Time
}

type CachedOptionsProvider struct {
	provider OptionsProvider
	cache    map[cacheKey]cachedOptions
	mu       sync.RWMutex
}

type cacheKey struct {
	chainId int
	action  Action
}

func NewCachedOptionsProvider(provider OptionsProvider) *CachedOptionsProvider {
	return &CachedOptionsProvider{
		provider: provider,
		cache:    make(map[cacheKey]cachedOptions),
	}
}

func (c *CachedOptionsProvider) GetOptions(chainId int, action Action) (map[int]SchemaOptions, error) {
	key := cacheKey{chainId: chainId, action: action}

	// Try to get from cache first
	c.mu.RLock()
	if cached, ok := c.cache[key]; ok && time.Now().Before(cached.expiry) {
		c.mu.RUnlock()
		return cached.options, nil
	}
	c.mu.RUnlock()

	// If not in cache or expired, get fresh options
	options, err := c.provider.GetOptions(chainId, action)
	if err != nil {
		return nil, fmt.Errorf("failed to get options: %w", err)
	}

	// Update cache
	c.mu.Lock()
	c.cache[key] = cachedOptions{
		options: options,
		expiry:  time.Now().Add(30 * time.Minute),
	}
	c.mu.Unlock()

	return options, nil
}

func (c *CachedOptionsProvider) PreWarmCache(chainId int, actions []Action) {
	go func() {
		for _, action := range actions {
			// Fetch options and store in cache
			if _, err := c.GetOptions(chainId, action); err != nil {
				// Just log the error and continue, don't block other actions
				fmt.Printf("Failed to pre-warm cache for action %s: %v\n", action, err)
			}
		}
	}()
}
