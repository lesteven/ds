package main

import (
    "net/http"
    "github.com/lesteven/goRes"
)

type Pub struct {
    Service string `json:"service"`
    Greeting string `json:"greeting"`
}

func Home(w http.ResponseWriter, r *http.Request) {
    data := Pub{"public", "this is the backend server!"}
    goRes.SendJson(w, 200, data)
}

