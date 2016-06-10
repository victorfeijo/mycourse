package main

type Grade struct {
        Id        int       `json:"id"`
        Name      string    `json:"name"`
        Completed bool      `json:"completed"`
        Notes     Notes     `json:"notes"`
}

type Status struct {
        Media float64    `json:"media"`
        Max   float64    `json:"max"`
        Min   float64    `json:"min"`
}

type Grades []Grade

func (g *Grade) Avarage() float64 {
        sum := 0.0
        for _, note := range g.Notes {
                sum = sum + note.Value
        }
        return sum/float64(len(g.Notes))
}

func (g *Grade) MaxNote() float64 {
        max := g.Notes[0].Value
        for _, note := range g.Notes {
                if note.Value > max {
                        max = note.Value
                }
        }

        return max
}

func (g *Grade) MinNote() float64 {
        min := g.Notes[0].Value
        for _, note := range g.Notes {
                if note.Value < min {
                        min = note.Value
                }
        }

        return min
}
