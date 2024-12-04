package main

import (
	"fmt"

	"SeedTradingSystem/currency"
)

func main() {
	storage := currency.NewStorage()
	fmt.Printf("You have %d items\n", len(storage.Items))
}
