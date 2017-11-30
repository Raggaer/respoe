package client

import (
	"net/http"
	"net/http/cookiejar"

	"golang.org/x/net/publicsuffix"
)

// Client client used for all path of exile website applications
type Client struct {
	HTTP   *http.Client
	Logged bool
}

// New returns a new client
func New() (*Client, error) {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}

	return &Client{
		HTTP: &http.Client{
			Jar: jar,
		},
		Logged: false,
	}, nil
}
