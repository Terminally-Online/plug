package actions

import (
	"solver/internal/utils"
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

type Options struct {
	Simple  []Option            `json:"simple,omitempty"`
	Complex map[string][]Option `json:"complex,omitempty"`
}

type OptionCacheKey struct {
	chainId uint64
	action  string
}

type CachedOptions struct {
	options map[int]Options
	expiry  time.Time
}

type OptionsProvider interface {
	GetOptions(chainId uint64, action string) (map[int]Options, error)
}

type CachedOptionsProvider struct {
	provider OptionsProvider
	cache    map[OptionCacheKey]CachedOptions
	mu       sync.RWMutex
}

func NewCachedOptionsProvider(provider OptionsProvider) *CachedOptionsProvider {
	return &CachedOptionsProvider{
		provider: provider,
		cache:    make(map[OptionCacheKey]CachedOptions, 32),
	}
}

func (c *CachedOptionsProvider) GetOptions(chainId uint64, action string) (map[int]Options, error) {
	key := OptionCacheKey{chainId: chainId, action: action}

	c.mu.RLock()
	if cached, ok := c.cache[key]; ok && time.Now().Before(cached.expiry) {
		c.mu.RUnlock()
		return cached.options, nil
	}
	c.mu.RUnlock()

	options, err := c.provider.GetOptions(chainId, action)
	if err != nil {
		return nil, utils.ErrOptions(err.Error())
	}

	if options == nil {
		options = make(map[int]Options)
	}

	c.mu.Lock()
	c.cache[key] = CachedOptions{
		options: options,
		expiry:  time.Now().Add(30 * time.Minute),
	}
	c.mu.Unlock()

	return options, nil
}

func (c *CachedOptionsProvider) PreWarmCache(chainId uint64, actions []string) {
	const maxWorkers = 4
	sem := make(chan struct{}, maxWorkers)

	completed := 0
	var mu sync.Mutex

	for _, action := range actions {
		sem <- struct{}{}
		go func(action string) {
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
