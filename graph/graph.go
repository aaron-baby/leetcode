package graph

type (
	ItemGraph struct {
		children     []*ItemGraph
		valueItem    map[string]ItemGraph
		Name         string
		dependencies int
	}
)

func (i *ItemGraph) AddNeighbor(item ItemGraph) {
	if _, ok := i.valueItem[item.Name]; !ok {
		i.children = append(i.children, &item)
		i.valueItem[item.Name] = item
		i.dependencies++
	}
}
