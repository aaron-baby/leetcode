package graph

import (
	"fmt"
	"strconv"
)

type (
	Vertex struct {
		Adjacencies []*Vertex
		Name        string
		// The number of edges leaving a vertex is its out-degree,
		// and the number of edges entering is the in-degree.
		InDegree int
	}

	Graph struct {
		Vertices  []*Vertex
		ValueItem map[string]*Vertex
	}
)

func (g *Graph) GetOrCreateNode(name string) *Vertex {
	if node, ok := g.ValueItem[name]; ok {
		return node
	}
	node := Vertex{
		Name: name,
	}
	g.Vertices = append(g.Vertices, &node)
	g.ValueItem[name] = &node

	return &node
}

// AddEdge adds an edge to the graph
func (g *Graph) String() {
	s := ""
	for i := 0; i < len(g.Vertices); i++ {
		s += g.Vertices[i].Name + "("+ strconv.Itoa(g.Vertices[i].InDegree) +") -> "
		near := g.Vertices[i].Adjacencies
		for j := 0; j < len(near); j++ {
			s += near[j].Name + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func (g *Graph) AddEdge(startName string, endName string) {
	start := g.GetOrCreateNode(startName)
	end := g.GetOrCreateNode(endName)
	start.AddNeighbor(end)
}

func (i *Vertex) AddNeighbor(item *Vertex) {
	i.Adjacencies = append(i.Adjacencies, item)
	item.InDegree++
}
