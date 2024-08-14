package main

import (
	"fmt"
	"task-scheduler/task"
	"time"
)

func main() {
	// Worker Thread Pools - Not a task scheduler
	// tpool := threads.NewPool(10)
	// start := time.Now()
	// for i := 0; i < 100; i++ {
	// 	count := i
	// 	tpool.Submit(func(workerID int) {
	// 		fmt.Println("Executing task: ", count, " in worker: ", workerID)
	// 		time.Sleep(1 * time.Second)
	// 	})
	// }
	// fmt.Println("It took a total of: ", time.Since(start))
	// tpool.Stop()

	// Task Scheduler
	scheduler := task.NewScheduler(10)

	scheduler.Start()

	for i := 0; i < 10; i++ {
		tc := time.Now().Add(time.Duration(i+2) * time.Second)
		scheduler.Schedule(func(x interface{}) {
			defer scheduler.Done()
			fmt.Println("Scheduled task running")
		}, tc)
	}
	scheduler.Wait()
	scheduler.Stop()
}
