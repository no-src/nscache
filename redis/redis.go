package redis

import (
	"context"
	"errors"
	"net/url"
	"sync"
	"time"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/extension"
	"github.com/redis/go-redis/v9"
)

const (
	// DriverName the unique name of the Redis driver for register
	DriverName = "redis"
)

type redisCache struct {
	nscache.NSCacheExt

	conn       *url.URL
	serializer encoding.Serializer
	mu         sync.RWMutex
	client     *redis.Client
}

func newCache(conn *url.URL) (nscache.NSCache, error) {
	opt, err := parseRedisConnection(conn)
	if err != nil {
		return nil, err
	}
	c := &redisCache{
		conn:       conn,
		serializer: encoding.DefaultSerializer,
		client:     redis.NewClient(opt),
	}
	c.NSCacheExt = extension.New(c)
	return c, nil
}

func (c *redisCache) Get(k string, v any) error {
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

func (c *redisCache) Set(k string, v any, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := c.serializer.Serialize(v)
	if err != nil {
		return err
	}
	stat := c.client.Set(context.Background(), k, data, expiration)
	return stat.Err()
}

func (c *redisCache) Remove(k string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.client.Del(context.Background(), k).Err()
}

func (c *redisCache) Close() error {
	return c.client.Close()
}

// parseRedisConnection parse the redis connection string
func parseRedisConnection(u *url.URL) (opt *redis.Options, err error) {
	if u == nil {
		return nil, errors.New("invalid redis connection string")
	}
	return redis.ParseURL(u.String())
}

func init() {
	nscache.Register(DriverName, newCache)
}
