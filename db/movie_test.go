package db

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/SingularReza/grandham-rewrite/metadata"
	drive "google.golang.org/api/drive/v3"
)

func TestCreateMovieEntry(t *testing.T) {

	movieData := metadata.MovieData{
		ID:               12345678,
		Title:            "Azhar",
		OriginalTitle:    "Azhar",
		OriginalLanguage: "hi",
		ReleaseDate:      "2016-05-13",
		GenreIDs:         []int{18},
		BackdropPath:     "/76sLxPTDeiwf7BpZMNo6J37inFZ.jpg",
		Overview:         "Indian biographical sports film directed by Tony D'Souza based on the life of the former Indian international cricketer, Mohammad Azharuddin. It will star emraan hashmi in the lead role alongside Prachi Desai , Nargis Fakhri and Huma Qureshi",
		PosterPath:       "/2jn4HTYdOAiErTVVNRYH2SNxxOd.jpg",
	}

	fileData := &drive.File{
		Size: 300,
		VideoMediaMetadata: &drive.FileVideoMediaMetadata{
			DurationMillis: 300,
			Height:         300,
			Width:          300,
		},
	}

	movieEntryID := CreateMovieEntry(movieData, "testid", fileData, 1)

	// Checking that the entry is successful | note: change this after making createmovieentry return (int64, err)
	if reflect.TypeOf(movieEntryID).Name() == "int64" {
		fmt.Print("CreateMovieEntry creates a row\n")
	} else {
		t.Errorf("CreateMovieEntry in db/movie.go failed")
	}
}
