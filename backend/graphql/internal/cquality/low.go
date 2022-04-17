package cquality

import (
	"sync"
)

// initLowQualityIndicators initializes the low quality indicators.
// this functions exists to lazy load the indicators.
func initLowQualityIndicators() {
	list := []string{
		// bloomberg
		"Are you a robot",
		// cloudflare
		"You have been blocked",
		// github
		"You can’t perform that action at this time.",
		// twitter
		"switch to a supported browser to continue using twitter.com",
		// youtube
		"https://www.youtube.com/howyoutubeworks",
		// 403
		"403 Forbidden",
		// 404
		"Page Not Found",
		// Cloudflare bot check
		"Checking your browser before accessing",
		// nytimes
		"Please enable JS and disable any ad blocker",
	}
	lowQualityIndicators = NewIndicators(list)
}

// nolint:gochecknoglobals // global for lazy loading
var (
	lowQualityIndicators Indicators
	lowQualityOnce       sync.Once
)

// LowQuality returns the indicators for low quality content.
func LowQuality() Indicators {
	lowQualityOnce.Do(initLowQualityIndicators)
	return lowQualityIndicators
}
