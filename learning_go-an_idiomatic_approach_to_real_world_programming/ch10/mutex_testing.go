package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := &Sum{
		total: 0,
	}

	var wg sync.WaitGroup

	concurrency := 1000
	wg.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func(number int) {
			sum.Add(number)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(sum.GetTotal())
}

type Sum struct {
	lock  sync.RWMutex
	total int
}

func (s *Sum) Add(i int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.total += i
}

func (s *Sum) GetTotal() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.total
}
