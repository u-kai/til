package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	client := http.Client{}
	body := []byte(`{"message":"hello","username":"gopher"}`)
	req, err := http.NewRequest("POST", "http://localhost:8080/comments", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	b := []byte{}
	_, err = resp.Body.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", b)
}
