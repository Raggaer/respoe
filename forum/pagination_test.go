package forum

import (
	"testing"

	"github.com/raggaer/respoe/client"
)

func TestGetPagination(t *testing.T) {
	c, err := client.New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	// Create dummy announcements forum
	f := &Forum{
		URL: "/forum/view-forum/news",
	}

	p, err := f.GetPagination(c)
	if err != nil {
		t.Fatalf("Unable to create forum pagination: %v", err)
	}

	// Atleast 55 pages should appear (since the time of making this test)
	if p.Last <= 54 {
		t.Fatalf("Wrong last page from pagination. Expected > 55 got %d", p.Last)
	}
}
