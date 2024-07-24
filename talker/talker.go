package main

import (
	"math/rand"
	"net"
	"time"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"go.elastic.co/ecslogrus"

	"github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

var stdFields = logrus.Fields{
	"service": "Talker",
	"file":    "talker.go",
}
var val int

type config struct {
	Port string `env:"LOG_PORT" envDefault:"5000"`
	Host string `env:"LOG_HOST" envDefault:"logstash"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		logrus.Error(err)
	}

	log := logrus.New()
	log.SetFormatter(&ecslogrus.Formatter{})
	conn, err := net.Dial("tcp", cfg.Host + ":" + cfg.Port)
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": "myappName"}))

	log.Hooks.Add(hook)
	logger := log.WithFields(stdFields).WithFields(logrus.Fields{"function": "main"})

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

		}
		logger.Info("val is ", val)
		time.Sleep(time.Second)
	}
}
