package johnny

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func downloadAndSaveFile(url string) (string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", errors.New("failed to download the file := " + err.Error())
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		log.Printf("bad status: %s\n", resp.Status)
		return "", errors.New("bad status: %s" + resp.Status)
	}

	// Writer the body to file
	tmpFile, err := ioutil.TempFile("", fmt.Sprintf("%s-", filepath.Base(os.Args[0])))
	if err != nil {
		return "", errors.New("not able to create tmp file while downloading ... " + err.Error())
	}
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		log.Println("error while writing to a tmp file.")
		return "", errors.New("error while writing to a tmp file." + err.Error())
	}

	return tmpFile.Name(), nil
}
