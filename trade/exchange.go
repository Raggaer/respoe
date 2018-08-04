package trade

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/raggaer/respoe/client"
)

const (
	exchangeURL       = "https://www.pathofexile.com/api/trade/exchange/%s"
	exchangeOffersURL = "https://www.pathofexile.com/api/trade/fetch/%s?query=%s&exchange"
)

// ExchangeQuery struct used to create exchange queries
type ExchangeQuery struct {
	Exchange Exchange `json:"exchange"`
}

// Exchange defines a exchange query
type Exchange struct {
	Have   []string       `json:"have"`
	Want   []string       `json:"want"`
	Status ExchangeStatus `json:"status"`
}

// ExchangeStatus the exchange online status
type ExchangeStatus struct {
	Option string `json:"option"`
}

// ExchangeEndpoints list of exchange endpoints
type ExchangeEndpoints struct {
	Endpoints []string
}

// ExchangeResponse struct used for exchange request responses
type ExchangeResponse struct {
	Result []string `json:"result"`
	Id     string   `json:"id"`
	Total  int      `json:"total"`
}

// ExchangeOffersResponse struct used for exchange offers
type ExchangeOffersResponse struct {
	Result []*ExchangeOffer `json:"result"`
}

// ExchangeOffer information about a currency exchange offer
type ExchangeOffer struct {
	Id      string
	Item    *client.Item
	Listing ExchangeOfferListing
}

// ExchangeOfferListing listing information about a currency exchange
type ExchangeOfferListing struct {
	Method    string
	Whisper   string
	Indexed   string
	IndexedAt *time.Time
	Price     ExchangeOfferPrice
}

// ExchangeOfferPrice price information about a currency exchange
type ExchangeOfferPrice struct {
	Want ExchangeOfferItem `json:"item"`
	Have ExchangeOfferItem `json:"exchange"`
}

// ExchangeOfferItem offer detailed information
type ExchangeOfferItem struct {
	Currency string
	Amount   int
	Stock    int
	Id       string
}

// RetrieveExchange retrieves the current exchange offers
func RetrieveExchange(league string, have, want []string, c *client.Client) ([]*ExchangeOffer, error) {
	// Encode exchange struct into JSON
	ex, err := json.Marshal(&ExchangeQuery{
		Exchange: Exchange{
			Have: have,
			Want: want,
			Status: ExchangeStatus{
				Option: "any",
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Create request object so we can add headers
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf(exchangeURL, league), bytes.NewReader(ex))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	exchangeResponse := ExchangeResponse{}
	if err := json.Unmarshal(respContent, &exchangeResponse); err != nil {
		return nil, err
	}

	// Retrieve deals response
	offers, err := retrieveExchangeOffers(&exchangeResponse, c)
	return offers, err
}

func retrieveExchangeOffers(e *ExchangeResponse, c *client.Client) ([]*ExchangeOffer, error) {
	d := ""
	i := 0
	for _, r := range e.Result {
		if i == 18 {
			d += r
			break
		} else {
			d += r + ","
		}
		i++
	}
	resp, err := c.HTTP.Get(fmt.Sprintf(exchangeOffersURL, d, e.Id))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	offers := ExchangeOffersResponse{}
	if err := json.Unmarshal(respBody, &offers); err != nil {
		return nil, err
	}
	return offers.Result, nil
}
