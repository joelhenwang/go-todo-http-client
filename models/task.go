package models

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priority"`
	BoardId     string `json:"boardId"`
}

func NewTask(title string, desc string, priority string, boardId string) *Task {
	return &Task{
		Id:          title,
		Title:       title,
		Description: desc,
		Status:      "To Do",
		Priority:    priority,
		BoardId:     boardId,
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

func (task *Task) EditPriority(newPriority string) {
	task.Priority = newPriority
}
