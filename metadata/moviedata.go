package metadata

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
	//scan "github.com/SingularReza/grandham-rewrite/gdrive"
)

// MovieSearchResults - gives an array of search results
type MovieSearchResults struct {
	Results []MovieData `json:"results,omitempty"`
}

// MovieData - movie info
type MovieData struct {
	ID               int      `json:"id,omitempty"`
	Title            string   `json:"title,omitempty"`
	OriginalTitle    string   `json:"original_title,omitempty"`
	OriginalLanguage string   `json:"original_language,omitempty"`
	ReleaseDate      string   `json:"release_date,omitempty"`
	GenreIDs         []int    `json:"genre_ids,omitempty"`
	BackdropPath     string   `json:"backdrop_path,omitempty"`
	Overview         string   `json:"overview,omitempty"`
	PosterPath       string   `json:"poster_path,omitempty"`
	Duration         int      `json:"duration,omitempty"`
	FileSize         int      `json:"filesize,omitempty"`
	Height           int      `json:"height,omitempty"`
	Width            int      `json:"width,omitempty"`
	FileID           string   `json:"fileid,omitempty"`
	FolderID         string   `json:"folderid,omitempty"`
	GenreNames       []string `json:"genre_names,omitempty"`
}

func getData(url string, target interface{}) error {
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	r, err := client.Get(url)
	checkErr(err)

	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

// GetMovieData - Gets metadata from themoviedb
func GetMovieData(moviename string) MovieData {
	baseURL, err := url.Parse("https://api.themoviedb.org/3/search/movie")
	checkErr(err)

	params := url.Values{}
	params.Add("api_key", "76a3ff4ff39ddb9509ee12400d0e7330")
	params.Add("query", moviename)
	baseURL.RawQuery = params.Encode()

	response := &MovieSearchResults{}
	getData(baseURL.String(), response)

	movieInfo := MovieData{}
	//fmt.Printf("", moviename)
	//fmt.Printf("moviename: %s, data: %+v\n", moviename, response)
	movieInfo = response.Results[0]

	// use scan to get runtime and filesize later
	go downloadImage("https://image.tmdb.org/t/p/w1280" + movieInfo.PosterPath)
	go downloadImage("https://image.tmdb.org/t/p/w1280" + movieInfo.BackdropPath)

	return movieInfo
}
