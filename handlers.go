package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "strconv"
    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to the grades system!")
}

func GradeIndex(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(grades); err != nil {
        panic(err)
    }
}

func GradeShow(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    vars := mux.Vars(r)
    gradeId, _ := strconv.Atoi(vars["gradeId"])
    grade := RepoFindGrade(gradeId)

    if err := json.NewEncoder(w).Encode(grade); err != nil {
        panic(err)
    }
}

func GradeCreate(w http.ResponseWriter, r* http.Request) {
    var grade Grade
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &grade); err != nil {
        w.Header().Set("Content-Type", "application/json;charset=UTF-8")
        w.WriteHeader(422)
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    g := RepoCreateGrade(grade)
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(g); err != nil {
        panic(err)
    }
}

func AddGradeNote(w http.ResponseWriter, r* http.Request) {
    var note Note
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &note); err != nil {
        w.Header().Set("Content-Type", "application/json;charset=UTF-8")
        w.WriteHeader(422)
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    vars := mux.Vars(r)
    gradeId, _ := strconv.Atoi(vars["gradeId"])
    grade := RepoFindGrade(gradeId)
    n := RepoAddNote(grade, note)

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(n); err != nil {
        panic(err)
    }
}

func GradeStatus(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    gradeId, _ := strconv.Atoi(vars["gradeId"])
    grade := RepoFindGrade(gradeId)
    status := Status{Media: grade.Avarage(), Max: grade.MaxNote(), Min: grade.MinNote()}

    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(status); err != nil {
        panic(err)
    }
}
