package main

import (
	"fmt"

	"SeedTradingSystem/currency"
)

func helpMenu() {
	fmt.Println("\nCommands:")
	fmt.Println("help: Show the help menu")
	fmt.Println("buy: Buy an item")
	fmt.Println("balance: Show the balance")
	fmt.Println("items: Show the items available for purchase")
	fmt.Println("inventory: Show the items you have in your inventory")
	fmt.Println("exit: Exit the program")
}

func main() {
	storage := currency.NewStorage(100)

	fmt.Println("Welcome to the trading system!")
	helpMenu()
	var cmd string

	for {
		fmt.Print("\nEnter a command: ")
		fmt.Scanln(&cmd)

		switch cmd {
		case "help":
			helpMenu()
		case "balance":
			fmt.Printf("Balance: %d\n", storage.Balance)
		case "items":
			fmt.Println("\nItems:")
			for i, item := range currency.Items {
				fmt.Printf("%d. %s - %d seeds\n", i+1, item.Name, item.Cost)
			}
		case "inventory":
			if len(storage.Items) == 0 {
				fmt.Println("Inventory is empty")
				break
			}

			fmt.Println("\nInventory:")
			i := 1
			for item, amount := range storage.Items {
				fmt.Printf("%d. %s - %d\n", i, item.Name, amount)
				i++
			}
		case "buy":
			var itemIdx int
			fmt.Print("Enter the item index: ")
			fmt.Scanln(&itemIdx)
			itemIdx--

			if itemIdx < 0 || itemIdx >= currency.ItemsAmount {
				fmt.Println("Invalid item index")
				break
			}

			item := currency.Items[itemIdx]
			if storage.BuyItem(item) {
				fmt.Printf("Bought %s for %d seeds\n", item.Name, item.Cost)
			} else {
				fmt.Println("Not enough seeds to buy the item")
			}
		case "exit":
			fmt.Println("Exiting the program...")
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}
