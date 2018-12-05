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

	// Check forum name
	if posts.ForumName != "Bug Reports" {
		t.Fatalf("Wrong forum name. Expected 'Bug Reports' got %s", posts.ForumName)
	}

	// Test admin post
	adminPost := Thread{
		URL: "/view-thread/2255460",
	}
	adminThreadPosts, err := adminPost.GetPostList(1, c)
	if err != nil {
		t.Fatalf("Unable to retrieve post list: %v", err)
	}
	if len(adminThreadPosts.List[0].Content) < 50000 {
		t.Fatalf("Invalid admin thread post content. Expected a length of '%d' got '%d'", 50000, len(adminThreadPosts.List[0].Content))
	}
}

func TestGetPostListWithItems(t *testing.T) {
	c, err := client.New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	// Create dummy thread
	f := &Thread{
		URL: "/view-thread/545343",
	}

	posts, err := f.GetPostList(1, c)
	if err != nil {
		t.Fatalf("Unable to retrieve post list: %v", err)
	}

	// Check for first post author
	if posts.List[0].Author != "mattc3303" {
		t.Fatalf("Wrong first post author. Expected 'mattc3303' got %s", posts.List[0].Author)
	}

	// Check for first post item
	if posts.Items[0].Name != "Mon'tregul's Grasp" {
		t.Fatalf("Wrong first post item name. Expected 'Mon'tregul's Grasp' got %s", posts.Items[0].Name)
	}

	// Check for second post item
	if posts.Items[1].Name != "Doedre's Damning" {
		t.Fatalf("Wrong second post item name. Expected 'Doedre's Damning' got %s", posts.Items[1].Name)
	}
}
