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
	if contains(g.vertices, key) {
		err := fmt.Errorf("Vertext %v not added because it is an existing key", key)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: key})
	}
}

func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %d : ", v.key)
		for _, k := range v.adjency {
			fmt.Printf(" %d ", k.key)
		}
	}
}

//AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to int) {
	//get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("Invalid Edge( %v --> %v )", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjency, to) {
		err := fmt.Errorf("Invalid Edge( %v --> %v )", from, to)
		fmt.Println(err.Error())
	} else {
		//add edge
		fromVertex.adjency = append(fromVertex.adjency, toVertex)
	}
}

//getVertex returns a pointer to the vertex with a key integer
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if v.key == k {
			return g.vertices[i]
		}
	}
	return nil
}

func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

func main() {
	graph := &Graph{}

	for i := 0; i < 5; i++ {
		graph.AddVertex(i)

	}

	// graph.AddVertex(0)
	// graph.AddVertex(0)
	// graph.AddVertex(0)

	// fmt.Println(graph.vertices[0])

	graph.AddEdge(1, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(6, 2)
	graph.Print()
}
