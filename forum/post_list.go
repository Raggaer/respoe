package forum

import (
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/raggaer/respoe/client"
)

// GetPostList returns the list of posts for the given thread
func (t *Thread) GetPostList(page int, c *client.Client) ([]*Post, error) {
	resp, err := c.HTTP.Get(forumIndexURL + t.URL + "/page/" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	postList := []*Post{}

	forumTable := doc.Find(".forumPostListTable").ChildrenFiltered("tbody")
	forumTable.Children().Each(func(i int, s *goquery.Selection) {
		p := &Post{}

		// Check if thread is a staff post
		if s.HasClass("staff") {
			p.Staff = true
		}

		// Retrieve post content
		postContent, err := s.Children().First().Children().NextFiltered("div.content").Html()
		if err != nil {
			return
		}

		p.Content = postContent

		// Retrieve post author
		p.Author = s.Children().Last().Children().First().Children().NextFiltered("div.posted-by").Children().NextFiltered("span.post_by_account").Text()

		// Retrieve post creation date
		postCreatedAt, err := time.Parse(
			"Jan 2, 2006 15:04:05 PM",
			s.Children().Last().Children().First().Children().NextFiltered("div.posted-by").Children().NextFiltered("span.post_date").Text(),
		)
		if err != nil {
			return
		}

		p.CreatedAt = postCreatedAt

		postList = append(postList, p)
	})

	return postList, nil
}
