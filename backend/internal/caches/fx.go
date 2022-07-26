package caches

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Provide(fx.Annotate(NewInMemoryCache, fx.As(new(Cache))))
}
