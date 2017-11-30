package client

import (
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Errorf("Unable to create http client: %v", err)
	}

	if err := c.Login(os.Getenv("RESPOE_EMAIL"), os.Getenv("RESPOE_PASSWORD")); err != nil {
		t.Errorf("Unable to login: %v", err)
	}
}
