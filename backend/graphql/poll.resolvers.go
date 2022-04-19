package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/generated"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/msgerr"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/model"
)

func (r *pollResolver) Type(ctx context.Context, obj *model.Poll) (string, error) {
	return obj.Type.String(), nil
}

func (r *pollResolver) Comments(ctx context.Context, obj *model.Poll, after *string, first *int) (*relays.Connection[*model.Comment], error) {
	relayResolver := r.NewRelayComments(ctx, obj.Kids)
	comments, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not get comments on the poll")
	}
	return comments, nil
}

func (r *pollResolver) PollOptions(ctx context.Context, obj *model.Poll, after *string, first *int) (*relays.Connection[*model.PollOption], error) {
	relayResolver := r.NewRelayPollOptions(ctx, obj.Parts)
	options, err := relayResolver.Resolve(nil, after, first, nil)
	if err != nil {
		return nil, msgerr.New(err, "Could not get poll options")
	}
	return options, nil
}

func (r *pollOptionResolver) Type(ctx context.Context, obj *model.PollOption) (string, error) {
	return obj.Type.String(), nil
}

func (r *pollOptionResolver) Poll(ctx context.Context, obj *model.PollOption) (string, error) {
	return strconv.Itoa(obj.Poll), nil
}

// Poll returns generated.PollResolver implementation.
func (r *Resolver) Poll() generated.PollResolver { return &pollResolver{r} }

// PollOption returns generated.PollOptionResolver implementation.
func (r *Resolver) PollOption() generated.PollOptionResolver { return &pollOptionResolver{r} }

type (
	pollResolver       struct{ *Resolver }
	pollOptionResolver struct{ *Resolver }
)
