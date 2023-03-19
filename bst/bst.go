package bst

import (
	"errors"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}

func (bst *BinarySearchTree) Insert(value int) (*Node, error) {
	new := &Node{
		Value: value,
	}

	if bst.Root == nil {
		bst.Root = new
		return bst.Root, nil
	} else {
		current := bst.Root

		if value == current.Value {
			return current, errors.New("incorrect_value")
		}

		for {

			if value > current.Value {
				if current.Right == nil {
					current.Right = new
					return current, nil
				} else {
					current = current.Right
				}
			} else if value < current.Value {
				if current.Left == nil {
					current.Left = new
					return current, nil
				} else {
					current = current.Left
				}
			}

		}
	}
}

func (bst *BinarySearchTree) Find(value int) bool {
	if bst.Root == nil {
		return false
	} else {
		current := bst.Root
		for value != current.Value {
			if value > current.Value {
				if current.Right != nil {
					current = current.Right
				} else {
					return false
				}
			} else if value < current.Value {
				if current.Left != nil {
					current = current.Left
				} else {
					return false
				}
			}
		}

		return true
	}

}

func (bst *BinarySearchTree) BFS() []int {
	if bst.Root == nil {
		return []int{}
	}

	node := bst.Root
	data := []int{}
	queue := []*Node{}

	queue = append(queue, node)
	for len(queue) > 0 {

		node := queue[0]  //shift
		queue = queue[1:] //shift

		data = append(data, node.Value)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return data
}

// you can recreate tree with method
func (bst *BinarySearchTree) DFSPreOrder() []int {

	data := []int{}

	var traverse func(node *Node)
	traverse = func(node *Node) {
		data = append(data, node.Value)
		if node.Left != nil {
			traverse(node.Left)
		}

		if node.Right != nil {
			traverse(node.Right)
		}
	}
	traverse(bst.Root)

	return data
}

func (bst *BinarySearchTree) DFSPostOrder() []int {

	data := []int{}
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node.Left != nil {
			traverse(node.Left)
		}

		if node.Right != nil {
			traverse(node.Right)
		}

		data = append(data, node.Value)
	}

	traverse(bst.Root)

	return data

}

func (bst *BinarySearchTree) DFSInOrder() []int {

	data := []int{}
	var traverse func(node *Node)
	traverse = func(node *Node) {
		if node.Left != nil {
			traverse(node.Left)
		}

		data = append(data, node.Value)

		if node.Right != nil {
			traverse(node.Right)
		}

	}

	traverse(bst.Root)

	return data

}
