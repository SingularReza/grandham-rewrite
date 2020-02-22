package scan

import (
    "fmt"
    "log"

    drive "google.golang.org/api/drive/v3"
)

// GetItemsList - reads movie library and parses the data correctly
func GetItemsList(folderid string) []*drive.File {
	srv := getService()
    r, err := srv.Files.List().
				Q("'"+ folderid +"' in parents").
				PageSize(1000).
                IncludeItemsFromAllDrives(true).
                SupportsAllDrives(true).
                Fields("files(id, name)").Do()
    if err != nil {
        log.Fatalf("Unable to retrieve files: %v", err)
    }

    if len(r.Files) == 0 {
        fmt.Println("No files found")
    }

    return r.Files
}