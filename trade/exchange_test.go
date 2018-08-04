package trade

import (
	"testing"

	"github.com/raggaer/respoe/client"
)

func TestRetrieveExchange(t *testing.T) {
	c, err := client.New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	// Retrieve exchange
	offers, err := RetrieveExchange("Standard", []string{"chaos"}, []string{"alch"}, c)
	if err != nil {
		t.Fatalf("Unable to retrieve currency exchange: %s", err)
	}

	// Do some offer validation
	if len(offers) <= 0 {
		t.Fatal("Unexpected number of exchange offers. There should atleast be one (in Standard)")
	}

	if offers[0].Listing.Price.Have.Currency != "chaos" {
		t.Fatalf("Unexpected first offer currency item name. Expected 'chaos' but got %s", offers[0].Listing.Price.Have.Currency)
	}

	if offers[0].Listing.Price.Want.Currency != "alch" {
		t.Fatalf("Unexpected first offer currency item name. Expected 'alch' but got %s", offers[0].Listing.Price.Want.Currency)
	}
}
