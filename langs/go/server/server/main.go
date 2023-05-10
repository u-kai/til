package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

type Comment struct {
	Message  string
	UserName string
}

func main() {
	var mutex = &sync.RWMutex{}
	comments := make([]Comment, 0, 100)

	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			mutex.RLock()
			if err := json.NewEncoder(w).Encode(comments); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			mutex.RUnlock()
		case http.MethodPost:
			var comment Comment
			if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			mutex.Lock()
			comments = append(comments, comment)
			mutex.Unlock()

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status":"created"}`))

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		go func() {

			time.Sleep(15 * time.Second)
			print("hello")
			w.Write([]byte("Hello World"))
		}()
	})
	println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	println("Server stopped")
	if err != nil {
		print("hello")
		panic(err)
	}
	//panic(err)

}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
