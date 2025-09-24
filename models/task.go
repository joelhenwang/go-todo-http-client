package models

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func NewTask(title string, desc string) *Task {
	return &Task{
		Id:          title,
		Title:       title,
		Description: desc,
		Status:      "To Do",
	}
}

func (task *Task) EditTitle(title string) {
	task.Title = title
}

func (task *Task) EditDescription(desc string) {
	task.Description = desc
}

func (task *Task) EditStatus(newStatus string) {
	task.Status = newStatus
}
