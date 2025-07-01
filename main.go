package main

import (
	"bufio"
	"fmt"
	"market-maker/market"
	"market-maker/player"
	"math/rand/v2"
	"os"
)

const (
	totalRounds = 10
)

func main() {
	rand.New(rand.NewPCG(100, 200)) // For reproducible results

	p := player.NewPlayer()
	m := market.NewMarket()

	fmt.Println("=== Market Maker Game ===")
	fmt.Printf("Starting with $%.2f\n", p.Cash)
	fmt.Printf("Playing %d rounds\n\n", totalRounds)

	for round := 1; round <= totalRounds; round++ {
		fmt.Printf("=== Round %d/%d ===\n", round, totalRounds)

		m.DisplayState()
		p.GetSpread()

		trades := m.ProcessOrders(p)
		if len(trades) > 0 {
			fmt.Println("\nTrades executed:")
			for _, trade := range trades {
				fmt.Println(" -", trade)
			}
		} else {
			fmt.Println("\nNo trades this round")
		}

		p.DisplayStatus()

		if p.CheckBankruptcy() {
			fmt.Println("\n☠️  BANKRUPTCY! Game over.")
			break
		}

		if round < totalRounds {
			waitForNextRound()
		}
	}

	p.ShowFinalResults()
}

func waitForNextRound() {
	fmt.Print("\nPress enter to continue...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
