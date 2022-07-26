package images

import (
	"go.uber.org/zap"
	"image/jpeg"
	"net/http"
)

type SocialPreviewHandler struct {
	generator *SocialPreviewGenerator
	logger    *zap.Logger
}

func NewSocialPreviewHandler(generator *SocialPreviewGenerator, logger *zap.Logger) *SocialPreviewHandler {
	return &SocialPreviewHandler{
		generator: generator,
		logger:    logger,
	}
}

func (h *SocialPreviewHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	title := req.URL.Query().Get("title")
	if title == "" {
		title = "Hackernews client focused on content and readability"
	}
	img, err := h.generator.Generate(title)
	if err != nil {
		h.logger.Error(
			"failed to encode generated social preview image",
			zap.String("title", title),
			zap.Error(err),
		)
		writeBlankImage(resp, http.StatusOK)
		return
	}
	err = jpeg.Encode(resp, img, &jpeg.Options{
		Quality: 85,
	})
	if err != nil {
		h.logger.Error(
			"failed to encode generated social preview image",
			zap.String("title", title),
			zap.Error(err),
		)
		writeBlankImage(resp, http.StatusOK)
		return
	}
}
