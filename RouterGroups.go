package mngo

import "net/http"

type RouterGroups struct {
	path string
	method string
	fn func(w http.ResponseWriter, r *http.Request)
}

func newRouterGroups() *RouterGroups {
	return &RouterGroups{}
}
