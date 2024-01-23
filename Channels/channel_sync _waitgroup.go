package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		var result = 0
		var value = 97

		goChan := make(chan int)
		mainChan := make(chan string, 1)

		calculateSquare := func() {
			time.Sleep(time.Second * 3)
			result = value * value
			goChan <- result
		}

		reportResult := func() {
			fmt.Println("The result of ", value, " squared is ", <-goChan)
			mainChan <- "You can quit now. I'm done."
		}

		go calculateSquare()
		go reportResult()
		<-mainChan
		wg.Done()
	}()

	wg.Wait()

}
