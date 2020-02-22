package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	db 		"github.com/SingularReza/grandham-rewrite/db"
	scan 	"github.com/SingularReza/grandham-rewrite/gdrive"
)

// LibraryRequest - generic request structure for any request related to library
type LibraryRequest struct {
	Name string				`json:"name,omitempty"`
	FolderIDs [] string		`json:"folderids,omitempty"`
	Type string				`json:"type,omitempty"`
	Range [] int			`json:"range,omitempty"`
}

// Item - generic structure for Items in a folder, conatins driveid and name
type Item struct {
	Name string
	FolderID string
}

func checkErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func sendResponse(w http.ResponseWriter, data interface{}) {
	data, ok := data.(LibraryRequest);

	if ok {
		response, err := json.Marshal(data)
		checkErr(w, err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	} else {
		fmt.Println("data is null")
	}
}


// CreateLibrary - Creates Library entry after scanning the relevant folder
func CreateLibrary(w http.ResponseWriter, r *http.Request) {
	library := LibraryRequest{}

	err := json.NewDecoder(r.Body).Decode(&library)
	if err != nil {
		panic(err)
	}

	libraryID := db.CreateLibrary(library.Name, library.Type)

	var items []Item

	for _, folderID := range library.FolderIDs {
		folderItems := scan.GetItemsList(folderID)
		for _, item := range folderItems {
			items = append(items, Item{item.Name, item.Id});
		}
	}

	/*
	if library.Type == "MOVIES" {
		for _, movie := range items {
			movieMetaData := metadata.GetMovieData(movie.Name)
			db.CreateMovieEntry(movieMetaData, movie.ID)
		}
	} else if library.Type == "ANIME" {
		for _, anime := range items {
			animeData := metadata.GetAnimeData(anime.Name)
			db.CreateAnimeEntry(animeData, anime.ID)
		}
	}*/

	for _, item := range items {
		fmt.Printf("{Name:\"%s\",Id:\"%s\"},", item.Name, item.FolderID)
	}
	//fmt.Println(items)
	fmt.Println(libraryID)

	sendResponse(w, library)
}