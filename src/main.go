package main

import (
	log "github.com/cihub/seelog"
	"time"
)

func main() {





	go func() {
		select {
			case <-time.After(time.Second):
				log.Flush()
		}
		return
	}()

	defer log.Flush()
    	log.Info("Hello from Seelog!")



}