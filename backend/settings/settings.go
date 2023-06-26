package settings

import (
	"encoding/json"
	"fmt"
	"os"
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

func Init() error {
	f, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&Cfg)
	if err != nil {
		fmt.Println(err)
	}

	// 正常終了すれば、nilを返す
	return nil
}
