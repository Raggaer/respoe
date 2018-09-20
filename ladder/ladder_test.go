package ladder

import (
	"testing"

	"github.com/raggaer/respoe/client"
)

func TestRetrieveLadder(t *testing.T) {
	c, err := client.New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	// Retrieve ladder for Hardcore
	ranking, err := RetrieveLadder("Hardcore", 0, 5, c)
	if err != nil {
		t.Fatalf("Unable to retrieve ladder information: %s", err)
	}

	// Check for errors
	if ranking.Error.Message != "" {
		t.Fatalf("Unable to retrieve ladder information: %s", ranking.Error.Message)
	}

	// Check if theres data inside ranking
	if len(ranking.Entries) <= 0 {
		t.Fatalf("Unable to retrieve ladder information: %s", "Theres no ranking data gathered")
	}
}
