package httputil

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// Get get
func Get(url string, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).Get(url, opts...)
}

// PostJSON post json
func PostJSON(url string, body any, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).PostJSON(url, body, opts...)
}

// PostForm post form
func PostForm(url string, data url.Values, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).PostForm(url, data, opts...)
}

// UploadFile upload file struct
type UploadFile struct {
	FieldName string
	FileName  string
	Body      io.Reader
}

// PostFormFiles post form files
func PostFormFiles(url string, data url.Values, files []*UploadFile, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).PostFormFiles(url, data, files, opts...)
}
