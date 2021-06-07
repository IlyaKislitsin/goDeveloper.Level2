package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	count := 1000

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			fmt.Printf("goroutine №%d started\n", i+1)
			defer func() {
				wg.Done()
				fmt.Printf("goroutine №%d ended\n", i+1)
			}()

			time.Sleep(time.Duration((i % 10)) * time.Second)
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}
