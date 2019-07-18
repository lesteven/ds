package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", Home)

    log.Fatal(http.ListenAndServe(":3000", r))
}
