package metadata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	//scan "github.com/SingularReza/grandham-rewrite/gdrive"
)

type animeRequest struct {
	Query     string      `json:"query"`
	Variables animeParams `json:"variables"`
}

type animeParams struct {
	Search string `json:"search"`
	Type   string `json:"type"`
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
	ID              int            `json:"id,omitempty"`
	Title           AnimeTitle     `json:"title,omitempty"`
	CoverImage      AnimeImage     `json:"coverImage,omitempty"`
	BannerImage     string         `json:"bannerImage,omitempty"`
	Description     string         `json:"description,omitempty"`
	Season          string         `json:"season,omitempty"`
	SeasonYear      int            `json:"seasonYear,omitempty"`
	Format          string         `json:"format,omitempty"`
	Episodes        int            `json:"episodes,omitempty"`
	EpisodeDuration int            `json:"duration,omitempty"`
	Genres          []string       `json:"genres,omitempty"`
	StartDate       AnimeDate      `json:"startDate,omitempty"`
	EndDate         AnimeDate      `json:"endDate,omitempty"`
	FolderID        string         `json:"folderid,omitempty"`
	EpisodeList     []AnimeEpisode `json:"episodelist,omitempty"`
}

// AnimeDate - holds anime dates for start and end
type AnimeDate struct {
	Year  int `json:"year,omitempty"`
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
}

// AnimeTitle - title key value
type AnimeTitle struct {
	Romaji  string `json:"romaji,omitempty"`
	English string `json:"english,omitempty"`
}

// AnimeImage - coverImage key value
type AnimeImage struct {
	Large string `json:"large,omitempty"`
}

type AnimeEpisode struct {
	FileID       string `json:"fileid,omitempty"`
	FileName     string `json:"filename,omitempty"`
	FileDuration int    `json:"fileduration,omitempty"`
	FileSize     int    `json:"filesize,omitempty"`
	FileHeight   int    `json:"fileheight,omitempty"`
	FileWidth    int    `json:"filewidth,omitempty"`
}

func checkErr(err error) {
	if err != nil {
		print(err)
	}
}

// GetAnimeData - gets metadata from anilist
func GetAnimeData(animeName string) AnimeMedia {
	// todo: expand query to get more data
	fmt.Printf("animename: %s\n", animeName)
	query := `query($search: String, $type: MediaType){
                Media(search: $search, type: $type){
                	id
                    title {
						romaji
						english
                    }
                    coverImage {
                        large
					}
					bannerImage
					startDate {
					  year
					  month
					  day
					}
					endDate {
					  year
					  month
					  day
					}
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
		Search: animeName,
		Type:   "ANIME",
	}

	reqBody := &animeRequest{
		Query:     query,
		Variables: params,
	}

	req, _ := json.Marshal(reqBody)

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	resp, err := client.Post("https://graphql.anilist.co/", "application/json", bytes.NewBuffer(req))
	checkErr(err)

	defer resp.Body.Close()

	respBody := AnimeInfo{}
	json.NewDecoder(resp.Body).Decode(&respBody)

	// use scan to get runtime and filesize later
	go downloadImage(respBody.Data.Media.CoverImage.Large)
	go downloadImage(respBody.Data.Media.BannerImage)

	return respBody.Data.Media
}
