package memcached

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"

	cortex_cache "github.com/cortexproject/cortex/pkg/chunk/cache"
	"github.com/grafana/tempo/tempodb/backend/cache"
)

type Config struct {
	ClientConfig cortex_cache.MemcachedClientConfig `yaml:",inline"`

	TTL time.Duration `yaml:"ttl"`
}

type Cache struct {
	client *cortex_cache.Memcached
}

func NewCache(cfg *Config, logger log.Logger) cache.Cache {
	if cfg.ClientConfig.MaxIdleConns == 0 {
		cfg.ClientConfig.MaxIdleConns = 16
	}
	if cfg.ClientConfig.Timeout == 0 {
		cfg.ClientConfig.Timeout = 100 * time.Millisecond
	}
	if cfg.ClientConfig.UpdateInterval == 0 {
		cfg.ClientConfig.UpdateInterval = time.Minute
	}

	client := cortex_cache.NewMemcachedClient(cfg.ClientConfig, "tempo", prometheus.DefaultRegisterer, logger)
	memcachedCfg := cortex_cache.MemcachedConfig{
		Expiration:  cfg.TTL,
		BatchSize:   0, // we are currently only requesting one key at a time, which is bad.  we could restructure Find() to batch request all blooms at once
		Parallelism: 0,
	}
	return &Cache{
		client: cortex_cache.NewMemcached(memcachedCfg, client, "tempo", prometheus.DefaultRegisterer, logger),
	}
}

func (m *Cache) Store(ctx context.Context, key string, val []byte) {
	m.client.Store(ctx, []string{key}, [][]byte{val})
}

func (m *Cache) Fetch(ctx context.Context, key string) []byte {
	found, vals, _ := m.client.Fetch(ctx, []string{key})
	if len(found) > 0 {
		return vals[0]
	}
	return nil
}

func (m *Cache) Shutdown() {
	m.client.Stop()
}
