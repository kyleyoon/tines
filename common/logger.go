package common

import (
    "os"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	logLevel, err := log.ParseLevel(os.Getenv("TF_LOG"))
	if err != nil {
		log.Errorf("Not Fount LogLevel: %s", err)
		log.SetLevel(log.InfoLevel)
		log.Println("Set LogLevel: info")
	} else {
		log.SetLevel(logLevel)
		log.Printf("Set LogLevel: %s", log.GetLevel())
	}
}