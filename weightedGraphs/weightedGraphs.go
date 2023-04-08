package weightedgraphs

import "errors"

type WeightedGraph struct {
	List map[string][]Node
}

type Node struct {
	Val    string
	Weight int
}

func (w *WeightedGraph) AddVertex(vertex string) error {
	if len(w.List[vertex]) == 0 {
		w.List[vertex] = []Node{}
	}

	return errors.New("vertex_already_exist")
}

func (w *WeightedGraph) AddEdge(vertex1 string, vertex2 string, weight int) {
	w.List[vertex1] = append(w.List[vertex1], Node{Val: vertex1, Weight: weight})
	w.List[vertex2] = append(w.List[vertex2], Node{Val: vertex2, Weight: weight})
}
