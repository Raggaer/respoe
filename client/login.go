package client

import (
	"net/url"

	"github.com/raggaer/respoe/util"
)

const loginFormURL = "https://www.pathofexile.com/login"

// Login logins the given account inside the path of exile website
func (c *Client) Login(email, password string) error {
	hashValue, err := util.GetFormHash(loginFormURL, c.HTTP)
	if err != nil {
		return err
	}

	resp, err := c.HTTP.PostForm(loginFormURL, url.Values{
		"login_email": {
			email,
		},
		"login_password": {
			password,
		},
		"remember_me": {
			"1",
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
	err = util.GetFormError(resp.Body)

	// Update logged field
	if err == nil {
		c.Logged = true
	}

	return err
}
