package main

import (
	"fmt"
	"net/http"

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

	fmt.Printf("Server starting at: %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Fatal error: %v", err)
	}

}
