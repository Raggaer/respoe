package forum

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/raggaer/respoe/client"
)

// Pagination forum pagination data
type Pagination struct {
	Current int
	First   int
	Last    int
}

const forumMainURL = "https://www.pathofexile.com"

// GetPagination returns the forum pagination struct
func (f *Forum) GetPagination(c *client.Client) (*Pagination, error) {
	resp, err := c.HTTP.Get(forumMainURL + f.URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	paginationDiv := doc.Find("div.pagination")

	// Get first element
	firstPage := paginationDiv.Children().First()

	pag := &Pagination{}

	// Check if first page is "Prev" button
	if firstPage.Text() == "Prev" {
		firstPage = firstPage.Next()
	}

	firstPageInt, err := strconv.Atoi(firstPage.Text())
	if err != nil {
		return nil, err
	}

	pag.First = firstPageInt

	// Get last page
	lastPage := paginationDiv.Children().Last()

	// Check if last page is "Next" button
	if lastPage.Text() == "Next" {
		lastPage = lastPage.Prev()
	}

	lastPageInt, err := strconv.Atoi(lastPage.Text())
	if err != nil {
		return nil, err
	}

	pag.Last = lastPageInt

	// Find current page
	currentPage := paginationDiv.ChildrenFiltered("a.current").Text()

	currentPageInt, err := strconv.Atoi(currentPage)
	if err != nil {
		return nil, err
	}

	pag.Current = currentPageInt

	return pag, nil
}
