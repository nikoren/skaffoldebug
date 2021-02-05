package main

import (
	"github.com/kelseyhightower/envconfig"
)


type appConfig struct {
	port string `required:"true"`
	clients map[string]int
}

func loadConfig()(appConfig,error){
	var c appConfig
	err := envconfig.Process("app", &c)
	if err != nil{
		return appConfig{}, err
	}
	return c, nil
}

