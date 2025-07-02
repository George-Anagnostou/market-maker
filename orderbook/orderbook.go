package orderbook

import (
	"errors"
	"sort"
	_ "time" // Force import for timestamp comparisons
)

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

func (ob *OrderBook) AddOrder(order Order) ([]Trade, error) {
	if order.Quantity <= 0 {
		return nil, errors.New("order quantity must be positive")
	}
	if order.Price <= 0 {
		return nil, errors.New("order price must be positive")
	}

	switch order.OrderType {
	case Buy:
		ob.bids = append(ob.bids, order)
		sort.Slice(ob.bids, func(i, j int) bool {
			if ob.bids[i].Price == ob.bids[j].Price {
				return ob.bids[i].Timestamp.Before(ob.bids[j].Timestamp)
			}
			return ob.bids[i].Price > ob.bids[j].Price // Descending for bids
		})
		return ob.matchOrders()
	case Sell:
		ob.asks = append(ob.asks, order)
		sort.Slice(ob.asks, func(i, j int) bool {
			if ob.asks[i].Price == ob.asks[j].Price {
				return ob.asks[i].Timestamp.Before(ob.asks[j].Timestamp)
			}
			return ob.asks[i].Price < ob.asks[j].Price // Ascending for asks
		})
		return ob.matchOrders()
	default:
		return nil, errors.New("invalid order type")
	}
}

func (ob *OrderBook) matchOrders() ([]Trade, error) {
	var trades []Trade

	for len(ob.bids) > 0 && len(ob.asks) > 0 {
		bestBid := ob.bids[0]
		bestAsk := ob.asks[0]

		if bestBid.Price < bestAsk.Price {
			break // No more matches possible
		}

		// Determine matched quantity
		qty := bestBid.Quantity
		if bestAsk.Quantity < qty {
			qty = bestAsk.Quantity
		}

		// Execute trade at ask price (price-time priority)
		trades = append(trades, Trade{
			BuyerID:  bestBid.ID,
			SellerID: bestAsk.ID,
			Price:    bestAsk.Price,
			Quantity: qty,
		})

		// Update order quantities
		ob.bids[0].Quantity -= qty
		ob.asks[0].Quantity -= qty

		// Remove filled orders
		if ob.bids[0].Quantity == 0 {
			ob.bids = ob.bids[1:]
		}
		if ob.asks[0].Quantity == 0 {
			ob.asks = ob.asks[1:]
		}
	}

	return trades, nil
}

func (ob *OrderBook) GetBids() []Order {
	return ob.bids
}

func (ob *OrderBook) GetAsks() []Order {
	return ob.asks
}
