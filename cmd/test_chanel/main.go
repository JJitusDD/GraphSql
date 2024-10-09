package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan string, results chan<- string) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j
	}
}

func main() {

}

// Goroutine Coordination: Channels help in coordinating multiple goroutines.
func TestGoroutine() {
	jobs := make(chan string, 100)
	results := make(chan string, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		switch j {
		case 1:
			jobs <- "Ninh"
		case 2:
			jobs <- "Ninh1"
		case 3:
			jobs <- "Ninh2"
		case 4:
			jobs <- "Ninh3"
		case 5:
			jobs <- "Ninh4"
		}

	}
	close(jobs)

	for a := 1; a <= 2; a++ {
		<-results
	}
}

// Timeouts: Using channels to implement timeouts.
func TestTimeOutHandler() {
	c := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c <- "Test process"
	}()

	select {
	case res := <-c:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}

// Pipeline Pattern: Creating a series of stages where the output of one stage is the input to the next.
func TestMultipleGoroutine() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
