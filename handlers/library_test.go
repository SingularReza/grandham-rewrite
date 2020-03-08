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
		Name:      "test-lib",
		FolderIDs: []string{"1FgWM14-uU6o8Dx12bev5SBWjRXk5uhP5"},
		Type:      "ANIME",
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
	expected := `{"name":"test-lib","folderids":["1FgWM14-uU6o8Dx12bev5SBWjRXk5uhP5"],"type":"ANIME"}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetLibraryList(t *testing.T) {
	data := LibraryRequest{
		Range: []int{1, 10},
	}

	reqData, err := json.Marshal(data)

	req, err := http.NewRequest("Get", "/library/list", bytes.NewBuffer(reqData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetLibraryList)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":11,"name":"movies-A","type":"MOVIES"}]`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetLibraryItems(t *testing.T) {
	data := LibraryRequest{
		ID:    76,
		Type:  "ANIME",
		Range: []int{1, 0},
	}

	reqData, err := json.Marshal(data)

	req, err := http.NewRequest("Get", "/library/items", bytes.NewBuffer(reqData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetLibraryItems)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"type":"ANIME","range":[1,0],"items":[{"ID":47,"Name":"Akira","PosterPath":"bx47-Sjkc8RDBjqwT.jpg"}],"id":76}`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
