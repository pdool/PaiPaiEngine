package DataStruct

import (
	"fmt"
)

// Queue queue
type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(element interface{}) {
	q.items = append(q.items, element)
}

func (q *Queue) Dequeue() interface{} {
	firstItem := q.items[0]
	q.items = q.items[1:]
	return firstItem
}

func (q *Queue) Front() interface{} {
	return q.items[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Clear() {
	q.items = []interface{}{}
}

func (q *Queue) Print() {
	fmt.Printf("%#v\n", q)
}
