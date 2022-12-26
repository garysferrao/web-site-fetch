package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/garysferrao/web-site-fetch/pkg"
)

var urlList []string

var errLog *log.Logger
var stdLog *log.Logger

func init() {
	errLog = log.New(os.Stderr, "", 0)
	stdLog = log.New(os.Stdout, "", 0)
	flag.Parse()
	urlList = flag.Args()
}

func main() {
	var err error
	for _, downloadURL := range urlList {
		var inputURL *url.URL
		inputURL, err = url.Parse(downloadURL)
		if err != nil || inputURL == nil {
			errLog.Printf("cannot parse URL '%s': %+v", downloadURL, err)
			continue
		}
		var outputFilePath string
		outputFilePath, err = pkg.DowloadHTMLPage(*inputURL)
		if err != nil || outputFilePath == "" {
			errLog.Printf("cannot download HTML page '%s': %+v", downloadURL, err)
			continue
		}
		var outputFile *os.File
		outputFile, err = os.Open(outputFilePath)
		if err != nil {
			errLog.Printf("cannot open downloaded HTML page '%s': %+v", outputFilePath, err)
			continue
		}
		defer func(f *os.File) {
			var err error = f.Close()
			if err != nil {
				errLog.Printf("cannot close file '%+v': %+v", f, err)
			}
		}(outputFile)
		var outputFileStat os.FileInfo
		outputFileStat, err = outputFile.Stat()
		if err != nil {
			errLog.Printf("cannot get file information for downloaded HTML page '%s': %+v", outputFilePath, err)
			continue
		}
		var lastModifiedDate time.Time = outputFileStat.ModTime()
		var numLinks, numImages int
		numLinks, numImages, err = pkg.GetInformationHTML(outputFile)
		if err != nil {
			errLog.Printf("cannot get information for downloaded HTML page '%s': %+v", outputFilePath, err)
			continue
		}
		stdLog.Printf("site: '%s'\nnum_links: %d\nnum_images: %d\nlast_fetch: %+v\n", inputURL.Host, numLinks, numImages, lastModifiedDate)
	}
}
