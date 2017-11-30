package client

import (
	"os"
	"testing"
)

func TestLogout(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	if err := c.Login(os.Getenv("RESPOE_EMAIL"), os.Getenv("RESPOE_PASSWORD")); err != nil {
		t.Fatalf("Unable to login: %v", err)
	}

	if err := c.Logout(); err != nil {
		t.Fatalf("Unable to logout: %v", err)
	}
}
