package main

import (
	"fmt"
	"testing"
)

func TestOrderProjects(t *testing.T) {
	projects := []string{"a", "b", "c", "d", "e", "f"}
	dependencies := [][]string{
		{"a", "d"},
		{"f", "b"},
		{"b", "d"},
		{"f", "a"},
		{"d", "c"},
	}
	g := buildGraph(projects, dependencies)
	g.String()
	o := OrderProjects(g.Vertices)
	fmt.Printf("%+v", o)
	//var orderedNames []string
	//for _, p := range o {
	//	orderedNames = append(orderedNames, p.Name)
	//}
	//want := []string{"f", "e", "a", "b", "d", "c"}
	//if !cmp.Equal(orderedNames, want) {
	//	t.Errorf("got %v want %v given", orderedNames, want)
	//}
}
