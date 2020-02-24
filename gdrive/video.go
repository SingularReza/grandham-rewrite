package scan

import (
	"log"

	drive "google.golang.org/api/drive/v3"
)

// FilterVideos - returns ids for all videos in a folder
func FilterVideos(folderID string) []string {
	srv := getService()
	response, err := srv.Files.List().
		Q("'" + folderID + "' in parents and mimeType != 'application/vnd.google-apps.folder'").
		PageSize(10).
		IncludeItemsFromAllDrives(true).
		SupportsAllDrives(true).
		Fields("files(id, name)").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	fileIDs := []string{}

	for _, file := range response.Files {
		fileIDs = append(fileIDs, file.Id)
	}

	return fileIDs
}

// GetVideoData - returns video metadat from gdrive api
func GetVideoData(videoID string) *drive.File {
	srv := getService()
	fileData, err := srv.Files.Get(videoID).
		SupportsAllDrives(true).
		Fields("id, name, fileExtension, size, videoMediaMetadata").
		Do()
	if err != nil {
		log.Fatalf("Unable to retrieve files: %v", err)
	}

	return fileData
}
