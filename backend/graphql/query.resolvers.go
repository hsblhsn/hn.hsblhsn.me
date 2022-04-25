package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/algolia"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/model"
)

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	var (
		nodeID    = id
		errOnUser error
	)
	if hackernews.IsUserName(nodeID) {
		user, err := r.hackerNews.GetUser(ctx, nodeID)
		if err != nil {
			errOnUser = err
			goto itemResolver
		}
		return &model.User{UserResponse: user}, nil
	}
itemResolver:
	idN, err := hackernews.GetIntID(nodeID)
	if err != nil {
		if errOnUser != nil {
			return nil, msgerr.New(errOnUser, "Could not find a user with the username")
		}
		return nil, msgerr.New(err, "Invalid ID")
	}
	result, err := r.hackerNews.GetItem(ctx, idN)
	if err != nil {
		return nil, msgerr.New(err, "Could not get item")
	}
	return ItemToNode(result)
}

func (r *queryResolver) Search(ctx context.Context, query string, after *string, first *int) (*relays.Connection[model.Node], error) {
	var (
		page    = 1
		perPage = 10
	)
	if after != nil {
		afterN, err := strconv.Atoi(*after)
		if err != nil {
			return nil, msgerr.New(err, "Invalid after")
		}
		page = afterN + 1
	}
	if first != nil {
		perPage = *first
	}
	result, err := r.algolia.Search(ctx, "story", query, &algolia.PaginationInput{
		Page:    page,
		PerPage: perPage,
	})
	if err != nil {
		return nil, msgerr.New(err, "Could not search")
	}
	conn, err := result.ToConnection()
	if err != nil {
		return nil, msgerr.New(err, "Could not build node connection")
	}
	return conn, nil
}

func (r *queryResolver) TopStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Story], error) {
	list, err := r.hackerNews.GetTopStories(ctx)
	if err != nil {
		return nil, msgerr.New(err, "Could not get top stories.")
	}
	relayResolver := r.NewRelayStories(ctx, list)
	stories, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate top stories.")
	}
	return stories, nil
}

func (r *queryResolver) NewStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Story], error) {
	list, err := r.hackerNews.GetNewStories(ctx)
	if err != nil {
		return nil, msgerr.New(err, "Could not get new stories.")
	}
	relayResolver := r.NewRelayStories(ctx, list)
	stories, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate new stories.")
	}
	return stories, nil
}

func (r *queryResolver) AskStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Story], error) {
	list, err := r.hackerNews.GetAskStories(ctx)
	if err != nil {
		return nil, msgerr.New(err, "Could not get ask HN.")
	}
	relayResolver := r.NewRelayStories(ctx, list)
	stories, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate ask HN.")
	}
	return stories, nil
}

func (r *queryResolver) ShowStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Story], error) {
	list, err := r.hackerNews.GetShowStories(ctx)
	if err != nil {
		return nil, msgerr.New(err, "Could not get show HN.")
	}
	relayResolver := r.NewRelayStories(ctx, list)
	stories, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate show HN.")
	}
	return stories, nil
}

func (r *queryResolver) JobStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Job], error) {
	list, err := r.hackerNews.GetJobStories(ctx)
	if err != nil {
		return nil, msgerr.New(err, "Could not get jobs.")
	}
	relayResolver := r.NewRelayJobs(ctx, list)
	jobs, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not paginate jobs.")
	}
	return jobs, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
