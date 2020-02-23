package metadata

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func downloadImage(imageURL string) {

	if len(imageURL) < 1 {
		return
	}

	client := http.Client{
		Timeout: 20 * time.Second,
	}

	resp, err := client.Get(imageURL)
	checkErr(err)

	fileName := filepath.Base(imageURL)

	defer resp.Body.Close()

	file, err := os.Create(filepath.Join("images", fileName))
	if os.IsNotExist(err) {
		os.Mkdir("images", 0700)
		file, err = os.Create(filepath.Join("images", fileName))
	}

	size, err := io.Copy(file, resp.Body)
	checkErr(err)

	defer file.Close()

	fmt.Printf("downloaded image %s of size %d\n", fileName, size)
}
