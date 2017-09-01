package utils

import (
	"net/http"
	log "github.com/cihub/seelog"
)

func InitWebHandler(conf WebConfig) error {
	addr := conf.Ip + ":" + conf.Port
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Errorf("listenandserve err, err:%v", err)
		return err
	}
}


func reloadLogConfHandler() {

}