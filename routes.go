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
        "GradeStatus",
        "GET",
        "/grades/{gradeId}/status",
        GradeStatus,
    },
    Route{
        "GradeCreate",
        "POST",
        "/grades",
        GradeCreate,
    },
    Route{
        "AddGradeNote",
        "POST",
        "/grades/{gradeId}",
        AddGradeNote,
    },
}
