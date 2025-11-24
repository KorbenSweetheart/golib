package wp

import (
	"sync"
)

// Job is simply a function that needs executing
type Job func()

type Pool struct {
	jobQueue chan Job
	wg       sync.WaitGroup
	quit     chan bool
}

// New creates the pool and immediately spins up the workers
func New(numWorkers int, bufferSize int) *Pool {
	p := &Pool{
		jobQueue: make(chan Job, bufferSize),
		quit:     make(chan bool),
	}

	// Spin up workers immediately (The "True" Pool)
	for i := 0; i < numWorkers; i++ {
		p.wg.Add(1)
		go func(id int) {
			defer p.wg.Done()
			for {
				select {
				case job, ok := <-p.jobQueue:
					if !ok {
						return // Channel closed
					}
					job() // Execute the closure
				case <-p.quit:
					return
				}
			}
		}(i)
	}

	return p
}

// Handle submits a job. It blocks if the buffer is full.
func (p *Pool) Handle(j Job) {
	p.jobQueue <- j
}

// Stop waits for submitted jobs to finish and shuts down.
func (p *Pool) Stop() {
	close(p.jobQueue) // Signal workers to drain the channel
	p.wg.Wait()       // Wait for them to finish
}
