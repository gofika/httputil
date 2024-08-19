package httputil

import (
	"context"
	"io"
	"net/http"
	"net/url"
)

// Get issues a GET to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func Get(url string, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).Get(url, opts...)
}

// Head issues a HEAD to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func Head(url string, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).Head(url, opts...)
}

// Post issues a POST to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func Post(url string, contentType string, body io.Reader, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).Post(url, contentType, body, opts...)
}

// Put issues a PUT to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func Put(url string, contentType string, body io.Reader, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).Put(url, contentType, body, opts...)
}

// Patch issues a PATCH to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func Patch(url string, contentType string, body io.Reader, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).Patch(url, contentType, body, opts...)
}

// Delete issues a DELETE to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func Delete(url string, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).Delete(url, opts...)
}

// PostForm issues a POST to the specified URL, with data's keys and
// values URL-encoded as the request body.
//
// The Content-Type header is set to application/x-www-form-urlencoded.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func PostForm(url string, data url.Values, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).PostForm(url, data, opts...)
}

// UploadFile upload file struct
type UploadFile struct {
	FieldName string
	FileName  string
	Body      io.Reader
}

// PostFormFiles issues a POST to the specified URL, with data's keys and
// values URL-encoded as the request body, and files as multipart form.
//
// The Content-Type header is set to multipart/form-data.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func PostFormFiles(url string, data url.Values, files []*UploadFile, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).PostFormFiles(url, data, files, opts...)
}

// PostJSON issues a POST to the specified URL with the given body as JSON.
// values JSON as the request body.
//
// The Content-Type header is set to application/json.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func PostJSON(url string, body any, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).PostJSON(url, body, opts...)
}

// PutJSON issues a PUT to the specified URL with the given body as JSON.
// values JSON as the request body.
//
// The Content-Type header is set to application/json.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func PutJSON(url string, body any, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).PutJSON(url, body, opts...)
}

// PatchJSON issues a PATCH to the specified URL with the given body as JSON.
// values JSON as the request body.
//
// The Content-Type header is set to application/json.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func PatchJSON(url string, body any, opts ...RequestOption) (resp *http.Response, err error) {
	return NewClient(context.Background()).PatchJSON(url, body, opts...)
}
