package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define the structure of the YouTube RSS feed
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

type Entry struct {
	ID        string `xml:"id"`
	VideoId   string `xml:"videoId"`
	ChannelId string `xml:"channelId"`
	Title     string `xml:"title"`
	Link      Link   `xml:"link"`
	Author    Author `xml:"author"`
	Published string `xml:"published"`
	Updated   string `xml:"updated"`
	Group     Group  `xml:"group"`
}

type Link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
}

type Author struct {
	Name string `xml:"name"`
	URI  string `xml:"uri"`
}

type Group struct {
	Title       string    `xml:"title"`
	Content     Content   `xml:"content"`
	Thumbnail   Thumbnail `xml:"thumbnail"`
	Description string    `xml:"description"`
	Community   Community `xml:"community"`
}

type Content struct {
	URL    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
}

type Thumbnail struct {
	URL    string `xml:"url,attr"`
	Width  string `xml:"width,attr"`
	Height string `xml:"height,attr"`
}

type Community struct {
	StarRating StarRating `xml:"starRating"`
	Statistics Statistics `xml:"statistics"`
}

type StarRating struct {
	Count   string `xml:"count,attr"`
	Average string `xml:"average,attr"`
	Min     string `xml:"min,attr"`
	Max     string `xml:"max,attr"`
}

type Statistics struct {
	Views string `xml:"views,attr"`
}

func main() {
	// YouTube channel url
	url := "https://www.youtube.com/feeds/videos.xml?channel_id=UC_O58Rr2DOskJvs9bArpLkQ"

	// Fetch the RSS feed
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching RSS feed: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	// Print the raw XML response for debugging
	// fmt.Println("Raw XML Response:")
	// fmt.Println(string(body))

	// Parse the RSS feed
	var feed Feed
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		log.Fatalf("Error parsing XML: %v", err)
	}

	// Print each entry
	for _, entry := range feed.Entries {
		fmt.Println("Author:", entry.Author.Name)
		fmt.Println("Title:", entry.Title)
		fmt.Println("Link:", entry.Link.Href)
		fmt.Println("Published:", entry.Published)
		fmt.Println("Updated:", entry.Updated)
		fmt.Println("Thumbnail URL:", entry.Group.Thumbnail.URL)
		fmt.Println("Description:", entry.Group.Description)
		fmt.Println()
	}
}
