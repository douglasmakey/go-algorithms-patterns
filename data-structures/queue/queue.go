package queue

import "sync"

type Queue struct {
	queue []interface{}
	mu  *sync.Mutex
}

func New() *Queue {
	q := &Queue{}
	q.queue = make([]interface{}, 0)
	q.mu = new(sync.Mutex)

	return q
}


func (q *Queue) Shift() (e interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	e, q.queue = q.queue[0], q.queue[1:]
	return
}

func (q *Queue) Push(e interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.queue = append(q.queue, e)
}

func (q *Queue) Peek() interface{} {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.queue[0]
}