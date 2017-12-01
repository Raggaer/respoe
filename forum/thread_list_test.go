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

	// We check for the first thread on the forum
	// If the sticky thread changes the test will fail
	if threads[0].Title != "Spelling mistakes and typos" || !threads[0].Staff || !threads[0].Sticky {
		t.Fatalf(
			`Unexpected first thread information. 
			Expected title 'Spelling mistakes and typos' got %s.
			Expected staff thread type got %t.
			Expected sticky thread type got %t`,
			threads[0].Title,
			threads[0].Staff,
			threads[0].Sticky,
		)
	}

	// We check for atleast 97 pages for the number of pages
	// Since the time this test was made
	if threads[0].Pagination.Last < 96 {
		t.Fatalf("Wrong last page for pagination. Expected > 96 got %d", threads[0].Pagination.Last)
	}
}
