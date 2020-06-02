package model

import (
	"encoding/json"
	"fmt"
)

type ComicResponse struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

// Formats the date into ISO standard
func (cr ComicResponse) FormattedDate() string {
	return fmt.Sprintf("%s-%s-%s", cr.Year, cr.Month, cr.Day)
}

// Converts ComicResponse into Comic object
func (cr ComicResponse) Comic() Comic {
	return Comic{
		Title:       cr.Title,
		Number:      cr.Num,
		Date:        cr.FormattedDate(),
		Description: cr.Alt,
		Image:       cr.Img,
	}
}

type Comic struct {
	Title       string `json:"title"`
	Number      int    `json:"number"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func (c Comic) PrettyString() string {
	p := fmt.Sprintf(
		"Title: %s\nComic No: %d\nDate: %s\nDescription:%s \nImage:%s \n",
		c.Title, c.Number, c.Description, c.Image)
	return p
}

func (c Comic) JSON() string {
	cJSON, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(cJSON)
}


