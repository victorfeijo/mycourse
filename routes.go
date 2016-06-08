package main

import "net/http"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "GradeIndex",
        "GET",
        "/grades",
        GradeIndex,
    },
    Route{
        "GradeShow",
        "GET",
        "/grades/{gradeId}",
        GradeShow,
    },
    Route{
        "GradeAvg",
        "GET",
        "/grades/{gradeId}/media",
        GradeAvg,
    },
}
