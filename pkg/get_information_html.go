package pkg

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func GetInformationHTML(contents io.Reader) (numLinks int, numImages int, err error) {
	var z *html.Tokenizer = html.NewTokenizer(contents)
	if z == nil {
		err = fmt.Errorf("cannot parse contents: could not create HTML Tokenizer")
		return
	}
	for {
		var tokenType html.TokenType = z.Next()
		if tokenType == html.ErrorToken {
			// nothing more to parse
			return
		} else if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken || tokenType == html.TextToken {
			switch z.Token().Data {
			case "a":
				numLinks++
			case "img":
				numImages++
			}
		}

	}
}
