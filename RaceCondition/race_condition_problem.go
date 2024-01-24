package main

import (
	"fmt"
	"sync"
)

var (
	wg             sync.WaitGroup
	inventoryCount int = 1000
)

func main() {
	fmt.Println("Inventory count at the start: ", inventoryCount)
	wg.Add(2)
	go makeSales()
	go newPurchase()
	wg.Wait()
	fmt.Println("Inventory count at the end: ", inventoryCount)

}

func makeSales() {
	for i := 0; i < 3000; i++ {
		inventoryCount -= 100
	}

	wg.Done()
}

func newPurchase() {
	for i := 0; i < 3000; i++ {
		inventoryCount += 100
	}

	wg.Done()
}
