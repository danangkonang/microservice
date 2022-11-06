package model

type UserLogin struct {
	UserId   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserRegister struct {
	UserId   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserLoginResponse struct {
	Token string `json:"token,omitempty"`
}
