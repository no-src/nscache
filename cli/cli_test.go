package cli

import (
	"errors"
	"strings"
	"testing"

	"github.com/no-src/nscache"
)

func TestHelp(t *testing.T) {
	startCli(t, "help")
}

func TestGet_Nil(t *testing.T) {
	startCli(t, "get hello")
}

func TestSet(t *testing.T) {
	startCli(t, "set hello world\nget hello")
}

func TestSet_WithExpiration(t *testing.T) {
	startCli(t, "set hello world 10s\nget hello")
}

func TestDelete(t *testing.T) {
	startCli(t, "set hello world\nget hello\ndel hello\nget hello")
}

func TestBlankLine(t *testing.T) {
	startCli(t, "\n\n")
}

func TestUnsupportedCommand(t *testing.T) {
	startCli(t, "go")
}

func startCli(t *testing.T, input string) {
	err := Start("memory:", strings.NewReader(input))
	if err != nil {
		t.Fatalf("start cli error => %s", err)
	}
}

func getCache() nscache.NSCache {
	cache, _ := nscache.NewCache("memory:")
	return cache
}

func TestGet_InvalidArgs(t *testing.T) {
	err := get([]string{"get"}, getCache())
	if !errors.Is(err, errNotEnoughArgs) {
		t.Fatalf("expect to get error %s, but actual get %s", errNotEnoughArgs, err)
	}
}

func TestSet_InvalidArgs(t *testing.T) {
	err := set([]string{"set", "hello"}, getCache())
	if !errors.Is(err, errNotEnoughArgs) {
		t.Fatalf("expect to get error %s, but actual get %s", errNotEnoughArgs, err)
	}
}

func TestDelete_InvalidArgs(t *testing.T) {
	err := delete([]string{"delete"}, getCache())
	if !errors.Is(err, errNotEnoughArgs) {
		t.Fatalf("expect to get error %s, but actual get %s", errNotEnoughArgs, err)
	}
}

func TestSet_InvalidExpiration(t *testing.T) {
	err := set([]string{"set", "hello", "world", "1x"}, getCache())
	if !errors.Is(err, errInvalidArg) {
		t.Fatalf("expect to get error %s, but actual get %s", errInvalidArg, err)
	}
}
