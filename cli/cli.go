package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"slices"
	"strings"

	"github.com/no-src/log"
	"github.com/no-src/nscache"
	_ "github.com/no-src/nscache/all"
)

var (
	errNotEnoughArgs = errors.New("not enough arguments")
	errInvalidArg    = errors.New("invalid argument")
)

// Start start the nscache cli with the connection string and input reader
func Start(conn string, r io.Reader) error {
	cache, err := nscache.NewCache(conn)
	if err != nil {
		return fmt.Errorf("%w connect to cache error", err)
	}
	defer cache.Close()
	log.Info("connect to cache success => %s", conn)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Split(line, " ")
		args = slices.DeleteFunc(args, func(s string) bool {
			return len(s) == 0
		})
		if len(args) == 0 {
			continue
		}
		cmdName := args[0]
		cmd := getCommand(cmdName)
		if cmd == nil {
			log.Error(nil, "invalid command =>%s, you can input the help to list all the supported commands", cmdName)
			continue
		}
		err = cmd.run(args, cache)
		if err != nil {
			log.Error(err, "execute [%s] command error", cmdName)
		}
	}
	return nil
}
