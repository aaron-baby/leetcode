package main

import "gitlab.com/aaw/leetcode/graph"

func buildGraph(projects []string, dependencies [][]string) graph.Graph {
	g := graph.Graph{
		ValueItem: make(map[string]*graph.Vertex),
	}
	for _, name := range projects {
		g.GetOrCreateNode(name)
	}
	for _, d := range dependencies {
		g.AddEdge(d[0], d[1])
	}

	return g
}

func OrderProjects(nodes []*graph.Vertex) []*graph.Vertex {
	// Create a queue and enqueue all vertices with
	// indegree 0
	var q, order []*graph.Vertex
	for _, n := range nodes {
		if n.InDegree == 0 {
			q = append(q, n)
		}
	}
	for len(q) > 0 {
		d, q := q[len(q)-1], q[:len(q)-1]
		order = append(order, d)
		for _, child := range d.Adjacencies {
			child.InDegree--
			if child.InDegree == 0 {
				q = append(q, child)
			}
		}
	}

	return order
}
