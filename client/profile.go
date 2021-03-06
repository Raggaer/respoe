package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const (
	viewProfileURL        = "http://www.pathofexile.com/account/view-profile/%s"
	profileCharactersURL  = "https://pathofexile.com/character-window/get-characters?accountName=%s"
	profileCharacterItems = "https://www.pathofexile.com/character-window/get-items"
)

// Profile website account profile
type Profile struct {
	GuildName   string
	GuildURL    string
	GuildID     int
	JoinedAt    time.Time
	ForumPosts  int
	LastVisited time.Time
	Badges      []*Badge
	Characters  []*Character
}

// Badge user profile badge
type Badge struct {
	Name string
	URL  string
}

// Character profile character
type Character struct {
	Name            string
	Level           int
	League          string
	Class           string
	AscendancyClass int `json:"ascendancyClass"`
	ClassID         int `json:"classId"`
	Items           []*Item
}

// CharacterItems items of the profile character
type CharacterItems struct {
	Items []*Item `json:"items"`
}

// GetAccountProfile retrieves the given account profile
func (c *Client) GetAccountProfile(account string) (*Profile, error) {
	response, err := c.HTTP.Get(fmt.Sprintf(viewProfileURL, account))
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	profile := Profile{}

	profileBoxes := doc.Find(".profile-boxes").Children()

	// Retrieve profile basic information
	basicBox := profileBoxes.First().Children().Last()

	// Guild name
	profile.GuildName = basicBox.Children().First().Next().Next().Text()

	if profile.GuildName != "None" {
		// Retrieve guild URL
		guildURL, ok := basicBox.Children().First().Next().Next().Attr("src")
		if ok {
			profile.GuildURL = guildURL

			// Retrieve guild ID from guild URL
			guildID, err := strconv.Atoi(strings.TrimPrefix(profile.GuildURL, "/guild/profile/"))
			if err == nil {
				profile.GuildID = guildID
			}
		}
	}

	// only one node in the selection, the "div"
	n := basicBox.Get(0)
	i := -1
	for nn := n.FirstChild; nn != nil; nn = nn.NextSibling {
		if nn.Type == html.TextNode {
			s := strings.TrimSpace(nn.Data)
			if s != "" {
				i++
				// Forum posts
				if i == 0 {
					forumPosts, err := strconv.Atoi(s)
					if err != nil {
						continue
					}
					profile.ForumPosts = forumPosts
				}

				// Joined date
				if i == 1 {
					joinedDate, err := time.Parse("Jan 2, 2006", s)
					if err != nil {
						continue
					}
					profile.JoinedAt = joinedDate
				}

				// Last visit date
				if i == 2 {
					lastVisitDate, err := time.Parse("Jan 2, 2006", s)
					if err != nil {
						continue
					}
					profile.LastVisited = lastVisitDate
				}
			}
		}
	}

	// Remove children elements leaving only floating text
	basicBox.Children().Remove()

	// Retrieve profile badges
	badgeList := []*Badge{}
	badges := doc.Find(".badges").Children()
	badges.Each(func(i int, s *goquery.Selection) {
		badge := s.Children().First()
		badgeURL, ok := badge.Attr("src")
		if !ok {
			return
		}
		badgeName, ok := badge.Attr("alt")
		if !ok {
			return
		}

		badgeList = append(badgeList, &Badge{
			URL:  badgeURL,
			Name: badgeName,
		})
	})

	// Check if characters are hidden
	characters := false
	doc.Find(".tab-links").Children().Each(func(i int, s *goquery.Selection) {
		if characters {
			return
		}
		if strings.TrimSpace(s.Text()) == "Characters" {
			characters = true
		}
	})

	// Retrieve account characters
	if characters {
		characterList, err := c.ProfileCharacters(account)
		if err != nil {
			return nil, err
		}
		for i, char := range characterList {
			items, err := c.CharacterItems(char.Name, account)
			if err != nil {
				continue
			}
			characterList[i].Items = items
		}
		profile.Characters = characterList
	}

	return &profile, nil
}

// ProfileCharacters retrieves all the characters of the given account
func (c *Client) ProfileCharacters(account string) ([]*Character, error) {
	resp, err := c.HTTP.Get(fmt.Sprintf(profileCharactersURL, account))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	characterList := []*Character{}
	if err := json.Unmarshal(responseBody, &characterList); err != nil {
		return nil, err
	}

	return characterList, nil
}

// CharacterItems retrieves all items of the given character
func (c *Client) CharacterItems(character, account string) ([]*Item, error) {
	resp, err := c.HTTP.PostForm(profileCharacterItems, url.Values{
		"accountName": {
			account,
		},
		"character": {
			character,
		},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	itemList := &CharacterItems{}
	if err := json.Unmarshal(responseBody, &itemList); err != nil {
		return nil, err
	}

	return itemList.Items, nil
}
