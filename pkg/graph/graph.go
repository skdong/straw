package graph

type Graph struct{
	Vertices []Vertice
	Edges []Edge
}

func CreateGraph()Graph{
	return Graph{}
}