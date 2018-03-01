package router

import (
	"net/http"
	"fmt"
	"../drive/engine"
	"../config"
	"../db/test"
)

func Get(w http.ResponseWriter, r *http.Request) (map[string]interface{}) {
	con := config.NewConfig()
	m := map[string]interface{}{"Path": con.Ip + con.Port, "Content": "write1"}
	test.FindAll()
	return m
}

func Post(w http.ResponseWriter, r *http.Request) {
	m := engine.ParseData(r)
	fmt.Println(m, m["Test"])
}

