package utils

import (
	"io/ioutil"
	log "github.com/cihub/seelog"
	"encoding/json"
)

func ParseConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Errorf("read file err, path:%v, err:%v", path, err)
		return nil, err
	}
	config := &Config{}
	err = json.Unmarshal(data, config)
	if err != nil {
		log.Errorf("unmarshal err, err:%v", err)
		return nil, err
	}

	GoCacheConfig = config
	return GoCacheConfig, nil
}
