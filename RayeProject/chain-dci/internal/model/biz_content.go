package model

type BizContent struct {
	OutTradeNo  string  `json:"out_trade_no"`
	ProductCode string  `json:"product_code"`
	TotalAmount float64 `json:"total_amount"`
	Subject     string  `json:"subject"`
	TimeExpire  string  `json:"time_expire"`
}
