package featureflags

import (
	"os"
)

func IsOn(feature Feature, def bool) bool {
	val := os.Getenv(feature.String())
	switch val {
	case "on":
		return true
	case "off":
		return false
	default:
		return def
	}
}
