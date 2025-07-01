package main

import (
	"fmt"
	"market-maker/market"
	"market-maker/player"
	"strconv"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func printTitle() {
	myFigure := figure.NewColorFigure("Market Maker", "", "green", true)
	myFigure.Print()
}

func printRoundStats(round int, p *player.Player, green func(a ...any) string) {
	printer := message.NewPrinter(language.English)
	fmt.Printf("\nRound %d:\n", round)
	fmt.Printf("\tCash = %s, Inventory = %s\n",
		green(printer.Sprintf("$%.2f", p.Cash)), printer.Sprintf("%.2f", p.Inventory))
}

func getBidPrice(red func(a ...any) string) (float64, error) {
	bidPrompt := promptui.Prompt{
		Label: "Enter bid price",
		Validate: func(input string) error {
			val, err := strconv.ParseFloat(input, 64)
			if err != nil || val <= 0 {
				return fmt.Errorf("invalid price")
			}
			return nil
		},
	}
	bidStr, err := bidPrompt.Run()
	if err != nil {
		fmt.Println(red("Error reading bid"), err)
		return 0.0, err
	}
	bid, _ := strconv.ParseFloat(bidStr, 64)
	return bid, nil
}

func getAskPrice(bid float64, red func(a ...any) string) (float64, error) {
	askPrompt := promptui.Prompt{
		Label: "Enter ask price",
		Validate: func(input string) error {
			val, err := strconv.ParseFloat(input, 64)
			if err != nil || val <= bid {
				return fmt.Errorf("ask must be greater than bid")
			}
			return nil
		},
	}
	askStr, err := askPrompt.Run()
	if err != nil {
		fmt.Println(red("Error reading ask"), err)
		return 0.0, err
	}
	ask, _ := strconv.ParseFloat(askStr, 64)
	return ask, nil
}

func printScore(avgPrice float64, p *player.Player, green func(a ...any) string) {
	printer := message.NewPrinter(language.English)
	score := p.Cash + float64(p.Inventory)*avgPrice
	fmt.Printf("\nGame Over! Final Cash = %s, Inventory = %s, Score = %s\n",
		green(printer.Sprintf("$%.2f", p.Cash)),
		printer.Sprintf("%.2f", p.Inventory),
		green(printer.Sprintf("$%.2f", score)),
	)
}

func main() {
	p := player.NewPlayer()
	m := market.NewMarket()
	rounds := 5

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	printTitle()

	for round := 1; round <= rounds; round++ {
		printRoundStats(round, p, green)

		// Get bid price
		bid, err := getBidPrice(red)
		if err != nil {
			fmt.Println("Failed to get bid")
			return
		}

		// Get ask price
		ask, err := getAskPrice(bid, red)
		if err != nil {
			fmt.Println("Failed to get ask")
			return
		}

		p.SetSpread(bid, ask)
		m.GenerateOrders()
		trades := m.ProcessOrders(p)

		fmt.Println("Trades:")
		if len(trades) == 0 {
			fmt.Println("  None")
		} else {
			for _, trade := range trades {
				fmt.Printf("  %s\n", trade)
			}
		}

		if p.CheckBankruptcy() {
			fmt.Println(red("Bankrupt! Game Over."))
			return
		}

		if round == rounds {
			avgPrice := 100.00 // placeholder value for now
			printScore(avgPrice, p, green)
		}
	}
}
