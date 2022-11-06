package model

type CardRequest struct {
	CartId     string `json:"cart_id"`
	UserId     string `json:"user_id"`
	ProductId  string `json:"product_id"`
	Qty        int64  `json:"qty"`
	IsCheckout bool   `json:"is_checkout"`
}

type CardResponse struct {
	CartId      string `json:"cart_id,omitempty"`
	ProductId   string `json:"product_id,omitempty"`
	ProductName string `json:"product_name,omitempty"`
	Qty         int64  `json:"qty,omitempty"`
	SubTotal    int64  `json:"sub_total,omitempty"`
	Price       int64  `json:"price,omitempty"`
}
