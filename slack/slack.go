package slack

import "encoding/json"

func Parse(payload string) (*SlackMessage, error) {
	m := new(SlackMessage)

	if err := json.Unmarshal([]byte(payload), m); err != nil {
		return nil, err
	}

	return m, nil
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

type Attachments struct {
	Text       string  `json:"text"`
	Title      string  `json:"title"`
	TitleLink  string  `json:"title_link"`
	Fallback   string  `json:"fallback"`
	Color      string  `json:"color"`
	AuthorName string  `json:"author_name"`
	AuthorIcon string  `json:"author_icon"`
	AuthorLink string  `json:"author_link"`
	Fields     []Field `json:"fields"`
	Footer     string  `json:"footer"`
	FooterIcon string  `json:"footer_icon"`
}

type SlackMessage struct {
	Channel     string        `json:"channel"`
	Username    string        `json:"username"`
	Fallback    string        `json:"fallback"`
	Text        string        `json:"text"`
	Attachments []Attachments `json:"attachments"`
}
