package router

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/danangkonang/gateway/middleware"
	"github.com/danangkonang/gateway/proxyserve"
	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	target, err := url.Parse("http://localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(target)
	// v1 := router.PathPrefix("").Subrouter()
	router.HandleFunc("/users", middleware.Auth(proxyserve.ProxyServe(proxy, target))).Methods("GET")
	router.HandleFunc("/products", middleware.Auth(proxyserve.ProxyServe(proxy, target))).Methods("POST")
	router.HandleFunc("/cart", middleware.Auth(proxyserve.ProxyServe(proxy, target))).Methods("DELETE")
}
