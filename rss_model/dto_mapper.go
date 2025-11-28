package rss_model

type Response struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"url"`
	Image       Image  `json:"image"`
	CreatedAt   string `json:"created_at"`
}

type Image struct {
	Url  string `json:"url"`
	Type string `json:"type"`
}

func ToResponse(items []*Item) []*Response {

	var itemsResponse []*Response

	for _, v := range items {

		resp := &Response{
			Title:       v.Title,
			Description: v.Description,
			Link:        v.Link,
			Image:       Image(v.Enclosure),
			CreatedAt:   v.PubDate,
		}

		itemsResponse = append(itemsResponse, resp)
	}

	return itemsResponse
}
