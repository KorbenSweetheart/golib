package lib

import (
	"fmt"
	"log"
	"sync"
	"time"
)

// semaphore with drop.
// can be used when we can drop some events or function that we wrap with it is not that important

var maxGoroutins = 10
var sem = make(chan struct{}, maxGoroutins)

func SemaphoreWithDrop(f func()) {
	select {
	case sem <- struct{}{}:
	default:
		return
	}
	go func() {
		f()

		<-sem
	}()
}

// Ideomatic Semaphore

func handleEvents(evts []string) error {
	var wg sync.WaitGroup

	concurrency := 10
	sem := make(chan struct{}, concurrency)

	for _, ev := range evts {
		wg.Add(1)
		sem <- struct{}{} // if we pass structs to chan
		go func(ev string) {
			defer wg.Done()
			defer func() { <-sem }()

			log.Printf("got new event: %s", ev)
			if err := process(ev); err != nil { // place neede function here
				log.Printf("can't handle event: %v", err)
			}
		}(ev)
	}

	wg.Wait()
	return nil
}

// just placeholder
func process(s string) error {
	fmt.Println(s)
	time.Sleep(1 * time.Second)
	return nil
}
