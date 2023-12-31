package app

import (
	"autossh/utils"
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

// 加载配置
func loadConfig(configFile string) (cfg *Config, err error) {
	configFile, err = utils.ParsePath(configFile)
	if err != nil {
		return cfg, err
	}

	if exists, _ := utils.FileIsExists(configFile); !exists {
		return cfg, errors.New("Can't read configFile file:" + configFile)
	}

	b, _ := ioutil.ReadFile(configFile)
	err = json.Unmarshal(b, &cfg)
	if err != nil {
		return cfg, err
	}

	cfg.file = configFile
	cfg.createServerIndex()

	return cfg, nil
}
