package ladder

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/raggaer/respoe/client"
)

const ladderURL = "https://www.pathofexile.com/api/ladders?offset=%d&limit=%d&id=%s&type=league"

// Ranking defines a ladder ranking response
type Ranking struct {
	Error   RankingError
	Total   int
	Entries []Entry
}

// RankingError defines a ladder error message
type RankingError struct {
	Message string
	Code    int
}

// Entry defines a ladder entry
type Entry struct {
	Rank      int
	Dead      bool
	Online    bool
	Character EntryCharacter
	Account   EntryAccount
}

// EntryAccount defines a character account
type EntryAccount struct {
	Name       string
	Challenges EntryAccountChallenges
	Twitch     EntryAccountTwitch
}

// EntryAccountTwitch defines a twitch account
type EntryAccountTwitch struct {
	Name string
}

// EntryAccountChallenges defines the total number of challenges
type EntryAccountChallenges struct {
	Total int
}

// EntryCharacter defines a ladder character
type EntryCharacter struct {
	Name       string
	Level      int
	Experience int64
	Class      string
	ID         string `json:"id"`
}

// RetrieveLadder retrieves a list of ladder entries for the given league
func RetrieveLadder(league string, offset, limit int, c *client.Client) (*Ranking, error) {
	resp, err := c.HTTP.Get(fmt.Sprintf(
		ladderURL,
		offset,
		limit,
		url.QueryEscape(league),
	))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var r Ranking
	if err := json.Unmarshal(respBody, &r); err != nil {
		return nil, err
	}

	// Check for error message
	if r.Error.Message != "" {
		return nil, errors.New(r.Error.Message)
	}

	return &r, nil
}
