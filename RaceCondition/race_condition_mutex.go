package main

import (
	"fmt"
	"sync"
)

var (
	wg                 = sync.WaitGroup{}
	mutex              = sync.Mutex{}
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
	for i := 0; i < 300000; i++ {
		mutex.Lock()
		inventoryCount -= 100
		mutex.Unlock()
	}

	wg.Done()
}

func newPurchase() {
	for i := 0; i < 300000; i++ {
		mutex.Lock()
		inventoryCount += 100
		mutex.Unlock()
	}

	wg.Done()
}
