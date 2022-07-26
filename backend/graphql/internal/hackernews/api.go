package hackernews

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/httpclient"
	"strings"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/readerviews"
	"github.com/pkg/errors"
)

type HackerNews struct {
	client *httpclient.Client
}

func NewHackerNews(client *httpclient.Client) *HackerNews {
	return &HackerNews{client: client}
}

func (h *HackerNews) API(ctx context.Context, endpoint string, result any) error {
	// Send the request
	resp, err := h.client.Get(ctx, endpoint)
	if err != nil {
		return errors.Wrap(err, "hackernews: could not send request")
	}
	defer resp.Body.Close()
	// Decode the response body to the target struct
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return errors.Wrap(err, "hackernews: could not decode response")
	}
	return nil
}

func (h *HackerNews) GetTopStories(ctx context.Context) ([]int, error) {
	var list []int
	endpoint := fmt.Sprintf("%s/v0/topstories.json", host)
	if err := h.API(ctx, endpoint, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (h *HackerNews) GetNewStories(ctx context.Context) ([]int, error) {
	var list []int
	endpoint := fmt.Sprintf("%s/v0/newstories.json", host)
	if err := h.API(ctx, endpoint, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (h *HackerNews) GetAskStories(ctx context.Context) ([]int, error) {
	var list []int
	endpoint := fmt.Sprintf("%s/v0/askstories.json", host)
	if err := h.API(ctx, endpoint, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (h *HackerNews) GetShowStories(ctx context.Context) ([]int, error) {
	var list []int
	endpoint := fmt.Sprintf("%s/v0/showstories.json", host)
	if err := h.API(ctx, endpoint, &list); err != nil {
		return nil, err
	}
	return list, nil
}

func (h *HackerNews) GetJobStories(ctx context.Context) ([]int, error) {
	var list []int
	endpoint := fmt.Sprintf("%s/v0/jobstories.json", host)
	if err := h.API(ctx, endpoint, &list); err != nil {
		return nil, err
	}
	return list, nil
}

// GetTypedItem http request to the endpoint and decode json data to the target struct,.
func (h *HackerNews) GetTypedItem(ctx context.Context, typ ItemType, itemID string) (*ItemResponse, error) {
	idN, idErr := GetIntID(itemID)
	if idErr != nil {
		return nil, idErr
	}
	out, err := h.GetItem(ctx, idN)
	if err != nil {
		return nil, err
	}
	if out.Type != typ {
		msg := fmt.Sprintf(
			"type: %q!=%q on id=%q",
			out.Type, typ, itemID,
		)
		return nil, errors.Wrap(ErrMismatchedType, msg)
	}
	return out, nil
}

func (h *HackerNews) GetItem(ctx context.Context, id int) (*ItemResponse, error) {
	var (
		endpoint = fmt.Sprintf("%s/v0/item/%d.json", host, id)
		out      = (*ItemResponse)(nil)
		err      = h.API(ctx, endpoint, &out)
	)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrItemNotFound
	}
	out.Text, err = readerviews.TransformHTML(out.URL, strings.NewReader(out.Text))
	if err != nil {
		return nil, errors.Wrap(err, "hackernews: could not transform item text")
	}
	return out, nil
}

func (h *HackerNews) GetUser(ctx context.Context, id string) (*UserResponse, error) {
	var (
		endpoint = fmt.Sprintf("%s/v0/user/%s.json", host, id)
		out      = (*UserResponse)(nil)
		err      = h.API(ctx, endpoint, &out)
	)
	if err != nil {
		return nil, err
	}
	if out == nil {
		return nil, ErrItemNotFound
	}
	return out, nil
}
