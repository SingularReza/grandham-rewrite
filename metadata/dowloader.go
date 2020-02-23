package metadata

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func downloadImage(imageURL string) {

	if len(imageURL) < 1 {
		return
	}

	fmt.Println(imageURL)
	resp, err := http.Get(imageURL)
	checkErr(err)

	defer resp.Body.Close()

	file, err := os.Create("images/" + filepath.Base(imageURL))
	checkErr(err)

	size, err := io.Copy(file, resp.Body)
	checkErr(err)

	defer file.Close()

	fmt.Printf("downloaded file %s of size %d\n", imageURL, size)
}
