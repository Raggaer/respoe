package forum

import (
	"encoding/json"
	"strings"
)

// Item defines a path of exile forum item
type Item struct {
	Verified             bool
	Support              bool
	W                    int
	H                    int
	SecondaryDescription string `json:"secDescrText"`
	Description          string `json:"descrText"`
	ItemLevel            int    `json:"ilvl"`
	Icon                 string
	League               string
	ID                   string `json:"id"`
	Name                 string
	TypeLine             string `json:"typeline"`
	Identified           bool
	FrameType            int            `json:"frameType"`
	ImplicitMods         []string       `json:"implicitMods"`
	CraftedMods          []string       `json:"craftedMods"`
	ExplicitMods         []string       `json:"explicitMods"`
	FlavourText          []string       `json:"flavourText"`
	Properties           []ItemProperty `json:"properties"`
}

// ItemProperty defines a set of item property values
type ItemProperty struct {
	Name        string
	DisplayMode int `json:"displayMode"`
	Type        int
	Values      [][]interface{}
}

// ParseItem parses a Path of Exile forum item
func ParseItem(data []byte) (*Item, error) {
	item := Item{}

	if err := json.Unmarshal(data, &item); err != nil {
		return nil, err
	}

	item.Name = strings.TrimPrefix(item.Name, "<<set:MS>><<set:M>><<set:S>>")
	item.TypeLine = strings.TrimPrefix(item.TypeLine, "<<set:MS>><<set:M>><<set:S>>")

	return &item, nil
}
