package graph

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateEdge(t * testing.T){
	vertices := [][]string{[]string{"1", "test1"}, []string{"2", "test2"}}
	src := CreateVertice(vertices[0][0], vertices[0][1])
	dest := CreateVertice(vertices[1][0], vertices[1][1])
	edge := CreateEdge(&src, &dest)
	assert.Equal(t, vertices[0][0], edge.Src.Id)
	assert.Equal(t, vertices[0][1], edge.Src.Name)
	assert.Equal(t, vertices[1][0], edge.Dest.Id)
	assert.Equal(t, vertices[1][1], edge.Dest.Name)
}