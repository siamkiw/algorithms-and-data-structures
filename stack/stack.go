package stack

import "errors"

type Node struct {
	value string
	next  *Node
}

type Stack struct {
	first *Node
	last  *Node
	size  int
}

func (s *Stack) Push(message string) int {

	new := &Node{
		value: message,
	}
	if s.first == nil {
		s.first = new
		s.last = new
	} else {
		old := s.first
		s.first = new
		s.first.next = old
	}
	s.size++
	return s.size
}

func (s *Stack) Pop() (*Node, error) {
	if s.first == nil {
		return &Node{}, errors.New("empty_list")
	}

	temp := s.first
	if s.size == 0 {
		s.first = &Node{}
		s.last = &Node{}
		return &Node{}, nil
	}

	s.first = s.first.next
	s.size--
	return temp, nil
}
