package redis

import (
	"context"
	"errors"
	"net/url"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
)

const (
	driverName = "redis"
)

type redisCache struct {
	conn       *url.URL
	mu         sync.Mutex
	client     *redis.Client
	serializer encoding.Serializer
}

func newCache(conn *url.URL) (nscache.NSCache, error) {
	opt, err := parseRedisConnection(conn)
	if err != nil {
		return nil, err
	}
	c := &redisCache{
		conn:       conn,
		client:     redis.NewClient(opt),
		serializer: encoding.DefaultSerializer,
	}
	return c, nil
}

func (c *redisCache) Get(k string, v any) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := c.client.Get(context.Background(), k).Bytes()
	if err != nil && err != redis.Nil {
		return err
	} else if err == redis.Nil {
		return nil
	}
	return c.serializer.Deserialize(data, &v)
}

func (c *redisCache) GetString(k string) (s string, ok bool) {
	var v any
	err := c.Get(k, &v)
	if err != nil || v == nil {
		return "", false
	}
	s, ok = v.(string)
	return s, ok
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

// parseRedisConnection parse the redis connection string
func parseRedisConnection(u *url.URL) (opt *redis.Options, err error) {
	if u == nil {
		return nil, errors.New("invalid redis url")
	}
	return redis.ParseURL(u.String())
}

func init() {
	nscache.Register(driverName, newCache)
}
