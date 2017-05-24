package main

import (
	"fmt"
	"sync"
)

type Queue struct {
	sync.Mutex
	queue []interface{}
	wg    sync.WaitGroup
}

func (q *Queue) Push(e interface{}) {
	go func() {
		defer q.wg.Done()
		defer q.Unlock()

		q.Lock()
		q.queue = append(q.queue, e)
	}()
}

func main() {
	q := Queue{}
	cycles := 10000

	q.wg.Add(cycles)

	for i := 0; i < cycles; i++ {
		q.Push(i)
	}

	// Wait what done all goroutine
	q.wg.Wait()

	// Print len queue and compare with cycles
	fmt.Println(len(q.queue))
	fmt.Println(len(q.queue) == cycles)
}
