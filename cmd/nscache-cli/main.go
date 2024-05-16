package main

import (
	"os"

	"github.com/no-src/log"
	"github.com/no-src/nscache/cli"
)

func main() {
	defer log.Close()
	if len(os.Args) < 2 {
		log.Error(nil, "please input cache connection string")
		return
	}
	log.ErrorIf(cli.Start(os.Args[1], os.Stdin), "start cli error")
}
