package concepts

type Comment struct {
	Author string `json:"author"`
	At     int64  `json:"at"`
	Text   string `json:"text"`
}
