package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph/model"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/msgerr"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/relays"
)

func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	idN, err := hackernews.GetIntID(id)
	if err != nil {
		return nil, msgerr.New(err, "Invalid ID")
	}
	result, err := r.hackernews.GetItem(ctx, idN)
	if err != nil {
		return nil, msgerr.New(err, "Could not get item")
	}
	switch result.Type {
	case hackernews.ItemTypeStory:
		return &model.Story{ItemResponse: result}, nil
	case hackernews.ItemTypeComment:
		return &model.Comment{ItemResponse: result}, nil
	case hackernews.ItemTypeJob:
		return &model.Job{ItemResponse: result}, nil
	case hackernews.ItemTypePoll:
		return &model.Poll{ItemResponse: result}, nil
	case hackernews.ItemTypePollOption:
		return &model.PollOption{ItemResponse: result}, nil
	default:
		msg := fmt.Sprintf("unknown item type: %q", result.Type)
		// nolint:goerr113 // we want to capture the result type here.
		return nil, msgerr.New(errors.New(msg), "Invalid item type received")
	}
}

func (r *queryResolver) Story(ctx context.Context, id string) (*model.Story, error) {
	result, err := r.hackernews.GetTypedItem(ctx, hackernews.ItemTypeStory, id)
	if err != nil {
		return nil, msgerr.New(err, "Could not get story")
	}
	return &model.Story{ItemResponse: result}, nil
}

func (r *queryResolver) Job(ctx context.Context, id string) (*model.Job, error) {
	result, err := r.hackernews.GetTypedItem(ctx, hackernews.ItemTypeJob, id)
	if err != nil {
		return nil, msgerr.New(err, "Could not get job")
	}
	return &model.Job{ItemResponse: result}, nil
}

func (r *queryResolver) Poll(ctx context.Context, id string) (*model.Poll, error) {
	result, err := r.hackernews.GetTypedItem(ctx, hackernews.ItemTypePoll, id)
	if err != nil {
		return nil, msgerr.New(err, "Could not get poll")
	}
	return &model.Poll{ItemResponse: result}, nil
}

func (r *queryResolver) PollOption(ctx context.Context, id string) (*model.PollOption, error) {
	result, err := r.hackernews.GetTypedItem(ctx, hackernews.ItemTypePollOption, id)
	if err != nil {
		return nil, msgerr.New(err, "Could not get poll option")
	}
	return &model.PollOption{ItemResponse: result}, nil
}

func (r *queryResolver) TopStories(ctx context.Context, after *string, first *int) (*relays.Connection[*model.Story], error) {
	list, err := r.hackernews.GetTopStories(ctx)
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
	list, err := r.hackernews.GetNewStories(ctx)
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
	list, err := r.hackernews.GetAskStories(ctx)
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
	list, err := r.hackernews.GetShowStories(ctx)
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
	list, err := r.hackernews.GetJobStories(ctx)
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