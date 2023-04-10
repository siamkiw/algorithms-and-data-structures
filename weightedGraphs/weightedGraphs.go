package weightedGraphs

import (
	"errors"
	"fmt"

	"github.com/siamkiw/algo/heaps/minHeaps"
)

type WeightedGraph struct {
	List map[string][]Node
}

type Node struct {
	Val    string
	Weight int
}

func NewWeightedGraph() *WeightedGraph {
	w := &WeightedGraph{
		List: map[string][]Node{},
	}

	return w
}

func (w *WeightedGraph) AddVertex(vertex string) error {
	if len(w.List[vertex]) == 0 {
		w.List[vertex] = []Node{}
	}

	return errors.New("vertex_already_exist")
}

func (w *WeightedGraph) AddEdge(vertex1 string, vertex2 string, weight int) error {
	if weight > 100 {
		return errors.New("weight_should_less_than_100")
	}
	w.List[vertex1] = append(w.List[vertex1], Node{Val: vertex2, Weight: weight})
	w.List[vertex2] = append(w.List[vertex2], Node{Val: vertex1, Weight: weight})

	return nil
}

func (w *WeightedGraph) Dijkstra(start string, finish string) []string {
	nodes := minHeaps.NewPriorityQueue()
	dist := map[string]int{}
	prevs := map[string]string{}
	var smallest string
	path := []string{}

	for vertex, _ := range w.List {
		if vertex == start {
			dist[vertex] = 0 // not need to do this zero value of map[string]int{} is 0
			nodes.Enqueue(vertex, 0)
		} else {
			dist[vertex] = 100 // we only allow max value to be 100
			nodes.Enqueue(vertex, 100)
		}
		prevs[vertex] = "" // not need to do this zero value of map[string]string{} is ""
	}

	// as long as there is something to visit

	for len(nodes.Value) > 0 {
		val, _ := nodes.Dequeue()
		smallest = val.Val

		if smallest == finish {
			// we are done
			// build up path return at the end
			for prevs[smallest] != "" {
				path = append(path, smallest)
				smallest = prevs[smallest]
			}
			break
		}
		if dist[smallest] != 100 {
			for neighbor, _ := range w.List[smallest] {
				// find neighboring node
				NextNode := w.List[smallest][neighbor]
				// calcuidate new distance to neighboring node
				candidate := dist[smallest] + NextNode.Weight
				if candidate < dist[NextNode.Val] {
					// updateing new smallest distance to neighbor
					dist[NextNode.Val] = candidate
					//  updateing previous How we got to neighbor
					prevs[NextNode.Val] = smallest
					// enqueue in priority quwuw with new priority
					nodes.Enqueue(NextNode.Val, candidate)
				}
			}
		}

	}

	// add A and reverse the result
	path = append(path, start)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}

func Exp() {
	w := NewWeightedGraph()
	w.AddVertex("A")
	w.AddVertex("B")
	w.AddVertex("C")
	w.AddVertex("D")
	w.AddVertex("E")
	w.AddVertex("F")

	w.AddEdge("A", "B", 4)
	w.AddEdge("A", "C", 2)
	w.AddEdge("B", "E", 3)
	w.AddEdge("C", "D", 2)
	w.AddEdge("C", "F", 4)
	w.AddEdge("D", "E", 3)
	w.AddEdge("D", "F", 1)
	w.AddEdge("E", "F", 1)

	for k, v := range w.List {
		fmt.Println(k, v)
	}

	res := w.Dijkstra("A", "E")
	fmt.Println("res :", res)
}
