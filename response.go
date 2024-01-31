package httputil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gofika/fileutil"
	"github.com/gofika/regexputil"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/ianaindex"
	"golang.org/x/text/transform"
)

// ReadAll read all data from resp
func ReadAll(resp *http.Response) (b []byte, err error) {
	b, err = io.ReadAll(resp.Body)
	return
}

// ReadString read string from resp.Body and auto convert encoding
func ReadString(resp *http.Response) (s string, err error) {
	var bodyReader io.Reader
	bodyReader = resp.Body
	// auto convert encoding
	contentType := resp.Header.Get("Content-Type")
	contentCharset, matched := regexputil.Match(`charset=([\(\):\.\w-]+)`, contentType)
	if matched && strings.ToLower(contentCharset) != "utf-8" { // convert for all not utf-8 charset
		var enc encoding.Encoding
		enc, err = ianaindex.IANA.Encoding(contentCharset)
		if err != nil {
			return
		}
		bodyReader = transform.NewReader(bodyReader, enc.NewDecoder())
	}
	b, err := io.ReadAll(bodyReader)
	s = string(b)
	return
}

// ReadAnyJSON read json from resp.Body
func ReadAnyJSON(resp *http.Response, v any) (err error) {
	err = json.NewDecoder(resp.Body).Decode(v)
	return
}

// ReadJSON read json from resp.Body
func ReadJSON[T any](resp *http.Response) (T, error) {
	var v T
	err := ReadAnyJSON(resp, &v)
	return v, err
}

// SaveFile save file from resp.Body
func SaveFile(resp *http.Response, name string) (written int64, err error) {
	var f *os.File
	f, err = fileutil.OpenWrite(name)
	if err != nil {
		return
	}
	written, err = io.Copy(f, resp.Body)
	if err != nil {
		return
	}
	err = f.Close()
	return
}
