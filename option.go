package httputil

import (
	"net/http"
	"strings"
	"time"
)

// ClientOptions http client options
type ClientOptions struct {
	userAgent           string
	timeout             time.Duration
	proxy               string
	dialTimeout         time.Duration
	keepAliveTimeout    time.Duration
	tlsHandshakeTimeout time.Duration
}

// ClientOption http client option
type ClientOption func(*ClientOptions)

// WithUserAgent sed to replace the default UserAgent when making a request.
func WithUserAgent(userAgent string) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.userAgent = userAgent
	}
}

// WithTimeout If a timeout is set, each HTTP request will use this time as the maximum limit for completing the operation.
// If not set, the default timeout will be 3 minutes.
// If you want to use a different time limit for specific requests, you can override the default timeout by using WithRequestTimeout during the request.
func WithTimeout(timeout time.Duration) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.timeout = timeout
	}
}

// WithProxy If a proxy is set, each HTTP request will use this proxy to make the request.
func WithProxy(proxy string) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.proxy = strings.TrimSpace(proxy)
	}
}

// WithDialTimeout If a dial timeout is set, each HTTP request will use this time as the maximum limit for establishing a connection.
func WithDialTimeout(dialTimeout time.Duration) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.dialTimeout = dialTimeout
	}
}

// WithKeepAliveTimeout If a keep-alive timeout is set, each HTTP request will use this time as the maximum limit for keeping the connection alive.
func WithKeepAliveTimeout(keepAliveTimeout time.Duration) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.keepAliveTimeout = keepAliveTimeout
	}
}

// WithTLSHandshakeTimeout If a TLS handshake timeout is set, each HTTP request will use this time as the maximum limit for completing the TLS handshake.
func WithTLSHandshakeTimeout(tlsHandshakeTimeout time.Duration) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.tlsHandshakeTimeout = tlsHandshakeTimeout
	}
}

// RequestOptions http request options
type RequestOptions struct {
	headers     http.Header
	referer     string
	contentType string
	timeout     time.Duration
}

// RequestOption http request option
type RequestOption func(*RequestOptions)

// WithHeaders If headers are set, each HTTP request will use these headers to make the request.
func WithHeaders(headers http.Header) func(*RequestOptions) {
	return func(options *RequestOptions) {
		options.headers = headers
	}
}

// WithReferer If a referer is set, each HTTP request will use this referer to make the request.
func WithReferer(referer string) func(*RequestOptions) {
	return func(options *RequestOptions) {
		options.referer = strings.TrimSpace(referer)
	}
}

// WithContentType If a content type is set, each HTTP request will use this content type to make the request.
func WithContentType(contentType string) func(*RequestOptions) {
	return func(options *RequestOptions) {
		options.contentType = strings.TrimSpace(contentType)
	}
}

// WithRequestTimeout If a timeout is set, each HTTP request will use this time as the maximum limit for completing the operation.
func WithRequestTimeout(timeout time.Duration) func(*RequestOptions) {
	return func(options *RequestOptions) {
		options.timeout = timeout
	}
}
