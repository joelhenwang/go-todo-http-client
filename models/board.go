package models

type Board struct {
	Id    string           `json:"id"`
	Title string           `json:"title"`
	Tasks map[string]*Task `json:"tasks"`
}

func NewBoard(title string) *Board {
	return &Board{
		Id:    title,
		Title: title,
		Tasks: map[string]*Task{},
	}
}

func (board *Board) AddTask(task Task) {
	board.Tasks[task.Id] = &task
}

func (board *Board) MoveTask(id string, newStatus string) {
	task := board.Tasks[id]

	if task.Status != newStatus {
		task.EditStatus(newStatus)
	}
}

func (board *Board) DeleteTask(id string) {
	delete(board.Tasks, id)
}
