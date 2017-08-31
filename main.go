package main

import (
	log "github.com/cihub/seelog"
	"flag"
	"time"
)



func initLog(path string) error {
	logger, err := log.LoggerFromConfigAsFile(path)
	if err != nil {
		log.Errorf("err parse logfile, err:%v", err)
		return err
	}
	log.ReplaceLogger(logger)
	return nil
}

func main() {

	confFile := flag.String("config", "config.json", "config file")
	logFile := flag.String("log", "seelog.xml", "seelog config file")

	flag.Parse()

	Conf := config.ParseConfig(*confFile)

	err := initLog(*logFile)
	if err != nil {
		log.Errorf("initLog fail, err:%v", err)
		goto out
	}





	go func() {
		time.Sleep(time.Second)
		log.Flush()
	}()


	out:
	return
}