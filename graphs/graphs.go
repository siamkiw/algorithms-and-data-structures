package graphs

import (
	"errors"
	"fmt"
)

type Graph struct {
	AdjacencyList map[string][]string
}

func NewGraph() *Graph {
	g := &Graph{
		AdjacencyList: map[string][]string{},
	}
	return g
}

func (g *Graph) AddVertex(vertex string) error {
	if len(g.AdjacencyList[vertex]) == 0 {
		g.AdjacencyList[vertex] = []string{}
		return nil
	}

	return errors.New("vertex_already_exist")
}

func (g *Graph) AddEdge(edge1 string, edge2 string) {
	g.AdjacencyList[edge1] = append(g.AdjacencyList[edge1], edge2)
	g.AdjacencyList[edge2] = append(g.AdjacencyList[edge2], edge1)
}

func (g *Graph) RemoveEdge(vertex1 string, vertex2 string) {

	for i, v := range g.AdjacencyList[vertex1] {
		if v == vertex2 {
			g.AdjacencyList[vertex1] = removeArrayAtIndex(g.AdjacencyList[vertex1], i)
			fmt.Println("g.AdjacencyList[vertex1] :", g.AdjacencyList[vertex1])
		}
	}

	for i, v := range g.AdjacencyList[vertex2] {
		if v == vertex1 {
			g.AdjacencyList[vertex2] = removeArrayAtIndex(g.AdjacencyList[vertex2], i)
			fmt.Println("g.AdjacencyList[vertex2] :", g.AdjacencyList[vertex2])
		}
	}

}

func (g *Graph) RemoveVertex(vertex string) {
	for len(g.AdjacencyList[vertex]) > 0 {
		lastIndex := len(g.AdjacencyList[vertex]) - 1
		adjacentVertex := g.AdjacencyList[vertex][lastIndex]
		g.AdjacencyList[vertex] = removeArrayAtIndex(g.AdjacencyList[vertex], lastIndex)
		g.RemoveEdge(vertex, adjacentVertex)
	}
	delete(g.AdjacencyList, vertex)
}

func removeArrayAtIndex(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// DFS :(
//
//	{
//		"A" : ["B", "C"],
//		"B" : ["A", "D"],
//		"C" : ["A", "E"],
//		"D" : ["B", "E", "F"],
//		"E" : ["C", "D", "F"],
//		"F" : ["D", "E"],
//		"G" : []
//	}
func (g *Graph) DepthFirstSearchRecursive(vertex string) []string {
	result := []string{}
	visited := map[string]bool{}

	var dfs func(vertex string)
	dfs = func(vertex string) {
		// todo add stoping point
		visited[vertex] = true
		result = append(result, vertex)

		for _, v := range g.AdjacencyList[vertex] {
			if !visited[v] {
				dfs(v)
			}
		}
	}

	dfs(vertex)
	return result
}

//	{
//		"A" : ["B", "C"],
//		"B" : ["A", "D"],
//		"C" : ["A", "E"],
//		"D" : ["B", "E", "F"],
//		"E" : ["C", "D", "F"],
//		"F" : ["D", "E"],
//		"G" : []
//	}
func (g *Graph) DepthFirstSearchIterative(vertex string) []string {
	stack := []string{vertex}
	result := []string{}
	visted := map[string]bool{
		vertex: true,
	}

	for len(stack) > 0 {
		currentVertex := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, currentVertex)

		for _, v := range g.AdjacencyList[currentVertex] {
			if !visted[v] {
				visted[v] = true
				stack = append(stack, v)
			}
		}
	}

	return result
}

func (g *Graph) BreadthFirstSearchIterative(vertex string) []string {
	queue := []string{vertex}
	result := []string{}
	visited := map[string]bool{
		vertex: true,
	}

	for len(queue) > 0 {
		// is there other way to do this ?
		currentVertex := queue[0] //shift
		queue = queue[1:]         //shift
		result = append(result, currentVertex)

		for _, v := range g.AdjacencyList[currentVertex] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, v)
			}
		}
	}

	return result
}
