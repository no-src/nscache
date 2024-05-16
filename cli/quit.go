package cli

import (
	"os"

	"github.com/no-src/nscache"
)

func quit(args []string, cache nscache.NSCache) error {
	// quit
	os.Exit(0)
	return nil
}
