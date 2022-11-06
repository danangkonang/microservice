package router

import (
	"github.com/danangkonang/user/config"
	"github.com/danangkonang/user/controller"
	"github.com/danangkonang/user/service"
	"github.com/gorilla/mux"
)

func ProfileRouter(router *mux.Router, db *config.DB) {
	rest := controller.NewUserController(
		service.NewServiceUser(db),
	)
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/user/login", rest.Login).Methods("POST")
	v1.HandleFunc("/user/register", rest.Register).Methods("POST")
}
