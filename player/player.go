package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const startingCash = 10_000.00

type Player struct {
	Cash      float64
	Inventory float64
	Bid       float64
	Ask       float64
}

func NewPlayer() *Player {
	return &Player{
		Cash:      startingCash,
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

func (p *Player) GetSpread() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nSet your spread (bid ask): ")
	input, _ := reader.ReadString('\n')
	values := strings.Fields(input)

	if len(values) != 2 {
		fmt.Println("Invalid input. Using previous spread.")
		return
	}

	bid, err1 := strconv.ParseFloat(values[0], 64)
	ask, err2 := strconv.ParseFloat(values[1], 64)

	if err1 != nil || err2 != nil || bid >= ask {
		fmt.Println("Invalid spread. Using previous spread.")
		return
	}

	p.SetSpread(bid, ask)
	fmt.Printf("Spread set: Bid $%.2f / Ask $%.2f\n", bid, ask)
}

func (p *Player) DisplayStatus() {
	fmt.Printf("\nYour Portfolio:\n")
	fmt.Printf(" - Cash: $%.2f\n", p.Cash)
	fmt.Printf(" - Inventory: %.2f\n", p.Inventory)
	fmt.Printf(" - Bid: $%.2f / Ask: $%.2f\n", p.Bid, p.Ask)
	fmt.Printf(" - Net Worth: $%.2f\n", p.Cash+p.Inventory*p.Bid) // Mark-to-market
}

func (p *Player) ShowFinalResults() {
	fmt.Println("\n=== Game Over ===")
	fmt.Printf("Final Net Worth: $%.2f\n",
		p.Cash+p.Inventory*p.Bid)
	fmt.Printf("Profit/Loss: $%.2f (%.2f%%)\n",
		p.Cash-startingCash,
		(p.Cash-startingCash)/startingCash*100)
}
