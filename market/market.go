package market

import (
	"fmt"
	"market-maker/orderbook"
	"market-maker/player"
	"math/rand/v2"
	"time"
)

type Market struct {
	orderBook          *orderbook.OrderBook
	generateOrdersFunc func() []*orderbook.Order // For testing
}

func NewMarket() *Market {
	return &Market{
		orderBook: orderbook.NewOrderBook(),
	}
}

func GenerateOrder() orderbook.Order {
	var orderType orderbook.OrderType
	if rand.IntN(2) == 0 {
		orderType = orderbook.Buy
	} else {
		orderType = orderbook.Sell
	}

	price := 90 + rand.Float64()*20   // $90 - $110
	quantity := 1 + rand.Float64()*10 // 1-10

	return orderbook.Order{
		ID:        fmt.Sprintf("%d", time.Now().UnixNano()),
		OrderType: orderType,
		Price:     price,
		Quantity:  quantity,
		Timestamp: time.Now(),
	}
}

func (m *Market) GenerateOrders() []*orderbook.Order {
	if m.generateOrdersFunc != nil {
		return m.generateOrdersFunc()
	}

	numOrders := rand.IntN(6) // 0-5 orders
	orders := make([]*orderbook.Order, 0, numOrders)
	for range numOrders {
		order := GenerateOrder()
		m.orderBook.AddOrder(order)
		orders = append(orders, &order)
	}
	return orders
}

func (m *Market) ProcessOrders(p *player.Player) []string {
	trades := []string{}
	orders := append(m.orderBook.GetBids(), m.orderBook.GetAsks()...)

	for _, ord := range orders {
		if ord.OrderType == orderbook.Buy && ord.Price >= p.Ask {
			p.Cash += ord.Price * ord.Quantity
			p.Inventory -= ord.Quantity
			trades = append(trades, fmt.Sprintf("✅ Sold %.2f units @ $%.2f (Total: $%.2f)",
				ord.Quantity, ord.Price, ord.Price*ord.Quantity))
		} else if ord.OrderType == orderbook.Sell && ord.Price <= p.Bid {
			p.Cash -= ord.Price * float64(ord.Quantity)
			p.Inventory += ord.Quantity
			trades = append(trades, fmt.Sprintf("✅ Bought %.2f units @ $%.2f (Total: $%.2f)",
				ord.Quantity, ord.Price, ord.Price*ord.Quantity))
		}
	}
	return trades
}

func (m *Market) DisplayState() {
	orders := append(m.orderBook.GetBids(), m.orderBook.GetAsks()...)
	fmt.Printf("\nMarket Orders (%d):\n", len(orders))
	for _, o := range orders {
		fmt.Printf(" - %s %s @ $%.2f (Qty: %.2f)\n",
			o.ID, o.OrderType, o.Price, o.Quantity)
	}
}
