package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//synchronous Task
	synchronousTask()

	//asynchronous Task
	var wg sync.WaitGroup
	wg.Add(1)
	go asynchronousTask(&wg)

	//wait for the asynchronous task to complete
	wg.Wait()

	fmt.Println("Main function completed")
}

func synchronousTask() {
	fmt.Println("Synchronous Task Completed")
}

func asynchronousTask(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("Asynchronous Task Completed")
}
