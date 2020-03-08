package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Library - generic structure for Library entry
type Library struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

// Item - simplified generic structure for a library item when a list of items are requested
type Item struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	PosterPath string `json:"posterpath,omitempty"`
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

// AddFolder - adds folder to folder table with the corresponding library id
func AddFolder(folderString string, libraryID int64) int64 {
	statement, err := database.Prepare("INSERT INTO FOLDERS (folder_string, library_id) VALUES (?, ?)")
	checkErr(err)

	result, err := statement.Exec(folderString, libraryID)
	checkErr(err)

	folderID, err := result.LastInsertId()
	checkErr(err)

	fmt.Println(folderID)

	return folderID
}

// GetLibraries - gets list of libraries
//note : think of passing library pointer from handler and returning
// it instead of creating a new struct solely for db
func GetLibraries(itemRange []int) []Library {
	query := `SELECT library_id, library_name, library_type FROM LIBRARIES
			  LIMIT ? OFFSET ?`
	rows, err := database.Query(query, itemRange[0], itemRange[1])
	checkErr(err)
	defer rows.Close()

	data := []Library{}

	for rows.Next() {
		var row Library
		err = rows.Scan(&row.ID, &row.Name, &row.Type)
		checkErr(err)

		data = append(data, row)
	}

	return data
}

// GetLibraryItems - gets list of items in the given library
func GetLibraryItems(libraryID int64, libraryType string, itemRange []int) []Item {
	var query string

	switch libraryType {
	case "ANIME":
		query = `SELECT anime_id, anime_title_romaji, anime_cover FROM ANIME
				 WHERE library_id = ? LIMIT ? OFFSET ?`
	case "MOVIES":
		query = `SELECT movie_id, movie_title, movie_poster FROM MOVIES
				 WHERE library_id = ? LIMIT ? OFFSET ?`
	}

	rows, err := database.Query(query, libraryID, itemRange[0], itemRange[1])
	checkErr(err)
	defer rows.Close()

	data := []Item{}

	for rows.Next() {
		var row Item
		err = rows.Scan(&row.ID, &row.Name, &row.PosterPath)
		checkErr(err)

		data = append(data, row)
	}

	return data
}
