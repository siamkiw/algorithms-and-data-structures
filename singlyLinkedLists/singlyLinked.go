package singlyLinkedLists

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  string
	Next *Node
}

type Singlylinkedlist struct {
	Head *Node
	Tail *Node
	Len  int
}

func (l *Singlylinkedlist) Push(value string) {
	newNode := &Node{
		Val: value,
	}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = l.Head
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Len++
}

func (l *Singlylinkedlist) Pop() (*Node, error) {
	if l.Len == 0 {
		return &Node{}, errors.New("empty_list")
	}
	current := l.Head
	newTail := current
	for current.Next != nil {
		newTail = current
		current = current.Next
	}
	l.Tail = newTail
	l.Tail.Next = nil
	l.Len--
	if l.Len == 0 {
		l.Head = nil
		l.Tail = nil
	}
	return current, nil
}

func (l *Singlylinkedlist) Shift() (*Node, error) {
	if l.Len == 0 {
		return &Node{}, errors.New("empty_list")
	}
	head := l.Head
	l.Head = head.Next
	l.Len--
	if l.Len == 0 {
		l.Tail = nil
	}
	return head, nil
}

func (l *Singlylinkedlist) Unshift(value string) {
	newNode := &Node{
		Val: value,
	}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = l.Head
	} else {
		newNode.Next = l.Head
		l.Head = newNode

	}
	l.Len++

}

func (l *Singlylinkedlist) Get(index int) (Node, error) {
	if index < 0 || index >= l.Len {
		return Node{}, errors.New("index_not_valid")
	}
	counter := 0
	current := l.Head
	for counter != index {
		current = current.Next
		counter++
	}
	return *current, nil
}

func (l *Singlylinkedlist) Set(index int, value string) bool {
	if index < 0 || index >= l.Len {
		return false
	}

	counter := 0
	current := l.Head
	for counter != index {
		current = current.Next
		counter++
	}
	current.Val = value

	return true
}

func (l *Singlylinkedlist) Insert(index int, value string) error {
	if index < 0 || index > l.Len {
		return errors.New("index_invalid")
	}

	if index == 0 {
		l.Unshift(value)
		return nil
	} else if index == l.Len {
		l.Push(value)
		return nil
	}

	counter := 1
	current := l.Head.Next
	prevNode := l.Head
	for counter != index {
		prevNode = current
		current = current.Next
		counter++
	}

	newNode := &Node{Val: value}
	newNode.Next = current
	prevNode.Next = newNode
	l.Len++

	return nil
}

func (l *Singlylinkedlist) Remove(index int) (*Node, error) {
	if index < 0 || index >= l.Len {
		return &Node{}, errors.New("index_invalid")
	}
	if index == 0 {
		_, err := l.Shift()
		return &Node{}, err
	} else if index == l.Len-1 {
		_, err := l.Pop()
		return &Node{}, err
	}

	counter := 1
	removed := l.Head.Next
	prevNode := l.Head
	for counter != index {
		prevNode = removed
		removed = removed.Next
		counter++
	}

	prevNode.Next = removed.Next
	return removed, nil
}

func (l *Singlylinkedlist) Reverse() error {
	if l.Len == 0 {
		return errors.New("empty_link")
	}
	node := l.Head
	l.Head = l.Tail
	l.Tail = node

	var next *Node
	var prve *Node

	for index := 0; index < l.Len; index++ {
		next = node.Next
		node.Next = prve
		prve = node
		node = next
	}
	return nil
}

// result "value -> value -> value"
func (l *Singlylinkedlist) Traverse() {
	current := l.Head
	node := ""
	for current != nil {
		node += fmt.Sprint(current)
		current = current.Next
		if current != nil {
			node += " -> "
		}
	}
	fmt.Println(node)
}
