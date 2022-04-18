package httpclient

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/caches"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	MaxResponseSize = 1024 * 1024 * 10 // 10MB
	UserAgent       = "HackerNews[bot] +https://hn.hsblhsn.me/"
)

type CachedClient struct {
	httpClient *http.Client
	cache      caches.Cache
	logger     *zap.Logger
}

func NewCachedClient(httpClient *http.Client, cache caches.Cache, logger *zap.Logger) *CachedClient {
	return &CachedClient{
		httpClient: httpClient,
		cache:      cache,
		logger:     logger.With(zap.String("component", "httpclient")),
	}
}

func (c *CachedClient) Get(ctx context.Context, uri string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, "httpclient: could not create request")
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "httpclient: could not send request")
	}
	return resp, nil
}

func (c *CachedClient) Do(request *http.Request) (*http.Response, error) {
	uri := request.URL.String()
	if cachedVal, err := c.cache.Get(uri); err == nil {
		c.logger.Debug("httpclient: found cached response", zap.String("uri", uri))
		reader := bufio.NewReader(bytes.NewReader(cachedVal))
		resp, err := http.ReadResponse(reader, nil)
		if err != nil {
			return nil, errors.Wrap(err, "httpclient: could not read cached response")
		}
		return resp, nil
	}

	c.logger.Debug("httpclient: sending http request", zap.String("uri", uri))
	request.Header.Set("Range", fmt.Sprintf("bytes=0-%d", MaxResponseSize))
	request.Header.Set("User-Agent", UserAgent)
	resp, err := c.httpClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "httpclient: could not send request")
	}

	resp.Body = http.MaxBytesReader(nil, resp.Body, MaxResponseSize)
	resp.ContentLength = -1

	body, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, errors.Wrap(err, "httpclient: could not dump response")
	}
	if err := c.cache.Set(uri, body); err != nil {
		return nil, errors.Wrap(err, "httpclient: could not cache response")
	}
	return c.Do(request)
}
