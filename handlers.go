package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the grades system!")
}

func GradeIndex(w http.ResponseWriter, r *http.Request) {
    grades := Grades{
        Grade{Name: "first_grade"},
        Grade{Name: "second_grande"},
    }

    if err := json.NewEncoder(w).Encode(grades); err != nil {
        panic(err)
    }
}

func GradeShow(w http.ResponseWriter, r *http.Request) {
    //vars == params(rails)
    vars := mux.Vars(r)
    gradeId := vars["gradeId"]
    fmt.Fprintln(w, "Grades Show: ", gradeId)
}
