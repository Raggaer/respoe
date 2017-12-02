package client

import (
	"errors"
	"net/url"

	"github.com/raggaer/respoe/util"
)

const sendPrivateMessageURL = "https://www.pathofexile.com/private-messages/compose"

// SendPrivateMessage sends a private message to the given recipients
func (c *Client) SendPrivateMessage(recipients []string, subject, content string) error {
	if !c.Logged {
		return errors.New("You need to be logged in to send a private message")
	}

	hashValue, err := util.GetFormHash(sendPrivateMessageURL, c.HTTP)
	if err != nil {
		return err
	}

	// Convert []string to valid textarea newline separated names
	recipientList := ""
	for _, r := range recipients {
		recipientList += r + "\r\n"
	}

	resp, err := c.HTTP.PostForm(sendPrivateMessageURL, url.Values{
		"add_recipient": {
			recipientList,
		},
		"subject": {
			subject,
		},
		"content": {
			content,
		},
		"hash": {
			hashValue,
		},
	})
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Get form error
	return util.GetFormError(resp.Body)
}
