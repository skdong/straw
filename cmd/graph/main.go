package main

import (
	"github.com/skdong/straw/pkg/graph"
	"fmt"
)

func main(){
	node := graph.CreateNode()
	fmt.Println(node.GetId())
}