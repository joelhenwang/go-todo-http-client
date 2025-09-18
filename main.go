package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joelhenwang/go-todo-http-client/models"
)

type ApiError struct {
	Error string `json:"error"`
}

func main() {
	mux := http.NewServeMux()

	board := models.NewBoard("test")

	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tasks := board.Tasks

			writeJson(w, http.StatusOK, tasks)
		}
	})

	mux.HandleFunc("GET /task/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		task, ok := board.Tasks[id]

		if !ok {
			writeJson(w, 403, ApiError{Error: fmt.Sprintf("Task not found with id '%s'", id)})
			return
		}

		writeJson(w, 200, task)
	})

	mux.HandleFunc("POST /task", func(w http.ResponseWriter, r *http.Request) {
		var newTask models.Task

		if err := decodeJson(w, r, &newTask); err != nil {
			writeJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}

		_, ok := board.Tasks[newTask.Id]
		fmt.Printf("ok: %v", ok)

		if ok {
			writeJson(w, http.StatusBadRequest, ApiError{Error: "Task already exists"})
			return
		}

		u := newTask
		board.Tasks[u.Id] = &u

		writeJson(w, http.StatusCreated, u)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "Ok"}`))
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	fmt.Printf("Server starting at: %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Fatal error: %v", err)
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
