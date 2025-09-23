package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joelhenwang/go-todo-http-client/models"
)

type Handler struct {
	url     string
	handler func(http.ResponseWriter, *http.Request)
}

type apiError struct {
	Error string `json:"error"`
}

func GetHandlers(board *models.Board) []Handler {
	return []Handler{
		GetTasks(board),
		GetTaskById(board),
		PostNewTask(board),
	}
}

// TODO Change in memory Board to get Tasks by BoardID
func GetTasks(board *models.Board) Handler {
	return Handler{
		url: "GET /tasks",
		handler: func(w http.ResponseWriter, r *http.Request) {
			tasks := board.Tasks
			writeJson(w, http.StatusOK, tasks)
		},
	}
}

func GetTaskById(board *models.Board) Handler {
	return Handler{
		url: "GET /task/{id}",
		handler: func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")
			task, ok := board.Tasks[id]

			if !ok {
				writeError(w, http.StatusNotFound, fmt.Sprintf("Task with id (%s) not found...", id))
				return
			}

			writeJson(w, 200, task)
		},
	}
}

func PostNewTask(board *models.Board) Handler {
	return Handler{
		url: "POST /task",
		handler: func(w http.ResponseWriter, r *http.Request) {
			var newTask models.Task

			if err := decodeJson(w, r, &newTask); err != nil {
				writeError(w, http.StatusBadRequest, fmt.Sprintf("Error decoding Task model: %s", err.Error()))
				return
			}

			_, ok := board.Tasks[newTask.Id]
			if ok {
				writeError(w, http.StatusConflict, fmt.Sprintf("Task with id (%s) already exists...", newTask.Id))
			}

			u := newTask
			board.Tasks[u.Id] = &u

			writeJson(w, http.StatusCreated, u)
		},
	}
}

func DeleteTaskById(board *models.Board) Handler {
	return Handler{
		url: "DELETE /task/{id}",
		handler: func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")
			_, ok := board.Tasks[id]

			if !ok {
				writeError(w, http.StatusNotFound, fmt.Sprintf("Task with id (%s) not found...", id))
				return
			}

			delete(board.Tasks, id)
		},
	}
}

func writeJson(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)

	if err := enc.Encode(data); err != nil {
		log.Printf("Error encoding to json: %v", err)
	}
}

func decodeJson(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return fmt.Errorf("Error decoding json: %w", err)
	}

	if dec.More() {
		return fmt.Errorf("Multiple json object found in body")
	}

	return nil
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeJson(w, status, apiError{Error: msg})
}
