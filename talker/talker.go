package main

import (
	"math/rand"
	"time"

	"github.com/caarlos0/env/v11"
	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	log "github.com/sirupsen/logrus"
)

var stdFields = log.Fields{
	"service": "Talker",
	"file":    "historysync.go",
}
var val int

type config struct {
	Port string `env:"LOG_PORT" envDefault:"9000"`
	Host string `env:"LOG_HOST" envDefault:"graylog"`
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	logger := log.WithFields(stdFields).WithFields(log.Fields{"function": "talker"})

	cfg, err := env.ParseAs[config]()
	if err != nil {
		log.Error(err)
	}
	hook := graylog.NewGraylogHook(cfg.Host+":"+cfg.Port, map[string]interface{}{"this": "is logged every time"})
	log.AddHook(hook)

	for {
		val = rand.Int()
		switch val % 5 {
		case 0:
			logger.Info("Setup redis orders...")
		case 1:
			logger.Debug("Unblock all users done!")
		case 2:
			logger.Warn("Block all users done!")
		case 3:
			logger.Error("done")
		default:
			logger.Trace("select me before i go")
			time.Sleep(time.Second)
		}
	}
}
