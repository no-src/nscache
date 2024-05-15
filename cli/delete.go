package cli

import (
	"github.com/no-src/log"
	"github.com/no-src/nscache"
)

func delete(args []string, cache nscache.NSCache) (err error) {
	// del key
	if err = checkArgs(args, 2); err != nil {
		return err
	}
	key := args[1]
	err = cache.Remove(key)
	if err != nil {
		return err
	}
	log.Info("remove cache success => key=%s", key)
	return nil
}
