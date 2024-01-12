// Graphs are used to solve many real-life problems.
// Graphs are used to represent networks.
// The networks may include paths in a city or telephone network or circuit network.
// Graphs are also used in social networks like linkedIn, Facebook.
// For example, in Facebook, each person is represented with a vertex(or node).
// Each node is a structure and contains information like person id, name, gender, locale etc.

package main

import (
	"fmt"
)

type Graph struct {
	graph map[string][]string
}

func NewGraph() Graph {
	return Graph{
		graph: make(map[string][]string),
	}
}

func (g Graph) AddVertex(vertex string) {
	if _, ok := g.graph[vertex]; !ok {
		g.graph[vertex] = []string{}
	}
}

func (g Graph) AddEdge(vertex1, vertex2 string) {
	if _, ok := g.graph[vertex1]; !ok {
		g.AddVertex(vertex1)
	}
	if _, ok := g.graph[vertex2]; !ok {
		g.AddVertex(vertex2)
	}
	g.graph[vertex1] = append(g.graph[vertex1], vertex2)
	g.graph[vertex2] = append(g.graph[vertex2], vertex1) // For undirected graph, add both ways.
}

func (g Graph) GetNeighbors(vertex string) []string {
	if neighbors, ok := g.graph[vertex]; ok {
		return neighbors
	}
	return nil
}

func main() {
	g := NewGraph()
	g.AddVertex("A")
	g.AddVertex("B")
	g.AddVertex("C")
	g.AddVertex("D")
	g.AddEdge("A", "B")
	g.AddEdge("A", "C")
	g.AddEdge("B", "D")
	g.AddEdge("C", "D")

	fmt.Println("Graph Representation:")
	for vertex, neighbors := range g.graph {
		fmt.Printf("%s: %v\\n", vertex, neighbors)
	}

	fmt.Println("Neighbors of 'A':", g.GetNeighbors("A"))
	fmt.Println("Neighbors of 'B':", g.GetNeighbors("B"))
	fmt.Println("Neighbors of 'C':", g.GetNeighbors("C"))
	fmt.Println("Neighbors of 'D':", g.GetNeighbors("D"))
}
