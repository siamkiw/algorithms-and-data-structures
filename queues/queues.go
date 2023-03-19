package queues

// add into the end on list and return move the head

type Node struct {
	value string
	next  *Node
}

type Queue struct {
	first *Node
	last  *Node
	size  int
}

func (q *Queue) Enqueue(message string) int {
	new := &Node{
		value: message,
	}
	if q.size == 0 {
		q.first = new
		q.last = new
	} else {
		q.last.next = new
		q.last = new
	}
	q.size++
	return q.size
}

func (q *Queue) Dequeue() *Node {
	if q.first == nil {
		return &Node{}
	}

	temp := q.first
	if q.size == 1 {
		q.last = nil
	}
	q.first = q.first.next
	q.size--

	return temp
}
