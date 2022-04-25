package graphql

import (
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/algolia"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/hackernews"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	hackerNews *hackernews.HackerNews
	algolia    *algolia.Algolia
}

// NewResolver returns a new Resolver instance.
// It takes a HackerNews instance as a dependency.
// This allows us to inject a mock HackerNews instance for testing.
func NewResolver(hackerNews *hackernews.HackerNews, algolia *algolia.Algolia) *Resolver {
	return &Resolver{
		hackerNews: hackerNews,
		algolia:    algolia,
	}
}
