package pkg_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/garysferrao/web-site-fetch/pkg"
)

func TestDownloadHTMLPage(t *testing.T) {
	var err error
	// set up temporary home directory
	var testHomeDir string
	testHomeDir, err = os.MkdirTemp(os.TempDir(), "test_download_html_page")
	if err != nil {
		t.Fatalf("cannot create temporary home directory: %+v", err)
		return
	}
	var oldHome = os.Getenv("HOME")
	err = os.Setenv("HOME", testHomeDir)
	if err != nil {
		t.Fatalf("cannot set home directory environment variable to temporary value for test: %+v", err)
		return
	}
	defer func() {
		var err error = os.Setenv("HOME", oldHome)
		if err != nil {
			t.Fatalf("cannot set home directory environment variable to its original value: %+v", err)
			return
		}
	}()
	t.Logf("home directory set is at %+v", testHomeDir)

	var tt = []struct {
		name            string
		mockHTTPHandler http.HandlerFunc
		expectContents  []byte
		expectError     bool
	}{
		{
			name: "success case",
			mockHTTPHandler: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "a success response")
			},
			expectContents: []byte("a success response"),
			expectError:    false,
		},
		// TODO (garysferrao): add more test cases
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			var mockServer *httptest.Server = httptest.NewServer(tc.mockHTTPHandler)
			defer mockServer.Close()
			var uri *url.URL
			uri, err = url.Parse(mockServer.URL)
			if err != nil || uri == nil {
				t.Fatalf("cannot parse URL for mock server: %+v", err)
				return
			}
			var outputFilePath string
			outputFilePath, err = pkg.DowloadHTMLPage(*uri)
			if (!tc.expectError && err != nil) || (tc.expectError && err == nil) {
				t.Errorf("expected error %t, got error: %+v", tc.expectError, err)
			}
			if (tc.expectContents == nil && outputFilePath != "") || (tc.expectContents != nil && outputFilePath == "") {
				t.Errorf("expected file pointer: %+v, got file path: %+v", tc.expectContents, outputFilePath)
				// do not dereference `outputFile` if it's not expected.
				return
			}
			var outputFile *os.File
			outputFile, err = os.Open(outputFilePath)
			if err != nil {
				t.Fatalf("cannot open output file: %+v", err)
				return
			}
			defer func() {
				var err error = outputFile.Close()
				if err != nil {
					t.Fatalf("cannot close output file: %+v", err)
					return
				}
			}()
			var outputFileContents []byte
			outputFileContents, err = io.ReadAll(outputFile)
			if err != nil {
				t.Fatalf("cannot read output file contents: %+v", err)
				return
			}
			if !bytes.Equal(tc.expectContents, outputFileContents) {
				t.Errorf("expected file contents %+v, got file contents: %+v", tc.expectContents, outputFileContents)
			}
		})
	}
}
