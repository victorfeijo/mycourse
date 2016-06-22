package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// RepoGetGrades by gradeID
func RepoGetGrades() Grades {
    db, err := sql.Open("mysql", "root:@/feijomlgo")
    CheckErr(err)
    rows, err := db.Query("SELECT * FROM Grade")
    CheckErr(err)
    var grades Grades
    for rows.Next() {
        var ID int
        var Name string
        var Completed bool
        err = rows.Scan(&ID, &Name, &Completed)
        CheckErr(err)
        db.Close()

        var grade = Grade{ID, Name, Completed, nil}
        grades = append(grades, grade)
    }

    db.Close()
    return grades
}

// RepoFindGrade by gradeID
func RepoFindGrade(id int) Grade {
    db, err := sql.Open("mysql", "root:@/feijomlgo")
    CheckErr(err)
    rows, err := db.Query("SELECT * FROM Grade WHERE ID=?", id)
    CheckErr(err)
    for rows.Next() {
        var ID int
        var Name string
        var Completed bool
        err = rows.Scan(&ID, &Name, &Completed)
        CheckErr(err)
        db.Close()
        Notes := RepoFindNotes(ID)

        return Grade{ID, Name, Completed, Notes}
    }

    db.Close()
    return Grade{}
}

// RepoCreateGrade create on database
func RepoCreateGrade(g Grade) {
    db, err := sql.Open("mysql", "root:@/feijomlgo")
    CheckErr(err)
    stmt, err := db.Prepare("INSERT Grade SET Name=?, Completed=?")
    CheckErr(err)
    _, err = stmt.Exec(g.Name, g.Completed)
    CheckErr(err)
    db.Close()
}

//RepoFindNotes find the notes from grade
func RepoFindNotes(gradeID int) Notes {
    db, err := sql.Open("mysql", "root:@/feijomlgo")
    CheckErr(err)
    rows, err := db.Query("SELECT * FROM Note WHERE GradeID=?", gradeID)
    CheckErr(err)
    var notes Notes
    for rows.Next() {
        var Name string
        var Value float64
        err = rows.Scan(&Name, &Value)
        CheckErr(err)
        db.Close()

        note := Note{Name, Value}
        notes = append(notes, note)
    }

    db.Close()
    return notes
}

// RepoAddNote to the Notes slice by gradeID
func RepoAddNote(g Grade, n Note) Note {
    g.Notes = append(g.Notes, n)
    return n
}

// CheckErr do something with the error
func CheckErr(err error) {
    if err != nil {
        panic(err)
    }
}
