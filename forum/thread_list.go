package forum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/raggaer/respoe/client"
	"github.com/raggaer/respoe/util"
)

// GetThreadList returns a list of threads for the given forum/page
func (f *Forum) GetThreadList(page int, c *client.Client) (*ThreadList, error) {
	resp, err := c.HTTP.Get(forumIndexURL + f.URL + "/page/" + strconv.Itoa(page))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	if util.IsMaintenance(doc) {
		return nil, errors.New("Forum is under maintenance")
	}

	// Parsing error
	var parsingError error

	// Retrieve forum name
	forumNamehrefs := doc.Find(".topBar.first .breadcrumb")
	forumName := forumNamehrefs.Text()
	forumNamehrefs.Children().Each(func(i int, s *goquery.Selection) {
		forumName = strings.Replace(forumName, s.Text(), "", 1)
	})

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
			parsingError = errors.New("Unable to retrieve thread title 'href' attribute")
			return
		}

		// Retrieve thread identifier from URL
		threadURLParts := strings.Split(threadTitleURL, "/")
		if len(threadURLParts) < 4 {
			parsingError = fmt.Errorf(
				"Unable to parse thread url parts. Expected length 4 got %d",
				len(threadURLParts),
			)
			return
		}

		threadID, err := strconv.ParseInt(threadURLParts[3], 10, 64)
		if err != nil {
			parsingError = fmt.Errorf(
				"Unable to parse thread identifier: %s",
				err,
			)
			return
		}

		t.ID = threadID
		t.URL = threadTitleURL
		t.Title = strings.TrimSpace(threadTitle.Text())

		// Retrieve pagination information
		pagination := thread.Children().NextFiltered("div.forum_pagination")

		if pagination.Length() > 0 {
			firstPage, err := strconv.Atoi(pagination.Children().First().Text())
			if err != nil {
				parsingError = fmt.Errorf(
					"Unable to parse thread first page: %s",
					err,
				)
				return
			}
			lastPage, err := strconv.Atoi(pagination.Children().Last().Text())
			if err != nil {
				parsingError = fmt.Errorf(
					"Unable to parse thread last page: %s",
					err,
				)
				return
			}

			t.Pagination = &util.Pagination{
				First:   firstPage,
				Current: firstPage,
				Last:    lastPage,
			}
		}

		// Retrieve thread author
		postBy := thread.Children().NextFiltered("div.postBy")
		t.Author = strings.TrimSpace(postBy.Children().First().Children().First().Text())

		// Retrieve thread creation date
		// Thread creation date starts with ', Date' so we need to remove ', '
		threadDate := postBy.Children().Last().Text()
		threadDate = strings.TrimSpace(strings.TrimPrefix(threadDate, ", "))

		creationDate, err := time.Parse(util.DateFormat, threadDate)
		if err != nil {
			parsingError = fmt.Errorf(
				"Unable to parse thread creation date: %s",
				err,
			)
			return
		}

		t.CreatedAt = creationDate

		// Retrieve views block
		viewBlock := s.Children().NextFiltered(".views")
		replies, err := strconv.Atoi(viewBlock.Children().First().ChildrenFiltered("span").Text())
		if err != nil {
			parsingError = fmt.Errorf(
				"Unable to parse thread replies number: %s",
				err,
			)
			return
		}
		views, err := strconv.Atoi(viewBlock.Children().Last().ChildrenFiltered("span").Text())
		if err != nil {
			parsingError = fmt.Errorf(
				"Unable to parse thread views number: %s",
				err,
			)
			return
		}

		t.Replies = replies
		t.Views = views

		threadList = append(threadList, t)
	})

	// Retrieve thread page pagination
	threadPagination, err := util.GetPaginationFromDoc(doc)
	if err != nil {
		return nil, err
	}

	return &ThreadList{
		ForumName:  forumName,
		List:       threadList,
		Pagination: threadPagination,
	}, parsingError
}
