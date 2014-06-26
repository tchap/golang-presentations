package main

import (
    "io"
    "net/http"
    "github.com/bmizerany/pat"
    "log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
}

func main() {
    mux := pat.New()
    mux.Get("/hello/:name", http.HandlerFunc(HelloServer))

    http.Handle("/", mux)
    if err := http.ListenAndServe(":12345", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
