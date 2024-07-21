package main

import (
	// "math/rand"
	// "time"

	"github.com/caarlos0/env/v11"
	logstash_logger "github.com/KaranJagtiani/go-logstash"
	log "github.com/sirupsen/logrus"
)

var stdFields = log.Fields{
	"service": "Talker",
	"file":    "historysync.go",
}
// /usr/share/metricbeat/metricbeat.yml
var val int

type config struct {
	Port int `env:"LOG_PORT" envDefault:"5228"`
	Host string `env:"LOG_HOST" envDefault:"logstash"`
}

func main() {
	// log.SetFormatter(&log.JSONFormatter{})
	// logger := log.WithFields(stdFields).WithFields(log.Fields{"function": "talker"})

	cfg, err := env.ParseAs[config]()
	if err != nil {
		log.Error(err)
	}

	logger := logstash_logger.Init("logstash", cfg.Port, "tcp", 5)

    payload := map[string]interface{}{
        "message": "TEST_MSG",
        "error":   false,
    }

    logger.Log(payload) // Generic log
    logger.Info(payload) // Adds "severity": "INFO"
    logger.Debug(payload) // Adds "severity": "DEBUG"
    logger.Warn(payload) // Adds "severity": "WARN"
    logger.Error(payload) // Adds "severity": "ERROR"

	// hook := graylog.NewGraylogHook(cfg.Host+":"+cfg.Port, map[string]interface{}{"this": "is logged every time"})
	// log.AddHook(hook)

	// for {
	// 	val = rand.Int()
	// 	switch val % 5 {
	// 	case 0:
	// 		logger.Info("Setup redis orders...")
	// 	case 1:
	// 		logger.Debug("Unblock all users done!")
	// 	case 2:
	// 		logger.Warn("Block all users done!")
	// 	case 3:
	// 		logger.Error("done")
	// 	default:
	// 		logger.Trace("select me before i go")
	// 		time.Sleep(time.Second)
	// 	}
	// }
}
