package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danangkonang/product/helper"
	"github.com/danangkonang/product/middleware"
	"github.com/danangkonang/product/model"
	"github.com/danangkonang/product/service"
)

type cartController struct {
	service service.ServiceCart
}

func NewCartController(sc service.ServiceCart) *cartController {
	return &cartController{
		service: sc,
	}
}

func (c *cartController) FindMyCard(w http.ResponseWriter, r *http.Request) {
	tk := middleware.ExtractToken(r)
	res, err := c.service.FindMyCard(fmt.Sprintf("%s", tk["UserId"]))
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", res)
}

func (c *cartController) CreateCart(w http.ResponseWriter, r *http.Request) {
	var prd *model.CardRequest
	err := json.NewDecoder(r.Body).Decode(&prd)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	tk := middleware.ExtractToken(r)

	prd.CartId = helper.UUID()
	prd.IsCheckout = false
	prd.UserId = fmt.Sprintf("%s", tk["UserId"])

	if err := c.service.CreateCart(prd); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", nil)
}
