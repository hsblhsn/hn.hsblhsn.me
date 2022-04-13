package images

import (
	"image"
	// for image.Decode.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"

	"github.com/nfnt/resize"
	"github.com/pkg/errors"
)

func Resize(content io.Reader, size ImageSize) (image.Image, error) {
	img, _, err := image.Decode(content)
	if err != nil {
		return nil, errors.Wrap(err, "images: could not decode image")
	}
	// resize image
	width, height := size.Dimension()
	img = resize.Thumbnail(width, height, img, resize.NearestNeighbor)
	return img, nil
}
