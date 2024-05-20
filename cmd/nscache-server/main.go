package main

import (
	"flag"
	"time"

	"github.com/no-src/log"
	"github.com/no-src/nscache/proxy/server"
)

var (
	addr    string
	conn    string
	logConf string
)

func main() {
	defer log.Close()

	flag.StringVar(&addr, "addr", ":8080", "the nscache server listen address")
	flag.StringVar(&conn, "conn", "memory:", "the nscache connection string")
	flag.StringVar(&logConf, "log_conf", "", "specified a log config file")
	flag.Parse()

	initLogger()

	go func() {
		time.Sleep(time.Second)
		log.Info("start nscache server, listen address: %s, connection: %s", addr, conn)
	}()

	log.ErrorIf(server.Start(addr, conn), "start nscache server error")
}

func initLogger() {
	if len(logConf) > 0 {
		logger, err := log.CreateLoggerFromConfig(logConf)
		if err != nil {
			log.Error(err, "init logger with config failed, fallback to use the default logger")
		} else {
			log.InitDefaultLogger(logger)
		}
	}
}
