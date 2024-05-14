package all

import (
	_ "github.com/no-src/nscache/bigcache"
	_ "github.com/no-src/nscache/buntdb"
	_ "github.com/no-src/nscache/etcd"
	_ "github.com/no-src/nscache/fastcache"
	_ "github.com/no-src/nscache/freecache"
	_ "github.com/no-src/nscache/memory"
	_ "github.com/no-src/nscache/redis"
	_ "github.com/no-src/nscache/redis_cluster"
)
