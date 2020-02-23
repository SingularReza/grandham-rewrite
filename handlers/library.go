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
		fmt.Println(items)
		for _, anime := range items {
			fmt.Print(anime.Name)
			animeData := metadata.GetAnimeData(anime.Name)
			animeEntryID := db.CreateAnimeEntry(animeData, anime.FolderID)
			fmt.Printf("%+v, %d\n", animeData, animeEntryID)
		}
	}

	sendResponse(w, library)
}

// AddFolder - Adds a folder to a library
func AddFolder(folderID string, libraryID int64, itemList []Item) []Item {
	folderItems := scan.GetItemsList(folderID)
	for _, item := range folderItems {
		itemList = append(itemList, Item{item.Name, item.Id})
		fmt.Print(item)
	}

	db.AddFolder(folderID, libraryID)

	//fmt.Print(itemList)
	return itemList
}
