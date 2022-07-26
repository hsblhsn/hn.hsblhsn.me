package httpclient

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/caches"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	MaxResponseSize = 1024 * 1024 * 10 // 10MB
	UserAgent       = "HackerNews +https://hn.hsblhsn.me/"
)

type Client struct {
	httpClient *http.Client
	cache      caches.Cache
	logger     *zap.Logger
}

func NewClient(httpClient *http.Client, cache caches.Cache, logger *zap.Logger) *Client {
	return &Client{
		httpClient: httpClient,
		cache:      cache,
		logger:     logger.With(zap.String("component", "httpclient")),
	}
}

func (c *Client) Get(ctx context.Context, uri string, opts ...FilterOption) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, "httpclient: could not create request")
	}
	resp, err := c.Send(req, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "httpclient: could not send request")
	}
	return resp, nil
}

func (c *Client) Send(request *http.Request, opts ...FilterOption) (*http.Response, error) {
	// return cached response in the request is cached before.
	if resp, err := c.getFromCache(request); err == nil {
		return resp, nil
	}

	// limit maximum response size.
	// set useragent.
	request.Header.Set("Range", fmt.Sprintf("bytes=0-%d", MaxResponseSize))
	if request.Header.Get("User-Agent") == "" {
		request.Header.Set("User-Agent", UserAgent)
	}

	// send the request.
	c.logger.Debug("httpclient: sending http request", zap.String("uri", request.URL.String()))
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "httpclient: could not send request")
	}

	// check for low quality contents and other filters.
	for _, v := range opts {
		if optErr := v.apply(resp); optErr != nil {
			return nil, errors.Wrap(optErr, "httpclient: response is filtered")
		}
	}

	// read response till the maximum allowed size.
	// this prevents malicious urls
	// from flooding our memory by sending a large response body.
	resp.Body = http.MaxBytesReader(nil, resp.Body, MaxResponseSize)
	resp.ContentLength = -1

	// cache the respose.
	if err := c.setToCache(request, resp); err != nil {
		return nil, err
	}

	// recursively send the request.
	// this will get the response cached a moment ago.
	return c.Send(request, opts...)
}

func (c *Client) getFromCache(request *http.Request) (*http.Response, error) {
	if request.Method != http.MethodGet {
		return nil, errors.New("httpclient: only GET requests can be cached")
	}
	cachedVal, err := c.cache.Get(request.URL.String())
	if err != nil {
		return nil, errors.Wrap(err, "httpclient: response not found in cache")
	}
	reader := bufio.NewReader(bytes.NewReader(cachedVal))
	resp, err := http.ReadResponse(reader, nil)
	if err != nil {
		return nil, errors.Wrap(err, "httpclient: could not read cached response")
	}
	return resp, nil
}

func (c *Client) setToCache(request *http.Request, resp *http.Response) error {
	if request.Method != http.MethodGet {
		return errors.New("httpclient: only GET requests can be cached")
	}
	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return errors.Wrap(err, "httpclient: could not dump response")
	}
	if err := c.cache.Set(request.URL.String(), body); err != nil {
		return errors.Wrap(err, "httpclient: could not cache response")
	}
	return nil
}
