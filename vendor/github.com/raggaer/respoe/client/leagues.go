package client

import (
	"encoding/json"
	"io/ioutil"
)

const leagueListURL = "http://www.pathofexile.com/api/trade/data/leagues"

// LeagueListResult the league list call result
type LeagueListResult struct {
	Result []*League
}

// League defines a path of exile active league
type League struct {
	Name string `json:"id"`
	Text string
}

func (c *Client) LeagueList() ([]*League, error) {
	resp, err := c.HTTP.Get(leagueListURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	list := &LeagueListResult{}
	if err := json.Unmarshal(respBody, &list); err != nil {
		return nil, err
	}
	return list.Result, nil
}
