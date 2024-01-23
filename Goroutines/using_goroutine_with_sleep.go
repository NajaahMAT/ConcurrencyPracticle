package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	go doSomething()
	go doSomethingElse()

	time.Sleep(time.Second * 5)

	fmt.Println("\n\nI guess I'm done")
	elapsed := time.Since(start)
	fmt.Println("Processes took elapsed: ", elapsed)

}

func doSomething() {
	time.Sleep(time.Second * 2)
	fmt.Println("\nI have done something")
}

func doSomethingElse() {
	time.Sleep(time.Second * 2)
	fmt.Println("\nI have done something else")
}
