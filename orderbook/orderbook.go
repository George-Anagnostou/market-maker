package orderbook

import "errors"

type OrderBook struct {
	bids []Order
	asks []Order
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		bids: make([]Order, 0),
		asks: make([]Order, 0),
	}
}

func (ob *OrderBook) AddOrder(order Order) error {
	if order.Quantity <= 0 {
		return errors.New("order quantity must be positive")
	}
	if order.Price <= 0 {
		return errors.New("order price must be positive")
	}

	switch order.OrderType {
	case Buy:
		ob.bids = append(ob.bids, order)
	case Sell:
		ob.asks = append(ob.asks, order)
	default:
		return errors.New("invalid order type")
	}

	// TODO: Implement order matching logic
	return nil
}

func (ob *OrderBook) GetBids() []Order {
	return ob.bids
}

func (ob *OrderBook) GetAsks() []Order {
	return ob.asks
}
