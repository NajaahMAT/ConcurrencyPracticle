package main

import (
	"fmt"
	"sync"
)

var (
	wg                 = sync.WaitGroup{}
	mutex              = sync.Mutex{}
	inventoryCount int = 1000
	newPurchase        = sync.NewCond(&mutex)
)

func main() {
	fmt.Println("Inventory count at the start: ", inventoryCount)
	wg.Add(2)
	go makeSales()
	go newPurchases()
	wg.Wait()
	fmt.Println("Inventory count at the end: ", inventoryCount)

}

func makeSales() {
	for i := 0; i < 300000; i++ {
		mutex.Lock()
		for inventoryCount-100 < 0 {
			newPurchase.Wait()
		}
		inventoryCount -= 100
		fmt.Println(inventoryCount)
		mutex.Unlock()
	}

	wg.Done()
}

func newPurchases() {
	for i := 0; i < 300000; i++ {
		mutex.Lock()
		inventoryCount += 100
		newPurchase.Signal()
		mutex.Unlock()
	}

	wg.Done()
}
