package main

// singly "github.com/siamkiw/algo/singlyLinkedLists"

// "github.com/siamkiw/algo/queues"

import (
	"fmt"
	"strings"

	// "github.com/siamkiw/algo/bst"
	// "github.com/siamkiw/algo/recursion"
	// "github.com/siamkiw/algo/heaps"
	"github.com/siamkiw/algo/graphs"
	"github.com/siamkiw/algo/heaps/minHeaps"
	"github.com/siamkiw/algo/weightedGraphs"
)

func main() {
	// isSame := frequency.Same([]int{1, 2, 3, 2, 5, 0}, []int{9, 1, 4, 4, 25, 0})
	// fmt.Println("isSame :", isSame)

	// isSame := frequency.Anagrams("js-algorithms-and-data-structures", "js-algorithms-structures-data-and")
	// fmt.Println("isSame :", isSame)

	// isSumZero, _ := pointers.SumZero([]int{-4, -3, 0, 3, 5, 10})
	// fmt.Println(isSumZero)

	// unique := pointers.CounterUniqueValues([]int{1, 1, 3, 4, 5, 5, 6, 7})
	// fmt.Println("unique :", unique)

	// maxSum, _ := pointers.MaxSumArray([]int{2, 6, 9, 2, 1, 8, 5, 6, 3}, 3)
	// fmt.Println(maxSum)
	// res := Solution([]int{1, 3, 6, 4, 1, 2})
	// fmt.Println("res :", res)

	// list := singly.Singlylinkedlist{}
	// list.Push("Hi")
	// list.Push("You")
	// list.Push("again")
	// list.Push("reverse")
	// list.Traverse()
	// fmt.Println("----")
	// list.Reverse()

	// list.Traverse()
	// len := list.Len
	// fmt.Println("len :", len)

	// s := &stack.Stack{}
	// s.Push("1")
	// s.Push("2")
	// s.Push("3")

	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())
	// fmt.Println(s.Pop())

	// q := queues.Queue{}
	// q.Enqueue("1")
	// q.Enqueue("2")
	// q.Enqueue("3")

	// fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())
	// fmt.Println(q.Dequeue())
	// q.Enqueue("4")
	// fmt.Println(q.Dequeue())

	// fmt.Println(recursion.SumRange(3))
	// fmt.Println(recursion.CollectOddValues([]int{1, 2, 3, 4, 5}))

	// b := &bst.BinarySearchTree{}
	// b.Insert(10)
	// b.Insert(6)
	// b.Insert(15)
	// b.Insert(3)
	// b.Insert(8)
	// b.Insert(20)

	// // fmt.Println(b.BFS())
	// fmt.Println(b.DFSInOrder())

	// res := romanToInt("MCMXCIV")
	// fmt.Println("res :", res)
	// [1,2,4]
	// list := &listLeet.ListNode{
	// 	Val: 1,
	// }
	// list2 := &listLeet.ListNode{
	// 	Val: 2,
	// }
	// list3 := &listLeet.ListNode{
	// 	Val: 4,
	// }

	// list.Next = list2
	// list2.Next = list3

	// list_1 := &listLeet.ListNode{
	// 	Val: 1,
	// }
	// list_2 := &listLeet.ListNode{
	// 	Val: 2,
	// }
	// list_3 := &listLeet.ListNode{
	// 	Val: 4,
	// }
	// list_1.Next = list_2
	// list_2.Next = list_3

	// res := listLeet.MergeTwoLists(list, list_1)

	// current := res
	// for current.Next != nil {
	// 	fmt.Println(current.Val)
	// 	current = current.Next
	// }

	// mh := &heaps.MaxHeap{}
	// mh.Value = []int{}
	// mh.Insert(41)
	// mh.Insert(39)
	// mh.Insert(33)
	// mh.Insert(18)
	// mh.Insert(27)
	// mh.Insert(12)
	// mh.Insert(55)
	// // mh.Insert(20)
	// r, _ := mh.Remove()
	// fmt.Println("r :", r)
	// fmt.Println("mh :", mh.Value)
	// fmt.Println("----")
	// r, _ = mh.Remove()
	// fmt.Println("r :", r)
	// fmt.Println("mh :", mh.Value)
	// r, _ = mh.Remove()
	// fmt.Println("r :", r)
	// fmt.Println("mh :", mh.Value)
	// r, _ = mh.Remove()
	// fmt.Println("r :", r)
	// fmt.Println("mh :", mh.Value)
	// r, _ = mh.Remove()
	// fmt.Println("r :", r)
	// fmt.Println("mh :", mh.Value)
	// r, _ = mh.Remove()
	// fmt.Println("r :", r)
	// fmt.Println("mh :", mh.Value)
	// r, _ = mh.Remove()
	// fmt.Println("r :", r)
	// fmt.Println("mh :", mh.Value)
	// r, _ = mh.Remove()
	// fmt.Println("r :", r)
	// fmt.Println("mh :", mh.Value)

	g := graphs.NewGraph()
	// g.AddVertex("hello world")
	// g.AddVertex("hello kiw")
	// g.AddVertex("test1")
	// g.AddVertex("test2")
	// g.AddEdge("hello world", "hello kiw")
	// g.AddEdge("hello world", "test1")
	// g.AddEdge("test2", "test1")
	// g.AddEdge("hello kiw", "test1")

	// g.RemoveVertex("hello world")
	// fmt.Println("before remove edge :", g.AdjacencyList)

	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddVertex("E")
	g.AddVertex("F")
	g.AddVertex("G")

	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "E")
	g.AddEdge("D", "E")
	g.AddEdge("D", "F")
	g.AddEdge("E", "F")

	// g.RemoveEdge("hello world", "hello kiw")
	dfsR := g.DepthFirstSearchRecursive("A")
	fmt.Println("dfsR :", dfsR)

	dfsI := g.DepthFirstSearchIterative("A")
	fmt.Println("dfsI :", dfsI)

	bfsI := g.BreadthFirstSearchIterative("A")
	fmt.Println("bfsI :", bfsI)

	fmt.Println("after remove edge :", g.AdjacencyList)

	fmt.Println("------minHeaps------")
	minHeaps.Exp()
	fmt.Println("------minHeaps------")

	fmt.Println("-----WeightedGraph------")
	weightedGraphs.Exp()
	fmt.Println("-----WeightedGraph------")

}

