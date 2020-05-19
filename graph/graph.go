package graph

type (
	ItemGraph struct {
		children     []*ItemGraph
		valueItem    map[string]ItemGraph
		Name         string
		dependencies int
	}

	Graph struct {
		nodes     []ItemGraph
		valueItem map[string]ItemGraph
	}
)

func (g *Graph) getOrCreateNode(name string) ItemGraph {
	if node, ok := g.valueItem[name]; ok {
		return node
	}
	node := ItemGraph{Name: name}
	g.nodes = append(g.nodes, node)
	g.valueItem[name] = node

	return node
}

func (g *Graph) addEdge(startName string, endName string) {
	start := g.getOrCreateNode(startName)
	end := g.getOrCreateNode(endName)
	start.AddNeighbor(end)
}

func (i *ItemGraph) AddNeighbor(item ItemGraph) {
	if _, ok := i.valueItem[item.Name]; !ok {
		i.children = append(i.children, &item)
		i.valueItem[item.Name] = item
		i.dependencies++
	}
}
