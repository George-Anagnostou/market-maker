package market

import (
	"fmt"
	"market-maker/order"
	"market-maker/player"
	"math/rand/v2"
)

type Market struct {
	Orders []*order.Order
}

func NewMarket() *Market {
	return &Market{}
}

func GenerateOrder() order.Order {
	var orderType string
	if rand.IntN(2) == 0 {
		orderType = "BUY"
	} else {
		orderType = "SELL"
	}

	price := 90 + rand.Float64()*20   // $90 - $110
	quantity := 1 + rand.Float64()*10 // 1-10

	return order.Order{
		OrderType: orderType,
		Price:     price,
		Quantity:  quantity,
	}
}

func (m *Market) GenerateOrders() []*order.Order {
	numOrders := rand.IntN(6) // 0-5 orders
	orders := make([]*order.Order, 0, numOrders)
	for range numOrders {
		order := GenerateOrder()
		orders = append(orders, &order)
	}
	m.Orders = orders
	return orders
}

func (m *Market) ProcessOrders(p *player.Player) []string {
	trades := []string{}
	for _, ord := range m.Orders {
		if ord.OrderType == "BUY" && ord.Price >= p.Ask {
			p.Cash += ord.Price * ord.Quantity
			p.Inventory -= ord.Quantity
			trades = append(trades, fmt.Sprintf("Sold %.2f at $%.2f", ord.Quantity, ord.Price))
		} else if !(ord.OrderType == "BUY") && ord.Price <= p.Bid {
			p.Cash -= ord.Price * float64(ord.Quantity)
			p.Inventory += ord.Quantity
			trades = append(trades, fmt.Sprintf("Bought %.2f at $%.2f", ord.Quantity, ord.Price))
		}
	}
	return trades
}
