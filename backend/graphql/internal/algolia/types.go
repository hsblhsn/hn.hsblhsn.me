package algolia

import (
	"strconv"
	"time"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/model"
	"github.com/pkg/errors"
)

// Response from algolia api.
// nolint:tagliatelle,govet
type Response struct {
	Hits             []*Hit `json:"hits"`
	Query            string `json:"query"`
	Params           string `json:"params"`
	NbHits           int    `json:"nbHits"`
	Page             int    `json:"page"`
	NbPages          int    `json:"nbPages"`
	HitsPerPage      int    `json:"hitsPerPage"`
	ProcessingTimeMS int    `json:"processingTimeMS"`
	ExhaustiveNbHits bool   `json:"exhaustiveNbHits"`
	ExhaustiveTypo   bool   `json:"exhaustiveTypo"`
}

func (r *Response) ToConnection() (*model.NodeConnection, error) {
	pageInfo := &relays.PageInfo{
		HasNextPage:     r.Page < r.NbPages,
		HasPreviousPage: r.Page > 1,
		StartCursor:     strconv.Itoa(r.Page),
		EndCursor:       strconv.Itoa(r.Page),
		PageCursor:      strconv.Itoa(r.Page),
	}
	edges := make([]*model.NodeEdge, len(r.Hits))
	for i := range r.Hits {
		node, err := r.Hits[i].ToNode()
		if err != nil {
			return nil, errors.Wrap(err, "algolia: failed to convert hit to node")
		}
		edges[i] = &model.NodeEdge{
			Cursor: strconv.Itoa(r.Page),
			Node:   node,
		}
	}
	return &model.NodeConnection{
		PageInfo: pageInfo,
		Edges:    edges,
	}, nil
}

// Hit represents a single hit from Algolia.
// nolint: tagliatelle
type Hit struct {
	CreatedAt      time.Time `json:"created_at"`
	ObjectID       string    `json:"objectID"`
	Title          string    `json:"title"`
	URL            string    `json:"url"`
	Author         string    `json:"author"`
	StoryText      *string   `json:"story_text"`
	CommentText    *string   `json:"comment_text"`
	StoryTitle     *string   `json:"story_title"`
	Tags           []string  `json:"_tags"`
	Points         int       `json:"points"`
	NumComments    int       `json:"num_comments"`
	CreatedAtI     int       `json:"created_at_i"`
	RelevancyScore int       `json:"relevancy_score,omitempty"`
}

func (h *Hit) ToNode() (model.Node, error) {
	if len(h.Tags) == 0 {
		return nil, errors.New("algolia: hit has no tags")
	}
	if h.Tags[0] == "story" {
		return h.ToStory()
	}
	return nil, nil
}

func (h *Hit) ToStory() (*model.Story, error) {
	databaseID, err := strconv.Atoi(h.ObjectID)
	if err != nil {
		return nil, errors.Wrap(err, "algolia: failed to parse database id from object id")
	}

	storyText := ""
	if h.StoryText != nil {
		storyText = *h.StoryText
	}

	return &model.Story{
		ItemResponse: &hackernews.ItemResponse{
			Deleted:     false,
			Dead:        false,
			DatabaseID:  databaseID,
			Time:        h.CreatedAtI,
			Descendants: h.NumComments,
			Score:       h.Points,
			Type:        "story",
			By:          h.Author,
			URL:         h.URL,
			Text:        storyText,
			Title:       h.Title,
			Kids:        nil,
		},
	}, nil
}
