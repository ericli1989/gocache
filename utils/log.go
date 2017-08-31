package utils

import (
	log "github.com/cihub/seelog"
	"time"
)

func initLog(path string) {
	logger, err := log.LoggerFromConfigAsFile(path)
	if err != nil {
		log.Errorf("err parse logconf, err:%v", err)
		return err
	}

	log.ReplaceLogger(logger)

	return nil
}
