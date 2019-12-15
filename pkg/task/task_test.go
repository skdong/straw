package task

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T){
	input := []string{"1", "testa"}
	task := CreateTask(input[0], input[1])
	assert.Equal(t, input[0], task.GetId())
	assert.Equal(t, input[1], task.GetName())
}