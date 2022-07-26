package readerviews

import (
	"bytes"
	"context"
	"fmt"
	images2 "github.com/hsblhsn/hn.hsblhsn.me/backend/images"
	"io"
	"net/url"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/grpc/readabilityclient"
	"github.com/pkg/errors"
)

const (
	ClientReadinessTimeout = time.Second * 3
)

func Convert(ctx context.Context, link string, content *bytes.Buffer) (string, error) {
	// convert the webpage to a readable html document
	ready := isReadabilityClientReady(ctx, ClientReadinessTimeout)
	if !ready {
		return "", errors.New("hackernews: readability client not ready")
	}
	rc := readabilityClient()
	resp, err := rc.GetReadableDocument(ctx, &readabilityclient.GetReadableDocumentRequest{
		Html:       content.String(),
		Identifier: link,
	})
	if err != nil {
		return "", errors.Wrap(err, "hackernews: could not call readability server")
	}
	return TransformHTML(link, strings.NewReader(resp.GetBody()))
}

func TransformHTML(link string, content io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(content)
	if err != nil {
		return "", errors.Wrap(err, "readerviews: could not parse html")
	}
	doc.Find("img").Each(func(i int, selection *goquery.Selection) {
		src, ok := selection.Attr("src")
		if !ok || src == "" {
			selection.Remove()
			return
		}
		absSrc := toAbs(link, src)
		proxied := images2.ProxiedURL(absSrc, images2.ImageSizeFull)
		selection.SetAttr("src", proxied)
	})
	doc.Find("a").Each(func(i int, selection *goquery.Selection) {
		href, ok := selection.Attr("href")
		if !ok || href == "" {
			selection.Remove()
			return
		}
		absLink, isHnItemLink := toHNLink(toAbs(link, href))
		selection.SetAttr("href", absLink)
		if !isHnItemLink {
			selection.SetAttr("target", "_blank")
		}
	})
	htmlContent, err := doc.Html()
	if err != nil {
		return "", errors.Wrap(err, "readerviews: error while rendering html")
	}
	return sanitizationPolicy.Sanitize(htmlContent), nil
}

// toAbs returns a absolute url if relpath is relative.
func toAbs(base, relpath string) string {
	baseURL, err := url.Parse(base)
	if err != nil || baseURL.Host == "" {
		return relpath
	}
	if baseURL.Scheme == "" {
		baseURL.Scheme = "https"
	}

	relURL, err := url.Parse(relpath)
	if err == nil && relURL.IsAbs() {
		return relURL.String()
	}
	if strings.HasPrefix(relpath, "//") {
		return fmt.Sprintf("%s:%s", baseURL.Scheme, relpath)
	}
	if strings.HasPrefix(relpath, "/") {
		return fmt.Sprintf("%s://%s%s", baseURL.Scheme, baseURL.Host, relpath)
	}
	return fmt.Sprintf("%s://%s%s", baseURL.Scheme, baseURL.Host, path.Join(baseURL.Path, relpath))
}

var hnItemLinkRx = regexp.MustCompile(`^https://news.ycombinator.com/item\?id=(\d{1,16})$`)

func toHNLink(link string) (string, bool) {
	if !hnItemLinkRx.MatchString(link) {
		return link, false
	}
	itemID := hnItemLinkRx.ReplaceAllString(link, "$1")
	return fmt.Sprintf("/item?id=%s", itemID), true
}
