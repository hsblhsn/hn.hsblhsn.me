package graphql

import (
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/algolia"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/caches"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/grpc/readabilityserver"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/httpclient"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/model"
	"go.uber.org/fx"
)

// Module for fx.
// It already includes all the internal dependencies.
func Module() fx.Option {
	return fx.Options(
		model.Module(),
		caches.Module(),
		httpclient.Module(),
		hackernews.Module(),
		algolia.Module(),
		readabilityserver.Module(),
		relays.Module(),
		fx.Provide(NewResolver),
		fx.Provide(NewImageHandler),
		fx.Provide(NewGQLHandler),
	)
}
