package client

import (
	"os"
	"testing"
)

func TestGetInbox(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	if err := c.Login(os.Getenv("RESPOE_EMAIL"), os.Getenv("RESPOE_PASSWORD")); err != nil {
		t.Fatalf("Unable to login: %v", err)
	}

	inbox, err := c.GetInbox(1)
	if err != nil {
		t.Fatalf("Unable to get inbox: %v", err)
	}

	t.Log(inbox[0])
}
