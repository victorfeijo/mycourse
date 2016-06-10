package main

import (
        "github.com/gorilla/mux"
)

// NewRouter wrapper from Router to muxRouter
func NewRouter() *mux.Router {

        router := mux.NewRouter().StrictSlash(true)
        for _, route := range routes {
                router.
                Methods(route.Method).
                Path(route.Pattern).
                Name(route.Name).
                Handler(route.HandlerFunc)
        }

        return router;
}
