package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewInMemoryStore()
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			createTodoHandler(store)(w, r)
		case http.MethodGet:
			getTodosHandler(store)(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getTodoByIDHandler(store)(w, r)
		case http.MethodPut:
			updateTodoHandler(store)(w, r)
		case http.MethodDelete:
			deleteTodoHandler(store)(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{
			"status": "ok",
		})
	})
	
	loggedMux := loggingMiddleware(mux)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", loggedMux))

}
