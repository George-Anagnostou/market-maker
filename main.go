package main

import (
	"bufio"
	"fmt"
	"market-maker/market"
	"market-maker/player"
	"math/rand/v2"
	"os"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const (
	totalRounds = 5
)

func main() {
	rand.New(rand.NewPCG(100, 200)) // For reproducible results

	printer := message.NewPrinter(language.English)

	p := player.NewPlayer()
	m := market.NewMarket()

	fmt.Println("=== Market Maker Game ===")
	printer.Printf("Starting with $%.2f\n", p.Cash)
	printer.Printf("Playing %d rounds\n\n", totalRounds)

	for round := 1; round <= totalRounds; round++ {
		printer.Printf("=== Round %d/%d ===\n", round, totalRounds)

		p.GetSpread()
		m.GenerateOrders()
		m.DisplayState()

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
