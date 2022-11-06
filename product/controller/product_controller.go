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

type productController struct {
	service service.ServiceProduct
}

func NewProductController(su service.ServiceProduct) *productController {
	return &productController{
		service: su,
	}
}

func (c *productController) FindProduct(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.Findproduct()
	if err != nil {
		helper.MakeRespon(w, 500, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", res)
}

func (c *productController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var prd *model.ProductRequest
	err := json.NewDecoder(r.Body).Decode(&prd)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	prd.ProductId = helper.UUID()
	tk := middleware.ExtractToken(r)

	prd.UserId = fmt.Sprintf("%s", tk["UserId"])

	if err := c.service.CreateProduct(prd); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", nil)
}
