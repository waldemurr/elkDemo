package main

import (
	"math/rand"
	"time"

	"go.elastic.co/ecslogrus"
	// "github.com/caarlos0/env/v11"
	"github.com/sirupsen/logrus"
)

var stdFields = logrus.Fields{
	"service": "Talker",
	"file":    "talker.go",
}

var val int

// type config struct {
// 	Port string `env:"LOG_PORT" envDefault:"5228"`
// 	Host string `env:"LOG_HOST" envDefault:"logstash"`
// }

func main() {
	log := logrus.New()
	log.SetFormatter(&ecslogrus.Formatter{})
	// log.SetFormatter(&logrus.JSONFormatter{})
	// Add custom fields.
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
