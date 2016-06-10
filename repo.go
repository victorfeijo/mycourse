package main

var currentId int

var grades Grades

func init() {
        notes1 := Notes{
                Note{Name: "p1", Value: 7.0},
        }
        notes2 := Notes{
                Note{Name: "p1", Value: 4.0},
                Note{Name: "p2", Value: 5.0},
        }
        notes3 := Notes{
                Note{Name: "t1", Value: 10.0},
        }
        notes4 := Notes{
                Note{Name: "p1", Value: 6.5},
        }
        notes5 := Notes{
                Note{Name: "p1", Value: 1.9},
                Note{Name: "p2", Value: 6.8},
        }

        RepoCreateGrade(Grade{Name: "Computação Distribuída", Notes: notes1})
        RepoCreateGrade(Grade{Name: "Engenharia de Software", Notes: notes2})
        RepoCreateGrade(Grade{Name: "Computação Gráfica", Notes: notes3})
        RepoCreateGrade(Grade{Name: "Redes de Computadores II", Notes: notes4})
        RepoCreateGrade(Grade{Name: "Bando de Dados I", Notes: notes5})
}

func RepoFindGrade(id int) Grade {
        for _, g := range grades {
                if g.Id == id {
                        return g
                }
        }

        return Grade{}
}

func RepoCreateGrade(g Grade) Grade {
        currentId += 1
        g.Id = currentId
        grades = append(grades, g)
        return g
}

func RepoAddNote(g Grade, n Note) Note {
        g.Notes = append(g.Notes, n)
        return n
}
