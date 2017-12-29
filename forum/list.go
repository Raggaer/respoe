package forum

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/raggaer/respoe/client"
)

const forumIndexURL = "https://www.pathofexile.com/forum"

// GetForumList returns a list of available forums
func GetForumList(c *client.Client) ([]*Forum, error) {
	resp, err := c.HTTP.Get(forumIndexURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parsing error
	var parsingError error

	forumList := []*Forum{}

	doc.Find("div.forum_name").Each(func(i int, s *goquery.Selection) {
		// Retrieve <a></a> node
		forumName := s.ChildrenFiltered("div.name").Children()

		// Retrieve <td class="stats"></td> node
		stats := s.Parent().NextFiltered("td.stats")

		threads, err := strconv.Atoi(stats.Children().First().ChildrenFiltered("span").Text())
		if err != nil {
			parsingError = fmt.Errorf(
				"Unable to parse number of threads: %s",
				err,
			)
			return
		}

		posts, err := strconv.Atoi(stats.Children().First().Next().ChildrenFiltered("span").Text())
		if err != nil {
			parsingError = fmt.Errorf(
				"Unable to parse number of posts: %s",
				err,
			)
			return
		}

		f := &Forum{
			Name:    forumName.Text(),
			URL:     forumName.AttrOr("href", "/"),
			Threads: threads,
			Posts:   posts,
		}

		forumList = append(forumList, f)
	})

	return forumList, parsingError
}
