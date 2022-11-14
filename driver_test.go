package nscache

import (
	"errors"
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
		} else if !errors.Is(err.(error), errCacheDriverFactoryIsNil) {
			t.Errorf("expect to get error => %v, but get error => %v", errCacheDriverFactoryIsNil, err)
		}
	}()
	Register("nil-cache", nil)
}

func TestRegister_WithInvalidDriverName(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("register a cache factory with an invalid name that expects to panic but not")
		} else if !errors.Is(err.(error), errInvalidCacheDriverName) {
			t.Errorf("expect to get error => %v, but get error => %v", errInvalidCacheDriverName, err)
		}
	}()
	Register("invalid_cache:", mockFactory)
}

func TestRegister_WithRepeatedCacheFactory(t *testing.T) {
	driverName := "repeated-cache"
	overwritten := Register(driverName, mockFactory)
	if overwritten {
		t.Errorf("register cache factory '%s' once, expect get overwritten false but get true", driverName)
		return
	}
	overwritten = Register(driverName, mockFactory)
	if !overwritten {
		t.Errorf("register cache factory '%s' twice, expect get overwritten true but get false", driverName)
	}
}

func TestRegister_WithConcurrent(t *testing.T) {
	driverName := "concurrent-cache"
	var stop int32
	go func() {
		for atomic.LoadInt32(&stop) == 0 {
			Register(driverName, mockFactory)
		}
	}()

	go func() {
		for atomic.LoadInt32(&stop) == 0 {
			Register(driverName, mockFactory)
		}
	}()

	<-time.After(time.Second * 3)
	atomic.StoreInt32(&stop, 1)
}

var mockFactory = func(conn *url.URL) (NSCache, error) {
	return nil, nil
}
