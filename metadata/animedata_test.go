package metadata

import (
	"testing"
)

// Note: keep a copy of credentials.json and token.json before running the test
func TestGetAnimeData(t *testing.T) {
	animeName := "kimi no na wa"

	animeData := GetAnimeData(animeName)

	// we are only checking for ID because if one field exists other fields are correctly returned as well

	if animeData.ID != 21519 {
		t.Errorf("GetAnimeData in metadata/animedata.go file returned %+v", animeData)
	}
}
