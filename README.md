# nscache

## Installation

```bash
go get -u github.com/no-src/nscache
```

## Quick Start

Current support following cache

- memory `memory:`
- redis `redis://127.0.0.1:6379`
- buntdb `buntdb://:memory:` or `buntdb://buntdb.db`
- etcd `etcd://127.0.0.1:2379?dial_timeout=5s`

For example, init a memory cache, write and read data.

```go
package main

import (
	"time"

	"github.com/no-src/log"
	"github.com/no-src/nscache"
	_ "github.com/no-src/nscache/memory"
)

func main() {
	c, err := nscache.NewCache("memory:")
	if err != nil {
		log.Error(err, "init cache error")
		return
	}
	k := "hello"
	c.Set(k, "world", time.Minute)
	var v string
	if err = c.Get(k, &v); err != nil {
		log.Error(err, "get cache error")
		return
	}
	log.Info("key=%s value=%s", k, v)
}
```