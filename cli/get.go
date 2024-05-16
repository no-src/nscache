package cli

import (
	"github.com/no-src/log"
	"github.com/no-src/nscache"
)

func get(args []string, cache nscache.NSCache) (err error) {
	// get key
	if err = checkArgs(args, 2); err != nil {
		return err
	}
	var v string
	key := args[1]
	err = cache.Get(key, &v)
	if err != nil {
		return err
	}
	log.Info("get cache success => \nkey=%s \nvalue=%s", key, v)
	return nil
}
