package route

import (
	"encoding/json"
	"net/http"
)

type Router struct {
	routes      []Route
	middlewares []func(http.HandlerFunc) http.HandlerFunc
}

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc func(http.ResponseWriter, *http.Request)
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Use(middleware func(http.HandlerFunc) http.HandlerFunc) {
	r.middlewares = append(r.middlewares, middleware)
}

func (r *Router) AddRoute(method, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	r.routes = append(r.routes, Route{Method: method, Pattern: pattern, HandlerFunc: handler})
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.Pattern == req.URL.Path && route.Method == req.Method {
			handler := route.HandlerFunc
			for _, mw := range r.middlewares {
				handler = mw(handler)
			}
			handler(w, req)
			return
		}
	}
	http.NotFound(w, req)
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}
