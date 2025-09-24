package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/joelhenwang/go-todo-http-client/db"
	"github.com/joelhenwang/go-todo-http-client/models"
)

func main() {
	mux := http.NewServeMux()

	board := models.NewBoard("test")

	handlers := GetHandlers(board)

	for _, handler := range handlers {
		mux.HandleFunc(handler.url, handler.handler)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "Ok"}`))
	})

	server := &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}

	db_conn := db.Init()
	defer db_conn.Close(context.Background())

	db.TestQueries(db_conn)

	fmt.Printf("Server starting at: %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Fatal error: %v", err)
	}

}
