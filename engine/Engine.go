package engine

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"flag"
	"sync"
)

var engine *Engine
var once sync.Once

type Engine struct {
	log bool
	debug bool
	Env string
	port string
	ip string
	routerGroup []RouterGroups
}

func NewEngine() *Engine {
	once.Do(func() {
		engine = &Engine{
			log : true,
			debug : false,
			Env: getEnv(),
		}
	})
	return engine
}

func (engine *Engine) Run(p string) {
	engine.port = p
	engine.logger(p)
	for _, v := range engine.routerGroup {
		http.HandleFunc(v.path, v.fn)
	}
	http.ListenAndServe(":" + p, nil)
}

func getEnv() string {
	env := flag.String("env", "develop", "server env")
	flag.Parse()
	return *env
}

func (engine *Engine) SetLog(log bool) {
	engine.log = log
}

func (engine *Engine) GET(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	engine.response(path, "GET", fn)
}

func (engine *Engine) POST(path string, fn func(w http.ResponseWriter, r *http.Request)) {
	engine.response(path, "POST", fn)
}

func (engine *Engine) response(path string, method string, fn func(w http.ResponseWriter, r *http.Request)) {
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

func Write(w http.ResponseWriter, content map[string]interface{}) {
	js, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (engine *Engine) logger(port string) {
	log.SetPrefix("【Engine】")
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
	if engine.log {
		log.Println("Engine starting")
		log.Println("Engine port --", port)
	}
}

func ParseData(r *http.Request) map[string]interface{} {
	result, _:= ioutil.ReadAll(r.Body)
	r.Body.Close()
	var f interface{}
	json.Unmarshal(result, &f)
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
