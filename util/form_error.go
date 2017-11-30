package util

import (
	"errors"
	"io"

	"github.com/PuerkitoBio/goquery"
)

// GetFormError returns the first form error
func GetFormError(body io.ReadCloser) error {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return err
	}

	// Find all form errors
	errs := doc.Find("ul.errors")

	if errs.Length() == 0 {
		return nil
	}

	// Retrieve first error
	f := errs.First()

	// Retrieve error input field
	// We assume the input field is previous to the ul.errors div
	inputField, ok := f.Prev().Attr("name")

	if !ok {
		return errors.New(f.Text())
	}

	return errors.New(inputField + "- " + f.Text())
}
