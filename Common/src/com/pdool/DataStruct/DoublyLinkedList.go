package DataStruct

import "errors"

// 双向链表
type DoublyLinkedList struct {
	length int
	head   *DNode
	tail   *DNode
}

// DNode ...
type DNode struct {
	Element interface{}
	Next    *DNode
	Prev    *DNode
}

// Head ...
func (d *DoublyLinkedList) Head() *DNode {
	return d.head
}

// Tail ...
func (d *DoublyLinkedList) Tail() *DNode {
	return d.tail
}

// Length ...
func (d *DoublyLinkedList) Length() int {
	return d.length
}

// Insert ...
func (d *DoublyLinkedList) Insert(position int, element interface{}) error {
	if position < 0 || position > d.length {
		return errors.New("position out of range")
	}
	node := DNode{}
	node.Element = element

	if position == 0 {
		if d.head == nil {
			d.tail = &node
		} else {
			d.head.Prev = &node
		}
		node.Next = d.head
		d.head = &node
		d.length++
		return nil
	}

	if position == d.length {
		node.Prev = d.tail
		if d.tail != nil {
			d.tail.Next = &node
		}
		d.tail = &node
		d.length++
		return nil
	}
	current := d.head.Next
	index := 1
	for {
		if current == nil {
			break
		}
		if position == index {
			node.Prev = current.Prev
			node.Next = current
			current.Prev.Next = &node
			current.Prev = &node
			d.length++
			return nil
		}
		index++
		current = current.Next
	}
	return nil
}

//RemoveAt ...
func (d *DoublyLinkedList) RemoveAt(position int) (interface{}, error) {
	if position < 0 || position >= d.length {
		return nil, errors.New("position out of range")
	}
	var item interface{}
	if position == 0 {
		item = d.head.Element
		d.head = d.head.Next
		if d.head != nil {
			d.head.Prev = nil
		}
		if position == 1 {
			d.tail = d.head
		}
		d.length--
		return item, nil
	}
	if position == d.length {
		item = d.tail.Element
		d.tail = d.tail.Prev
		d.tail.Next = nil
		if position == 1 {
			d.head = d.tail
		}
		d.length--
		return item, nil
	}

	current := d.head.Next
	index := 1

	for {
		if current == nil {
			break
		}

		if position == index {
			item = current.Element
			current.Prev.Next = current.Next
			current.Next.Prev = current.Prev
			break
		}
		current = current.Next
		index++
	}
	d.length--
	return item, nil
}
