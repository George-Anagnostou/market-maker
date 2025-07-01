package orderbook

import (
	"testing"
	"time"
)

func TestOrderBook(t *testing.T) {

	type testCase struct {
		name        string
		orders      []Order
		wantBids    []Order
		wantAsks    []Order
		wantMatches [][2]string // pairs of matched order IDs
		wantRejects []string    // order IDs that should be rejected
	}

	tests := []testCase{
		{
			name: "Single BUY order",
			orders: []Order{
				{ID: "1", OrderType: Buy, Price: 100, Quantity: 10, Timestamp: time.Now()},
			},
			wantBids: []Order{
				{ID: "1", OrderType: Buy, Price: 100, Quantity: 10},
			},
			wantAsks:    []Order{},
			wantMatches: [][2]string{},
			wantRejects: []string{},
		},
		{
			name: "Single SELL order",
			orders: []Order{
				{ID: "1", OrderType: Sell, Price: 100, Quantity: 10, Timestamp: time.Now()},
			},
			wantBids: []Order{},
			wantAsks: []Order{
				{ID: "1", OrderType: Sell, Price: 100, Quantity: 10},
			},
			wantMatches: [][2]string{},
			wantRejects: []string{},
		},
		{
			name: "Exact price match",
			orders: []Order{
				{ID: "1", OrderType: Buy, Price: 100, Quantity: 10, Timestamp: time.Now()},
				{ID: "2", OrderType: Sell, Price: 100, Quantity: 10, Timestamp: time.Now().Add(time.Second)},
			},
			wantBids:    []Order{},
			wantAsks:    []Order{},
			wantMatches: [][2]string{{"1", "2"}},
			wantRejects: []string{},
		},
		{
			name: "Partial quantity match",
			orders: []Order{
				{ID: "1", OrderType: Buy, Price: 100, Quantity: 10, Timestamp: time.Now()},
				{ID: "2", OrderType: Sell, Price: 100, Quantity: 5, Timestamp: time.Now().Add(time.Second)},
			},
			wantBids: []Order{
				{ID: "1", OrderType: Buy, Price: 100, Quantity: 5},
			},
			wantAsks:    []Order{},
			wantMatches: [][2]string{{"1", "2"}},
			wantRejects: []string{},
		},
		{
			name: "Price-time priority",
			orders: []Order{
				{ID: "1", OrderType: Buy, Price: 100, Quantity: 10, Timestamp: time.Now()},
				{ID: "2", OrderType: Sell, Price: 100, Quantity: 5, Timestamp: time.Now().Add(time.Second)},
				{ID: "3", OrderType: Sell, Price: 100, Quantity: 5, Timestamp: time.Now().Add(2 * time.Second)},
			},
			wantBids:    []Order{},
			wantAsks:    []Order{},
			wantMatches: [][2]string{{"1", "2"}},
			wantRejects: []string{},
		},
		{
			name: "Zero quantity order",
			orders: []Order{
				{ID: "1", OrderType: Buy, Price: 100, Quantity: 0, Timestamp: time.Now()},
			},
			wantBids:    []Order{},
			wantAsks:    []Order{},
			wantMatches: [][2]string{},
			wantRejects: []string{"1"},
		},
		{
			name: "Negative price order",
			orders: []Order{
				{ID: "1", OrderType: Buy, Price: -100, Quantity: 10, Timestamp: time.Now()},
			},
			wantBids:    []Order{},
			wantAsks:    []Order{},
			wantMatches: [][2]string{},
			wantRejects: []string{"1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ob := NewOrderBook()

			for _, order := range tt.orders {
				err := ob.AddOrder(order)
				if contains(tt.wantRejects, order.ID) {
					if err == nil {
						t.Errorf("AddOrder() = nil, want error for order %s", order.ID)
					}
					continue
				}
				if err != nil {
					t.Errorf("AddOrder() error = %v, want nil", err)
				}
			}

			// TODO: Implement checks for bids, asks, and matches
			// once OrderBook implementation is complete
		})
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
