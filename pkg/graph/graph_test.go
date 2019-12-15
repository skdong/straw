package graph

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateGraphVertices(t * testing.T){
	graph := CreateGraph()
	vs := [][]string{[]string{"1", "test1"},
					[]string{"2", "test2"}}
	for i := range(vs){
		v := CreateVertice(vs[i][0], vs[i][1])
		graph.Vertices = append(graph.Vertices, v)
	}

	for i := range(vs){
		assert.Equal(t, vs[i][0], graph.Vertices[i].Id)
		assert.Equal(t, vs[i][1], graph.Vertices[i].Name)
	}
}