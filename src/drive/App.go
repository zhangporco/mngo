package drive

import (
	"./engine"
	"net/http"
	"../router"
	"../config"
)

func App() {
	con := config.NewConfig()
	app := engine.NewEngine()
	app.GET("/get", func(w http.ResponseWriter, r *http.Request) {
		res := router.Get(w, r)
		engine.Write(w, res)
	})
	app.POST("/post", func(w http.ResponseWriter, r *http.Request) {
		router.Post(w, r)})
	app.Run(con.Port)
}
