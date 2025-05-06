package domain

type PriceAlert struct {
	UserId   int64
	Symbol   string
	Price    float64
	Above    bool
	Notified bool
}
