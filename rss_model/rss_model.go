package rss_model

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Channel struct {
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	Description string  `xml:"description"`
	Language    string  `xml:"language"`
	CopyRight   string  `xml:"copyright"`
	PubDate     string  `xml:"pubDate"`
	TimeToLive  int32   `xml:"ttl"`
	Items       []*Item `xml:"item"`
}

type Item struct {
	Title       string    `xml:"title"`
	GUID        string    `xml:"guid"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Enclosure   Enclosure `xml:"enclosure"`
	PubDate     string    `xml:"pubDate"`
}

type Enclosure struct {
	Url  string `xml:"url,attr"`
	Type string `xml:"type,attr"`
}

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

func RssGenerator(c *fiber.Ctx) error {

	timeFormat := time.Now().Format(time.RFC3339)

	items := []*Item{
		{
			Title:       "Article 1",
			Link:        "https://example.com/article",
			Description: "This is the description article 1",
			PubDate:     timeFormat,
		},
		{
			Title:       "Article 1",
			Link:        "https://example.com/article",
			Description: "This is the description article 1",
			PubDate:     timeFormat,
		},
	}

	channel := Channel{
		Title:       "Sample RSS Feed",
		Link:        "https://techmovement.co.th",
		Description: "This is a sample RSS feed generated using Golang",
		PubDate:     timeFormat,
		Items:       items,
	}

	feed := RSS{
		Version: "2.0",
		Channel: channel,
	}

	xmlData, err := xml.MarshalIndent(feed, "", "    ")
	if err != nil {
		return err
	}

	rssFeed := []byte(xml.Header + string(xmlData))

	file, err := os.Create("./rss_model/feed.xml")
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(rssFeed)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": fiber.StatusOK,
	})
}

func ServeRss(c *fiber.Ctx) error {

	short := c.Query("link")

	rssURL, ok := Links[short]
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid link",
		})
	}

	resp, err := http.Get(rssURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	rss := new(RSS)

	xml.Unmarshal(data, &rss)

	response := ToResponse(rss.Channel.Items)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data":    response,
		"message": fiber.StatusOK,
	})
}
