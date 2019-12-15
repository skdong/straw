package task

type Task struct{
	id string
	name string
}

func (t * Task)GetId()string{
	return t.id
}

func (t *Task)GetName()string{
	return t.name
}

func CreateTask(id, name string) Task{
	return Task{id: id, name: name}
}