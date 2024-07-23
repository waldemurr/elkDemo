package main

import (
	"github.com/caarlos0/env/v11"
	logstash_logger "github.com/KaranJagtiani/go-logstash"
	log "github.com/sirupsen/logrus"
)

var stdFields = log.Fields{
	"service": "Talker",
	"file":    "talker.go",
}
var val int

type config struct {
	Port int `env:"LOG_PORT" envDefault:"5044"`
	Host string `env:"LOG_HOST" envDefault:"logstash"`
}

func main() {
	
	cfg, err := env.ParseAs[config]()
	if err != nil {
		log.Error(err)
	}

	logger := logstash_logger.Init(cfg.Host, 5044, "tcp", 5)

    payload := map[string]interface{}{
        "message": "TEST_MSG",
        "error":   false,
    }

    logger.Log(payload) // Generic log
    logger.Info(payload) // Adds "severity": "INFO"
    logger.Debug(payload) // Adds "severity": "DEBUG"
    logger.Warn(payload) // Adds "severity": "WARN"
    logger.Error(payload) // Adds "severity": "ERROR"
}