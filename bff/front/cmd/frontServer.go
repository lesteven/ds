package main

import (
    "log"
    "net/http"
    //"github.com/gorilla/mux"
)

// ParseGlob equal to ParseFiles w/ list of files
//var t = template.Must(template.ParseGlob("templates/*/*.html"))

func main() {
    //r := mux.NewRouter()

    //r.HandleFunc("/", Home)

    mux := http.NewServeMux()

    mux.HandleFunc("/", Home)
    fileServer := http.FileServer(http.Dir("../ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Fatal(http.ListenAndServe(":8080", mux))
}
