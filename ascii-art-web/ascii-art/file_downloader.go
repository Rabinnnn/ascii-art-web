package asciiart

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Acts by replacing a non-existing banner file when called
// it does this by fetching the raw file from a remote server and copying the body into a recreated empty file of the same name
func Download(fileName string) error {
	response, err := http.Get("https://learn.zone01kisumu.ke/git/root/public/raw/subjects/ascii-art/" + fileName[18:])
	if err != nil {
		return fmt.Errorf("URL failed")
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed; status code %d", response.StatusCode)
	}
	file, er := os.Create(fileName)
	if er != nil {
		return fmt.Errorf("error creating banner file")
	}
	defer file.Close()
	_, e := io.Copy(file, response.Body)
	return e
}
