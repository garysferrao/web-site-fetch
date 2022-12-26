package pkg

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func DowloadHTMLPage(inputURL url.URL) (outputFilePath string, err error) {
	var userHomeDir string
	userHomeDir, err = os.UserHomeDir()
	if err != nil {
		outputFilePath = ""
		err = fmt.Errorf("cannot get user home directory: %w", err)
		return
	}
	outputFilePath = filepath.Join(userHomeDir, inputURL.Host+".html")
	var outputFile *os.File
	outputFile, err = os.Create(outputFilePath)
	if err != nil {
		outputFilePath = ""
		err = fmt.Errorf("cannot create file for downloading: %w", err)
		return
	}
	var httpResponse *http.Response
	httpResponse, err = http.Get(inputURL.String())
	if err != nil || httpResponse == nil {
		outputFilePath = ""
		err = fmt.Errorf("cannot download HTML file: %w", err)
		return
	}

	var writtenBytes int64
	writtenBytes, err = io.Copy(outputFile, httpResponse.Body)
	if err != nil {
		outputFilePath = ""
		err = fmt.Errorf("cannot write output file: %w", err)
		return
	} else if writtenBytes == 0 {
		outputFilePath = ""
		err = fmt.Errorf("zero written bytes: %w", err)
		return
	}

	err = outputFile.Close()
	if err != nil {
		outputFilePath = ""
		err = fmt.Errorf("cannot close output file: %w", err)
		return
	}

	return
}
