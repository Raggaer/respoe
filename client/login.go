package client

import (
	"errors"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const loginFormURL = "https://www.pathofexile.com/login"

// Login logins the given account inside the path of exile website
func (c *Client) Login(email, password string) error {
	getResp, err := c.c.Get(loginFormURL)
	if err != nil {
		return err
	}

	defer getResp.Body.Close()

	loginDocument, err := goquery.NewDocumentFromReader(getResp.Body)
	if err != nil {
		return err
	}

	hashValue, found := loginDocument.Find("input[name='hash']").Attr("value")
	if !found {
		return errors.New("Unable to find login hash field")
	}

	resp, err := c.c.PostForm(loginFormURL, url.Values{
		"login_email": {
			email,
		},
		"login_password": {
			password,
		},
		"hash": {
			hashValue,
		},
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if logged was successfull
	loginDocument, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	// We try to find <span class="profile-link"> to check if we are logged in
	spanSelector := loginDocument.Find("span.profile-link")
	if spanSelector.Text() == "" {
		return errors.New("Unable to login. Probably wrong email or password")
	}

	return nil
}
