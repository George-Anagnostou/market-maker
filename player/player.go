package player

type Player struct {
	Cash      float64
	Inventory float64
	Bid       float64
	Ask       float64
}

func NewPlayer() *Player {
	return &Player{
		Cash:      10_000,
		Inventory: 0.00,
	}
}

func (p *Player) SetSpread(bid, ask float64) {
	p.Bid = bid
	p.Ask = ask
}

func (p *Player) CheckBankruptcy() bool {
	if p.Cash < 0 {
		return true
	}
	return false
}
