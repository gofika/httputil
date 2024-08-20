package httputil

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"golang.org/x/net/publicsuffix"
)

type Client struct {
	client *http.Client
	ctx    context.Context
	opts   *ClientOptions
}

// NewClient new client
func NewClient(ctx context.Context, opts ...ClientOption) *Client {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	options := &ClientOptions{
		dialTimeout:         1 * time.Minute,
		keepAliveTimeout:    0,
		tlsHandshakeTimeout: 1 * time.Minute,
	}
	for _, opt := range opts {
		opt(options)
	}
	var transport http.RoundTripper
	if options.proxy != "" {
		transport = &http.Transport{
			Proxy: func(_ *http.Request) (*url.URL, error) {
				return url.Parse(options.proxy)
			},
			DialContext: (&net.Dialer{
				Timeout:   options.dialTimeout,
				KeepAlive: options.keepAliveTimeout,
			}).DialContext,
			TLSHandshakeTimeout: options.tlsHandshakeTimeout,
		}
	}
	return &Client{
		client: &http.Client{
			Jar:       jar,
			Transport: transport,
		},
		ctx:  ctx,
		opts: options,
	}
}

// Close close
func (c *Client) Close() {
	c.client.CloseIdleConnections()
}

func (c *Client) fillHeader(header http.Header, opts *RequestOptions) {
	if c.opts.userAgent != "" {
		header.Set("User-Agent", c.opts.userAgent)
	}
	if opts.referer != "" {
		header.Set("Referer", opts.referer)
	}
	if opts.contentType != "" {
		header.Set("Content-Type", opts.contentType)
	}
	for key, values := range opts.headers {
		for _, value := range values {
			if header.Get(key) != "" {
				continue
			}
			header.Add(key, value)
		}
	}
}

// Do sends an HTTP request and returns an HTTP response, following
// policy (such as redirects, cookies, auth) as configured on the
// client.
//
// An error is returned if caused by client policy (such as
// CheckRedirect), or failure to speak HTTP (such as a network
// connectivity problem). A non-2xx status code doesn't cause an
// error.
//
// If the returned error is nil, the [Response] will contain a non-nil
// Body which the user is expected to close. If the Body is not both
// read to EOF and closed, the [Client]'s underlying [RoundTripper]
// (typically [Transport]) may not be able to re-use a persistent TCP
// connection to the server for a subsequent "keep-alive" request.
//
// The request Body, if non-nil, will be closed by the underlying
// Transport, even on errors. The Body may be closed asynchronously after
// Do returns.
//
// On error, any Response can be ignored. A non-nil Response with a
// non-nil error only occurs when CheckRedirect fails, and even then
// the returned [Response.Body] is already closed.
//
// Generally [Get], [Post], or [PostForm] will be used instead of Do.
//
// If the server replies with a redirect, the Client first uses the
// CheckRedirect function to determine whether the redirect should be
// followed. If permitted, a 301, 302, or 303 redirect causes
// subsequent requests to use HTTP method GET
// (or HEAD if the original request was HEAD), with no body.
// A 307 or 308 redirect preserves the original HTTP method and body,
// provided that the [Request.GetBody] function is defined.
// The [NewRequest] function automatically sets GetBody for common
// standard library body types.
//
// Any returned error will be of type [*url.Error]. The url.Error
// value's Timeout method will report true if the request timed out.
func (c *Client) Do(req *http.Request, opts ...RequestOption) (resp *http.Response, err error) {
	options := &RequestOptions{}
	for _, opt := range opts {
		opt(options)
	}
	c.fillHeader(req.Header, options)
	return c.client.Do(req)
}

// Get issues a GET to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) Get(url string, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}
	return c.Do(req, opts...)
}

// Head issues a HEAD to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) Head(url string, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodHead, url, nil)
	if err != nil {
		return
	}
	return c.Do(req, opts...)
}

// Post issues a POST to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) Post(url string, contentType string, body io.Reader, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, url, body)
	if err != nil {
		return
	}
	return c.Do(req, append([]RequestOption{WithContentType(contentType)}, opts...)...)
}

// Put issues a PUT to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) Put(url string, contentType string, body io.Reader, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodPut, url, body)
	if err != nil {
		return
	}
	return c.Do(req, append([]RequestOption{WithContentType(contentType)}, opts...)...)
}

// Patch issues a PATCH to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) Patch(url string, contentType string, body io.Reader, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodPatch, url, body)
	if err != nil {
		return
	}
	return c.Do(req, append([]RequestOption{WithContentType(contentType)}, opts...)...)
}

// Delete issues a DELETE to the specified URL.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) Delete(url string, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodDelete, url, nil)
	if err != nil {
		return
	}
	return c.Do(req, opts...)
}

// PostForm issues a POST to the specified URL, with data's keys and
// values URL-encoded as the request body.
//
// The Content-Type header is set to application/x-www-form-urlencoded.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) PostForm(url string, data url.Values, opts ...RequestOption) (resp *http.Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()), opts...)
}

// PostFormFiles issues a POST to the specified URL, with data's keys and
// values URL-encoded as the request body, and files as multipart form.
//
// The Content-Type header is set to multipart/form-data.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) PostFormFiles(url string, data url.Values, files []*UploadFile, opts ...RequestOption) (resp *http.Response, err error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	contentType := writer.FormDataContentType()
	for k, i := range data {
		for _, j := range i {
			err = writer.WriteField(k, j)
			if err != nil {
				return nil, err
			}
		}
	}
	for _, file := range files {
		part, err := writer.CreateFormFile(file.FieldName, file.FileName)
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(part, file.Body); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}
	return c.Do(req, append([]RequestOption{WithContentType(contentType)}, opts...)...)
}

// PostJSON issues a POST to the specified URL with the given body as JSON.
// values JSON as the request body.
//
// The Content-Type header is set to application/json.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) PostJSON(url string, body any, opts ...RequestOption) (resp *http.Response, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return
	}
	return c.Post(url, "application/json", bytes.NewBuffer(payload), opts...)
}

// PutJSON issues a PUT to the specified URL with the given body as JSON.
// values JSON as the request body.
//
// The Content-Type header is set to application/json.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) PutJSON(url string, body any, opts ...RequestOption) (resp *http.Response, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return
	}
	return c.Put(url, "application/json", bytes.NewBuffer(payload), opts...)
}

// PatchJSON issues a PATCH to the specified URL with the given body as JSON.
// values JSON as the request body.
//
// The Content-Type header is set to application/json.
//
// When err is nil, resp always contains a non-nil resp.Body.
// Caller should close resp.Body when done reading from it.
func (c *Client) PatchJSON(url string, body any, opts ...RequestOption) (resp *http.Response, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return
	}
	return c.Patch(url, "application/json", bytes.NewBuffer(payload), opts...)
}
