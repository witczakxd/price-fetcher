package types

type PriceResponse struct {
	Ticker string `json:"ticker"`
	Price float64 `json:"price"`
	ReqID int `json:"req_id"`
}