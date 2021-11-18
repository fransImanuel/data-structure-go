package main

import "fmt"

type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	key     int
	adjency []*Vertex
}

func (g *Graph) AddVertex(key int) {
	newVertex := &Vertex{key: key}
	g.vertices = append(g.vertices, newVertex)
}

func main() {
	graph := &Graph{}

	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	fmt.Println(graph.vertices[0])
}
