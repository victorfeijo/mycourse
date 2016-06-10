package main

// Grade struct
type Grade struct {
        ID        int       `json:"id"`
        Name      string    `json:"name"`
        Completed bool      `json:"completed"`
        Notes     Notes     `json:"notes"`
}

// Status to show some math with notes
type Status struct {
        Media float64    `json:"media"`
        Max   float64    `json:"max"`
        Min   float64    `json:"min"`
}

// Grades slice
type Grades []Grade

// Avarage of notes
func (g *Grade) Avarage() float64 {
        sum := 0.0
        for _, note := range g.Notes {
                sum = sum + note.Value
        }
        return sum/float64(len(g.Notes))
}

// MaxNote of the notes slice
func (g *Grade) MaxNote() float64 {
        max := g.Notes[0].Value
        for _, note := range g.Notes {
                if note.Value > max {
                        max = note.Value
                }
        }

        return max
}

// MinNote of the notes slice
func (g *Grade) MinNote() float64 {
        min := g.Notes[0].Value
        for _, note := range g.Notes {
                if note.Value < min {
                        min = note.Value
                }
        }

        return min
}
