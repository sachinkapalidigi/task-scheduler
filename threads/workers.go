package threads

import (
	"fmt"
	"sync"
)

type Task func(workerID int)

type Pool struct {
	wg        sync.WaitGroup
	capacity  int
	taskQueue chan Task
	stopCh    chan struct{}
}

func NewPool(capacity int) *Pool {
	p := &Pool{
		wg:        sync.WaitGroup{},
		capacity:  capacity,
		taskQueue: make(chan Task, capacity),
		stopCh:    make(chan struct{}),
	}

	p.wg.Add(capacity)
	for i := 0; i < capacity; i++ {
		go p.worker(i)
	}

	return p
}

func (p *Pool) worker(w int) {
	fmt.Println("Starting worker routine: ", w)
	defer p.wg.Done()
	defer fmt.Println("Stopping worker routine: ", w)
	for {
		select {
		case task := <-p.taskQueue:
			task(w)
		case <-p.stopCh:
			return // exit the worker routine
		}
	}
}

func (p *Pool) Submit(t Task) {
	p.taskQueue <- t
}

func (p *Pool) Stop() {
	close(p.stopCh)
	p.wg.Wait()
	close(p.taskQueue) // close all task queues
}
