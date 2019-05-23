package client

import (
	"os"
	"testing"
)

func TestSendPrivateMessage(t *testing.T) {
	t.SkipNow()
	c, err := New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	if err := c.Login(os.Getenv("RESPOE_EMAIL"), os.Getenv("RESPOE_PASSWORD")); err != nil {
		t.Fatalf("Unable to login: %v", err)
	}

	if err := c.SendPrivateMessage([]string{
		os.Getenv("RESPOE_RECIPIENT"),
	}, "Hello from Respoe", "Hi brother"); err != nil {
		t.Fatalf("Unable to send private message: %v", err)
	}
}
