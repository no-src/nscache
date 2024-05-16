package main

import (
	"os"

	"github.com/no-src/log"
	"github.com/no-src/nscache/cli"
	"github.com/no-src/nsgo/fsutil"
)

func main() {
	defer log.Close()

	initLogger("log.yaml")

	if len(os.Args) < 2 {
		log.Error(nil, "please input cache connection string")
		return
	}
	log.ErrorIf(cli.Start(os.Args[1], os.Stdin), "start cli error")
}

func initLogger(logConf string) {
	exist, err := fsutil.FileExist(logConf)
	if err != nil || !exist {
		return
	}
	logger, err := log.CreateLoggerFromConfig(logConf)
	if err == nil {
		log.InitDefaultLogger(logger)
		log.Info("init logger with log config file => %s", logConf)
	}
}
