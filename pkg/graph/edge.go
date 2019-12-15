package graph

type Edge struct{
	Src *Vertice
	Dest *Vertice
}

func CreateEdge(src, dest *Vertice)Edge{
	return Edge{Src: src, Dest: dest}
}