package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	req := NewReq(
		WithAddress("https://yakuter.com"),
		WithTimeout(30*time.Second),
		WithContentType("application/json"),
	)

	fmt.Printf("%+v", req)
	// Output: &{request:0xc00014a000 client:0xc00007ac90 address:https://yakuter.com}
}

// Req is the main struct for requests
type Req struct {
	request *http.Request
	client  *http.Client
	address string
}

// ReqFunc is the mutator func
type ReqFunc func(*Req)

// NewReq initiates new Req instrance from given ReqFuncs
func NewReq(opts ...ReqFunc) *Req {
	req := &Req{
		request: &http.Request{
			Header: make(http.Header),
		},
		client: &http.Client{},
	}

	// The mutator loop
	for _, opt := range opts {
		opt(req)
	}

	return req
}

// WithTimeout changes the request timeout
func WithAddress(a string) ReqFunc {
	return func(r *Req) {
		r.address = a
	}
}

// WithTimeout changes the request timeout
func WithTimeout(d time.Duration) ReqFunc {
	return func(r *Req) {
		r.client.Timeout = d
	}
}

// WithHeaders sets request headers
func WithHeaders(headers map[string]string) ReqFunc {
	return func(r *Req) {
		if len(headers) > 0 {
			for k, v := range headers {
				r.request.Header.Set(k, v)
			}
		}
	}
}

// WithContentType sets content type of request
func WithContentType(contentType string) ReqFunc {
	return func(r *Req) {
		r.request.Header.Set("Content-Type", contentType)
	}
}

// WithCookie sets a cookie to the request
func WithCookie(c *http.Cookie) ReqFunc {
	return func(r *Req) {
		r.request.AddCookie(c)
	}
}

// WithTLSConfig changes the request TLS client configuration
func WithTLSConfig(c *tls.Config) ReqFunc {
	return func(r *Req) {
		r.client.Transport.(*http.Transport).TLSClientConfig = c
	}
}

//WithTransport sets transport configuration of request
func WithTransport(transport *http.Transport) ReqFunc {
	return func(r *Req) {
		r.client.Transport = transport
	}
}

// WithBody sets request body
func WithBody(data []byte) ReqFunc {
	return func(r *Req) {
		r.request.Body = ioutil.NopCloser(bytes.NewReader(data))
		r.request.GetBody = func() (io.ReadCloser, error) {
			return ioutil.NopCloser(bytes.NewReader(data)), nil
		}

		r.request.ContentLength = int64(len(data))
	}
}

//WithBodyXML sets content type as XML.
func WithBodyXML() ReqFunc {
	return func(r *Req) {
		WithContentType("application/xml; charset=UTF-8")
	}
}
