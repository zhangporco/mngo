package engine

import "net/http"

type RouterGroups struct {
	path string
	method string
	fn func(w http.ResponseWriter, r *http.Request)
}

func newRouterGroups() *RouterGroups {
	rg := &RouterGroups{}
	return rg
}
