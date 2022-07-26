package opengraphs

import (
	"bytes"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/images"
	"strings"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/readerviews"
	"github.com/otiai10/opengraph/v2"
	"github.com/pkg/errors"
)

type (
	OpenGraph = opengraph.OpenGraph
	Image     = opengraph.Image
	Favicon   = opengraph.Favicon
)

func GetOpengraphData(uri string, content *bytes.Buffer) (*OpenGraph, error) {
	data := opengraph.New(uri)
	if err := data.Parse(content); err != nil {
		return nil, errors.Wrap(err, "opengraph: could not parse opengraph data")
	}
	if err := data.ToAbs(); err != nil {
		return nil, errors.Wrap(err, "opengraph: could not convert uri to absolute")
	}
	var err error
	data.Description, err = readerviews.TransformHTML(uri, strings.NewReader(data.Description))
	if err != nil {
		return nil, errors.Wrap(err, "opengraph: could not transform description")
	}
	for i := range data.Image {
		data.Image[i].URL = images.ProxiedURL(data.Image[i].URL, images.ImageSizeThumbnail)
	}
	if len(data.Image) == 0 {
		data.Image = []Image{
			{URL: images.SocialPreviewURL(data.Title)},
		}
	}
	if data.URL == "" {
		data.URL = uri
	}

	return data, nil
}
