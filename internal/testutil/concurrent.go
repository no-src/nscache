package testutil

import (
	"sync"
	"testing"
	"time"

	"github.com/no-src/nscache"
)

func testCacheConcurrent(t *testing.T, conn string, expiration time.Duration) {
	c, err := nscache.NewCache(conn)
	if err != nil {
		t.Errorf("init cache error, err=%v", err)
		return
	}
	defer c.Close()

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(3)
		go func() {
			c.Set(TestKey, TestValue, expiration)
			wg.Done()
		}()
		go func() {
			c.GetString(TestKey)
			wg.Done()
		}()
		go func() {
			c.Remove(TestKey)
			wg.Done()
		}()
	}
	wg.Wait()
}
