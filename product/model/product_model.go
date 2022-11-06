package model

type ProductRequest struct {
	ProductId   string `json:"product_id,omitempty"`
	UserId      string `json:"user_id,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	Qty         int64  `json:"qty,omitempty"`
	Price       int64  `json:"price,omitempty"`
}

type ProductResponse struct {
	ProductId   string `json:"product_id,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	Qty         int64  `json:"qty,omitempty"`
	Price       int64  `json:"price,omitempty"`
}
