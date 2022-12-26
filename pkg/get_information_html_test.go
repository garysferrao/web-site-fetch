package pkg_test

import (
	"io"
	"strings"
	"testing"

	"github.com/garysferrao/web-site-fetch/pkg"
)

func TestGetInformationHTML(t *testing.T) {
	var tt = []struct {
		name            string
		inputContents   io.Reader
		expectNumLinks  int
		expectNumImages int
		expectError     bool
	}{
		{
			name:            "HTML page contains one link and one image",
			inputContents:   strings.NewReader(`<html><body><a href="https://www.example.com"></a><img src="https://www.example.com/image.png" /></body></html>`),
			expectNumLinks:  1,
			expectNumImages: 1,
			expectError:     false,
		},
		// TODO (garysferrao): add more test cases
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var outNumLinks, outNumImages int
			var err error
			outNumLinks, outNumImages, err = pkg.GetInformationHTML(tc.inputContents)
			if (!tc.expectError && err != nil) || (tc.expectError && err == nil) {
				t.Errorf("expected error %t, got error: %+v", tc.expectError, err)
			}
			if tc.expectNumLinks != outNumLinks {
				t.Errorf("expected number of links %d, got number of links: %d", tc.expectNumLinks, outNumLinks)
			}
			if tc.expectNumImages != outNumImages {
				t.Errorf("expected number of images %d, got number of images: %d", tc.expectNumImages, outNumImages)
			}
		})
	}
}
