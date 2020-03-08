package handler

import (
	"encoding/json"
	"net/http"

	db "github.com/SingularReza/grandham-rewrite/db"
)

func getAnimeInfo(animeID int64) db.Anime {
	animeInfo := db.GetAnimeInfo(animeID)
}

func getMovieInfo(movieID int64) db.Movie {
	movieInfo := db.GetMovieInfo(movieID)
}

// GetItemInfo - retrieves all info about an item
func GetItemInfo(w http.ResponseWriter, r *http.Request) {
	item := db.Item{}

	err := json.NewDecoder(r.Body).Decode(&item)
	checkErr(err)

	switch item.Type {
	case "ANIME":
		animeInfo := getAnimeInfo(item.ID)
		sendResponse(w, animeInfo)
	case "MOVIES":
		movieInfo := getMovieInfo(item.ID)
		sendResponse(w, movieInfo)
	}
}
