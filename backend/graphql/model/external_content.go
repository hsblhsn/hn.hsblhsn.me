package model

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"sync"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/cquality"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/httpclient"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/opengraphs"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/readerviews"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// nolint:gochecknoglobals // global client to call from ExternalContentLoader.
var (
	client *httpclient.CachedClient
	logger *zap.Logger
)

func registerDependencies(c *httpclient.CachedClient, l *zap.Logger) {
	client = c
	logger = l.With(zap.String("component", "model_external_content"))
}

type ExternalContentLoadable struct {
	ext  *ExternalContentLoader
	once sync.Once
}

func (s *ExternalContentLoadable) GetLoader(url string) *ExternalContentLoader {
	s.once.Do(func() {
		s.ext = NewExternalContentLoader(url)
	})
	return s.ext
}

type ExternalContentLoader struct {
	url     string
	content []byte
	once    sync.Once
	mu      sync.Mutex
}

func NewExternalContentLoader(url string) *ExternalContentLoader {
	return &ExternalContentLoader{url: url}
}

func (s *ExternalContentLoader) getContentFromURL(ctx context.Context) error {
	if s.url == "" {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	var err error
	s.once.Do(func() {
		var resp *http.Response
		resp, err = client.Get(ctx, s.url)
		if err != nil {
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		s.content, err = io.ReadAll(resp.Body)
	})
	if err != nil {
		return errors.Wrap(err, "model: could not get content from url")
	}
	quality := cquality.LowQuality()
	if quality.Indicates(s.content) {
		return errors.Errorf("model: content quality is low on: %q", s.url)
	}
	return nil
}

func (s *ExternalContentLoader) Opengraph(ctx context.Context) *opengraphs.OpenGraph {
	if err := s.getContentFromURL(ctx); err != nil {
		logger.Error(
			"model: could not get content from url",
			zap.Error(err),
			zap.String("url", s.url),
		)
		return nil
	}

	if s.url == "" {
		return nil
	}
	out, err := opengraphs.GetOpengraphData(s.url, bytes.NewBuffer(s.content))
	if err != nil {
		logger.Error("model: could not parse opengraph content",
			zap.Error(err),
			zap.String("url", s.url),
		)
		return nil
	}
	return out
}

func (s *ExternalContentLoader) HTML(ctx context.Context) *string {
	if err := s.getContentFromURL(ctx); err != nil {
		logger.Error("model: could not get content from url",
			zap.Error(err),
			zap.String("url", s.url),
		)
		return nil
	}

	if s.url == "" {
		return nil
	}
	out, err := readerviews.Convert(ctx, s.url, bytes.NewBuffer(s.content))
	if err != nil {
		logger.Error("model: could not convert readerview content",
			zap.Error(err),
			zap.String("url", s.url),
		)
		return nil
	}
	return &out
}
