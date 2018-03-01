package config

import (
	"../drive/engine"
)


type Config struct {
	Ip string
	Port string
	Mongo Mongo
}

type Mongo struct {
	Url string
	DbName string
}

func NewConfig() *Config {
	app := engine.NewEngine()
	con := Config{}
	if app.Env == "production" {
		con.Ip = "127.0.0.1"
		con.Port = "8081"
		con.Mongo.Url = "mongodb://misg-dw:xxx@127.0.0.1:27017/mini-inquiry"
		con.Mongo.DbName = "mini-inquiry"
	} else {
		con.Ip = "127.0.0.1"
		con.Port = "8080"
		con.Mongo.Url = "mongodb://misg-dw:xxx@127.0.0.1:27017/mini-inquiry"
		con.Mongo.DbName = "mini-inquiry"
	}
	return &con
}
