package main

import (
	"fmt"
	"time"
)

func main() {
	const numJobs = 10
	jobsChan := make(chan int, numJobs)
	completedJobsChan := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobsChan, completedJobsChan)
	}

	for j := 1; j <= numJobs; j++ {
		jobsChan <- j
	}
	close(jobsChan)

	for a := 1; a <= numJobs; a++ {
		<-completedJobsChan
	}
}

func worker(id int, jobsChan <-chan int, completedJobsChan chan<- int) {
	for j := range jobsChan {
		fmt.Println("worker", id, "started job", j, "with", len(jobsChan), "jobs left to process")
		time.Sleep(time.Second * 2)
		fmt.Println("worker", id, "      finished job", j)
		completedJobsChan <- j
	}
}
