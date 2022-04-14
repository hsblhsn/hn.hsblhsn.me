package featureflags

import "strings"

const (
	FeatureImgProxy Feature = "image_proxy"
	FeatureBionify  Feature = "bionify"
)

type Feature string

func (f Feature) String() string {
	return "FEATURE_FLAG_" + strings.ToUpper(string(f))
}
