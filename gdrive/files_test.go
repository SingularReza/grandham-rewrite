package db

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetItemsList(t *testing.T) {
	folderID := "test_library"
	libraryType := "ANIME"

    libraryID := CreateLibrary(libraryName, libraryType)

	// Check that the response body is what we expect.

    if reflect.TypeOf(libraryID).Name() == "int64"  {
        fmt.Printf("Library created and id is %d", libraryID)
    } else {
		t.Errorf("handler returned unexpected body: got %d want an int64", libraryID)
	}
}
