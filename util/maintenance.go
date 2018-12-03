package util

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// IsMaintenance checks if the document is on maintenance mode
func IsMaintenance(doc *goquery.Document) bool {
	return strings.Contains(doc.Text(), "Down For Maintenance")
}
