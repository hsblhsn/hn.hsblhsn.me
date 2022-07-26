package images

import (
	"image"
	"image/jpeg"
	"net/http"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/httpclient"
	"go.uber.org/zap"
)

const quality = 85

type ImageProxyHandler struct {
	client *httpclient.Client
	logger *zap.Logger
}

func NewImageProxyHandler(client *httpclient.Client, logger *zap.Logger) *ImageProxyHandler {
	return &ImageProxyHandler{
		client: client,
		logger: logger.With(zap.String("component", "images_resize_handler")),
	}
}

func (h *ImageProxyHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// get image size from query string
	size := ImageSize(req.URL.Query().Get("size"))
	if size == "" {
		size = ImageSizeFull
	}
	// get image src from query string
	src := req.URL.Query().Get("src")
	if src == "" {
		h.logger.Error("no src provided", zap.String("size", string(size)))
		writeBlankImage(resp, http.StatusOK)
		return
	}
	// get image from src
	imgResp, err := h.client.Get(req.Context(), src)
	if err != nil {
		h.logger.Error("failed to get image", zap.String("size", string(size)), zap.String("src", src), zap.Error(err))
		writeBlankImage(resp, http.StatusOK)
		return
	}
	// resize image
	resized, err := Resize(imgResp.Body, size)
	if err != nil {
		h.logger.Error("failed to resize image", zap.String("size", string(size)), zap.String("src", src), zap.Error(err))
		writeBlankImage(resp, http.StatusOK)
		return
	}

	err = jpeg.Encode(resp, resized, &jpeg.Options{
		Quality: quality,
	})
	if err != nil {
		h.logger.Error("failed to encode image", zap.String("size", string(size)), zap.String("src", src), zap.Error(err))
		writeBlankImage(resp, http.StatusOK)
		return
	}
}

func blankImage() image.Image {
	return image.NewRGBA(image.Rect(0, 0, 1, 1))
}

func writeBlankImage(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	if err := jpeg.Encode(w, blankImage(), &jpeg.Options{Quality: quality}); err != nil {
		panic(err)
	}
}
