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

// GetPostList returns the list of posts for the given thread
func (t *Thread) GetPostList(page int, c *client.Client) (*PostList, error) {
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

	// Parsing error
	var parsingError error

	// Post item information place holder
	itemList := []*Item{}

	// Retrieve item information JSON data
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		// Check if its the script tag we are looking for
		if !strings.Contains(
			s.Text(), "require([\"PoE/Item/DeferredItemRenderer\"]",
		) {
			return
		}

		// Grab JSON data and remove JavaScript calls
		v := s.Text()
		v = strings.TrimSpace(v[135 : len(v)-40])

		count := 1
		nextExit := false

		// Parse every item
		for {
			s := strings.Split(v, ",[]],["+strconv.Itoa(count))
			if len(s) <= 1 {
				nextExit = true
			}

			// Parse current item
			item, err := ParseItem([]byte(s[0]))
			if err != nil {
				parsingError = fmt.Errorf(
					"Unable to parse item %d: %s",
					count,
					err,
				)
				return
			}

			// Add item to the post list
			itemList = append(itemList, item)

			// Remove item from the feed
			v = v[len(s[0]):len(v)]
			v = strings.TrimPrefix(v, ",[]],["+strconv.Itoa(count)+",")

			count++

			if nextExit {
				break
			}
		}
	})

	// Retrieve forum url
	forumURL, ok := doc.Find(".topBar.first .breadcrumb").Children().Next().Next().Attr("href")
	if !ok {
		forumURL = ""
	}

	// Retrieve forum name
	forumName := strings.TrimSuffix(strings.TrimPrefix(doc.Find(".topBar.first .breadcrumb").Text(), "Forum Index»"), "»View Thread")

	// Retrieve thread name
	threadName := doc.Find(".topBar.last.layoutBoxTitle").Text()

	forumTable := doc.Find(".forumPostListTable").ChildrenFiltered("tbody")
	forumTable.Children().Each(func(i int, s *goquery.Selection) {
		p := &Post{}

		// Check if post is made by staff
		if s.HasClass("staff") {
			p.Staff = true
		}

		// Check if post is made by a valuable poster
		if s.HasClass("valued-poster") {
			p.ValuedPoster = true
		}

		// Retrieve post content
		postContent, err := s.Children().First().Children().NextFiltered("div.content").Html()
		if err != nil {
			parsingError = fmt.Errorf(
				"Unable to retrieve post content HTML: %s",
				err,
			)
			return
		}

		p.Content = postContent

		// Retrieve post avatar
		avatarURL, avatarFound := s.Children().Last().Children().First().Children().NextFiltered("div.avatar").Children().First().Attr("src")
		if avatarFound {
			p.Avatar = avatarURL
		}

		// Retrieve post author
		p.Author = s.Children().Last().Children().First().Children().NextFiltered("div.posted-by").Children().NextFiltered("span.post_by_account").Text()

		// Retrieve post achievements
		achievementsDiv := s.Children().Last().Children().First().Children().NextFiltered("div.posted-by").Children().NextFiltered("span.post_by_account").First().Children().First()
		achievementsURL, urlFound := achievementsDiv.Attr("src")
		achievementsAlt, altFound := achievementsDiv.Attr("alt")
		if urlFound && altFound {
			p.Achievement = PostAchievement{
				URL: achievementsURL,
				Alt: achievementsAlt,
			}
		}

		// Retrieve post creation date
		postCreatedAt, err := time.Parse(
			"Jan 2, 2006 15:04:05 PM",
			s.Children().Last().Children().First().Children().NextFiltered("div.posted-by").Children().NextFiltered("span.post_date").Text(),
		)
		if err != nil {
			parsingError = fmt.Errorf(
				"Unable to parse post creation date: %s",
				err,
			)
			return
		}

		p.CreatedAt = postCreatedAt

		// Retrieve post badges
		s.Children().Last().Children().First().Children().NextFiltered("div.posted-by").Children().NextFiltered("div.badges").Children().Each(
			func(i int, b *goquery.Selection) {
				badge := b.Children().First()
				alt, altFound := badge.Attr("alt")
				url, urlFound := badge.Attr("src")
				if altFound && urlFound {
					p.Badges = append(p.Badges, PostBadge{
						Name: alt,
						URL:  url,
					})
				}
			},
		)

		postList = append(postList, p)
	})

	// Retrieve thread pagination
	threadPagination, err := util.GetPaginationFromDoc(doc)
	if err != nil {
		return nil, err
	}

	return &PostList{
		Items:      itemList,
		ForumName:  forumName,
		Title:      threadName,
		List:       postList,
		ForumURL:   forumURL,
		Pagination: threadPagination,
	}, parsingError
}
