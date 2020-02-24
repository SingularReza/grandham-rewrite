package db

import (
	"fmt"
	"strings"

	metadata "github.com/SingularReza/grandham-rewrite/metadata"
	_ "github.com/mattn/go-sqlite3"
)

// CreateMovieEntry - creates an entry in table MOVIES
func CreateMovieEntry(movieData metadata.MovieData, folderID string) int64 {
	statement, err := database.Prepare(`INSERT INTO MOVIES (movie_id, movie_title, movie_original_title,
										movie_language, release_date, movie_genres, movie_backdrop,
										movie_poster) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	checkErr(err)

	var genres, genreName string
	for _, genreid := range movieData.GenreIDs {
		query := `Select genre_name from MOVIEGENRES WHERE genre_id = ?`
		newStatement, err := database.Prepare(query)
		checkErr(err)

		err = newStatement.QueryRow(genreid).Scan(&genreName)
		checkErr(err)

		genres += genreName + ","
	}
	genres = strings.TrimLeft(genres, ",")
	//fmt.Printf("movieid: %d\n", movieData.ID)
	result, err := statement.Exec(movieData.ID, movieData.Title, movieData.OriginalTitle,
		movieData.OriginalLanguage, movieData.ReleaseDate, genres,
		movieData.BackdropPath, movieData.PosterPath)
	checkErr(err)

	movieID, err := result.LastInsertId()
	checkErr(err)

	infoID := addMovieExtraInfo(movieData.Overview, movieID)

	fmt.Printf("movieid: %d, movieinfoid: %d", movieID, infoID)

	return movieID
}

func addMovieExtraInfo(description string, movieID int64) int64 {
	statement, err := database.Prepare(`INSERT INTO MOVIEINFO (movie_overview, movie_id) VALUES (?, ?)`)
	checkErr(err)

	result, err := statement.Exec(description, movieID)
	checkErr(err)

	infoID, err := result.LastInsertId()
	checkErr(err)

	fmt.Println(infoID)

	return infoID
}
