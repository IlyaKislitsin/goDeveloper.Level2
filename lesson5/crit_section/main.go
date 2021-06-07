package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		count   = 1000
		testMap = make(map[int]int, count/100)
		wg      = sync.WaitGroup{}
		mutex   sync.Mutex
	)

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func(i int) {
			mutex.Lock()
			defer mutex.Unlock()
			defer wg.Done()
			testMap[i%10] += 1
		}(i)
	}

	wg.Wait()
	fmt.Println(testMap)
	fmt.Println("done")
}
