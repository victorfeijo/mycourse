package main

type Note struct {
        Name      string    `json:"name"`
        Value     float64   `json:"value"`
}

type Notes []Note
