package hackernews

import (
	"strconv"

	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql/internal/relays"
	"github.com/pkg/errors"
)

const host = "https://hacker-news.firebaseio.com"

var ErrMismatchedType = errors.New("hackernews: mismatched item type")

type ItemType string

func (t ItemType) String() string {
	return string(t)
}

const (
	ItemTypeJob        = "job"
	ItemTypeStory      = "story"
	ItemTypeComment    = "comment"
	ItemTypePoll       = "poll"
	ItemTypePollOption = "pollopt"
)

// nolint:govet // i have no idea why this is happening.
type ItemResponse struct {
	Deleted     bool     `json:"deleted"`
	Dead        bool     `json:"dead"`
	DatabaseID  int      `json:"id"`
	Time        int      `json:"time"`
	Descendants int      `json:"descendants"`
	Score       int      `json:"score"`
	Poll        int      `json:"poll"`
	Parent      int      `json:"parent"`
	Type        ItemType `json:"type"`
	By          string   `json:"by"`
	URL         string   `json:"url"`
	Text        string   `json:"text"`
	Title       string   `json:"title"`
	Kids        []int    `json:"kids"`
	Parts       []int    `json:"parts"`
}

func (i *ItemResponse) ID() string {
	return NewID(i.DatabaseID)
}

func (ItemResponse) IsNode() {}

func NewID(id int) string {
	return relays.NewID(id)
}

func GetIntID(id string) (int, error) {
	idN, err := strconv.Atoi(id)
	if err != nil {
		return -1, errors.Wrap(err, "hackernews: invalid ID")
	}
	return idN, nil
}
