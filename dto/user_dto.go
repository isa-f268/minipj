package dto

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResp struct {
	Email     string
	Full_name string
}

type RegisterReq struct {
	Email    string
	Password string
	Name     string
}
type LoginResp struct {
	Token string `json:"token"`
}

type TopUpReq struct {
	Balance int `json:"balance"`
}

type TopUpResp struct {
	Balance       int    `json:"balance"`
	Total_balance int    `json:"total"`
	Topup_date    string `json:"date"`
}

type UserData struct {
	Full_name string `json:"full_name"`
	Weight    int    `json:"weight"`
	Height    int    `json:"height"`
}

type ErrorResponse struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}
