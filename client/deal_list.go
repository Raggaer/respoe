package client

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const specialDealsURL = "https://www.pathofexile.com/shop/category/daily-deals"

// GetDealList returns the list of current running deals
func (c *Client) GetDealList() ([]*Deal, error) {
	resp, err := c.HTTP.Get(specialDealsURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	dealList := []*Deal{}

	shopItems := doc.Find("div.shopItems").Children()
	shopItems.Each(func(i int, s *goquery.Selection) {
		d := &Deal{}

		// Retrieve deal name
		d.Name = s.Children().NextFiltered("a.name").Text()

		// Retrieve deal price
		dealPrice, err := strconv.Atoi(s.Children().NextFiltered("div.price").Text())
		if err != nil {
			return
		}

		d.Price = dealPrice

		// Retrieve deal description
		d.Description = strings.TrimSpace(s.Children().NextFiltered("div.description").Text())

		dealList = append(dealList, d)
	})

	return dealList, nil
}
