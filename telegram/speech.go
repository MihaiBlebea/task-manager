package telegram

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func downloadFile(url string) io.Reader {
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	return bytes.NewReader(b)
}

func downloadFileTmp(url string) (filepath string, err error) {
	log.Printf("> downloading voice file: %s\n", url)

	var file *os.File
	if file, err = ioutil.TempFile("/tmp", "downloaded_"); err == nil {
		filepath = file.Name()

		defer file.Close()

		var response *http.Response
		if response, err = http.Get(url); err == nil {
			defer response.Body.Close()

			if _, err = io.Copy(file, response.Body); err == nil {
				log.Printf("> finished downloading voice file: %s\n", filepath)
			}
		}
	}

	return filepath, err
}

func oggToMp3(oggFilepath string) (mp3Filepath string, err error) {
	mp3Filepath = fmt.Sprintf("%s.mp3", oggFilepath)

	// $ ffmpeg -i input.ogg -ac 1 output.mp3
	params := []string{"-i", oggFilepath, "-ac", "1", mp3Filepath}
	cmd := exec.Command("ffmpeg", params...)

	if _, err = cmd.CombinedOutput(); err != nil {
		mp3Filepath = ""
	}

	return mp3Filepath, err
}
