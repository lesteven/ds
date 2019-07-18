package main

import (
    "net/http"
    "html/template"
)

type Person struct {
    Name string
}

func Home(w http.ResponseWriter, r *http.Request) {

    p := Person{"Steven"}
    t, _ := template.ParseFiles("home.html")
    t.Execute(w, p)
}

