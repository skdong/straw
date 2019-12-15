package graph

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreatVertice(t *testing.T){
	input := map[string]string{"id":"123", "name": "test"}

	vertice := CreateVertice(input["id"], input["name"])

	expected_id := input["id"]
	expected_name := input["name"]
	assert.Equal(t, expected_id, vertice.Id)
	assert.Equal(t, expected_name, vertice.Name)
}
