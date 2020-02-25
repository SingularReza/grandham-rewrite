package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/SingularReza/grandham-rewrite/db"
	scan "github.com/SingularReza/grandham-rewrite/gdrive"
	metadata "github.com/SingularReza/grandham-rewrite/metadata"
)

// LibraryRequest - generic request structure for any request related to library
type LibraryRequest struct {
	Name      string   `json:"name,omitempty"`
	FolderIDs []string `json:"folderids,omitempty"`
	Type      string   `json:"type,omitempty"`
	Range     []int    `json:"range,omitempty"`
}

// Item - generic structure for Items in a folder, conatins driveid and name
type Item struct {
	Name     string
	FolderID string
}

func checkErr(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func sendResponse(w http.ResponseWriter, data interface{}) {
	data, ok := data.(LibraryRequest)

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

func recordAnimeData(anime Item, libraryID int64) {
	animeData := metadata.GetAnimeData(anime.Name)
	animeEntryID := db.CreateAnimeEntry(animeData, anime.FolderID, libraryID)
	fmt.Printf("%+v, %d\n", animeData, animeEntryID)
}

func recordMovieData(movie Item, libraryID int64) {
	movieData := metadata.GetMovieData(movie.Name)
	movieFile := scan.FilterVideos(movie.FolderID)[0]
	movieFileData := scan.GetVideoData(movieFile)
	movieEntryID := db.CreateMovieEntry(movieData, movie.FolderID, movieFileData, libraryID)
	fmt.Printf("%+v, %d\n", movieData, movieEntryID)
}

// CreateLibrary - Creates Library entry after scanning the relevant folder
func CreateLibrary(w http.ResponseWriter, r *http.Request) {
	library := LibraryRequest{}

	err := json.NewDecoder(r.Body).Decode(&library)
	if err != nil {
		panic(err)
	}

	libraryID := db.CreateLibrary(library.Name, library.Type)

	items := []Item{}

	for _, folderID := range library.FolderIDs {
		items = AddFolder(folderID, libraryID, items)
	}

	if library.Type == "ANIME" {
		for _, anime := range items {
			go recordAnimeData(anime, libraryID)
		}
	} else if library.Type == "MOVIE" {
		for _, movie := range items {
			go recordMovieData(movie, libraryID)
		}
	}

	sendResponse(w, library)
}

// AddFolder - Adds a folder to a library
func AddFolder(folderID string, libraryID int64, itemList []Item) []Item {
	folderItems := scan.GetItemsList(folderID)
	for _, item := range folderItems {
		itemList = append(itemList, Item{item.Name, item.Id})
	}

	db.AddFolder(folderID, libraryID)

	return itemList
}
