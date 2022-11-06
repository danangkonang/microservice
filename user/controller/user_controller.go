package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/danangkonang/user/helper"
	"github.com/danangkonang/user/model"
	"github.com/danangkonang/user/service"
)

type userController struct {
	service service.ServiceUser
}

func NewUserController(su service.ServiceUser) *userController {
	return &userController{
		service: su,
	}
}

func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	var user *model.UserLogin

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.MakeRespon(w, 400, "form body required", nil)
		return
	}
	defer r.Body.Close()

	res, err := c.service.Login(user.UserName)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}

	if err := helper.VeryfiPassword(user.Password, res.Password); err != nil {
		helper.MakeRespon(w, 400, "invalid password", nil)
		return
	}
	fmt.Println(res.UserId)
	token, err := helper.GenerateToken(res.UserId)
	if err != nil {
		helper.MakeRespon(w, 500, "internal server error", nil)
		return
	}
	now := time.Now()
	data := map[string]interface{}{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   int64(now.Add(time.Hour * 24).Sub(now).Seconds()),
	}
	helper.MakeRespon(w, 200, "success", data)
}

func (c *userController) Register(w http.ResponseWriter, r *http.Request) {
	var user *model.UserRegister
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	user.UserId = helper.UUID()
	user.Password = helper.HashPassword(user.Password)

	if err := c.service.Register(user); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", nil)
}
