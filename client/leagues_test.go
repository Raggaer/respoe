package client

import (
	"testing"
)

func TestLeagueList(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	leagues, err := c.LeagueList()
	if err != nil {
		t.Fatalf("Unable to retrieve league list: %v", err)
	}

	// Find Hardcore league since it will always exist
	f := false
	for _, l := range leagues {
		if l.Name == "Hardcore" {
			f = true
		}
	}

	if !f {
		t.Fatal("Unable to find 'Hardcore' league inside league list")
	}
}
