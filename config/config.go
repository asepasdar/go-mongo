package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//Configuration : configuration object
var Configuration *_configuration

type _configuration struct {
	Server struct {
		Port         int           `json:"port"`
		ReadTimeout  time.Duration `json:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout"`
	} `json:"server"`

	Database struct {
		URI  string `json:"uri"`
		Name string `json:"name"`
	} `json:"database"`
}

func init() {
	if Configuration != nil {
		return
	}

	basePath, err := os.Getwd()
	if err != nil {
		panic(err)
		return
	}

	bts, err := ioutil.ReadFile(filepath.Join(basePath, "config", "config.json"))
	if err != nil {
		panic(err)
		return
	}

	Configuration = new(_configuration)
	err = json.Unmarshal(bts, &Configuration)
	if err != nil {
		panic(err)
		return
	}
}
