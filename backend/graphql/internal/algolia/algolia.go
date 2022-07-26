package algolia

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/httpclient"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

type Algolia struct {
	Client *httpclient.Client
}

func NewAlgolia(client *httpclient.Client) *Algolia {
	return &Algolia{
		Client: client,
	}
}

type PaginationInput struct {
	Page    int
	PerPage int
}

func (a *Algolia) Search(ctx context.Context, typ string, query string, pagination *PaginationInput) (*Response, error) {
	page := pagination.Page
	if page < 1 || page > 100 {
		page = 1
	}
	first := pagination.PerPage
	if first < 1 || first > 10 {
		first = 10
	}
	apiEP := "https://hn.algolia.com/api/v1/search"
	values := url.Values{}
	values.Add("query", query)
	values.Add("tags", typ)
	values.Add("page", strconv.Itoa(page))
	values.Add("hitsPerPage", strconv.Itoa(first))
	ep := fmt.Sprintf("%s?%s", apiEP, values.Encode())
	resp, err := a.Client.Get(ctx, ep)
	if err != nil {
		return nil, errors.Wrap(err, "algolia: failed to send request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.Errorf("algolia: received response code: %v", resp.StatusCode)
	}

	result := (*Response)(nil)
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, errors.Wrap(err, "algolia: failed to decode response")
	}
	if result == nil {
		return nil, errors.Errorf("algolia: received empty or unknown response")
	}
	return result, nil
}
