package httpclient

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

var ErrContentTypeNotAccepted = errors.New("httpclient: content type is not accepted")

type FilterOption interface {
	apply(r *http.Response) error
}

type FilterOptionFunc func(r *http.Response) error

var _ FilterOption = (FilterOptionFunc)(nil)

func (f FilterOptionFunc) apply(r *http.Response) error {
	return f(r)
}

func WithAcceptedContentTypes(cTypes []string) FilterOption {
	const maxSplitCount = 2
	return FilterOptionFunc(func(r *http.Response) error {
		cType := r.Header.Get("Content-Type")
		cType = strings.SplitN(cType, ";", maxSplitCount)[0]
		for _, v := range cTypes {
			if strings.EqualFold(v, cType) {
				return nil
			}
		}
		return errors.Wrapf(ErrContentTypeNotAccepted, "httpclient: got content type %q", cType)
	})
}
