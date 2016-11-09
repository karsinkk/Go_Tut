package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func workThatShit(w int, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for i := range jobs {
		fmt.Printf("Worker %d - Processing Job : %d \n", w, i)
		time.Sleep(time.Second * 1)
		results <- i * 32
	}
}

func main() {

	jobs := make(chan int)
	results := make(chan int, 100)

	for i := 1; i < 4; i++ {
		wg.Add(1)
		go workThatShit(i, jobs, results)
	}

	for j := 1; j < 10; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for res := range results {
			fmt.Println("Results:", res)
		}
	}()
	close(results)
	wg.Wait()

}
