package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	wg                   = sync.WaitGroup{}
	inventoryCount int64 = 1000
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
		atomic.AddInt64(&inventoryCount, -100)
	}

	wg.Done()
}

func newPurchase() {
	for i := 0; i < 300000; i++ {
		atomic.AddInt64(&inventoryCount, 100)
	}

	wg.Done()
}
