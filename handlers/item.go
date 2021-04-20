package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"

	db "github.com/SingularReza/grandham-rewrite/db"
	metadata "github.com/SingularReza/grandham-rewrite/metadata"
)

func getAnimeInfo(animeID int) metadata.AnimeMedia {
	animeInfo := db.GetAnimeInfo(animeID)
	return animeInfo
}

func getMovieInfo(movieID int) metadata.MovieData {
	movieInfo := db.GetMovieInfo(movieID)
	return movieInfo
}

// GetItemInfo - retrieves all info about an item
func GetItemInfo(w http.ResponseWriter, r *http.Request) {
	item := db.Item{}

	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))

	err = json.NewDecoder(r.Body).Decode(&item)
	checkErr(w, err)

	switch item.Type {
	case "ANIME":
		animeInfo := getAnimeInfo(item.ID)
		sendResponse(w, animeInfo)
	case "MOVIES":
		movieInfo := getMovieInfo(item.ID)
		sendResponse(w, movieInfo)
	}
}
