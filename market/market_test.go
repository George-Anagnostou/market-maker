package market

import (
	"market-maker/orderbook"
	"market-maker/player"
	"math/rand/v2"
	"testing"
	"time"
)

func TestGenerateOrders(t *testing.T) {
	var seed1 uint64 = 100
	var seed2 uint64 = 200
	rand.New(rand.NewPCG(seed1, seed2)) // reproducible results

	m := NewMarket()
	orders := m.GenerateOrders()

	// check order count
	if len(orders) <= 0 {
		t.Errorf("GenerateOrders count = %v, want > 0", len(orders))
	}

	// check order properties
	for _, ord := range orders {
		if ord.Price == 0 {
			t.Errorf("Order.Price = %v, want non-zero", ord.Price)
		}
		if ord.Quantity == 0 {
			t.Errorf("Order.Quantity = %v, want non-zero", ord.Quantity)
		}
		if ord.ID == "" {
			t.Errorf("Order.ID is empty, want non-empty")
		}
	}
}

func TestProcessOrders(t *testing.T) {
	p := player.NewPlayer()
	p.SetSpread(99, 101)

	now := time.Now()
	m := &Market{
		orderBook: orderbook.NewOrderBook(),
		generateOrdersFunc: func() []*orderbook.Order {
			return []*orderbook.Order{
				{ID: "1", OrderType: orderbook.Buy, Price: 102, Quantity: 5, Timestamp: now},
				{ID: "2", OrderType: orderbook.Sell, Price: 98, Quantity: 3, Timestamp: now},
				{ID: "3", OrderType: orderbook.Buy, Price: 100, Quantity: 2, Timestamp: now},
				{ID: "4", OrderType: orderbook.Sell, Price: 100, Quantity: 4, Timestamp: now},
			}
		},
	}

	trades := m.ProcessOrders(p)

	expectedTrades := []string{
		"Sold 5.00 at $102.00",
		"Bought 3.00 at $98.00",
	}
	if len(trades) != len(expectedTrades) {
		t.Errorf("ProcessOrders trades count = %v, want %v", len(trades), len(expectedTrades))
	}
	for i, trade := range trades {
		if trade != expectedTrades[i] {
			t.Errorf("ProcessOrders trade %d = %v, want %v", i, trade, expectedTrades[i])
		}
	}

	// check player state
	wantCash := 10216.0
	wantInventory := -2.0
	if p.Cash != wantCash {
		t.Errorf("Player.Cash = %v, want = %v", p.Cash, wantCash)
	}
	if p.Inventory != wantInventory {
		t.Errorf("Player.Inventory = %v, want %v", p.Inventory, wantInventory)
	}
}
