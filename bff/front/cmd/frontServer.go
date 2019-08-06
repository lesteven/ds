package main

import (
    "log"
    "net/http"
    //"github.com/gorilla/mux"
    "flag"
)

// ParseGlob equal to ParseFiles w/ list of files
//var t = template.Must(template.ParseGlob("templates/*/*.html"))

func main() {
    //r := mux.NewRouter()
    //r.HandleFunc("/", Home)

    // value of addr is pointer to flag
    addr := flag.String("addr", ":8080", "HTTP port")
    flag.Parse()

    mux := http.NewServeMux()

    mux.HandleFunc("/", Home)
    fileServer := http.FileServer(http.Dir("../ui/static/"))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Fatal(http.ListenAndServe(*addr, mux))
}
