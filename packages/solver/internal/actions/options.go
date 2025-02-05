package actions

import (
	"solver/internal/utils"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

const (
	ActionUnstake = "unstake"
	ActionClaim   = "claim"
	ActionBridge  = "bridge"
)

type OptionInfo struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type Option struct {
	Label string     `json:"label"`
	Name  string     `json:"name"`
	Value string     `json:"value"`
	Icon  string     `json:"icon,omitempty"`
	Info  OptionInfo `json:"info,omitempty"`
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
}

type OptionsProvider interface {
	GetOptions(chainId uint64, from common.Address, action string) (map[int]Options, error)
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

func (c *CachedOptionsProvider) GetOptions(chainId uint64, from common.Address, action string) (map[int]Options, error) {
	key := OptionCacheKey{chainId: chainId, action: action}

	c.mu.RLock()
	if cached, ok := c.cache[key]; ok {
		c.mu.RUnlock()
		return cached.options, nil
	}
	c.mu.RUnlock()

	options, err := c.provider.GetOptions(chainId, from, action)
	if err != nil {
		return nil, utils.ErrOptions(err.Error())
	}

	if options == nil {
		options = make(map[int]Options)
	}

	c.mu.Lock()
	c.cache[key] = CachedOptions{
		options: options,
	}
	c.mu.Unlock()

	return options, nil
}

func (c *CachedOptionsProvider) PreWarmCache(chainId uint64, from common.Address, actions []string) {
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
			if _, err := c.GetOptions(chainId, from, action); err != nil {
				// Silently continue on error
			}
		}(action)
	}

	for i := 0; i < maxWorkers; i++ {
		sem <- struct{}{}
	}
}

// DefaultOptionsProvider is a basic implementation of OptionsProvider
type DefaultOptionsProvider struct{}

// GetOptions implements OptionsProvider
func (p *DefaultOptionsProvider) GetOptions(chainId uint64, from common.Address, action string) (map[int]Options, error) {
	return make(map[int]Options), nil
}
