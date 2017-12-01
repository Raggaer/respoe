package forum

import (
	"testing"

	"github.com/raggaer/respoe/client"
)

func TestGetThreadList(t *testing.T) {
	c, err := client.New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	// Create dummy announcements forum
	f := &Forum{
		URL: "/view-forum/bug-reports",
	}

	threads, err := f.GetThreadList(1, c)
	if err != nil {
		t.Fatalf("Unable to retrieve thread list: %v", err)
	}

	t.Log(threads[0])
}
