package nscache

import (
	"net/url"
	"sync/atomic"
	"testing"
	"time"
)

func TestRegister_WithNilCacheFactory(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("register a nil cache factory expect to panic but not")
		}
	}()
	Register("nil_cache_factory", nil)
}

func TestRegister_WithRepeatedCacheFactory(t *testing.T) {
	overwritten := Register("repeated_cache_factory", mockFactory)
	if overwritten {
		t.Errorf("register cache factory 'repeated_cache_factory' once, expect get overwritten false but get true")
		return
	}
	overwritten = Register("repeated_cache_factory", mockFactory)
	if !overwritten {
		t.Errorf("register cache factory 'repeated_cache_factory' twice, expect get overwritten true but get false")
	}
}

func TestRegister_WithConcurrent(t *testing.T) {
	var stop atomic.Bool
	go func() {
		for !stop.Load() {
			Register("concurrent_cache_factory", mockFactory)
		}
	}()

	go func() {
		for !stop.Load() {
			Register("concurrent_cache_factory", mockFactory)
		}
	}()

	<-time.After(time.Second * 3)
	stop.Store(true)
}

var mockFactory = func(conn *url.URL) (NSCache, error) {
	return nil, nil
}
