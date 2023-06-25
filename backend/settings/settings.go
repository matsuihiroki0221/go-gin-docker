package settings

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	DB DBConfig `json:"db"`
}

type DBConfig struct {
	Dsn      string `json:"dsn"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	HostName string `json:"hostName"`
	DBName   string `json:"dbName"`
	Port     int    `json:"port"`
}

var Cfg *Config

func Init(configByte []byte) {
	fmt.Println(string(configByte))

	Cfg = &Config{}
	err := json.Unmarshal(configByte, Cfg)
	if err != nil {
		fmt.Println(err)
	}
}
