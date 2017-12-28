package forum

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/raggaer/respoe/client"
	"github.com/raggaer/respoe/util"
)

const threadReplyURL = "https://www.pathofexile.com/forum/post-reply/%d"

// Reply posts a new reply for the given thread
func (t *Thread) Reply(msg string, c *client.Client) error {
	if !c.Logged {
		return errors.New("You need to be logged in to reply a thread")
	}

	replyURL := fmt.Sprintf(threadReplyURL, t.ID)

	// Retrieve XSRF value
	hashValue, err := util.GetReplyHash(replyURL, c.HTTP)
	if err != nil {
		return err
	}

	// Submit form with the reply message
	resp, err := c.HTTP.PostForm(replyURL, url.Values{
		"content": {
			msg,
		},
		"forum_post": {
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
