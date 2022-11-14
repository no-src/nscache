package nscache

import (
	"net/url"
	"testing"
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
	factory := func(conn *url.URL) (NSCache, error) {
		return nil, nil
	}
	overwritten := Register("repeated_cache_factory", factory)
	if overwritten {
		t.Errorf("register cache factory 'repeated_cache_factory' once, expect get overwritten false but get true")
		return
	}
	overwritten = Register("repeated_cache_factory", factory)
	if !overwritten {
		t.Errorf("register cache factory 'repeated_cache_factory' twice, expect get overwritten true but get false")
	}
}
