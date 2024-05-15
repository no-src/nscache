package cli

import (
	"github.com/no-src/log"
	"github.com/no-src/nscache"
)

func help(_ []string, _ nscache.NSCache) error {
	// help
	log.Info(`the nscache cli commands:
quit                      quit the program
help                      print the help info
get key                   get the cache value with the specified key
set key value [expire]    add or update the cache
del key                   delete the specified key
`)
	return nil
}
