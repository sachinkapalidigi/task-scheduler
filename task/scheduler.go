package task

import (
	"container/heap"
	"fmt"
	"sync"
	"task-scheduler/collection"
	"task-scheduler/threads"
	"time"
)

type Scheduler struct {
	pq   collection.PriorityTaskHeap
	tp   *threads.Pool
	stop chan struct{}
	wg   sync.WaitGroup
}

type Task func(interface{})

func NewScheduler(maxWorkers int) *Scheduler {
	pq := collection.PriorityTaskHeap{}
	heap.Init(&pq)
	tp := threads.NewPool(maxWorkers)
	stop := make(chan struct{})
	sh := Scheduler{
		pq:   pq,
		tp:   tp,
		stop: stop,
	}
	return &sh
}

func (s *Scheduler) Schedule(t Task, at time.Time) {
	fmt.Println("Scheduling task to run at: ", at.String())
	heap.Push(&s.pq, collection.PriorityTask{
		Task:     t,
		Priority: at.UnixNano(),
	})
	s.wg.Add(1)
}

func (s *Scheduler) Start() {
	go func() {
		for {
			select {
			case <-s.stop:
				return
			default:
				// process task
				now := time.Now()
				fmt.Println("Checking at: ", now.String())
				for !s.pq.Empty() && s.pq.Peek().Priority <= now.UnixNano() {
					pt := heap.Pop(&s.pq).(collection.PriorityTask)
					s.tp.Submit(func(workerID int) {
						fmt.Println("Sumbmitting task for execution")
						pt.Task(workerID)
					})
				}
				time.Sleep(1 * time.Second)
			}
		}
	}()
}

func (s *Scheduler) Wait() {
	s.wg.Wait()
}

func (s *Scheduler) Done() {
	s.wg.Done()
}

func (s *Scheduler) Stop() {
	close(s.stop)
	s.tp.Stop()
}
