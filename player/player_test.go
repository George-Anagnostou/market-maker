package player

import "testing"

func TestNewPlayer(t *testing.T) {
	p := NewPlayer()
	if p.Cash != 10_000 {
		t.Errorf("NewPlayer.Cash = %v, want 10,000", p.Cash)
	}
	if p.Inventory != 0 {
		t.Errorf("NewPlayer.Inventory = %v, want 0", p.Inventory)
	}
	if p.Bid != 0 {
		t.Errorf("NewPlayer.Bid = %v, want 0", p.Bid)
	}
	if p.Ask != 0 {
		t.Errorf("NewPlayer.Ask = %v, want 0", p.Ask)
	}
}

func TestSetSpread(t *testing.T) {
	p := NewPlayer()
	bid, ask := 99.0, 101.0
	p.SetSpread(bid, ask)
	if p.Bid != bid {
		t.Errorf("Player.Bid = %v, want %v", p.Bid, bid)
	}
	if p.Ask != ask {
		t.Errorf("Player.Ask = %v, want %v", p.Ask, ask)
	}
	if p.Cash != 10_000 {
		t.Errorf("Player.Cash = %v, want 10,000 (unchanged)", p.Cash)
	}
	if p.Inventory != 0 {
		t.Errorf("Player.Inventory = %v, want 0 (unchanged)", p.Inventory)
	}
}
