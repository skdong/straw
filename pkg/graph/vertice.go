package graph

type Vertice struct {
	Id string
	Name string
	Info string
}

func (v *Vertice)GetId()string{
	return v.Id
}

func CreateVertice(id, name string) Vertice{
	point := Vertice{Id:id, Name:name}
	return point
}