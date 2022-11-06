package app

import (
	"net/http"

	"github.com/danangkonang/product/config"
	"github.com/danangkonang/product/helper"
	"github.com/danangkonang/product/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Run() {
	r := mux.NewRouter().StrictSlash(true)

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.MakeRespon(w, 404, "page not found", nil)
	})

	r.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.MakeRespon(w, http.StatusMethodNotAllowed, "Method NotAllowed", nil)
	})

	router.ProductRouter(r, config.Connection())
	router.CartRouter(r, config.Connection())

	header := []string{
		"X-Requested-With",
		"Access-Control-Allow-Origin",
		"Content-Type",
		"Authorization",
		"Accept",
	}
	method := []string{"GET", "POST", "PUT", "DELETE"}
	origin := []string{"*"}
	http.ListenAndServe(":9001", handlers.CORS(
		handlers.AllowedHeaders(header),
		handlers.AllowedMethods(method),
		handlers.AllowedOrigins(origin),
	)(r))
}
