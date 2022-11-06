package router

import (
	"github.com/danangkonang/product/config"
	"github.com/danangkonang/product/controller"
	"github.com/danangkonang/product/middleware"
	"github.com/danangkonang/product/service"
	"github.com/gorilla/mux"
)

func CartRouter(router *mux.Router, db *config.DB) {
	rest := controller.NewCartController(
		service.NewServiceCart(db),
	)
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/carts", middleware.Auth(rest.FindMyCard)).Methods("GET")
	v1.HandleFunc("/cart", middleware.Auth(rest.CreateCart)).Methods("POST")
}
