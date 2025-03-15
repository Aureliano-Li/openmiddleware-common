package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func GetLogger() log.Logger {
	var logger = log.New()
	logger.Hooks.Add(NewContextHook())
	logger.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(log.DebugLevel)
	return *logger
}
