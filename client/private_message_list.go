package client

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const inboxURL = "https://www.pathofexile.com/private-messages/inbox"

// GetInbox returns the list of private messages
func (c *Client) GetInbox(page int) ([]*PrivateMessage, error) {
	if !c.Logged {
		return nil, errors.New("You need to be logged in to change your client password")
	}

	resp, err := c.HTTP.Get(inboxURL + "/page/" + strconv.Itoa(page))
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

	inbox := []*PrivateMessage{}

	messageList := doc.Find(".pm-list").ChildrenFiltered("tbody").Children()
	messageList.Each(func(i int, s *goquery.Selection) {
		d := &PrivateMessage{}

		// Retrieve message status
		// Read or Unread
		status, ok := s.Children().First().Children().First().Attr("alt")
		if !ok {
			parsingError = errors.New("Unable to retrieve message status")
			return
		}

		d.Unread = status == "Unread"

		// Retrieve subject
		d.Subject = strings.TrimSpace(s.Children().NextFiltered(".message-details").Children().First().Text())

		// Retrieve message URL
		messageURL, ok := s.Children().NextFiltered(".message-details").Children().First().Children().First().Attr("href")
		if !ok {
			parsingError = errors.New("Unable to retrieve message url 'href' attribute")
			return
		}

		d.URL = messageURL

		// Retrieve message sender
		d.Sender = s.Children().NextFiltered(".message-details").Children().NextFiltered(".profile-link").Text()

		// Retrieve message date
		messageDate, err := time.Parse("Jan 2, 2006 15:04:05 PM", s.Children().NextFiltered(".message-details").Children().Last().Text())
		if err != nil {
			parsingError = fmt.Errorf(
				"Unable to retrieve message date: %s",
				err,
			)
			return
		}

		d.ReceivedAt = messageDate

		inbox = append(inbox, d)
	})

	return inbox, parsingError
}
