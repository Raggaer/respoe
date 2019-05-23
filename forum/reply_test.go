package forum

import (
	"os"
	"testing"

	"github.com/raggaer/respoe/client"
)

func TestReply(t *testing.T) {
	t.SkipNow()
	c, err := client.New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	if err := c.Login(os.Getenv("RESPOE_EMAIL"), os.Getenv("RESPOE_PASSWORD")); err != nil {
		t.Fatalf("Unable to login: %v", err)
	}

	// Create dummy thread
	// Thread 2058135 is on the Off-Topic board
	tt := Thread{
		ID: 2058135,
	}

	// Reply to thread
	if err := tt.Reply("I will have more replies soon! :-)", c); err != nil {
		t.Fatalf("Unable to reply to thread: %s", err)
	}
}
