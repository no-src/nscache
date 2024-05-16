package cli

import (
	"fmt"
	"time"

	"github.com/no-src/log"
	"github.com/no-src/nscache"
)

func set(args []string, cache nscache.NSCache) (err error) {
	// set key value 10s
	if err = checkArgs(args, 3); err != nil {
		return err
	}
	key := args[1]
	value := args[2]
	var expire time.Duration
	if len(args) > 3 {
		expire, err = time.ParseDuration(args[3])
		if err != nil {
			return fmt.Errorf("%w %w parse expiration time error", errInvalidArg, err)
		}
	}
	err = cache.Set(key, value, expire)
	if err != nil {
		return err
	}
	log.Info("set cache success => key=%s", key)
	return nil
}
