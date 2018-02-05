package DataStruct

// LinkedList ...
type LinkedList struct {
	length int
	head   *Node
}

// Node ...
type Node struct {
	Element interface{}
	Next    *Node
}

// Append add element to LinkedList
func (l *LinkedList) Append(element interface{}) {
	n := Node{}
	n.Element = element

	if l.head == nil {
		l.head = &n
	} else {
		current := l.head
		for {
			if current.Next == nil {
				current.Next = &n
				break
			}
			current = current.Next
		}
	}
	l.length++
}

// Insert element at position
func (l *LinkedList) Insert(position int, element interface{}) {
	if position >= 0 && position <= l.length {
		n := Node{}
		n.Element = element

		if position == 0 {
			n.Next = l.head
			l.head = &n
		} else {
			previous := l.head
			index := 1
			for {
				if index == position {
					break
				}
				index++
				previous = previous.Next
			}
			n.Next = previous.Next
			previous.Next = &n
		}

		l.length++
	}
}

//Remove remove element
func (l *LinkedList) Remove(element interface{}) int {
	current := l.head
	var previous *Node
	index := 0
	for {
		if current.Element == element {
			previous = current.Next
			l.length--
			if index == 0 {
				l.head = previous
			}
			break
		}
		if current.Next == nil {
			break
		}
		previous = current
		current = current.Next
		index++
	}
	return index
}

// RemoveAt remove at position
func (l *LinkedList) RemoveAt(position int) interface{} {
	if position >= 0 && position < l.length {
		current := l.head
		index := 0
		for {
			if position == index {
				ele := current.Element
				current = current.Next
				l.length--
				if index == 0 {
					l.head = current
				}
				return ele
			}
			if current == nil {
				break
			}
			index++
			current = current.Next
		}
	}
	return nil
}

//IndexOf ...
func (l *LinkedList) IndexOf(element interface{}) int {
	current := l.head
	index := 0
	for {
		if current.Element == element {
			return index
		}
		if current == nil {
			return -1
		}
		current = current.Next
	}
}

// IsEmpty ...
func (l *LinkedList) IsEmpty() bool {
	return l.length == 0
}

// Size ...
func (l *LinkedList) Size() int {
	return l.length
}

// Length ...
func (l *LinkedList) Length() int {
	return l.length
}

// Head ...
func (l *LinkedList) Head() *Node {
	return l.head
}
func (l *LinkedList) Get(position int) interface{} {
	if position >= 0 && position < l.length {
		current := l.head
		index := 0
		for {
			if position == index {
				ele := current.Element
				current = current.Next
				l.length--
				if index == 0 {
					l.head = current
				}
				return ele
			}
			if current == nil {
				break
			}
			index++
			current = current.Next
		}
	}
	return nil
}
