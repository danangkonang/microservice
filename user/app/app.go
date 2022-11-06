package app

import (
	"net/http"

	"github.com/danangkonang/user/config"
	"github.com/danangkonang/user/helper"
	"github.com/danangkonang/user/router"
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

	router.ProfileRouter(r, config.Connection())

	header := []string{
		"X-Requested-With",
		"Access-Control-Allow-Origin",
		"Content-Type",
		"Authorization",
		"Accept",
	}
	method := []string{"GET", "POST", "PUT", "DELETE"}
	origin := []string{"*"}

	http.ListenAndServe(":9000", handlers.CORS(
		handlers.AllowedHeaders(header),
		handlers.AllowedMethods(method),
		handlers.AllowedOrigins(origin),
	)(r))
}
