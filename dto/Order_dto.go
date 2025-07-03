package dto

type OrderReq struct {
	Book_id    int `json:"book_id"`
	Total_days int `json:"total_day"`
}

type PaymentReq struct {
	Order_id int `json:"order_id"`
}

type MidtransReq struct {
	Order_id int `json:"order_id"`
	Amount   int `json:"amount"`
}
