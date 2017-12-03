package client

import (
	"testing"
)

func TestGetDealList(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	deals, err := c.GetDealList()
	if err != nil {
		t.Fatalf("Unable to retrieve deal list: %v", err)
	}

	// We asumme there are always atleast one offer
	if len(deals) <= 0 {
		t.Fatal("There should be atleast one deal running")
	}
}
