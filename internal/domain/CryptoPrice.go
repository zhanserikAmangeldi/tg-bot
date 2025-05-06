package domain

type CryptoPrice struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price,string"`
}
