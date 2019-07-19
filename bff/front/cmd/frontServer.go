package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

// ParseGlob equal to ParseFiles w/ list of files
//var t = template.Must(template.ParseGlob("templates/*/*.html"))

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", Home)

    log.Fatal(http.ListenAndServe(":8080", r))
}
