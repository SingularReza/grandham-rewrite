package db

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateLibrary(t *testing.T) {
	libraryName := "test_library"
	libraryType := "ANIME"

	libraryID := CreateLibrary(libraryName, libraryType)

	// Check that the response body is what we expect.

	if reflect.TypeOf(libraryID).Name() == "int64" {
		fmt.Printf("Library created and id is %d", libraryID)
	} else {
		t.Errorf("handler returned unexpected body: got %d want an int64", libraryID)
	}
}

func TestGetLibraries(t *testing.T) {
	libraryList := GetLibraries([]int{1, 0})

	// Check that the response body is what we expect.
	expected := []Library{Library{ID: 1, Name: "test", Type: "ANIME"}}

	if libraryList[0] != expected[0] {
		t.Errorf("handler returned unexpected body: got %v expected %v", libraryList, expected)
	}
}

func TestGetLibraryItems(t *testing.T) {
	libraryItems := GetLibraryItems(76, "ANIME", []int{1, 0})

	// Check that the response body is what we expect.

	expected := Item{ID: 11, Name: "Akira", PosterPath: "bx47-Sjkc8RDBjqwT.jpg"}

	if libraryItems[0] != expected {
		t.Errorf("handler returned unexpected body: got %v expected %v", libraryItems, expected)
	}
}
