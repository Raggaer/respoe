package util

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetFormHash returns a XSRF hash of the given url form
func GetFormHash(url string, c *http.Client) (string, error) {
	getResp, err := c.Get(url)
	if err != nil {
		return "", err
	}

	defer getResp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(getResp.Body)
	if err != nil {
		return "", err
	}

	hashValue, found := doc.Find("input[name='hash']").Attr("value")
	if !found {
		return "", errors.New("Unable to find form hash field")
	}

	return hashValue, nil
}
