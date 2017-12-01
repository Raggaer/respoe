package forum

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/raggaer/respoe/client"
)

// GetThreadList returns a list of threads for the given forum/page
func (f *Forum) GetThreadList(page int, c *client.Client) ([]*Thread, error) {
	resp, err := c.HTTP.Get(forumIndexURL + f.URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	threadList := []*Thread{}

	forumTable := doc.Find("#view_forum_table").ChildrenFiltered("tbody")
	forumTable.Children().Each(func(i int, s *goquery.Selection) {
		t := &Thread{}

		// Retrieve thread flags
		flags := s.Children().First()

		if flags.Length() > 0 {
			flags.Children().Each(func(i int, f *goquery.Selection) {
				if f.HasClass("sticky") {
					t.Sticky = true
				}
				if f.HasClass("staff") {
					t.Staff = true
				}
				if f.HasClass("locked") {
					t.Locked = true
				}
				if f.HasClass("support") {
					t.Support = true
				}
			})
		}

		// Retrieve thread block
		thread := s.Children().NextFiltered(".thread")
		threadTitle := thread.Children().First().Children().NextFiltered("div.title").Children().First()
		threadTitleURL, ok := threadTitle.Attr("href")

		if !ok {
			return
		}

		t.URL = threadTitleURL
		t.Title = strings.TrimSpace(threadTitle.Text())

		// Retrieve pagination information
		pagination := thread.Children().NextFiltered("div.forum_pagination")
		firstPage, err := strconv.Atoi(pagination.Children().First().Text())
		if err != nil {
			return
		}
		lastPage, err := strconv.Atoi(pagination.Children().Last().Text())
		if err != nil {
			return
		}

		t.Pagination = &Pagination{
			First:   firstPage,
			Current: firstPage,
			Last:    lastPage,
		}

		// Retrieve views block
		viewBlock := s.Children().NextFiltered(".views")
		replies, err := strconv.Atoi(viewBlock.Children().First().ChildrenFiltered("span").Text())
		if err != nil {
			return
		}
		views, err := strconv.Atoi(viewBlock.Children().Last().ChildrenFiltered("span").Text())
		if err != nil {
			return
		}

		t.Replies = replies
		t.Views = views

		threadList = append(threadList, t)
	})

	return threadList, nil
}
