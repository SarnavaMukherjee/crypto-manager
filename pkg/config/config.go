/**
 * File: config.go
 * Author: Sarnava Mukherjee
 * Contact: (sarnavamukherjee20@gmail.com)
 */

package config

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	SERVER_PORT    int    `default:"9090"`
	SERVER_MODE    string `default:"debug"`
	LOG_LEVEL      string `default:"info"`
	MONGO_USER     string `default:"mongoadmin"`
	MONGO_PASSWORD string `default:"secret"`
	MONGO_HOST     string `default:"127.0.0.1"`
}

var dial sync.Once
var config Config

func newConfig(app string) {
	err := envconfig.Process(app, &config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("[PMS-CONFIG] created config object")
}

func CreateConfig() {
	dial.Do(func() {
		newConfig("pms")
	})
}

func GetConfig() Config {
	return config
}
