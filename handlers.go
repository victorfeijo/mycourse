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
    notes := Notes{
        Note{Name: "note1", Value: 6.6},
        Note{Name: "note2", Value: 9.5},
    }

    grades := Grades{
        Grade{Name: "first_grade", Notes: notes},
        Grade{Name: "second_grande"},
    }

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

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

func GradeAvg(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Media: ", 10)
}
