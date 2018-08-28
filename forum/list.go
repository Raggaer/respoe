package forum

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/raggaer/respoe/client"
	"github.com/raggaer/respoe/util"
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

		// Retrieve forum description
		forumDescription := strings.TrimPrefix(strings.TrimSpace(s.Parent().Text()), strings.TrimSpace(forumName.Text()))

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

		// Retrieve last post <td class="last_post"></td> node
		lastPost := s.Parent().Parent().Children().NextFiltered("td.last_post")

		postAuthor := lastPost.Children().First().Children().First().Text()

		postURL, ok := lastPost.Children().NextFiltered("div.post_date").Children().First().Attr("href")
		if !ok {
			return
		}

		postDate, err := time.Parse(
			util.DateFormat,
			lastPost.Children().NextFiltered("div.post_date").Text(),
		)
		if err != nil {
			parsingError = fmt.Errorf(
				"Unable to parse last post date: %s",
				err,
			)
			return
		}

		postAchievement := lastPost.Children().First().Children().First()

		lastPostAchiv, ok := postAchievement.Children().First().Children().First().Attr("src")
		if !ok {
			return
		}

		f := &Forum{
			Name:        forumName.Text(),
			URL:         forumName.AttrOr("href", "/"),
			Threads:     threads,
			Description: forumDescription,
			Posts:       posts,
			LastPost: LastPost{
				Author:    postAuthor,
				CreatedAt: postDate,
				URL:       postURL,
				Achievement: PostAchievement{
					URL: lastPostAchiv,
				},
			},
		}

		forumList = append(forumList, f)
	})

	return forumList, parsingError
}
