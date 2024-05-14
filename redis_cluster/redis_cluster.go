package redis_cluster

import (
	"context"
	"errors"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/extension"
	"github.com/redis/go-redis/v9"
)

const (
	// DriverName the unique name of the Redis Cluster driver for register
	DriverName = "redis-cluster"
)

type redisClusterCache struct {
	nscache.NSCacheExt

	conn       *url.URL
	serializer encoding.Serializer
	mu         sync.RWMutex
	client     *redis.ClusterClient
}

func newCache(conn *url.URL) (nscache.NSCache, error) {
	opt, err := parseRedisClusterConnection(conn)
	if err != nil {
		return nil, err
	}
	c := &redisClusterCache{
		conn:       conn,
		serializer: encoding.DefaultSerializer,
		client:     redis.NewClusterClient(opt),
	}
	c.NSCacheExt = extension.New(c)
	return c, nil
}

func (c *redisClusterCache) Get(k string, v any) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, err := c.client.Get(context.Background(), k).Bytes()
	if err == redis.Nil {
		err = nscache.ErrNil
	}
	if err != nil {
		return err
	}
	return c.serializer.Deserialize(data, &v)
}

func (c *redisClusterCache) Set(k string, v any, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := c.serializer.Serialize(v)
	if err != nil {
		return err
	}
	stat := c.client.Set(context.Background(), k, data, expiration)
	return stat.Err()
}

func (c *redisClusterCache) Remove(k string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.client.Del(context.Background(), k).Err()
}

func (c *redisClusterCache) Close() error {
	return c.client.Close()
}

// parseRedisClusterConnection parse the redis cluster connection string
func parseRedisClusterConnection(u *url.URL) (opt *redis.ClusterOptions, err error) {
	if u == nil {
		return nil, errors.New("invalid redis cluster connection string")
	}
	redisURL := "redis" + strings.TrimPrefix(u.String(), DriverName)
	return redis.ParseClusterURL(redisURL)
}

func init() {
	nscache.Register(DriverName, newCache)
}
