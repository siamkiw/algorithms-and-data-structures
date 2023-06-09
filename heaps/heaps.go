package heaps

import "errors"

// for any inedx of an array n..
// left child is 2n+1
// right child is 2n+2

// for any index of an array n..
// parent is (n-1)/2 floored

// [41, 39, 33, 18, 27, 12]
type MaxHeap struct {
	Value []int
}

func (mh *MaxHeap) Insert(num int) {

	mh.Value = append(mh.Value, num)
	mh.bubbleUp()

}

// [55, 39, 41, 18, 27, 12, 33, 20]
//
//	7
//
// currentIdx = 7
// parentIdx = 3
func (mh *MaxHeap) bubbleUp() {
	currentIdx := len(mh.Value) - 1
	current := mh.Value[currentIdx]

	for currentIdx > 0 {
		num := float64((currentIdx - 1) / 2)
		parentIdx := int(num)
		parent := mh.Value[parentIdx]
		if current > parent {
			mh.Value[parentIdx] = current
			mh.Value[currentIdx] = parent
			currentIdx = parentIdx
		} else {
			break
		}
	}
}

func (mh *MaxHeap) Remove() (int, error) {
	if len(mh.Value) == 1 {
		max := mh.Value[0]
		mh.Value = []int{}
		return max, nil
	} else if len(mh.Value) > 0 {
		max := mh.Value[0]
		mh.Value[0] = mh.Value[len(mh.Value)-1]
		mh.Value = mh.Value[:len(mh.Value)-1]

		mh.bubbleDown()

		return max, nil
	}

	return -1, errors.New("unexpected_error")

}

func (mh *MaxHeap) bubbleDown() {
	idx := 0
	lenght := len(mh.Value)
	element := mh.Value[0]

	for {
		leftIdx := 2*idx + 1
		rightIdx := 2*idx + 2
		var left int
		var right int
		swap := -1
		if leftIdx < lenght { // check that length is not out of bound
			left = mh.Value[leftIdx]
			if left > element {
				swap = leftIdx
			}
		}
		// if left is less than element(which did not change the swap value) and right is more than element
		// or if left is more than element (which changed th swap value) and right's value is more than left's value
		// swap to the left
		if rightIdx < lenght {
			right = mh.Value[rightIdx]
			if (swap == -1 && right > element) ||
				(swap != -1 && right > left) {
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
