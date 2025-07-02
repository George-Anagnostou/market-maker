package orderbook

import "time"

type OrderType int

const (
	Buy OrderType = iota
	Sell
)

func (ot OrderType) String() string {
	return [...]string{"BUY", "SELL"}[ot]
}

type Order struct {
	ID        string
	OrderType OrderType
	Price     float64
	Quantity  float64
	Timestamp time.Time
}

type Trade struct {
	BuyerID  string
	SellerID string
	Price    float64
	Quantity float64
}
