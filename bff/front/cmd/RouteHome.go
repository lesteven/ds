package main

import (
    "net/http"
    "log"
    "html/template"
)

type Person struct {
    Name string
}

func Home(w http.ResponseWriter, r *http.Request) {

    p := Person{"Steven"}
    //err := t.ExecuteTemplate(w, "home.html", p)

    ts, _ := template.ParseFiles("../ui/html/home.html")
    err := ts.Execute(w, p)

    if err != nil {
        log.Fatal("front server home ", err)
    }
}

