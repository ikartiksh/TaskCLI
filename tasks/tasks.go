package tasks

type task struct {
	ID          string
	Name 		string
	Status      string
}

var TaskStore = make(map[string]*task)

func NewTask(id, name string) *task {
	return &task {
		ID: id,
		Name: name,
		Status: "Pending",
	}
}

func AddTasks(ts *task) {
	TaskStore[ts.ID] = ts
}

func ListAllTasks() []*task {
	var tasks []*task
	for _, t := range TaskStore {
		tasks = append (tasks,t)
	}
	return tasks
}