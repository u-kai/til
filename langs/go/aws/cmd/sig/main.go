package main

import (
	"aws/lib"
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := os.Getenv("TEST_API_GATEWAY_URL")
	clietnt := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	data, err := lib.AddAwsSignature(clietnt, req)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
