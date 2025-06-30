package market

import (
	"market-maker/order"
	"market-maker/player"
	"math/rand/v2"
	"testing"
)

func TestGenerateOrders(t *testing.T) {
	var seed1 uint64 = 100
	var seed2 uint64 = 200
	rand.New(rand.NewPCG(seed1, seed2)) // reproducable results

	m := NewMarket()
	orders := m.GenerateOrders()

	// check order count
	if len(orders) <= 0 {
		t.Errorf("GenerateOrders count = %v, want > 0", len(orders))
	}

	// check order properties
	// these properties should be better specified, but for now
	// we will check that they are simply non-zero
	for _, ord := range orders {
		if ord.Price == 0 {
			t.Errorf("Order.Price = %v, want non-zero", ord.Price)
		}
		if ord.Quantity == 0 {
			t.Errorf("Order.Quantity = %v, want non-zero", ord.Quantity)
		}
	}
}

func TestProcessOrders(t *testing.T) {
	p := player.NewPlayer()
	m := NewMarket()

	p.SetSpread(99, 101)

	m.Orders = []*order.Order{
		{OrderType: "BUY", Price: 102, Quantity: 5},  // should match
		{OrderType: "SELL", Price: 98, Quantity: 3},  // should match
		{OrderType: "BUY", Price: 100, Quantity: 2},  // no match
		{OrderType: "SELL", Price: 100, Quantity: 4}, // no match
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
	// sold 5 at $102 -> 5*102=$510, -5 inventory
	// bought 3 at $98 -> -3*98=-@294, +3 inventory
	// Net cash = 10,000 + 510 - 294 = $10,216
	// Net inventory = 0 - 5 + 3 = -2
	wantCash := 10216.0
	wantInventory := -2.0
	if p.Cash != wantCash {
		t.Errorf("Player.Cash = %v, want = %v", p.Cash, wantCash)
	}
	if p.Inventory != wantInventory {
		t.Errorf("Player.Inventory = %v, want %v", p.Inventory, wantInventory)
	}
}
