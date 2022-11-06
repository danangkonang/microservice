package router

import (
	"github.com/danangkonang/product/config"
	"github.com/danangkonang/product/controller"
	"github.com/danangkonang/product/middleware"
	"github.com/danangkonang/product/service"
	"github.com/gorilla/mux"
)

func ProductRouter(router *mux.Router, db *config.DB) {
	rest := controller.NewProductController(
		service.NewServiceProduct(db),
	)
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/products", rest.FindProduct).Methods("GET")
	v1.HandleFunc("/product", middleware.Auth(rest.CreateProduct)).Methods("POST")
}
