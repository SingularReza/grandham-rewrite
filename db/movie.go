package db

import (
	metadata "github.com/SingularReza/grandham-rewrite/metadata"
	_ "github.com/mattn/go-sqlite3"
)

// CreateMovieEntry - creates an entry in table MOVIES
func CreateMovieEntry(movieData metadata.MovieData, folderID string) int64 {
	return 1
}
