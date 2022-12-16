package model

type Order struct {
	OrderNo           int    `json:"order_no"`
	OrderCustomerName string `json:"order_customerName"`
	OrderAddress      string `json:"order_address"`
	OrderTime         string `json:"order_time"`
	OrderAge          int    `json:"order_age"`
	OrderGender       string `json:"order_gender"`
}

type PostOrder struct {
	PostId string  `json:"post_id"`
	Order  []Order `json:"data"`
}
