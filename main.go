package main

import (
	log "github.com/cihub/seelog"
	"./utils"
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

	Conf, err := utils.ParseConfig(*confFile)
	if err != nil {
		log.Errorf("parseConfig err, confFile:%v, err:%v", confFile, err)
		goto out
	}

	err = initLog(*logFile)
	if err != nil {
		log.Errorf("initLog fail, err:%v", err)
		goto out
	}

	err = utils.InitWebHandler(Conf.Web)
	if err != nil {
		log.Errorf("init webhandler err, err:%v", err)
		goto out
	}





	go func() {
		time.Sleep(time.Second)
		log.Flush()
	}()


	out:
	return
}