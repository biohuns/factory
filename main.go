package main

import (
	"fmt"
	"time"

	"github.com/biohuns/factory/pkg/building"
)

func main() {
	ironOreDeposit := building.NewIronOreDeposit()
	storage := building.NewStorage()
	ironOreDeposit.SetOutput(storage)

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		ironOreDeposit.Tick(100)
		storage.Tick(100)
		fmt.Println("ironOreDeposite amount:", ironOreDeposit.Amount())
		fmt.Println("storage item:", storage.Item(), "amount:", storage.Amount())
	}
}
