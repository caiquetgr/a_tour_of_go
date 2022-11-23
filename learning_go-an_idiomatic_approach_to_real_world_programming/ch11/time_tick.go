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

	go func() {
		for {
			select {
			// should never use time.Tick outside trivial programs, because it's internal time.Ticker cannot be shutdown (therefore, not garbage collected)
			// instead, should use time.Ticker and use the ticker.Stop() to release the resources
			case <-time.Tick(1 * time.Second):
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

	fmt.Println("Finished!")

}
