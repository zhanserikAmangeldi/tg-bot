package domain

type CryptoPrice struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}
