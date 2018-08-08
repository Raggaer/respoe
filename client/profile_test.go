package client

import (
	"testing"
)

func TestGetAccountProfile(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	profile, err := c.GetAccountProfile("Punckanuto")
	if err != nil {
		t.Fatalf("Unable to retrieve account profile: %v", err)
	}

	if len(profile.Characters) <= 0 {
		t.Fatalf("Unexpected profile character list length. Got %d", len(profile.Characters))
	}

	char := profile.Characters[0]
	if char.Name != "FacebreakerBreach" {
		t.Fatalf("Unexpected first profile character. Expected %s Got %s", "FacebreakerBreach", char.Name)
	}

	if len(profile.Characters[1].Items) <= 0 {
		t.Fatalf("Unexpected second profile character %s. Should atleast have 1 item", profile.Characters[1].Name)
	}
}
