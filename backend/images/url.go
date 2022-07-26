package images

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/featureflags"
)

func ProxiedURL(src string, size ImageSize) string {
	if !featureflags.IsOn(featureflags.FeatureImgProxy, false) {
		return src
	}
	if src == "" {
		return ""
	}
	if strings.HasPrefix(src, "data:image") {
		return src
	}
	host := os.Getenv("DOMAIN")
	if host != "" {
		host = "https://" + host
	}
	escaped := url.QueryEscape(src)
	return fmt.Sprintf(
		"%s/images.jpeg?src=%s&size=%s",
		host,
		unescape(escaped),
		size,
	)
}

func SocialPreviewURL(title string) string {
	if !featureflags.IsOn(featureflags.FeatureImgSocialPreview, false) {
		return ""
	}
	host := os.Getenv("DOMAIN")
	if host != "" {
		host = "https://" + host
	}
	escaped := url.QueryEscape(title)
	return fmt.Sprintf(
		"%s/social_preview.jpeg?title=%s",
		host,
		unescape(escaped),
	)
}

func unescape(str string) string {
	str = strings.ReplaceAll(str, "%2C", `,`)
	str = strings.ReplaceAll(str, "%25", `%`)
	return str
}
