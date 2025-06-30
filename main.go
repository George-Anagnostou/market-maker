package main

import (
	"fmt"
	"market-maker/market"
	"market-maker/player"
	"strconv"

	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

func main() {
	p := player.NewPlayer()
	m := market.NewMarket()
	rounds := 5

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	for round := 1; round <= rounds; round++ {
		fmt.Printf("\nRound %d: Cash=%s, Inventory=%f\n",
			round, green(fmt.Sprintf("$%.2f", p.Cash)), p.Inventory)

		// Get bid price
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
			return
		}
		bid, _ := strconv.ParseFloat(bidStr, 64)

		// Get ask price
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
			return
		}
		ask, _ := strconv.ParseFloat(askStr, 64)

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

		if p.Cash < 0 {
			fmt.Println(red("Bankrupt! Game Over."))
			break
		}

		if round == rounds {
			avgPrice := 100.0 // Placeholder average market price
			score := p.Cash + float64(p.Inventory)*avgPrice
			fmt.Printf("\nGame Over! Final Cash=%s, Inventory=%f, Score=%s\n",
				green(fmt.Sprintf("$%.2f", p.Cash)), p.Inventory, green(fmt.Sprintf("$%.2f", score)))
		}
	}
}
