package client

import (
	"os"
	"testing"
)

func TestChangePassword(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	// We need to be logged in to change the password
	if err := c.Login(os.Getenv("RESPOE_EMAIL"), os.Getenv("RESPOE_PASSWORD")); err != nil {
		t.Fatalf("Unable to login: %v", err)
	}

	if err := c.ChangePassword(os.Getenv("RESPOE_PASSWORD"), os.Getenv("RESPOE_NEW_PASSWORD")); err != nil {
		t.Fatalf("Unable to change password: %v", err)
	}
}
