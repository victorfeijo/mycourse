package main

import "time"

type Grade struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
    Notes     Notes     `json:"notes"`
}

type Grades []Grade

func (g *Grade) avarage() float64 {
    sum := 0.0
    for _, note := range g.Notes {
        sum = sum + note.Value
    }
    return sum
}
