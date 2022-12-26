package main

import (
	"flag"
	"log"
	"net/url"
	"os"

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
		stdLog.Printf("downloaded '%s' to %s", downloadURL, outputFilePath)
	}
}
