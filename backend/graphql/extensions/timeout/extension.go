package timeout

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Extension struct {
	timeout time.Duration
}

var _ interface {
	graphql.ResponseInterceptor
	graphql.HandlerExtension
} = &Extension{}

func NewExtension(timeout time.Duration) *Extension {
	return &Extension{
		timeout: timeout,
	}
}

func (c *Extension) ExtensionName() string {
	return "ComplexityExtension"
}

func (c *Extension) Validate(_ graphql.ExecutableSchema) error {
	return nil
}

func (c *Extension) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	ctx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	started := time.Now()
	resp := next(ctx)
	timeTook := time.Since(started)
	if resp.Extensions == nil {
		resp.Extensions = make(map[string]interface{})
	}
	resp.Extensions["timeTook"] = timeTook.String()
	return resp
}
