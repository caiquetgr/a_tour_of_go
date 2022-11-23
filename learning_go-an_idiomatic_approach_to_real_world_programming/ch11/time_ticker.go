package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// time.AFter
	var wg sync.WaitGroup
	wg.Add(5)

	done := make(chan struct{})

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-ticker.C: // should never use time.Tick outside trivial programs, because it's internal time.Ticker cannot be shutdown (therefore, not garbage collected)
				fmt.Println("Ping!")
				wg.Done()
			case <-done:
				fmt.Println("Done!")
				return
			}
		}
	}()

	wg.Wait()
	close(done)
	ticker.Stop()

	fmt.Println("Finished ticker!")

}
