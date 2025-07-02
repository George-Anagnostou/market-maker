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

	fmt.Println("\n=== Set Market Spread ===")
	fmt.Println("Enter your bid and ask prices (separated by space)")
	fmt.Println("Example: 95.50 96.50")
	fmt.Print("Your spread: ")
	input, _ := reader.ReadString('\n')
	values := strings.Fields(input)

	if len(values) != 2 {
		fmt.Println("❌ Invalid input format. Please enter two numbers separated by space.")
		return
	}

	bid, err1 := strconv.ParseFloat(values[0], 64)
	ask, err2 := strconv.ParseFloat(values[1], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("❌ Invalid numbers. Please enter valid prices.")
		return
	}
	if bid >= ask {
		fmt.Println("❌ Bid must be lower than ask price.")
		return
	}

	p.SetSpread(bid, ask)
	fmt.Printf("✅ Spread set successfully:\n")
	fmt.Printf("   - Bid: $%.2f (You'll buy at this price)\n", bid)
	fmt.Printf("   - Ask: $%.2f (You'll sell at this price)\n", ask)
}

func (p *Player) DisplayStatus() {
	fmt.Println("\n=== Portfolio Summary ===")
	fmt.Printf("💰 Cash Balance: $%.2f\n", p.Cash)
	fmt.Printf("📦 Inventory: %.2f units\n", p.Inventory)
	fmt.Printf("📊 Current Spread:\n")
	fmt.Printf("   - Bid: $%.2f\n", p.Bid)
	fmt.Printf("   - Ask: $%.2f\n", p.Ask)
	fmt.Printf("📈 Net Worth: $%.2f\n", p.Cash+p.Inventory*p.Bid)
}

func (p *Player) ShowFinalResults() {
	fmt.Println("\n=== Game Over ===")
	fmt.Printf("Final Net Worth: $%.2f\n",
		p.Cash+p.Inventory*p.Bid)
	fmt.Printf("Profit/Loss: $%.2f (%.2f%%)\n",
		p.Cash-startingCash,
		(p.Cash-startingCash)/startingCash*100)
}
