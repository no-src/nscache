# nscache

[![Chat](https://img.shields.io/discord/936876326722363472)](https://discord.gg/AycM3RpMjw)
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