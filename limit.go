package main

import (
"fmt"
"sync"
)
func main() {
	var wg sync.WaitGroup
	maxGoroutines := 1000
	guard := make(chan struct{}, maxGoroutines)

	for i := 0; i < 30; i++ {
		guard <- struct{}{}
		wg.Add(1)
		go func(n int) { 
			defer func() {
			wg.Done()
			<-guard
			}()
			worker(n)
		}(i)
	}
	wg.Wait()
}

func worker(i int) { fmt.Println("doing work on", i) }
