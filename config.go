package main

import (
	"github.com/kelseyhightower/envconfig"
)


type appConfig struct {
	Port string `required:"true"`
	Clients map[string]int `required:"true"`
}

func loadConfig()(appConfig,error){
	var c appConfig
	err := envconfig.Process("app", &c)
	if err != nil{
		return appConfig{}, err
	}
	return c, nil
}

