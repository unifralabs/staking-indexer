package config

import (
	"fmt"
)

const (
	defaultCacheSize         = 100
	defaultConfirmationDepth = 6
)

type BTCScannerConfig struct {
	CacheSize         uint64 `long:"cachesize" description:"max number of BTC blocks in the cache"`
	ConfirmationDepth uint64 `long:"confirmationdepth" description:"the confirmation depth to consider a BTC block as confirmed"`
}

func (cfg *BTCScannerConfig) Validate() error {
	if cfg.CacheSize < defaultCacheSize {
		return fmt.Errorf("btc-cache-size should not be less than %v", defaultCacheSize)
	}
	if cfg.ConfirmationDepth < defaultConfirmationDepth {
		return fmt.Errorf("btc-confirmation-depth should not be less than %d", defaultConfirmationDepth)
	}
	return nil
}

func DefaultBTCScannerConfig() *BTCScannerConfig {
	return &BTCScannerConfig{
		CacheSize:         defaultCacheSize,
		ConfirmationDepth: defaultConfirmationDepth,
	}
}
