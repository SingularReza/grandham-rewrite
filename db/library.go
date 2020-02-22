package db

import (
	"database/sql"
	"fmt"

	_"github.com/mattn/go-sqlite3"
)

// Library - generic structure for Library entry
type Library struct {
	ID int
	Name string
	Type string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

var database, err = sql.Open("sqlite3", "./local.db")

// CreateLibrary - creates a library entry in LIBRARY table
func CreateLibrary(libraryName string, libraryType string) int64 {
	statement, err := database.Prepare("INSERT INTO LIBRARIES (library_name, library_type) VALUES (?, ?)")
	checkErr(err)

	result, err := statement.Exec(libraryName, libraryType)
	checkErr(err)

	libraryID, err := result.LastInsertId()
	checkErr(err)

	fmt.Println(libraryID)

	return libraryID
}