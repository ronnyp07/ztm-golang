package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func wait() {
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
}

type Hits struct {
	count int
	sync.Mutex
}

func serve(wg *sync.WaitGroup, hits *Hits, interation int) {
	wait()
	hits.Lock()
	defer hits.Unlock()
	defer wg.Done()
	hits.count += 1
	fmt.Println("server interation", interation)

}

func main() {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	hitcounter := Hits{}

	for i := 0; i < 20; i++ {
		interation := i
		wg.Add(1)
		go serve(&wg, &hitcounter, interation)
	}

	fmt.Printf("waiting for goroutines....\n\n")
	wg.Wait()

	hitcounter.Lock()
	totalHits := hitcounter.count
	hitcounter.Unlock()

	fmt.Printf("\ntotal hits = %d\n", totalHits)

}
