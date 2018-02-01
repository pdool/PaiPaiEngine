package Event

import (
	"container/list"
	"sync"
)

type Queue struct {
	l *list.List
	m sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{l: list.New()}
}

func (q *Queue) Push(v interface{}) {
	if v == nil {
		return
	}
	q.m.Lock()
	defer q.m.Unlock()
	q.l.PushBack(v)
}

func (q *Queue) Pop() *list.Element {
	q.m.Lock()
	defer q.m.Unlock()
	element := q.l.Front()
	if element == nil {
		return nil
	}
	q.l.Remove(element)
	return element
}

func (q *Queue) Len() int {
	q.m.Lock()
	defer q.m.Unlock()
	return q.l.Len()
}
