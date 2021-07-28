package main

import (
	"net/http"
	"regexp"
)

func (s *server) routes() {
	s.r = &router{}

	s.r.get("/", s.handleHome())
}

type router struct {
	routes []route
}

type route struct {
	method  string
	pattern string
	handler http.HandlerFunc
}

func (r *router) get(path string, h http.HandlerFunc) {
	r.routes = append(r.routes, route{
		pattern: path,
		method:  "GET",
		handler: h,
	})
}

func (r *router) getHandler(method, path string) http.HandlerFunc {
	for _, route := range r.routes {
		re := regexp.MustCompile(route.pattern)
		if route.method == method && re.MatchString(path) {
			return route.handler
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var (
		path    = req.URL.Path
		method  = req.Method
		handler = r.getHandler(method, path)
	)
	handler.ServeHTTP(w, req)
}
