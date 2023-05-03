package main

import "net/http"

func main() {
	http.HandleFunc("/", Hello)
	println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
