package metadata

import (
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
)

type animeRequest struct {
	Query string `json:"query"`
	Variables animeParams `json:"variables"`
}

type animeParams struct {
	Search string `json:"search"`
	Type string `json:"type"`
}

// AnimeInfo - data key value
type AnimeInfo struct {
	Data AnimeData `json:"data,omitempty"`
}

// AnimeData - movie info
type AnimeData struct {
	Media AnimeMedia `json:"Media,omitempty"`
}

// AnimeMedia - Media key value
type AnimeMedia struct {
	ID int `json:"id,omitempty"`
	Title AnimeTitle `json:"title,omitempty"`
	CoverImage AnimeImage `json:"coverImage,omitempty"`
	BannerImage string `json:"bannerImage,omitempty"`
	Description string `json:"description,omitempty"`
	Season string `json:"season,omitempty"`
	SeasonYear int `json:"seasonYear,omitempty"`
	Format string `json:"format,omitempty"`
	Episodes int `json:"episodes,omitempty"`
	EpisodeDuration int `json:"duration,omitempty"`
	Genres []string `json:"genres,omitempty"`
	StartDate AnimeDate `json:"startDate,omitempty"`
	EndDate AnimeDate `json:"endDate,omitempty"`
}

// AnimeDate - holds anime dates for start and end
type AnimeDate struct {
	Year int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day int `json:"day,omitempty"`
}

// AnimeTitle - title key value
type AnimeTitle struct {
	Romaji string `json:"romaji,omitempty"`
	English string `json:"english,omitempty"`
}

// AnimeImage - coverImage key value
type AnimeImage struct {
	Large string `json:"large,omitempty"`
}

// Anime - final anime data object
type Anime struct {
	Title string
	ReleaseDate string
	GenreIDs []int
	BackdropPath string
	Overview string
	PosterPath string
	// created related struct
}

func checkErr(err error) {
    if err != nil {
        print(err)
    }
}

// GetAnimeData - gets metadata from anilist
func GetAnimeData(animename string) AnimeInfo {
	// todo: expand query to get more data
	query := `query($search: String, $type: MediaType){
                Media(search: $search, type: $type){
                    title {
						romaji
						english
                    }
                    coverImage {
                        large
					}
					startDate
					endDate
					description
					season
					seasonYear
					format
					episodes
					duration
					genres
                }
            }`

    params := animeParams{
    	Search: animename,
  		Type: "ANIME",
    }

	reqBody := &animeRequest{
		Query: query,
		Variables: params,
	}

	req, _ := json.Marshal(reqBody)

	resp, err := http.Post("https://graphql.anilist.co/", "application/json", bytes.NewBuffer(req))
	checkErr(err)

	defer resp.Body.Close()

	fmt.Print(resp.Body)

	b := AnimeInfo{}
	json.NewDecoder(resp.Body).Decode(&b)

	//return b.Data.Media

	//animeObject := Anime{} //convert b.Data.Media to anime object
	// download poster and backdrop
	return b
}