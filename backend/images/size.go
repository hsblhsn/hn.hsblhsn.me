package images

type ImageSize string

func (f ImageSize) String() string {
	switch f {
	case ImageSizeThumbnail, ImageSizeFull:
		return string(f)
	default:
		return string(ImageSizeFull)
	}
}

// nolint:gomnd // dimensions are magic numbers.
func (f ImageSize) Dimension() (height, width uint) {
	switch f {
	case ImageSizeThumbnail:
		return 180, 180
	case ImageSizeFull:
		return 560, 560
	default:
		return 560, 560
	}
}

const (
	// ImageSizeThumbnail is the size of the thumbnail image.
	ImageSizeThumbnail ImageSize = "thumbnail"
	// ImageSizeFull is the size of the full image.
	ImageSizeFull ImageSize = "full"
)
