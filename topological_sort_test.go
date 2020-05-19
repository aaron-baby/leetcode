package main

import (
	"github.com/google/go-cmp/cmp"
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
	var orderedNames []string
	for _, p := range OrderProjects(g.Nodes) {
		orderedNames = append(orderedNames, p.Name)
	}
	want := []string{"f", "e", "a", "b", "d", "c"}
	if !cmp.Equal(orderedNames, want) {
		t.Errorf("got %v want %v given", orderedNames, want)
	}
}
