package main

import (
	"log"
	"net/http"

	"github.com/danangkonang/gateway/router"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	// target, err := url.Parse("http://localhost:3000")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// http.Handle("/satu", http.StripPrefix("/", httputil.NewSingleHostReverseProxy(target)))

	router.Router(r)

	log.Fatal(http.ListenAndServe(":7000", r))
}
