package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateLibrary(t *testing.T) {
	data := LibraryRequest{
		Name: "library_name",
		FolderIDs: []string{"folder1", "folder2"},
		Type: "ANIME",
	}

	reqData, err := json.Marshal(data)

    req, err := http.NewRequest("POST", "/library/create", bytes.NewBuffer(reqData))
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreateLibrary)

    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method
    // directly and pass in our Request and ResponseRecorder.
    handler.ServeHTTP(rr, req)

    // Check that the status code is what we expect.
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check that the response body is what we expect.
    expected := `{"name":"library_name","folderids":["folder1","folder2"],"type":"ANIME"}`

    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}