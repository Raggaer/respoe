package client

import "errors"

const logoutURL = "https://www.pathofexile.com/logout"

// Logout closes the running session
// Path of Exile website uses a GET request for this task (not the best choice)
func (c *Client) Logout() error {
	if !c.Logged {
		return errors.New("You need to be logged in to change your client password")
	}

	if _, err := c.HTTP.Get(logoutURL); err != nil {
		return err
	}

	return nil
}
