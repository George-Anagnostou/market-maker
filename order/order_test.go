package order

import "testing"

func TestOrderInitialization(t *testing.T) {
	tests := []struct {
		name      string
		OrderType string
		Price     float64
		Quantity  float64
	}{
		{
			name:      "Buy order",
			OrderType: "BUY",
			Price:     100.50,
			Quantity:  5.00,
		},
		{
			name:      "Sell order",
			OrderType: "SELL",
			Price:     99.75,
			Quantity:  3.00,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ord := Order{
				OrderType: tt.OrderType,
				Price:     tt.Price,
				Quantity:  tt.Quantity,
			}

			if ord.OrderType != tt.OrderType {
				t.Errorf("Order.OrderType = %v, want %v", ord.OrderType, tt.OrderType)
			}
			if ord.Price != tt.Price {
				t.Errorf("Order.Price = %v, want %v", ord.Price, tt.Price)
			}
			if ord.Quantity != tt.Quantity {
				t.Errorf("Order.Quantity = %v, want %v", ord.Quantity, tt.Quantity)
			}
		})
	}
}
