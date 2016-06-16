package main

// Note represents a grades note
type Note struct {
    Name      string    `json:"name"`
    Value     float64   `json:"value"`
}

// Notes slice
type Notes []Note
