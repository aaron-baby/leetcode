package graph

type (
	ItemGraph struct {
		Children     []*ItemGraph
		ValueItem    map[string]ItemGraph
		Name         string
		Dependencies int
	}

	Graph struct {
		Nodes     []*ItemGraph
		ValueItem map[string]ItemGraph
	}
)

func (g *Graph) GetOrCreateNode(name string) ItemGraph {
	if node, ok := g.ValueItem[name]; ok {
		return node
	}
	node := ItemGraph{
		Name:      name,
		ValueItem: make(map[string]ItemGraph),
	}
	g.Nodes = append(g.Nodes, &node)
	g.ValueItem[name] = node

	return node
}

func (g *Graph) AddEdge(startName string, endName string) {
	start := g.GetOrCreateNode(startName)
	end := g.GetOrCreateNode(endName)
	start.AddNeighbor(end)
}

func (i *ItemGraph) AddNeighbor(item ItemGraph) {
	if _, ok := i.ValueItem[item.Name]; !ok {
		i.Children = append(i.Children, &item)
		i.ValueItem[item.Name] = item
		i.Dependencies++
	}
}
