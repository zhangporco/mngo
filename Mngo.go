package mngo

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"flag"
	"sync"
)

var engine *Mngo
var once sync.Once

/**
	mngo 启动引擎
	@author Porco
 */
type Mngo struct {
	log bool
	debug bool
	Env string
	port string
	ip string
	routerGroup []RouterGroups
}

func NewMngo() *Mngo {
	once.Do(func() {
		engine = &Mngo{
			log : true,
			debug : false,
			//Env: getEnv(),
		}
	})
	return engine
}

func (engine *Mngo) Run(p string) {
	engine.port = p
	engine.logger(p)
	for _, v := range engine.routerGroup {
		http.HandleFunc(v.path, v.fn)
	}
	http.ListenAndServe(":" + p, nil)
}

/**
此函数可以从 go run shell 中获取 env 参数
@example: go run ./App.go --env="develop"
 */
func getEnv() string {
	env := flag.String("env", "develop", "server env")
	flag.Parse()
	return *env
}

func (engine *Mngo) SetLog(log bool) {
	engine.log = log
}

func (engine *Mngo) GET(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	engine.response(path, "GET", fn)
}

func (engine *Mngo) POST(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	engine.response(path, "POST", fn)
}

func (engine *Mngo) response(path string, method string, fn func(w http.ResponseWriter, r *http.Request)) {
	rg := newRouterGroups()
	rg.path = path
	rg.method = method

	rg.fn = func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			fn(w, r)
		} else {
			res := map[string]interface{}{"Host": r.Host, "Content": "it's not a " + method + " request"}
			Write(w, res)
		}
	}
	engine.routerGroup = append(engine.routerGroup, *rg)
}

func Write(w http.ResponseWriter, content interface{}) {
	js, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (engine *Mngo) logger(port string) {
	log.SetPrefix("【Mngo engine】")
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	if engine.log {
		log.Println("Mngo starting")
		log.Println("Mngo port --", port)
	}
}

/**
	解析参数
 */
func ParseData(r *http.Request) interface{} {
	result, _:= ioutil.ReadAll(r.Body)
	r.Body.Close()
	var f interface{}
	json.Unmarshal(result, &f)
	if f == nil {
		return f
	}
	m := f.(map[string]interface{})
	return m
}

//func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
//	var validPath = regexp.MustCompile(helpers.UrlRegexp)
//	return func(w http.ResponseWriter, r *http.Request) {
//		m := validPath.FindStringSubmatch(r.URL.Path)
//		if m == nil {
//			http.NotFound(w, r)
//			return
//		}
//		fn(w, r, m[2])
//	}
//}
