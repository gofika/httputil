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
		timeout:               time.Minute,
		dialTimeout:           30 * time.Second,
		keepAliveTimeout:      30 * time.Second,
		maxIdleConns:          100,
		idleConnTimeout:       90 * time.Second,
		tlsHandshakeTimeout:   10 * time.Second,
		expectContinueTimeout: 1 * time.Second,
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
			MaxIdleConns:          options.maxIdleConns,
			IdleConnTimeout:       options.idleConnTimeout,
			TLSHandshakeTimeout:   options.tlsHandshakeTimeout,
			ExpectContinueTimeout: options.expectContinueTimeout,
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

// Do do
func (c *Client) Do(req *http.Request, opts ...RequestOption) (resp *http.Response, err error) {
	options := &RequestOptions{}
	for _, opt := range opts {
		opt(options)
	}
	c.fillHeader(req.Header, options)
	timeout := c.opts.timeout
	var timeoutZero time.Duration
	if options.timeout != timeoutZero {
		timeout = options.timeout
	}
	if timeout == timeoutZero {
		return c.client.Do(req)
	}
	ctx, cancel := context.WithTimeout(c.ctx, timeout)
	defer cancel()
	return c.client.Do(req.WithContext(ctx))
}

// Get get
func (c *Client) Get(url string, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}
	return c.Do(req, opts...)
}

// Post post
func (c *Client) Post(url string, contentType string, body io.Reader, opts ...RequestOption) (resp *http.Response, err error) {
	req, err := http.NewRequestWithContext(c.ctx, http.MethodPost, url, body)
	if err != nil {
		return
	}
	return c.Do(req, append([]RequestOption{WithContentType(contentType)}, opts...)...)
}

// PostForm post form
func (c *Client) PostForm(url string, data url.Values, opts ...RequestOption) (resp *http.Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded", bytes.NewBufferString(data.Encode()), opts...)
}

// PostJSON post json
func (c *Client) PostJSON(url string, body any, opts ...RequestOption) (resp *http.Response, err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return
	}
	return c.Post(url, "application/json", bytes.NewBuffer(payload), opts...)
}

// PostFormFiles post form files
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
	return c.Do(req, append([]RequestOption{WithContentType(contentType), WithRequestTimeout(time.Hour)}, opts...)...)
}
