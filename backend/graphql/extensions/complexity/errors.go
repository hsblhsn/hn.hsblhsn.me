package complexity

import (
	"github.com/pkg/errors"
)

var (
	ErrLimitExceeded  = errors.New("complexity: limit exceeded")
	ErrNoCounterInCtx = errors.New("complexity: no counter in context")
	ErrCounterIsNil   = errors.New("complexity: counter is nil")
	ErrCounterExists  = errors.New("complexity: counter already exists in the context")
)

const (
	MsgComplexityLimitExceeded = "complexity: limit exceeded"
	MsgComplexityCounterError  = "complexity: counter error"
)
