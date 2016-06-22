package main

import (
    "log"
    //"os"
    "net/http"
    "fmt"
)

func main() {
    router := NewRouter()
    //bind := fmt.Sprintf("%s:%s", os.Getenv("OPENSHIFT_GO_IP"),
    //                   os.Getenv("OPENSHIFT_GO_PORT"))
    bind := fmt.Sprintf("%s:%s", "localhost", "8080")
    log.Fatal(http.ListenAndServe(bind, router))
}