func romanToInt(s string) int {

	type data struct {
		value    int
		subtract map[string]bool
	}

	roman := map[string]data{
		"I": data{
			value: 1,
			subtract: map[string]bool{
				"V": true,
				"X": true,
			},
		},
		"V": data{
			value:    5,
			subtract: map[string]bool{},
		},
		"X": data{
			value: 10,
			subtract: map[string]bool{
				"L": true,
				"C": true,
			},
		},
		"L": data{
			value:    50,
			subtract: map[string]bool{},
		},
		"C": data{
			value: 100,
			subtract: map[string]bool{
				"D": true,
				"M": true,
			},
		},
		"D": data{
			value:    500,
			subtract: map[string]bool{},
		},
		"M": data{
			value:    1000,
			subtract: map[string]bool{},
		},
	}

	result := 0

	for i := 0; i < len(s); i++ {
		fmt.Println("i :", i)
		if i+1 == len(s) {
			result += roman[strings.ToUpper(string(s[i]))].value
			return result
		}

		current := roman[strings.ToUpper(string(s[i]))]

		fmt.Println("current :", string(s[i]), current)

		if current.value == 0 {
			return 0
		}

		next := string(s[i+1])
		fmt.Println("next :", next)
		isSubtract := current.subtract[next]
		if isSubtract {
			tempResult := roman[next].value - current.value
			fmt.Println("tempResult :", tempResult)
			result += tempResult
			i++
		} else {
			result += current.value
		}

		fmt.Println("result :", result)
		fmt.Println("-----")

	}

	return result

}
