package main

import "net/http"

// Route represents a http route
type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

// Routes slice
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
            "/grades/{gradeID}",
            GradeShow,
    },
    Route{
            "GradeStatus",
            "GET",
            "/grades/{gradeID}/status",
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
            "/grades/{gradeID}",
            AddGradeNote,
    },
}
