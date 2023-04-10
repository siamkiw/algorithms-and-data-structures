package minHeaps

import (
	"errors"
	"fmt"
)

// for any inedx of an array n..
// left child is 2n+1
// right child is 2n+2

// for any index of an array n..
// parent is (n-1)/2 floored

type MinHeap struct {
	Value []Node
}

type Node struct {
	Val      string
	Priority int
}

func NewPriorityQueue() *MinHeap {
	return &MinHeap{
		Value: []Node{},
	}
}

func (mh *MinHeap) Enqueue(val string, pri int) {
	node := Node{Val: val, Priority: pri}
	mh.Value = append(mh.Value, node)
	mh.bubbleUp()
}

func (mh *MinHeap) bubbleUp() {
	currentIdx := len(mh.Value) - 1
	current := mh.Value[currentIdx]

	for currentIdx > 0 {
		num := float64((currentIdx - 1) / 2)
		parentIdx := int(num)
		parent := mh.Value[parentIdx]
		if current.Priority < parent.Priority {
			mh.Value[parentIdx] = current
			mh.Value[currentIdx] = parent
			currentIdx = parentIdx
		} else {
			break
		}
	}
}

func (mh *MinHeap) Dequeue() (Node, error) {
	if len(mh.Value) == 1 {
		max := mh.Value[0]
		mh.Value = []Node{}
		return max, nil
	} else if len(mh.Value) > 0 {
		max := mh.Value[0]
		mh.Value[0] = mh.Value[len(mh.Value)-1]
		mh.Value = mh.Value[:len(mh.Value)-1]

		mh.bubbleDown()

		return max, nil
	}

	return Node{}, errors.New("unexpected_error")

}

func (mh *MinHeap) bubbleDown() {
	idx := 0
	lenght := len(mh.Value)
	element := mh.Value[0]

	for {
		leftIdx := 2*idx + 1
		rightIdx := 2*idx + 2
		var left Node
		var right Node
		swap := -1
		if leftIdx < lenght { // check that length is not out of bound
			left = mh.Value[leftIdx]
			if left.Priority < element.Priority {
				swap = leftIdx
			}
		}
		// if left is less than element(which did not change the swap value) and right is more than element
		// or if left is more than element (which changed th swap value) and right's value is more than left's value
		// swap to the left
		if rightIdx < lenght {
			right = mh.Value[rightIdx]
			if (swap == -1 && right.Priority < element.Priority) ||
				(swap != -1 && right.Priority < left.Priority) {
				swap = rightIdx
			}
		}

		if swap == -1 {
			break // if swap is not haapen break the loop
		}

		// swap the value between the greatest child
		mh.Value[idx] = mh.Value[swap] // set the element to the swap value
		mh.Value[swap] = element       // set swap value to element
		idx = swap                     // next loop start at the swaped index
	}
}

func Exp() {
	mh := MinHeap{}
	mh.Enqueue("commom cold", 1)
	mh.Enqueue("hunshot wound", 5)
	mh.Enqueue("high fever", 2)

	fmt.Println("mh :", mh)
	res, _ := mh.Dequeue()
	fmt.Println("printing Dequeue 1:", res)
	res, _ = mh.Dequeue()
	fmt.Println("printing Dequeue 2:", res)
	res, _ = mh.Dequeue()
	fmt.Println("printing Dequeue 2:", res)

	fmt.Println("mh :", mh)
}
