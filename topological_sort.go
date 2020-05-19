package main

import "gitlab.com/aaw/leetcode/graph"

func buildGraph(projects []string, dependencies [][]string) graph.Graph {
	g := graph.Graph{
		ValueItem: make(map[string]graph.ItemGraph),
	}
	for _, name := range projects {
		g.GetOrCreateNode(name)
	}
	for _, d := range dependencies {
		g.AddEdge(d[0], d[1])
	}

	return g
}

func OrderProjects(nodes []*graph.ItemGraph) []*graph.ItemGraph {
	order := make([]*graph.ItemGraph, len(nodes))
	endOffset := addNonDependent(order, nodes, 0)

	toBeProcessed := 0
	for toBeProcessed < len(order) {
		current := order[toBeProcessed]

		// circular dependency
		if current == nil {
			return order
		}
		// decrement dependencies
		for _, child := range current.Children {
			child.Dependencies--
		}
		endOffset = addNonDependent(order, current.Children, endOffset)
		toBeProcessed++
	}

	return order
}

func addNonDependent(order []*graph.ItemGraph, nodes []*graph.ItemGraph, offset int) int {
	for _, n := range nodes {
		if n.Dependencies == 0 {
			order[offset] = n
			offset++
		}
	}

	return offset
}
