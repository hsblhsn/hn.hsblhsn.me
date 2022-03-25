package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"
	"sync"

	"go.uber.org/zap"
)

// ErrorMsg represents an JSON error message.
type ErrorMsg struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// String implements fmt.Stringer interface.
func (m ErrorMsg) String() string {
	m.Error = true
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// HTTPError returns an error message as a JSON response.
func HTTPError(w http.ResponseWriter, err error, code int, msg string) {
	if err != nil {
		zap.L().Error(
			"error occurred in handler",
			zap.Error(err),
			zap.Int("code", code),
			zap.String("msg", msg),
		)
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	fmt.Fprintln(w, ErrorMsg{Message: msg})
}

var (
	_ImgErrBytes []byte
	_ImgErrOnce  sync.Once
)

func ImgErr(w http.ResponseWriter, err error, code int, msg string) {
	_ImgErrOnce.Do(func() {
		img := image.NewRGBA(image.Rectangle{})
		writer := bytes.NewBuffer(_ImgErrBytes)
		if err := jpeg.Encode(writer, img, nil); err != nil {
			panic(err)
		}
	})
	w.Header().Set("Content-Type", "image/jpeg")
	_, _ = w.Write(_ImgErrBytes)
}
