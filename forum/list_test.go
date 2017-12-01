package forum

import (
	"testing"

	"github.com/raggaer/respoe/client"
)

func TestGetForumList(t *testing.T) {
	c, err := client.New()
	if err != nil {
		t.Fatalf("Unable to create http client: %v", err)
	}

	list, err := GetForumList(c)
	if err != nil {
		t.Fatalf("Unable to get forum list: %v", err)
	}

	if len(list) <= 0 {
		t.Fatalf("Unexpected empty forum list")
	}

	if list[0].Name != "Announcements" {
		t.Fatalf("Unexpected forum list element at index 0. Expected Announcements got %s", list[0].Name)
	}
}
