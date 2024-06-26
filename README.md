# nscache

[![Build](https://img.shields.io/github/actions/workflow/status/no-src/nscache/go.yml?branch=main)](https://github.com/no-src/nscache/actions)
[![License](https://img.shields.io/github/license/no-src/nscache)](https://github.com/no-src/nscache/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/no-src/nscache.svg)](https://pkg.go.dev/github.com/no-src/nscache)
[![Go Report Card](https://goreportcard.com/badge/github.com/no-src/nscache)](https://goreportcard.com/report/github.com/no-src/nscache)
[![codecov](https://codecov.io/gh/no-src/nscache/branch/main/graph/badge.svg?token=ol5hru7WCf)](https://codecov.io/gh/no-src/nscache)
[![Release](https://img.shields.io/github/v/release/no-src/nscache)](https://github.com/no-src/nscache/releases)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

## Installation

```bash
go get -u github.com/no-src/nscache
```

## Quick Start

First, you need to import the cache driver, then create your cache component instance with the specified connection
string and use it.

Current support following cache drivers

| Driver        | Import Driver Package                     | Connection String Example                                                |
|---------------|-------------------------------------------|--------------------------------------------------------------------------|
| Memory        | `github.com/no-src/nscache/memory`        | `memory:`                                                                |
| Redis         | `github.com/no-src/nscache/redis`         | `redis://127.0.0.1:6379`                                                 |
| Redis Cluster | `github.com/no-src/nscache/redis_cluster` | `redis-cluster://127.0.0.1:7001?addr=127.0.0.1:7002&addr=127.0.0.1:7003` |
| BuntDB        | `github.com/no-src/nscache/buntdb`        | `buntdb://:memory:` or `buntdb://buntdb.db`                              |
| Etcd          | `github.com/no-src/nscache/etcd`          | `etcd://127.0.0.1:2379?dial_timeout=5s`                                  |
| BoltDB        | `github.com/no-src/nscache/boltdb`        | `boltdb://boltdb.db`                                                     |
| FreeCache     | `github.com/no-src/nscache/freecache`     | `freecache://?cache_size=50mib`                                          |
| BigCache      | `github.com/no-src/nscache/bigcache`      | `bigcache://?eviction=10m`                                               |
| FastCache     | `github.com/no-src/nscache/fastcache`     | `fastcache://?max_bytes=50mib`                                           |
| Memcached     | `github.com/no-src/nscache/memcached`     | `memcached://127.0.0.1:11211`                                            |
| Proxy         | `github.com/no-src/nscache/proxy/client`  | `proxy://127.0.0.1:8080`                                                 |

For example, initial a memory cache and write, read and remove data.

```go
package main

import (
	"time"

	_ "github.com/no-src/nscache/memory"

	"github.com/no-src/log"
	"github.com/no-src/nscache"
)

func main() {
	// initial cache driver
	c, err := nscache.NewCache("memory:")
	if err != nil {
		log.Error(err, "init cache error")
		return
	}
	defer c.Close()

	// write data
	k := "hello"
	c.Set(k, "world", time.Minute)

	// read data
	var v string
	if err = c.Get(k, &v); err != nil {
		log.Error(err, "get cache error")
		return
	}
	log.Info("key=%s value=%s", k, v)

	// remove data
	if err = c.Remove(k); err != nil {
		log.Error(err, "remove cache error")
		return
	}
}
```

## Command Line Tool

You can use the command line tool to operate all the cache drivers that `nscache` supports.

Install the command line tool `nscache-cli`.

```bash
go install github.com/no-src/nscache/cmd/nscache-cli@latest
```

Run the cli tool with specified cache driver connection string to operate cache.

```bash
nscache-cli memory:
```

## Proxy

### Proxy Server

Install the proxy server `nscache-server`.

```bash
go install github.com/no-src/nscache/cmd/nscache-server@latest
```

Run the proxy server with specified listen address and cache driver connection string.

```bash
nscache-server -addr=:8080 -conn=memory:
```

### Proxy Client

You can use the code to operate cache proxy server like the example code in the [Quick Start](#quick-start) section.

Or use the `nscache-cli` to connect to the proxy server and operate cache.

```bash
nscache-cli proxy://127.0.0.1:8080
```