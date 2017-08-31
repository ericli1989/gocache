package utils

import (
	"net/http"
	log "github.com/cihub/seelog"
	"jd.com/tp/tp-storage/config"
)

func initWebHandler(config.Web) {
	err := http.ListenAndServe()
	if err != nil {
		log.Errorf("listenandserve err, err:%v", err)
		return err
	}
}


func reloadLogConfHandler() {

}