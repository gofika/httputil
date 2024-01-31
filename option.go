package httputil

import (
	"net/http"
	"strings"
	"time"
)

// ClientOptions http client options
type ClientOptions struct {
	userAgent             string
	timeout               time.Duration
	proxy                 string
	dialTimeout           time.Duration
	keepAliveTimeout      time.Duration
	maxIdleConns          int
	idleConnTimeout       time.Duration
	tlsHandshakeTimeout   time.Duration
	expectContinueTimeout time.Duration
}

type ClientOption func(*ClientOptions)

func WithUserAgent(userAgent string) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.userAgent = userAgent
	}
}

func WithTimeout(timeout time.Duration) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.timeout = timeout
	}
}

func WithProxy(proxy string) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.proxy = strings.TrimSpace(proxy)
	}
}

func WithDialTimeout(dialTimeout time.Duration) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.dialTimeout = dialTimeout
	}
}

func WithKeepAliveTimeout(keepAliveTimeout time.Duration) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.keepAliveTimeout = keepAliveTimeout
	}
}

func WithMaxIdleConns(maxIdleConns int) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.maxIdleConns = maxIdleConns
	}
}

func WithIdleConnTimeout(idleConnTimeout time.Duration) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.idleConnTimeout = idleConnTimeout
	}
}

func WithTLSHandshakeTimeout(tlsHandshakeTimeout time.Duration) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.tlsHandshakeTimeout = tlsHandshakeTimeout
	}
}

func WithExpectContinueTimeout(expectContinueTimeout time.Duration) func(*ClientOptions) {
	return func(options *ClientOptions) {
		options.expectContinueTimeout = expectContinueTimeout
	}
}

type RequestOptions struct {
	headers     http.Header
	referer     string
	contentType string
	timeout     time.Duration
}

type RequestOption func(*RequestOptions)

func WithHeaders(headers http.Header) func(*RequestOptions) {
	return func(options *RequestOptions) {
		options.headers = headers
	}
}

func WithReferer(referer string) func(*RequestOptions) {
	return func(options *RequestOptions) {
		options.referer = strings.TrimSpace(referer)
	}
}

func WithContentType(contentType string) func(*RequestOptions) {
	return func(options *RequestOptions) {
		options.contentType = strings.TrimSpace(contentType)
	}
}

func WithRequestTimeout(timeout time.Duration) func(*RequestOptions) {
	return func(options *RequestOptions) {
		options.timeout = timeout
	}
}
