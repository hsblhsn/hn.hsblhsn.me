package complexity

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/msgerr"
)

type Extension struct {
	entries Map
	limit   int
}

var _ interface {
	graphql.ResponseInterceptor
	graphql.FieldInterceptor
	graphql.HandlerExtension
} = &Extension{}

func NewExtension(limit int, m Map) *Extension {
	return &Extension{
		limit:   limit,
		entries: m,
	}
}

func (c *Extension) ExtensionName() string {
	return "ComplexityExtension"
}

func (c *Extension) Validate(_ graphql.ExecutableSchema) error {
	return nil
}

func (c *Extension) InterceptResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	ctx = NewCounterCtx(ctx, c.limit)
	resp := next(ctx)
	counter, err := CounterFromCtx(ctx)
	if err != nil {
		panic(err)
	}
	counter.MakeReady()
	resp.Extensions = map[string]interface{}{
		"complexity":      counter.Value(),
		"complexityLimit": counter.Limit(),
	}
	return resp
}

func (c *Extension) InterceptField(ctx context.Context, next graphql.Resolver) (interface{}, error) {
	counter, err := CounterFromCtx(ctx)
	if err != nil {
		return nil, msgerr.New(err, MsgComplexityCounterError)
	}

	// ignore introspection fields.
	fieldCtx := graphql.GetFieldContext(ctx)
	switch fieldCtx.Object {
	case "__Field", "__Type", "__Directive", "__InputValue", "__Schema":
		return next(ctx)
	}
	if err := counter.Add(c.entries.Get(fieldCtx.Object, 0)); err != nil {
		return nil, msgerr.New(err, MsgComplexityLimitExceeded)
	}
	fieldName := fmt.Sprintf("%s.%s", fieldCtx.Object, fieldCtx.Field.Name)
	if err := counter.Add(c.entries.Get(fieldName, 0)); err != nil {
		return nil, msgerr.New(err, MsgComplexityLimitExceeded)
	}
	return next(ctx)
}
