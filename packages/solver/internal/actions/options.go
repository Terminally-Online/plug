package actions

import (
	"solver/internal/utils"

	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ActionUnstake = "unstake"
	ActionClaim   = "claim"
	ActionBridge  = "bridge"
)

type OptionIcon struct {
	Default   string `json:"default,omitempty"`
	Secondary string `json:"secondary,omitempty"`
}

type OptionInfo struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Option struct {
	Label string     `json:"label"`
	Name  string     `json:"name"`
	Value string     `json:"value"`
	Icon  OptionIcon `json:"icon,omitempty"`
	Info  OptionInfo `json:"info,omitempty"`
}

type Options struct {
	Simple  []Option            `json:"simple,omitempty"`
	Complex map[string][]Option `json:"complex,omitempty"`
}

type OptionCacheKey struct {
	chainId uint64
	from    common.Address
	action  string
}

type CachedOptions struct {
	options     map[int]Options
	lastUpdated time.Time
	refreshing  bool
}

const cacheDuration = 5 * time.Minute

type OptionsProvider interface {
	GetOptions(chainId uint64, from common.Address, search map[int]string, action string) (map[int]Options, error)
}

type CachedOptionsProvider struct {
	provider OptionsProvider
	cache    map[OptionCacheKey]CachedOptions
	mu       sync.RWMutex
}

var (
	defaultProvider *CachedOptionsProvider
	providerMu      sync.RWMutex
)

func GetCachedOptionsProvider() *CachedOptionsProvider {
	providerMu.RLock()
	defer providerMu.RUnlock()
	return defaultProvider
}

func SetCachedOptionsProvider(provider *CachedOptionsProvider) {
	providerMu.Lock()
	defaultProvider = provider
	providerMu.Unlock()
}

func GetSupportedActions() []string {
	return []string{
		ActionDeposit,
		ActionWithdraw,
		ActionBorrow,
		ActionRepay,
		ActionSwap,
		ActionStake,
		ActionUnstake,
		ActionClaim,
		ActionRenew,
		ActionBridge,
	}
}

func NewCachedOptionsProvider(provider OptionsProvider) *CachedOptionsProvider {
	return &CachedOptionsProvider{
		provider: provider,
		cache:    make(map[OptionCacheKey]CachedOptions, 32),
	}
}

func (c *CachedOptionsProvider) GetOptions(chainId uint64, from common.Address, search map[int]string, action string) (map[int]Options, error) {
	if (from == common.Address{}) {
		from = utils.ZeroAddress
	}

	key := OptionCacheKey{
		chainId: chainId,
		from:    from,
		action:  action,
	}

	options, err := c.GetOrCreateCachedOptions(key)
	if err != nil {
		return nil, err
	}

	if len(search) == 0 {
		return options, nil
	}

	return c.FilterOptions(options, search), nil
}

func (c *CachedOptionsProvider) GetOrCreateCachedOptions(key OptionCacheKey) (map[int]Options, error) {
	// Check cache first with read lock
	c.mu.RLock()
	cached, exists := c.cache[key]
	if exists {
		isStale := time.Since(cached.lastUpdated) >= cacheDuration
		c.mu.RUnlock()

		if isStale {
			go c.RefreshCache(key)
		}
		return cached.options, nil
	}
	c.mu.RUnlock()

	// No cached data exists, need to get fresh options
	c.mu.Lock()
	defer c.mu.Unlock()

	options, err := c.provider.GetOptions(key.chainId, key.from, nil, key.action)
	if err != nil {
		return nil, utils.ErrOptions(err.Error())
	}

	if options == nil {
		options = make(map[int]Options)
	}

	c.cache[key] = CachedOptions{
		options:     options,
		lastUpdated: time.Now(),
		refreshing:  false,
	}

	return options, nil
}

func (c *CachedOptionsProvider) RefreshCache(key OptionCacheKey) {
	c.mu.Lock()
	cached, exists := c.cache[key]
	if !exists || cached.refreshing || time.Since(cached.lastUpdated) < cacheDuration {
		c.mu.Unlock()
		return
	}

	c.cache[key] = CachedOptions{
		options:     cached.options,
		lastUpdated: cached.lastUpdated,
		refreshing:  true,
	}
	c.mu.Unlock()

	options, err := c.provider.GetOptions(key.chainId, key.from, nil, key.action)

	c.mu.Lock()
	defer c.mu.Unlock()

	if err != nil {
		c.cache[key] = CachedOptions{
			options:     cached.options,
			lastUpdated: cached.lastUpdated,
			refreshing:  false,
		}
		return
	}

	if options == nil {
		options = make(map[int]Options)
	}

	c.cache[key] = CachedOptions{
		options:     options,
		lastUpdated: time.Now(),
		refreshing:  false,
	}
}

func (c *CachedOptionsProvider) FilterOptions(options map[int]Options, search map[int]string) map[int]Options {
	filtered := make(map[int]Options, len(options))
	for k, v := range options {
		filtered[k] = v
		if search[k] != "" {
			var matchedOpts []Option
			searchTerm := strings.ToLower(search[k])

			for _, opt := range v.Simple {
				if strings.Contains(strings.ToLower(opt.Label), searchTerm) ||
					strings.Contains(strings.ToLower(opt.Name), searchTerm) ||
					strings.Contains(strings.ToLower(opt.Value), searchTerm) {
					matchedOpts = append(matchedOpts, opt)
				}
			}

			filtered[k] = Options{Simple: matchedOpts}
		}
	}
	return filtered
}

func (c *CachedOptionsProvider) PreWarmCache(chainId uint64, from common.Address, actions []string) {
	if (from == common.Address{}) {
		from = utils.ZeroAddress
	}

	const maxWorkers = 8
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
			if _, err := c.GetOptions(chainId, from, map[int]string{}, action); err != nil {
				// Silently continue on error
			}
		}(action)
	}

	for i := 0; i < maxWorkers; i++ {
		sem <- struct{}{}
	}
}

type DefaultOptionsProvider struct{}

func (p *DefaultOptionsProvider) GetOptions(chainId uint64, from common.Address, search map[int]string, action string) (map[int]Options, error) {
	return make(map[int]Options), nil
}
