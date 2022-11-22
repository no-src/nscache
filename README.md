# nscache

[![Build](https://img.shields.io/github/workflow/status/no-src/nscache/Go)](https://github.com/no-src/nscache/actions)
[![License](https://img.shields.io/github/license/no-src/nscache)](https://github.com/no-src/nscache/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/no-src/nscache.svg)](https://pkg.go.dev/github.com/no-src/nscache)
[![Go Report Card](https://goreportcard.com/badge/github.com/no-src/nscache)](https://goreportcard.com/report/github.com/no-src/nscache)
[![codecov](https://codecov.io/gh/no-src/nscache/branch/main/graph/badge.svg?token=ol5hru7WCf)](https://codecov.io/gh/no-src/nscache)
[![Release](https://img.shields.io/github/v/release/no-src/nscache)](https://github.com/no-src/nscache/releases)

## Installation

```bash
go get -u github.com/no-src/nscache
```

## Quick Start

First, you need to import the cache driver, then create your cache component instance with the specified connection
string and use it.

Current support following cache drivers

| Driver | Import Driver Package              | Connection String Example                   |
|--------|------------------------------------|---------------------------------------------|
| Memory | `github.com/no-src/nscache/memory` | `memory:`                                   |
| Redis  | `github.com/no-src/nscache/redis`  | `redis://127.0.0.1:6379`                    |
| BuntDB | `github.com/no-src/nscache/buntdb` | `buntdb://:memory:` or `buntdb://buntdb.db` |
| Etcd   | `github.com/no-src/nscache/etcd`   | `etcd://127.0.0.1:2379?dial_timeout=5s`     |

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