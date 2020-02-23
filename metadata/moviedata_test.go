package metadata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Note: keep a copy of credentials.json and token.json before running the test
func TestGetAnimeData(t *testing.T) {
	movieName := "kimi no na wa"

	movieData := GetAnimeData(movieName)

	expected := MovieData{}

	if !cmp.Equal(movieData, expected) {
		t.Errorf("GetMovieData in metadata/moviedata.go file returned unexpected response")
	}
}
