package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Db struct {
		Host           string `json:"host"`
		Port           int    `json:"port"`
		Password       string `json:"password"`
		User           string `json:"user"`
		Table          string `json:"table"`
		Database       string `json:"database"`
		PackageName    string `json:"packageName"`
		JSONAnnotation bool   `json:"jsonAnnotation"`
		DBAnnotation   bool   `json:"dbAnnotation"`
		XMLAnnotation  bool   `json:"xmlAnnotation"`
		GormAnnotation bool   `json:"gormAnnotation"`
	} `json:"db"`
	GenPath string `json:"genPath"`
}

func GetConf(configPath string) (*Config, error) {
	conf := Config{}

	jsonFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		return &conf, err
	}

	err = json.Unmarshal(jsonFile, &conf)

	if err != nil {
		fmt.Println(err)
		return &conf, err
	}

	return &conf, nil
}
