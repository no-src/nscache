package etcd

import (
	"net/url"
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = testutil.EtcdConnectionString
	expiration       = testutil.DefaultExpiration
)

func TestEtcdCache(t *testing.T) {
	testutil.TestCache(t, connectionString, expiration)
}

func TestEtcdCache_NewCache_WithNilURL(t *testing.T) {
	_, err := newCache(nil)
	if err == nil {
		t.Errorf("expect get an error, but get nil")
	}
}

func TestEtcdCache_NewCache_WithInvalidURL(t *testing.T) {
	testCases := []struct {
		conn string
	}{
		{"etcd://127.0.0.1:2379?dial_timeout=5z"},
		{"etcd://127.0.0.1:2379?dial_timeout=5s&max_call_send_msg_size=z"},
		{"etcd://127.0.0.1:2379?dial_timeout=5s&max_call_recv_msg_size=z"},
		{"etcd://127.0.0.1:2379?dial_timeout=5s&max_call_send_msg_size=2&max_call_recv_msg_size=1"},
	}
	for _, tc := range testCases {
		t.Run(tc.conn, func(t *testing.T) {
			u, err := url.Parse(tc.conn)
			if err != nil {
				t.Errorf("invalid url => %s", tc.conn)
				return
			}
			_, err = newCache(u)
			if err == nil {
				t.Errorf("expect get an error, but get nil")
			}
		})
	}
}
