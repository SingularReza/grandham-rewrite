package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// LibraryRequest - generic request structure for any request related to library
type LibraryRequest struct {
	Name string				`json:"name,omitempty"`
	FolderIDs [] string		`json:"folderids,omitempty"`
	Type string				`json:"type,omitempty"`
	Range [] int			`json:"range,omitempty"`
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

	sendResponse(w, library)
}