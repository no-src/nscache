package cli

import (
	"fmt"

	"github.com/no-src/nscache"
)

var commands = map[string]*command{
	"help": newCommand(help),
	"quit": newCommand(quit),
	"set":  newCommand(set),
	"get":  newCommand(get),
	"del":  newCommand(delete),
}

type command struct {
	fn func(args []string, cache nscache.NSCache) error
}

func (c *command) run(args []string, cache nscache.NSCache) error {
	if c.fn == nil {
		return nil
	}
	return c.fn(args, cache)
}

func newCommand(fn func(args []string, cache nscache.NSCache) error) *command {
	return &command{
		fn: fn,
	}
}

func checkArgs(args []string, expectArgLen int) error {
	argLen := len(args)
	if argLen < expectArgLen {
		return fmt.Errorf("%w got %d but expectd %d", errNotEnoughArgs, argLen, expectArgLen)
	}
	return nil
}
