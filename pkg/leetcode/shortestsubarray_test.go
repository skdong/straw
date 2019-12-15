package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestsubarray(t *testing.T) {
	testCases := [][]int{
		{3, 4, -1, 1, 2, 1},
		{3, 3, 2, -1, 2},
	}

	for i := range testCases {
		assert.Equal(t, testCases[i][0],
			shortestSubArray(testCases[i][2:], testCases[i][1]))
	}
}
