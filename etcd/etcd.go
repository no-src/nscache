package etcd

import (
	"context"
	"errors"
	"net/url"
	"sync"
	"time"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/extension"
	"github.com/no-src/nscache/internal/unit"
	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	// DriverName the unique name of the Etcd driver for register
	DriverName = "etcd"
)

type etcdCache struct {
	nscache.NSCacheExt

	conn       *url.URL
	serializer encoding.Serializer
	mu         sync.RWMutex
	client     *clientv3.Client
}

func newCache(conn *url.URL) (nscache.NSCache, error) {
	conf, err := parseEtcdConnection(conn)
	if err != nil {
		return nil, err
	}
	client, err := clientv3.New(conf)
	if err != nil {
		return nil, err
	}
	c := &etcdCache{
		conn:       conn,
		serializer: encoding.DefaultSerializer,
		client:     client,
	}
	c.NSCacheExt = extension.New(c)
	return c, nil
}

func (c *etcdCache) Get(k string, v any) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	resp, err := c.client.Get(context.Background(), k)
	if err == nil {
		if resp.Count == 0 {
			return nscache.ErrNil
		}
		data := resp.Kvs[0].Value
		err = c.serializer.Deserialize(data, &v)
	}
	return err
}

func (c *etcdCache) Set(k string, v any, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := c.serializer.Serialize(v)
	if err != nil {
		return err
	}
	lease := clientv3.NewLease(c.client)
	ctx := context.Background()
	lgr, err := lease.Grant(ctx, int64(expiration.Seconds()))
	if err == nil {
		_, err = c.client.Put(ctx, k, string(data), clientv3.WithLease(lgr.ID))
	}
	return err
}

func (c *etcdCache) Remove(k string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, err := c.client.Delete(context.Background(), k)
	return err
}

func (c *etcdCache) Close() error {
	return c.client.Close()
}

// parseEtcdConnection parse the etcd connection string
func parseEtcdConnection(u *url.URL) (c clientv3.Config, err error) {
	if u == nil {
		return c, errors.New("invalid etcd connection string")
	}
	c.Endpoints = []string{u.Host}

	dialTimeoutValue := u.Query().Get("dial_timeout")
	if len(dialTimeoutValue) > 0 {
		dialTimeout, err := time.ParseDuration(dialTimeoutValue)
		if err != nil {
			return c, err
		}
		c.DialTimeout = dialTimeout
	}
	maxCallSendMsgSizeValue := u.Query().Get("max_call_send_msg_size")
	if len(maxCallSendMsgSizeValue) > 0 {
		maxCallSendMsgSize, err := unit.ParseBytes(maxCallSendMsgSizeValue)
		if err != nil {
			return c, err
		}
		c.MaxCallSendMsgSize = maxCallSendMsgSize
	}
	maxCallRecvMsgSizeValue := u.Query().Get("max_call_recv_msg_size")
	if len(maxCallRecvMsgSizeValue) > 0 {
		maxCallRecvMsgSize, err := unit.ParseBytes(maxCallRecvMsgSizeValue)
		if err != nil {
			return c, err
		}
		c.MaxCallRecvMsgSize = maxCallRecvMsgSize
	}
	return c, nil
}

func init() {
	nscache.Register(DriverName, newCache)
}
