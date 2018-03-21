package forum

import (
	"testing"

	"github.com/raggaer/respoe/client"
)

func TestGetPostList(t *testing.T) {
	c, err := client.New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	// Create dummy thread
	f := &Thread{
		URL: "/view-thread/19091",
	}

	posts, err := f.GetPostList(1, c)
	if err != nil {
		t.Fatalf("Unable to retrieve post list: %v", err)
	}

	// Check for first post author
	if posts.List[0].Author != "Cristo9FP" {
		t.Fatalf("Wrong first post author. Expected 'Cristo9FP' got %s", posts.List[0].Author)
	}

	// Check for staff post at second post
	if !posts.List[1].Staff {
		t.Fatalf("Wrong second post staff type got %t", posts.List[1].Staff)
	}

	// Check for badges at thrid post
	if len(posts.List[2].Badges) <= 0 {
		t.Fatalf("Wrong third post badges amount got %d", len(posts.List[2].Badges))
	}
}
