package main

import "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

mux.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("Hello World!\n"))
}
