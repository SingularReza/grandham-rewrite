package metadata

import (
	"os"
	"testing"
)

func TestdownloadImage(t *testing.T) {
	downloadImage("https://s4.anilist.co/file/anilistcdn/media/anime/cover/medium/nx21519-F2z1QPS5GmpC.jpg")

	if _, err := os.Stat("images/nx21519-F2z1QPS5GmpC.jpg"); os.IsNotExist(err) {
		t.Errorf("downloadImage is not downloading the file")
	}
}
