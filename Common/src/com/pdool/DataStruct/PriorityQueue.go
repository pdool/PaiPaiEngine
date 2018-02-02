package DataStruct

import (
	"fmt"
)

// PriotityQueue queue with priotity
type PriorityQueue struct {
	items []item
}

type item struct {
	Value    interface{}
	Priority int
}

func (p *PriorityQueue) Enqueue(element interface{}, priority int) {
	ele := item{}
	ele.Value = element
	ele.Priority = priority
	p.items = append(p.items, ele)
	for i := len(p.items) - 1; i > 0; i-- {
		if p.items[i-1].Priority < priority {
			p.items[i-1], p.items[i] = p.items[i], p.items[i-1]
		}
	}
}

func (p *PriorityQueue) Dequeue() interface{} {
	firstValue := p.items[0].Value
	p.items = p.items[1:]
	return firstValue
}

func (p *PriorityQueue) Front() interface{} {
	return p.items[0].Value
}

func (p *PriorityQueue) IsEmpty() bool {
	return len(p.items) == 0
}

func (p *PriorityQueue) Size() int {
	return len(p.items)
}

func (p *PriorityQueue) Clear() {
	p.items = []item{}
}
func (p *PriorityQueue) Print() {
	fmt.Printf("%#v\n", p)
}
