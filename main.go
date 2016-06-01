package main

import (
    "html"
    "log"
    "os"
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"), os.Getenv("OPENSHIFT_GO_PORT"))
    log.Fatal(http.ListenAndServe(bind, router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

