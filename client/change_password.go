package client

import (
	"errors"
	"net/url"

	"github.com/raggaer/respoe/util"
)

const changePasswordFormURL = "https://www.pathofexile.com/my-account/change-password"

// ChangePassword changes the client password
func (c *Client) ChangePassword(current, new string) error {
	if !c.Logged {
		return errors.New("You need to be logged in to change your client password")
	}

	hashValue, err := util.GetFormHash(changePasswordFormURL, c.HTTP)
	if err != nil {
		return err
	}

	resp, err := c.HTTP.PostForm(changePasswordFormURL, url.Values{
		"old_password": {
			current,
		},
		"new_password": {
			new,
		},
		"password_confirmation": {
			new,
		},
		"hash": {
			hashValue,
		},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return util.GetFormError(resp.Body)
}
